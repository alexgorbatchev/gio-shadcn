/*
Package alert provides a callout alert box component for gio-shadcn applications.

Alerts display important messages or callouts following
shadcn/ui design principles.
*/
package alert

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Alert represents a callout alert box component.
type Alert struct {
	Title       string
	Description string
	Variant     theme.Variant
	Classes     string
}

// Config represents configuration for creating an Alert.
type Config struct {
	Title       string
	Description string
	Variant     theme.Variant
	Classes     string
}

// New creates a new Alert component with the given configuration.
func New(config Config) *Alert {
	v := config.Variant
	if v == "" {
		v = theme.VariantDefault
	}
	return &Alert{
		Title:       config.Title,
		Description: config.Description,
		Variant:     v,
		Classes:     config.Classes,
	}
}

// Layout renders the alert box with exact bounding background.
func (a *Alert) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	bgColor := th.Colors.Card
	fgColor := th.Colors.CardFg
	borderColor := th.Colors.Border

	if a.Variant == theme.VariantDestructive {
		borderColor = th.Colors.Destructive
		fgColor = th.Colors.Destructive
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
	})
	callOp := macro.Stop()

	rect := image.Rectangle{Max: dims.Size}
	radius := gtx.Dp(th.Radius.RadiusMD)
	rr := clip.UniformRRect(rect, radius)

	paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

	stroke := clip.Stroke{
		Path:  rr.Path(gtx.Ops),
		Width: 1.0,
	}
	paint.FillShape(gtx.Ops, borderColor, stroke.Op())

	callOp.Add(gtx.Ops)

	return dims
}

func _(c color.NRGBA) {} // unused color guard
