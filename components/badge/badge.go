/*
Package badge provides a badge component for gio-shadcn applications.

Badges display small status indicators, tags, or count labels following
shadcn/ui design principles.

# Quick Start

Create a default badge:

	bg := badge.New(badge.Config{
		Text: "Default",
	})

	dims := bg.Layout(gtx, th)

# Variants

• VariantDefault - Primary background badge
• VariantSecondary - Muted background badge
• VariantOutline - Border only badge
• VariantDestructive - Dangerous/Error indicator badge
*/
package badge

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
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

// Layout renders the badge with the given graphics context and theme.
func (b *Badge) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	// Get variant colors
	bgColor, fgColor, borderColor, borderWidth := b.getVariantColors(&th.Colors)

	// Inset padding
	padding := layout.Inset{
		Top:    th.Spacing.Space1,
		Bottom: th.Spacing.Space1,
		Left:   th.Spacing.Space2,
		Right:  th.Spacing.Space2,
	}

	// Calculate custom classes if present
	styles := utils.ParseClasses(b.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// Render label text
		lbl := material.Label(material.NewTheme(), th.Typography.FontSizeXS, b.Text)
		lbl.Color = fgColor
		lbl.Alignment = text.Middle
		textDims := lbl.Layout(gtx)

		// Border radius
		radius := gtx.Dp(th.Radius.RadiusFull)

		// Expand area for background & border
		rect := image.Rectangle{Max: textDims.Size}
		rr := clip.UniformRRect(rect, radius)

		// Record operations
		paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

		if borderWidth > 0 {
			stroke := clip.Stroke{
				Path:  rr.Path(gtx.Ops),
				Width: borderWidth,
			}
			paint.FillShape(gtx.Ops, borderColor, stroke.Op())
		}

		return textDims
	})
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
