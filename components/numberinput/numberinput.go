/*
Package numberinput provides a numeric stepper input component for gio-shadcn applications.

NumberInputs allow users to adjust numeric quantities with increment/decrement buttons following
shadcn/ui design principles.
*/
package numberinput

import (
	"fmt"
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

// NumberInput represents a numeric stepper input component.
type NumberInput struct {
	Value    float32
	Step     float32
	Min      float32
	Max      float32
	OnChange func(float32)

	decBtn *widget.Clickable
	incBtn *widget.Clickable
}

// Config represents configuration for creating a NumberInput.
type Config struct {
	Value    float32
	Step     float32
	Min      float32
	Max      float32
	OnChange func(float32)
}

// New creates a new NumberInput stepper component.
func New(config Config) *NumberInput {
	step := config.Step
	if step <= 0 {
		step = 1.0
	}
	minVal := config.Min
	maxVal := config.Max
	if maxVal <= minVal {
		maxVal = 100.0
		minVal = 0.0
	}
	val := config.Value
	if val < minVal {
		val = minVal
	} else if val > maxVal {
		val = maxVal
	}

	return &NumberInput{
		Value:    val,
		Step:     step,
		Min:      minVal,
		Max:      maxVal,
		OnChange: config.OnChange,
		decBtn:   new(widget.Clickable),
		incBtn:   new(widget.Clickable),
	}
}

// Layout renders the stepper buttons and value display.
func (ni *NumberInput) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if ni.decBtn.Clicked(gtx) && ni.Value > ni.Min {
		ni.Value -= ni.Step
		if ni.Value < ni.Min {
			ni.Value = ni.Min
		}
		if ni.OnChange != nil {
			ni.OnChange(ni.Value)
		}
	}

	if ni.incBtn.Clicked(gtx) && ni.Value < ni.Max {
		ni.Value += ni.Step
		if ni.Value > ni.Max {
			ni.Value = ni.Max
		}
		if ni.OnChange != nil {
			ni.OnChange(ni.Value)
		}
	}

	mTheme := material.NewTheme()

	containerDims := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		// Decrement Button
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ni.decBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return ni.layoutBtn(gtx, th, mTheme, "-", ni.Value <= ni.Min)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
		}),
		// Value Display Box
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			padding := layout.Inset{
				Top:    th.Spacing.Space2,
				Bottom: th.Spacing.Space2,
				Left:   th.Spacing.Space3,
				Right:  th.Spacing.Space3,
			}
			valDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(mTheme, th.Typography.FontSizeSM, fmt.Sprintf("%.0f", ni.Value))
				lbl.Color = th.Colors.Foreground
				lbl.Font.Weight = font.Bold
				lbl.Alignment = text.Middle
				return lbl.Layout(gtx)
			})

			rect := image.Rectangle{Max: valDims.Size}
			radius := gtx.Dp(th.Radius.RadiusSM)
			rr := clip.UniformRRect(rect, radius)

			paint.FillShape(gtx.Ops, th.Colors.Background, rr.Op(gtx.Ops))

			stroke := clip.Stroke{
				Path:  rr.Path(gtx.Ops),
				Width: 1.0,
			}
			paint.FillShape(gtx.Ops, th.Colors.Border, stroke.Op())

			return valDims
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
		}),
		// Increment Button
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ni.incBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return ni.layoutBtn(gtx, th, mTheme, "+", ni.Value >= ni.Max)
			})
		}),
	)

	return containerDims
}

func (ni *NumberInput) layoutBtn(gtx layout.Context, th *theme.Theme, mTheme *material.Theme, label string, disabled bool) layout.Dimensions {
	bgColor := th.Colors.Muted
	fgColor := th.Colors.MutedFg

	if disabled {
		fgColor.A = 100
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space2,
		Bottom: th.Spacing.Space2,
		Left:   th.Spacing.Space3,
		Right:  th.Spacing.Space3,
	}

	btnDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(mTheme, th.Typography.FontSizeSM, label)
		lbl.Color = fgColor
		lbl.Font.Weight = font.Bold
		return lbl.Layout(gtx)
	})

	rect := image.Rectangle{Max: btnDims.Size}
	radius := gtx.Dp(th.Radius.RadiusSM)
	rr := clip.UniformRRect(rect, radius)

	paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

	return btnDims
}
