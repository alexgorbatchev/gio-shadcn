/*
Package button provides a versatile button component for gio-shadcn applications.
*/
package button

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

type Button struct {
	clickable *widget.Clickable

	Text     string
	Variant  theme.Variant
	Size     theme.Size
	Icon     *widget.Icon
	Disabled bool
	Classes  string
	OnClick  func()

	cachedStyles     utils.StyleUtility
	cachedClasses    string
	stylesCacheValid bool
}

type Option func(*Button)

func WithVariant(variant theme.Variant) Option {
	return func(b *Button) {
		b.Variant = variant
	}
}

func WithSize(size theme.Size) Option {
	return func(b *Button) {
		b.Size = size
	}
}

func WithText(text string) Option {
	return func(b *Button) {
		b.Text = text
	}
}

func WithIcon(icon *widget.Icon) Option {
	return func(b *Button) {
		b.Icon = icon
	}
}

func WithOnClick(onClick func()) Option {
	return func(b *Button) {
		b.OnClick = onClick
	}
}

func WithDisabled(disabled bool) Option {
	return func(b *Button) {
		b.Disabled = disabled
	}
}

func WithClasses(classes string) Option {
	return func(b *Button) {
		b.Classes = classes
	}
}

func NewButton(options ...Option) *Button {
	b := &Button{
		clickable: &widget.Clickable{},
		Variant:   theme.VariantDefault,
		Size:      theme.SizeDefault,
	}

	for _, option := range options {
		option(b)
	}

	return b
}

func ValidateButton(b *Button) error {
	if b == nil {
		return fmt.Errorf("button cannot be nil")
	}

	if b.clickable == nil {
		return fmt.Errorf("button must have a clickable widget")
	}

	if b.Text == "" && b.Icon == nil {
		return fmt.Errorf("button must have either text or icon")
	}

	return nil
}

func (b *Button) SafeLayout(gtx layout.Context, th *theme.Theme) (layout.Dimensions, error) {
	if err := ValidateButton(b); err != nil {
		return layout.Dimensions{}, err
	}

	if th == nil {
		return layout.Dimensions{}, fmt.Errorf("theme cannot be nil")
	}

	return b.Layout(gtx, th), nil
}

type Config struct {
	Text     string
	Variant  theme.Variant
	Size     theme.Size
	Icon     *widget.Icon
	Disabled bool
	Classes  string
	OnClick  func()
}

func New(config Config) *Button {
	return &Button{
		clickable: new(widget.Clickable),
		Text:      config.Text,
		Variant:   config.Variant,
		Size:      config.Size,
		Icon:      config.Icon,
		Disabled:  config.Disabled,
		Classes:   config.Classes,
		OnClick:   config.OnClick,
	}
}

func (b *Button) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if b.clickable.Clicked(gtx) && !b.Disabled && b.OnClick != nil {
		b.OnClick()
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

	return b.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return b.drawButton(gtx, th, bgColor, fgColor, variant, padding, minHeight, fontSize, styles)
	})
}

func (b *Button) Update(gtx layout.Context) theme.ComponentState {
	return &State{
		active:   b.clickable.Clicked(gtx),
		hovered:  b.clickable.Hovered(),
		pressed:  b.clickable.Pressed(),
		disabled: b.Disabled,
	}
}

type State struct {
	active   bool
	hovered  bool
	pressed  bool
	disabled bool
}

func (bs *State) IsActive() bool   { return bs.active }
func (bs *State) IsHovered() bool  { return bs.hovered }
func (bs *State) IsPressed() bool  { return bs.pressed }
func (bs *State) IsDisabled() bool { return bs.disabled }

func (b *Button) drawButton(gtx layout.Context, th *theme.Theme, bgColor, fgColor color.NRGBA, variant theme.VariantConfig, padding layout.Inset, minHeight unit.Dp, fontSize unit.Sp, styles utils.StyleUtility) layout.Dimensions {
	radius := th.Radius.RadiusMD
	if styles.Radius > 0 {
		radius = styles.Radius
	}

	// Measure content with unconstrained max bounds to prevent 800px width expansion
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
	switch {
	case b.Icon != nil && b.Text != "":
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return b.Icon.Layout(gtx, fgColor)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return b.layoutText(gtx, th, fgColor, fontSize)
			}),
		)
	case b.Icon != nil:
		return b.Icon.Layout(gtx, fgColor)
	default:
		return b.layoutText(gtx, th, fgColor, fontSize)
	}
}

func (b *Button) layoutText(gtx layout.Context, th *theme.Theme, fgColor color.NRGBA, fontSize unit.Sp) layout.Dimensions {
	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}
	label := material.Label(mTheme, fontSize, b.Text)
	label.Color = fgColor
	label.Alignment = text.Start
	return label.Layout(gtx)
}

func (b *Button) getSizeConfig(th *theme.Theme) (layout.Inset, unit.Dp, unit.Sp) {
	switch b.Size {
	case theme.SizeSM:
		return layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space3,
			Right:  th.Spacing.Space3,
		}, unit.Dp(32), th.Typography.FontSizeSM

	case theme.SizeLG:
		return layout.Inset{
			Top:    th.Spacing.Space3,
			Bottom: th.Spacing.Space3,
			Left:   th.Spacing.Space8,
			Right:  th.Spacing.Space8,
		}, unit.Dp(44), th.Typography.FontSizeBase

	case theme.SizeIcon:
		return layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space2,
			Right:  th.Spacing.Space2,
		}, unit.Dp(36), th.Typography.FontSizeSM

	default:
		return layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space4,
			Right:  th.Spacing.Space4,
		}, unit.Dp(36), th.Typography.FontSizeSM
	}
}

func (b *Button) Clicked(gtx layout.Context) bool {
	return b.clickable.Clicked(gtx) && !b.Disabled
}

func (b *Button) SetDisabled(disabled bool) {
	b.Disabled = disabled
}

func (b *Button) SetText(text string) {
	b.Text = text
}

func (b *Button) SetVariant(variant theme.Variant) {
	b.Variant = variant
}

func (b *Button) SetOnClick(onClick func()) {
	b.OnClick = onClick
}
