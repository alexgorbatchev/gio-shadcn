/*
Package label provides typography and text labeling components for gio-shadcn applications.
*/
package label

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

type Label struct {
	Text      string
	TextStyle theme.TextStyle
	Classes   string
	Variant   theme.Variant
	Size      theme.Size
}

type Option func(*Label)

func WithLabelText(text string) Option {
	return func(l *Label) {
		l.Text = text
	}
}

func WithTextStyle(style theme.TextStyle) Option {
	return func(l *Label) {
		l.TextStyle = style
	}
}

func WithLabelClasses(classes string) Option {
	return func(l *Label) {
		l.Classes = classes
	}
}

func WithLabelVariant(variant theme.Variant) Option {
	return func(l *Label) {
		l.Variant = variant
	}
}

func WithLabelSize(size theme.Size) Option {
	return func(l *Label) {
		l.Size = size
	}
}

func NewLabel(options ...Option) *Label {
	l := &Label{
		Variant: theme.VariantDefault,
		Size:    theme.SizeDefault,
	}

	for _, option := range options {
		option(l)
	}

	return l
}

func (l *Label) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	styles := utils.ParseClasses(l.Classes)

	textStyle := l.getDefaultTextStyle(th)
	if l.TextStyle.Size > 0 {
		textStyle.Size = l.TextStyle.Size
	}
	if l.TextStyle.Color != nil {
		textStyle.Color = l.TextStyle.Color
	}
	if l.TextStyle.Weight > 0 {
		textStyle.Weight = l.TextStyle.Weight
	}

	if l.Size != "" {
		textStyle = l.applySizeToTextStyle(textStyle, th)
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	label := material.Label(mTheme, textStyle.Size, l.Text)
	label.Color = th.Colors.Foreground

	if textStyle.Color != nil {
		label.Color = textStyle.Color.Foreground
	}

	label.Alignment = textStyle.Alignment
	label.Font.Weight = textStyle.Weight
	label.Font.Style = textStyle.Style

	if styles.Background.A > 0 {
		label.Color = styles.Background
	}

	dims := label.Layout(gtx)

	// Reset active GPU paint color state
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (l *Label) Update(_ layout.Context) theme.ComponentState {
	return &State{}
}

type State struct {
	active   bool
	hovered  bool
	pressed  bool
	disabled bool
}

func (ls *State) IsActive() bool   { return ls.active }
func (ls *State) IsHovered() bool  { return ls.hovered }
func (ls *State) IsPressed() bool  { return ls.pressed }
func (ls *State) IsDisabled() bool { return ls.disabled }

func (l *Label) getDefaultTextStyle(th *theme.Theme) theme.TextStyle {
	switch l.Variant {
	case theme.VariantDefault:
		return th.Typography.Body(&th.Colors)
	case theme.VariantSecondary:
		return theme.TextStyle{
			Size:          th.Typography.FontSizeBase,
			Color:         &th.Colors,
			Weight:        th.Typography.Body(&th.Colors).Weight,
			Style:         th.Typography.Body(&th.Colors).Style,
			Alignment:     th.Typography.Body(&th.Colors).Alignment,
			LineHeight:    th.Typography.Body(&th.Colors).LineHeight,
			LetterSpacing: th.Typography.Body(&th.Colors).LetterSpacing,
		}
	default:
		return th.Typography.Body(&th.Colors)
	}
}

func (l *Label) applySizeToTextStyle(textStyle theme.TextStyle, th *theme.Theme) theme.TextStyle {
	switch l.Size {
	case theme.SizeSM:
		textStyle.Size = th.Typography.FontSizeSM
	case theme.SizeLG:
		textStyle.Size = th.Typography.FontSizeLG
	default:
		textStyle.Size = th.Typography.FontSizeBase
	}
	return textStyle
}

func (l *Label) SetText(text string) {
	l.Text = text
}

func (l *Label) SetTextStyle(style theme.TextStyle) {
	l.TextStyle = style
}

type Typography struct {
	Text      string
	Element   TypographyElement
	Classes   string
	TextStyle theme.TextStyle
}

type TypographyElement string

const (
	H1    TypographyElement = "h1"
	H2    TypographyElement = "h2"
	H3    TypographyElement = "h3"
	H4    TypographyElement = "h4"
	P     TypographyElement = "p"
	Small TypographyElement = "small"
	Lead  TypographyElement = "lead"
	Large TypographyElement = "large"
	Muted TypographyElement = "muted"
)

func NewTypography(text string, element TypographyElement, classes string) *Typography {
	return &Typography{
		Text:    text,
		Element: element,
		Classes: classes,
	}
}

func (t *Typography) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	styles := utils.ParseClasses(t.Classes)
	textStyle := t.getTextStyleForElement(th)

	// Merge individual properties instead of wiping out size with zero-value fields
	if t.TextStyle.Size > 0 {
		textStyle.Size = t.TextStyle.Size
	}
	if t.TextStyle.Color != nil {
		textStyle.Color = t.TextStyle.Color
	}
	if t.TextStyle.Weight > 0 {
		textStyle.Weight = t.TextStyle.Weight
	}
	if t.TextStyle.Alignment != 0 {
		textStyle.Alignment = t.TextStyle.Alignment
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	label := material.Label(mTheme, textStyle.Size, t.Text)

	// Set default color based on typography element
	label.Color = t.getColorForElement(th)

	// If explicit color set in TextStyle, respect it
	if textStyle.Color != nil {
		label.Color = textStyle.Color.Foreground
	}

	label.Alignment = textStyle.Alignment
	label.Font.Weight = textStyle.Weight
	label.Font.Style = textStyle.Style

	if styles.Background.A > 0 {
		label.Color = styles.Background
	}

	dims := label.Layout(gtx)

	// Reset active GPU paint color state
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (t *Typography) getTextStyleForElement(th *theme.Theme) theme.TextStyle {
	switch t.Element {
	case H1:
		return th.Typography.H1(&th.Colors)
	case H2:
		return th.Typography.H2(&th.Colors)
	case H3:
		return th.Typography.H3(&th.Colors)
	case H4:
		return th.Typography.H4(&th.Colors)
	case P:
		return th.Typography.Body(&th.Colors)
	case Small:
		return th.Typography.BodySmall(&th.Colors)
	case Lead:
		style := th.Typography.Body(&th.Colors)
		style.Size = th.Typography.FontSizeLG
		return style
	case Large:
		style := th.Typography.Body(&th.Colors)
		style.Size = th.Typography.FontSizeXL
		return style
	case Muted:
		style := th.Typography.BodySmall(&th.Colors)
		return style
	default:
		return th.Typography.Body(&th.Colors)
	}
}

func (t *Typography) getColorForElement(th *theme.Theme) color.NRGBA {
	switch t.Element {
	case H1, H2, H3, H4, P, Small, Lead, Large:
		return th.Colors.Foreground
	case Muted:
		return th.Colors.MutedFg
	default:
		return th.Colors.Foreground
	}
}

func (t *Typography) SetText(text string) {
	t.Text = text
}

func (t *Typography) SetElement(element TypographyElement) {
	t.Element = element
}

func (t *Typography) SetTextStyle(style theme.TextStyle) {
	t.TextStyle = style
}
