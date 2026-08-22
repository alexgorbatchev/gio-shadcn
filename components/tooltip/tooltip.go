/*
Package tooltip provides a tooltip hover label component for gio-shadcn applications.

Tooltips display informative text popups following
shadcn/ui design principles.
*/
package tooltip

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Tooltip represents a tooltip label component.
type Tooltip struct {
	Text    string
	Classes string
}

// Config represents configuration for creating a Tooltip.
type Config struct {
	Text    string
	Classes string
}

// New creates a new Tooltip component with the given configuration.
func New(config Config) *Tooltip {
	return &Tooltip{
		Text:    config.Text,
		Classes: config.Classes,
	}
}

// Layout renders the tooltip popover box.
func (t *Tooltip) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	bgColor := th.Colors.Primary
	fgColor := th.Colors.PrimaryFg

	styles := utils.ParseClasses(t.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space1,
		Bottom: th.Spacing.Space1,
		Left:   th.Spacing.Space3,
		Right:  th.Spacing.Space3,
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	macro := op.Record(gtx.Ops)
	dims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(mTheme, th.Typography.FontSizeXS, t.Text)
		lbl.Color = fgColor
		lbl.Alignment = text.Middle
		return lbl.Layout(gtx)
	})
	callOp := macro.Stop()

	rect := image.Rectangle{Max: dims.Size}
	radius := gtx.Dp(th.Radius.RadiusSM)
	rr := clip.UniformRRect(rect, radius)

	paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

	callOp.Add(gtx.Ops)

	return dims
}
