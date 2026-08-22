/*
Package card provides a flexible container component for gio-shadcn applications.
*/
package card

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

type Card struct {
	Variant theme.Variant
	Classes string
	Padding layout.Inset
}

type Option func(*Card)

func WithCardVariant(variant theme.Variant) Option {
	return func(c *Card) {
		c.Variant = variant
	}
}

func WithCardClasses(classes string) Option {
	return func(c *Card) {
		c.Classes = classes
	}
}

func WithCardPadding(padding layout.Inset) Option {
	return func(c *Card) {
		c.Padding = padding
	}
}

func NewCard(options ...Option) *Card {
	c := &Card{
		Variant: theme.VariantDefault,
		Padding: layout.Inset{Top: 24, Right: 24, Bottom: 24, Left: 24},
	}

	for _, option := range options {
		option(c)
	}

	return c
}

type Config struct {
	Variant theme.Variant
	Classes string
	Padding layout.Inset
}

func New(config Config) *Card {
	return &Card{
		Variant: config.Variant,
		Classes: config.Classes,
		Padding: config.Padding,
	}
}

// Layout renders the card background FIRST, then content ON TOP using op.Record macro.
func (c *Card) Layout(gtx layout.Context, th *theme.Theme, content layout.Widget) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	variant := theme.GetCardVariant(c.Variant, &th.Colors)
	styles := utils.ParseClasses(c.Classes)

	padding := c.Padding
	if padding == (layout.Inset{}) {
		padding = layout.Inset{
			Top:    th.Spacing.Space6,
			Bottom: th.Spacing.Space6,
			Left:   th.Spacing.Space6,
			Right:  th.Spacing.Space6,
		}
	}

	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}

	bgColor := variant.Background
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	radius := th.Radius.RadiusLG
	if styles.Radius > 0 {
		radius = styles.Radius
	}

	// 1. Record content operations into a macro to measure dimensions
	macro := op.Record(gtx.Ops)
	dims := padding.Layout(gtx, content)
	callOp := macro.Stop()

	// 2. Draw card background FIRST using exact measured bounds
	rect := image.Rectangle{Max: dims.Size}
	radiusPx := gtx.Dp(radius)
	theme.DrawRRectBackground(gtx, rect, radiusPx, bgColor)

	if variant.BorderWidth > 0 {
		rr := clip.UniformRRect(rect, radiusPx)
		theme.DrawStroke(gtx, rr.Path(gtx.Ops), float32(gtx.Dp(unit.Dp(variant.BorderWidth))), variant.Border)
	}

	// 3. Play recorded content operations ON TOP of background
	callOp.Add(gtx.Ops)

	return dims
}

func (c *Card) Update(_ layout.Context) theme.ComponentState {
	return &State{}
}

type State struct {
	active   bool
	hovered  bool
	pressed  bool
	disabled bool
}

func (cs *State) IsActive() bool   { return cs.active }
func (cs *State) IsHovered() bool  { return cs.hovered }
func (cs *State) IsPressed() bool  { return cs.pressed }
func (cs *State) IsDisabled() bool { return cs.disabled }

type Header struct {
	Classes string
	Padding layout.Inset
}

func NewHeader(classes string) *Header {
	return &Header{Classes: classes}
}

func (h *Header) Layout(gtx layout.Context, th *theme.Theme, content layout.Widget) layout.Dimensions {
	styles := utils.ParseClasses(h.Classes)
	padding := h.Padding
	if padding == (layout.Inset{}) {
		padding = layout.Inset{Top: th.Spacing.Space6, Bottom: th.Spacing.Space6, Left: th.Spacing.Space6, Right: th.Spacing.Space6}
	}
	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}
	return padding.Layout(gtx, content)
}

type Content struct {
	Classes string
	Padding layout.Inset
}

func NewContent(classes string) *Content {
	return &Content{Classes: classes}
}

func (c *Content) Layout(gtx layout.Context, th *theme.Theme, content layout.Widget) layout.Dimensions {
	styles := utils.ParseClasses(c.Classes)
	padding := c.Padding
	if padding == (layout.Inset{}) {
		padding = layout.Inset{Top: th.Spacing.Space6, Bottom: th.Spacing.Space6, Left: th.Spacing.Space6, Right: th.Spacing.Space6}
	}
	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}
	return padding.Layout(gtx, content)
}

type Footer struct {
	Classes string
	Padding layout.Inset
}

func NewFooter(classes string) *Footer {
	return &Footer{Classes: classes}
}

func (f *Footer) Layout(gtx layout.Context, th *theme.Theme, content layout.Widget) layout.Dimensions {
	styles := utils.ParseClasses(f.Classes)
	padding := f.Padding
	if padding == (layout.Inset{}) {
		padding = layout.Inset{Top: th.Spacing.Space6, Bottom: th.Spacing.Space6, Left: th.Spacing.Space6, Right: th.Spacing.Space6}
	}
	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}
	return padding.Layout(gtx, content)
}

type Title struct {
	Text    string
	Classes string
}

func NewTitle(text string, classes string) *Title {
	return &Title{Text: text, Classes: classes}
}

func (t *Title) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	textStyle := th.Typography.H3(&th.Colors)
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return renderText(gtx, textStyle, t.Text) }),
	)
}

type Description struct {
	Text    string
	Classes string
}

func NewDescription(text string, classes string) *Description {
	return &Description{Text: text, Classes: classes}
}

func (d *Description) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	textStyle := th.Typography.BodySmall(&th.Colors)
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return renderText(gtx, textStyle, d.Text) }),
	)
}

func renderText(gtx layout.Context, style theme.TextStyle, text string) layout.Dimensions {
	thMat := material.NewTheme()
	label := material.Label(thMat, style.Size, text)
	if style.Color != nil {
		label.Color = style.Color.Foreground
	}
	if style.Weight > 0 {
		label.Font.Weight = style.Weight
	}
	return label.Layout(gtx)
}
