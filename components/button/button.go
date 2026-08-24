/*
Package button provides an interactive button component for gio-shadcn applications.

Buttons trigger actions or events following
shadcn/ui design principles.
*/
package button

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

type Button struct {
	Text      string
	Variant   theme.Variant
	Size      theme.Size
	Disabled  bool
	Classes   string
	Icon      *lucide.Icon
	OnClick   func()
	clickable widget.Clickable

	cachedClasses    string
	cachedStyles     utils.StyleUtility
	stylesCacheValid bool
}

type Config struct {
	Text     string
	Variant  theme.Variant
	Size     theme.Size
	Disabled bool
	Classes  string
	Icon     *lucide.Icon
	OnClick  func()
}

func New(config Config) *Button {
	v := config.Variant
	if v == "" {
		v = theme.VariantDefault
	}
	s := config.Size
	if s == "" {
		s = theme.SizeDefault
	}
	return &Button{
		Text:     config.Text,
		Variant:  v,
		Size:     s,
		Disabled: config.Disabled,
		Classes:  config.Classes,
		Icon:     config.Icon,
		OnClick:  config.OnClick,
	}
}

func (b *Button) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	variant := theme.GetButtonVariant(b.Variant, &th.Colors)
	padding, minHeight, fontSize := b.getSizeConfig(th)

	var styles utils.StyleUtility
	if !b.stylesCacheValid || b.cachedClasses != b.Classes {
		styles = utils.ParseClasses(b.Classes)
		b.cachedStyles = styles
		b.cachedClasses = b.Classes
		b.stylesCacheValid = true
	} else {
		styles = b.cachedStyles
	}

	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}

	bgColor := variant.Background
	fgColor := variant.Foreground

	switch {
	case b.Disabled:
		bgColor = variant.DisabledBg
		fgColor = variant.DisabledFg
	case b.clickable.Pressed():
		bgColor = variant.ActiveBg
		fgColor = variant.ActiveFg
	case b.clickable.Hovered():
		bgColor = variant.HoverBg
		fgColor = variant.HoverFg
	}

	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	dims := b.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return b.drawButton(gtx, th, bgColor, fgColor, variant, padding, minHeight, fontSize, styles)
	})

	if b.clickable.Clicked(gtx) && !b.Disabled && b.OnClick != nil {
		b.OnClick()
	}

	return dims
}

func (b *Button) Update(gtx layout.Context) theme.ComponentState {
	return &State{
		active:   b.clickable.Clicked(gtx),
		hovered:  b.clickable.Hovered(),
		pressed:  b.clickable.Pressed(),
		disabled: b.Disabled,
	}
}

func (b *Button) SetText(text string) {
	b.Text = text
}

func (b *Button) drawButton(gtx layout.Context, th *theme.Theme, bgColor, fgColor color.NRGBA, variant theme.VariantConfig, padding layout.Inset, minHeight unit.Dp, fontSize unit.Sp, styles utils.StyleUtility) layout.Dimensions {
	radius := th.Radius.RadiusMD
	if styles.Radius > 0 {
		radius = styles.Radius
	}

	gtxContent := gtx
	gtxContent.Constraints = layout.Constraints{
		Min: image.Pt(0, 0),
		Max: image.Pt(1e6, 1e6),
	}

	contentDims := padding.Layout(gtxContent, func(gtx layout.Context) layout.Dimensions {
		return b.layoutContent(gtx, th, fgColor, fontSize)
	})

	finalHeight := contentDims.Size.Y
	if finalHeight < gtx.Dp(minHeight) {
		finalHeight = gtx.Dp(minHeight)
	}

	btnSize := image.Pt(contentDims.Size.X, finalHeight)
	gtx.Constraints = layout.Exact(btnSize)

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: btnSize}
			radiusPx := gtx.Dp(radius)

			theme.DrawRRectBackground(gtx, rect, radiusPx, bgColor)

			if variant.BorderWidth > 0 {
				rr := clip.UniformRRect(rect, radiusPx)
				theme.DrawStroke(gtx, rr.Path(gtx.Ops), variant.BorderWidth, variant.Border)
			}

			return layout.Dimensions{Size: btnSize}
		}),

		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return b.layoutContent(gtx, th, fgColor, fontSize)
				})
			})
		}),
	)
}

func (b *Button) layoutContent(gtx layout.Context, th *theme.Theme, fgColor color.NRGBA, fontSize unit.Sp) layout.Dimensions {
	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	iconSize := unit.Dp(16)
	if b.Size == theme.SizeSM {
		iconSize = unit.Dp(14)
	} else if b.Size == theme.SizeLG {
		iconSize = unit.Dp(20)
	}

	if b.Icon != nil && b.Text != "" {
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return b.Icon.LayoutSize(gtx, iconSize, fgColor)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(mTheme, fontSize, b.Text)
				lbl.Color = fgColor
				lbl.Alignment = text.Start
				return lbl.Layout(gtx)
			}),
		)
	}

	if b.Icon != nil {
		return b.Icon.LayoutSize(gtx, iconSize, fgColor)
	}

	lbl := material.Label(mTheme, fontSize, b.Text)
	lbl.Color = fgColor
	lbl.Alignment = text.Start

	return lbl.Layout(gtx)
}

func (b *Button) getSizeConfig(th *theme.Theme) (padding layout.Inset, minHeight unit.Dp, fontSize unit.Sp) {
	switch b.Size {
	case theme.SizeSM:
		return layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space3,
			Right:  th.Spacing.Space3,
		}, unit.Dp(32), th.Typography.FontSizeXS
	case theme.SizeLG:
		return layout.Inset{
			Top:    th.Spacing.Space3,
			Bottom: th.Spacing.Space3,
			Left:   th.Spacing.Space6,
			Right:  th.Spacing.Space6,
		}, unit.Dp(48), th.Typography.FontSizeBase
	case theme.SizeIcon:
		return layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space2,
			Right:  th.Spacing.Space2,
		}, unit.Dp(40), th.Typography.FontSizeSM
	default: // SizeDefault
		return layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space4,
			Right:  th.Spacing.Space4,
		}, unit.Dp(40), th.Typography.FontSizeSM
	}
}

type State struct {
	active   bool
	hovered  bool
	pressed  bool
	disabled bool
}

func (s *State) IsActive() bool   { return s.active }
func (s *State) IsHovered() bool  { return s.hovered }
func (s *State) IsPressed() bool  { return s.pressed }
func (s *State) IsDisabled() bool { return s.disabled }
