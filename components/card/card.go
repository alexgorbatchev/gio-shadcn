/*
Package card provides a card container component for gio-shadcn applications.

Cards are flexible content containers with customizable headers, content areas,
footers, and variants following shadcn/ui design principles.
*/
package card

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/paint"
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

// Layout renders the card background FIRST, then content ON TOP using layout.Stack.
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

	dims := layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: gtx.Constraints.Min}
			radiusPx := gtx.Dp(radius)

			theme.DrawRRectBackground(gtx, rect, radiusPx, bgColor)

			if variant.BorderWidth > 0 {
				rr := theme.RRect(rect, radiusPx)
				theme.DrawStroke(gtx, rr.Path(gtx.Ops), float32(gtx.Dp(unit.Dp(variant.BorderWidth))), variant.Border)
			}

			return layout.Dimensions{Size: gtx.Constraints.Min}
		}),

		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return padding.Layout(gtx, content)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (c *Card) Update(_ layout.Context) theme.ComponentState {
	return &State{}
}

type State struct {
	hovered bool
	focused bool
	active  bool
}

func (s *State) IsActive() bool   { return s.active }
func (s *State) IsHovered() bool  { return s.hovered }
func (s *State) IsPressed() bool  { return false }
func (s *State) IsDisabled() bool { return false }

type Header struct {
	Title       string
	Description string
	Classes     string
}

type HeaderConfig struct {
	Title       string
	Description string
	Classes     string
}

func NewHeader(config HeaderConfig) *Header {
	return &Header{
		Title:       config.Title,
		Description: config.Description,
		Classes:     config.Classes,
	}
}

func (h *Header) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}
	defer func() { paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops) }()

	return utils.ParseClasses(h.Classes).Padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if h.Title == "" {
					return layout.Dimensions{}
				}
				return layout.Inset{Bottom: th.Spacing.Space1}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return NewTitle(TitleConfig{Text: h.Title}).Layout(gtx, th)
				})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if h.Description == "" {
					return layout.Dimensions{}
				}
				return NewDescription(DescriptionConfig{Text: h.Description}).Layout(gtx, th)
			}),
		)
	})
}

type Title struct {
	Text    string
	Classes string
}

type TitleConfig struct {
	Text    string
	Classes string
}

func NewTitle(config TitleConfig) *Title {
	return &Title{
		Text:    config.Text,
		Classes: config.Classes,
	}
}

func (t *Title) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}
	defer func() { paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops) }()
	return cardText(gtx, th, t.Text, t.Classes, true)
}

type Description struct {
	Text    string
	Classes string
}

type DescriptionConfig struct {
	Text    string
	Classes string
}

func NewDescription(config DescriptionConfig) *Description {
	return &Description{
		Text:    config.Text,
		Classes: config.Classes,
	}
}

func (d *Description) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}
	defer func() { paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops) }()
	return cardText(gtx, th, d.Text, d.Classes, false)
}

func cardText(gtx layout.Context, th *theme.Theme, value, classes string, title bool) layout.Dimensions {
	m := th.MaterialTheme
	if m == nil {
		m = material.NewTheme()
	}
	lbl := material.Label(m, th.Typography.FontSizeSM, value)
	lbl.Color = th.Colors.MutedFg
	if title {
		lbl.TextSize = th.Typography.FontSizeBase
		lbl.Font.Weight = font.SemiBold
		lbl.Color = th.Colors.CardFg
	}
	styles := utils.ParseClasses(classes)
	if styles.Background.A > 0 {
		lbl.Color = styles.Background
	}
	return styles.Padding.Layout(gtx, lbl.Layout)
}

type Content struct {
	Classes string
}

type ContentConfig struct {
	Classes string
}

func NewContent(config ContentConfig) *Content {
	return &Content{
		Classes: config.Classes,
	}
}

func (c *Content) Layout(gtx layout.Context, th *theme.Theme, children layout.Widget) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}
	defer func() { paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops) }()
	return utils.ParseClasses(c.Classes).Padding.Layout(gtx, children)
}

type Footer struct {
	Classes string
}

type FooterConfig struct {
	Classes string
}

func NewFooter(config FooterConfig) *Footer {
	return &Footer{
		Classes: config.Classes,
	}
}

func (f *Footer) Layout(gtx layout.Context, th *theme.Theme, children layout.Widget) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}
	defer func() { paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops) }()
	return utils.ParseClasses(f.Classes).Padding.Layout(gtx, children)
}
