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

// Layout renders the badge using live layout.Stack matching button.go layout.
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

	// Measure content dimensions with unconstrained max bounds
	gtxContent := gtx
	gtxContent.Constraints = layout.Constraints{
		Min: image.Pt(0, 0),
		Max: image.Pt(1e6, 1e6),
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	contentDims := padding.Layout(gtxContent, func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(mTheme, th.Typography.FontSizeXS, b.Text)
		lbl.Color = fgColor
		lbl.Alignment = text.Start
		return lbl.Layout(gtx)
	})

	badgeSize := contentDims.Size
	gtx.Constraints = layout.Exact(badgeSize)

	dims := layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: badgeSize}
			radiusPx := badgeSize.Y / 2
			if radiusPx <= 0 {
				radiusPx = 1
			}

			bgClip := clip.Rect{Max: badgeSize}.Push(gtx.Ops)
			theme.DrawRRectBackground(gtx, rect, radiusPx, bgColor)

			if borderWidth > 0 {
				rr := clip.UniformRRect(rect, radiusPx)
				theme.DrawStroke(gtx, rr.Path(gtx.Ops), borderWidth, borderColor)
			}
			bgClip.Pop()

			return layout.Dimensions{Size: badgeSize}
		}),

		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(mTheme, th.Typography.FontSizeXS, b.Text)
				lbl.Color = fgColor
				lbl.Alignment = text.Start
				return lbl.Layout(gtx)
			})
		}),
	)

	// Reset active GPU paint color state back to theme background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
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
