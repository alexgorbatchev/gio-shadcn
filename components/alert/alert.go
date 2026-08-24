/*
Package alert provides an alert callout box component for gio-shadcn applications.

Alerts display important messages, notices, warnings, and error callouts following
shadcn/ui design principles.
*/
package alert

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Alert represents a shadcn/ui alert callout box.
type Alert struct {
	Title       string
	Description string
	Variant     theme.Variant
	Classes     string
	Icon        *lucide.Icon
}

// Config represents configuration for creating an Alert.
type Config struct {
	Title       string
	Description string
	Variant     theme.Variant
	Classes     string
	Icon        *lucide.Icon
}

// New creates a new Alert component with the given configuration.
func New(config Config) *Alert {
	v := config.Variant
	if v == "" {
		v = theme.VariantDefault
	}
	ic := config.Icon
	if ic == nil {
		if v == theme.VariantDestructive {
			ic = lucide.TriangleAlert
		} else {
			ic = lucide.Info
		}
	}
	return &Alert{
		Title:       config.Title,
		Description: config.Description,
		Variant:     v,
		Classes:     config.Classes,
		Icon:        ic,
	}
}

// Layout renders the alert box with exact bounding background and clean color resets.
func (a *Alert) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	bgColor := th.Colors.Background
	fgColor := th.Colors.Foreground
	borderColor := th.Colors.Border

	if a.Variant == theme.VariantDestructive {
		bgColor = th.Colors.Destructive
		borderColor = th.Colors.Destructive
		fgColor = th.Colors.DestructiveFg
	}

	styles := utils.ParseClasses(a.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space4,
		Bottom: th.Spacing.Space4,
		Left:   th.Spacing.Space4,
		Right:  th.Spacing.Space4,
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	macro := op.Record(gtx.Ops)
	dims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Start}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if a.Icon != nil {
					return layout.Inset{Right: th.Spacing.Space3, Top: unit.Dp(2)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return a.Icon.LayoutSize(gtx, unit.Dp(18), fgColor)
					})
				}
				return layout.Dimensions{}
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if a.Title == "" {
							return layout.Dimensions{}
						}
						lbl := material.Label(mTheme, th.Typography.FontSizeBase, a.Title)
						lbl.Color = fgColor
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if a.Title != "" && a.Description != "" {
							return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
						}
						return layout.Dimensions{}
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if a.Description == "" {
							return layout.Dimensions{}
						}
						lbl := material.Label(mTheme, th.Typography.FontSizeSM, a.Description)
						lbl.Color = th.Colors.MutedFg
						if a.Variant == theme.VariantDestructive {
							lbl.Color = fgColor
						}
						return lbl.Layout(gtx)
					}),
				)
			}),
		)
	})
	callOp := macro.Stop()

	rect := image.Rectangle{Max: dims.Size}
	radius := gtx.Dp(th.Radius.RadiusMD)

	theme.DrawRRectBackground(gtx, rect, radius, bgColor)

	rr := clip.UniformRRect(rect, radius)
	theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

	callOp.Add(gtx.Ops)

	// Reset color state back to foreground
	paint.ColorOp{Color: th.Colors.Foreground}.Add(gtx.Ops)

	return dims
}
