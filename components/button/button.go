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
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
	"github.com/vibrantgio/effects/spring"
	"github.com/vibrantgio/effects/tween"
)

type Position int

const (
	PositionSingle Position = iota
	PositionFirst
	PositionMiddle
	PositionLast
)

type Button struct {
	Text      string
	Variant   theme.Variant
	Size      theme.Size
	Disabled  bool
	Classes   string
	Icon      *lucide.Icon
	Position  Position
	OnClick   func()
	clickable widget.Clickable

	hoverSpring      *spring.Spring
	initialized      bool
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
	Position Position
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
		Position: config.Position,
		OnClick:  config.OnClick,
	}
}

// ButtonGroup represents a connected segmented group of buttons with optional single selection.
type ButtonGroup struct {
	Buttons       []*Button
	SelectedIndex int
	OnSelect      func(index int)
}

// NewGroup creates a new connected ButtonGroup with connected left/middle/right corner positions.
func NewGroup(buttons ...*Button) *ButtonGroup {
	bg := &ButtonGroup{
		Buttons:       buttons,
		SelectedIndex: -1,
	}
	for i, btn := range buttons {
		idx := i
		if len(buttons) == 1 {
			btn.Position = PositionSingle
		} else if i == 0 {
			btn.Position = PositionFirst
		} else if i == len(buttons)-1 {
			btn.Position = PositionLast
		} else {
			btn.Position = PositionMiddle
		}

		origOnClick := btn.OnClick
		btn.OnClick = func() {
			bg.SelectedIndex = idx
			if origOnClick != nil {
				origOnClick()
			}
			if bg.OnSelect != nil {
				bg.OnSelect(idx)
			}
		}
	}
	return bg
}

// Select programmatically selects the button at the specified index.
func (bg *ButtonGroup) Select(index int) {
	if index >= 0 && index < len(bg.Buttons) {
		bg.SelectedIndex = index
		if bg.Buttons[index].OnClick != nil {
			bg.Buttons[index].OnClick()
		}
	}
}

// Layout renders the connected buttons horizontally without gaps and highlights the active selection.
func (bg *ButtonGroup) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	children := make([]layout.FlexChild, len(bg.Buttons))
	for i, btn := range bg.Buttons {
		b := btn
		isSelected := (i == bg.SelectedIndex)
		children[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if isSelected {
				origVariant := b.Variant
				b.Variant = theme.VariantSecondary
				dims := b.Layout(gtx, th)
				b.Variant = origVariant
				return dims
			}
			return b.Layout(gtx, th)
		})
	}
	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)
}

func (b *Button) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	variant := theme.GetButtonVariant(b.Variant, &th.Colors)
	padding, targetHeight, fontSize := b.getSizeConfig(th)

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

	if !b.initialized {
		b.hoverSpring = spring.New(0, 0, spring.Options{
			Stiffness: 300.0,
			Damping:   26.0,
		})
		b.initialized = true
	}

	targetHover := 0.0
	if b.clickable.Hovered() && !b.Disabled {
		targetHover = 1.0
	}
	b.hoverSpring.SetTarget(targetHover)
	b.hoverSpring.Tick(2.0)
	hoverVal := b.hoverSpring.Value()
	if !b.hoverSpring.Settled(0.005) {
		gtx.Execute(op.InvalidateCmd{})
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
	case hoverVal > 0.001:
		bgColor = tween.LerpNRGBA(variant.Background, variant.HoverBg, hoverVal)
		fgColor = tween.LerpNRGBA(variant.Foreground, variant.HoverFg, hoverVal)
	}

	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	for b.clickable.Clicked(gtx) {
		if !b.Disabled && b.OnClick != nil {
			b.OnClick()
		}
	}

	return b.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return b.drawButton(gtx, th, bgColor, fgColor, variant, padding, targetHeight, fontSize, styles)
	})
}

func (b *Button) Click() {
	b.clickable.Click()
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

func (b *Button) drawButton(gtx layout.Context, th *theme.Theme, bgColor, fgColor color.NRGBA, variant theme.VariantConfig, padding layout.Inset, targetHeight unit.Dp, fontSize unit.Sp, styles utils.StyleUtility) layout.Dimensions {
	radius := th.Radius.RadiusMD
	if styles.Radius > 0 {
		radius = styles.Radius
	}

	heightPx := gtx.Dp(targetHeight)

	gtxContent := gtx
	gtxContent.Constraints = layout.Constraints{
		Min: image.Pt(0, 0),
		Max: image.Pt(1e6, 1e6),
	}

	horizPadding := layout.Inset{
		Left:  padding.Left,
		Right: padding.Right,
	}

	// 1. Record layout operations into a macro so no operations leak onto gtx.Ops during measurement
	macro := op.Record(gtx.Ops)
	contentDims := horizPadding.Layout(gtxContent, func(gtx layout.Context) layout.Dimensions {
		return b.layoutContent(gtx, th, fgColor, fontSize)
	})
	callOp := macro.Stop()

	btnWidth := contentDims.Size.X
	if b.Size == theme.SizeIcon && btnWidth < heightPx {
		btnWidth = heightPx
	}

	btnSize := image.Pt(btnWidth, heightPx)
	gtx.Constraints = layout.Exact(btnSize)

	return layout.Stack{Alignment: layout.Center}.Layout(gtx,
		// Background & Border drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: btnSize}
			radiusPx := gtx.Dp(radius)

			var rr clip.RRect
			rr.Rect = rect

			switch b.Position {
			case PositionFirst:
				rr.NW = radiusPx
				rr.SW = radiusPx
				rr.NE = 0
				rr.SE = 0
			case PositionMiddle:
				rr.NW = 0
				rr.SW = 0
				rr.NE = 0
				rr.SE = 0
			case PositionLast:
				rr.NW = 0
				rr.SW = 0
				rr.NE = radiusPx
				rr.SE = radiusPx
			default: // PositionSingle
				rr.NW = radiusPx
				rr.SW = radiusPx
				rr.NE = radiusPx
				rr.SE = radiusPx
			}

			if bgColor.A > 0 {
				cl := rr.Op(gtx.Ops).Push(gtx.Ops)
				paint.ColorOp{Color: bgColor}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
				cl.Pop()
			}

			if variant.BorderWidth > 0 {
				theme.DrawStroke(gtx, rr.Path(gtx.Ops), variant.BorderWidth, variant.Border)
			}

			return layout.Dimensions{Size: btnSize}
		}),

		// Recorded Content played EXACTLY ONCE, centered within button bounds
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				callOp.Add(gtx.Ops)
				return contentDims
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

func (b *Button) getSizeConfig(th *theme.Theme) (padding layout.Inset, targetHeight unit.Dp, fontSize unit.Sp) {
	switch b.Size {
	case theme.SizeSM:
		return layout.Inset{
			Left:  th.Spacing.Space3,
			Right: th.Spacing.Space3,
		}, unit.Dp(32), th.Typography.FontSizeXS
	case theme.SizeLG:
		return layout.Inset{
			Left:  th.Spacing.Space6,
			Right: th.Spacing.Space6,
		}, unit.Dp(40), th.Typography.FontSizeBase
	case theme.SizeIcon:
		return layout.Inset{
			Left:  th.Spacing.Space2,
			Right: th.Spacing.Space2,
		}, unit.Dp(36), th.Typography.FontSizeSM
	default: // SizeDefault
		return layout.Inset{
			Left:  th.Spacing.Space4,
			Right: th.Spacing.Space4,
		}, unit.Dp(36), th.Typography.FontSizeSM
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
