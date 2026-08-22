/*
Package badge provides a badge component for gio-shadcn applications.

Badges display small status indicators, tags, or count labels following
shadcn/ui design principles.
*/
package badge

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Badge represents a shadcn/ui badge component.
type Badge struct {
	Text    string
	Variant theme.Variant
	Classes string
}

// Config represents configuration for creating a Badge.
type Config struct {
	Text    string
	Variant theme.Variant
	Classes string
}

// New creates a new Badge with the given configuration.
func New(config Config) *Badge {
	v := config.Variant
	if v == "" {
		v = theme.VariantDefault
	}
	return &Badge{
		Text:    config.Text,
		Variant: v,
		Classes: config.Classes,
	}
}

// Layout renders the badge with background painted before text using exact tight bounding size and proper clip popping.
func (b *Badge) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	bgColor, fgColor, borderColor, borderWidth := b.getVariantColors(&th.Colors)

	padding := layout.Inset{
		Top:    th.Spacing.Space1,
		Bottom: th.Spacing.Space1,
		Left:   th.Spacing.Space2,
		Right:  th.Spacing.Space2,
	}

	styles := utils.ParseClasses(b.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	// Unconstrain BOTH Min and Max so text measures its exact natural dimensions
	gtxText := gtx
	gtxText.Constraints = layout.Constraints{
		Min: image.Pt(0, 0),
		Max: image.Pt(1e6, 1e6),
	}

	// 1. Record text content operations into a macro
	macro := op.Record(gtx.Ops)
	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}
	lbl := material.Label(mTheme, th.Typography.FontSizeXS, b.Text)
	lbl.Color = fgColor
	lbl.Alignment = text.Start
	lblDims := lbl.Layout(gtxText)
	callOp := macro.Stop()

	// Calculate exact badge bounds from text dimensions + padding
	padLeft := gtx.Dp(padding.Left)
	padRight := gtx.Dp(padding.Right)
	padTop := gtx.Dp(padding.Top)
	padBottom := gtx.Dp(padding.Bottom)

	badgeSize := image.Pt(
		lblDims.Size.X+padLeft+padRight,
		lblDims.Size.Y+padTop+padBottom,
	)

	// 2. Draw background first using push/pop clip stack
	rect := image.Rectangle{Max: badgeSize}
	radius := gtx.Dp(th.Radius.RadiusFull)

	theme.DrawRRectBackground(gtx, rect, radius, bgColor)

	if borderWidth > 0 {
		rr := clip.UniformRRect(rect, radius)
		theme.DrawStroke(gtx, rr.Path(gtx.Ops), borderWidth, borderColor)
	}

	// 3. Offset and play recorded text operation on top
	offStack := op.Offset(image.Pt(padLeft, padTop)).Push(gtx.Ops)
	callOp.Add(gtx.Ops)
	offStack.Pop()

	return layout.Dimensions{Size: badgeSize}
}

func (b *Badge) getVariantColors(cs *theme.ColorScheme) (bg, fg, border color.NRGBA, borderWidth float32) {
	switch b.Variant {
	case theme.VariantSecondary:
		return cs.Secondary, cs.SecondaryFg, color.NRGBA{}, 0
	case theme.VariantOutline:
		return color.NRGBA{}, cs.Foreground, cs.Border, 1.0
	case theme.VariantDestructive:
		return cs.Destructive, cs.DestructiveFg, color.NRGBA{}, 0
	default: // VariantDefault
		return cs.Primary, cs.PrimaryFg, color.NRGBA{}, 0
	}
}
