/*
Package numberinput provides a numeric stepper input component for gio-shadcn applications.

NumberInputs allow users to increment and decrement numeric values following
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

// NumberInput represents a numeric stepper input widget with increment and decrement buttons.
type NumberInput struct {
	Value    float32
	Step     float32
	Min      float32
	Max      float32
	Classes  string
	OnChange func(float32)

	decBtn *widget.Clickable
	incBtn *widget.Clickable
}

// Config represents configuration for creating a NumberInput component.
type Config struct {
	Value    float32
	Step     float32
	Min      float32
	Max      float32
	Classes  string
	OnChange func(float32)
}

// New creates a new NumberInput component.
func New(config Config) *NumberInput {
	stepVal := config.Step
	if stepVal <= 0 {
		stepVal = 1.0
	}
	val := config.Value
	if config.Max > config.Min {
		if val < config.Min {
			val = config.Min
		} else if val > config.Max {
			val = config.Max
		}
	}
	return &NumberInput{
		Value:    val,
		Step:     stepVal,
		Min:      config.Min,
		Max:      config.Max,
		Classes:  config.Classes,
		OnChange: config.OnChange,
		decBtn:   new(widget.Clickable),
		incBtn:   new(widget.Clickable),
	}
}

// Layout renders the decrement button, numeric value display box, and increment button.
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

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

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

			gtxContent := gtx
			gtxContent.Constraints.Min = image.Pt(0, 0)

			renderVal := func(gtx layout.Context) layout.Dimensions {
				return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, fmt.Sprintf("%.0f", ni.Value))
					lbl.Color = th.Colors.Foreground
					lbl.Font.Weight = font.Bold
					lbl.Alignment = text.Middle
					return lbl.Layout(gtx)
				})
			}

			valDims := renderVal(gtxContent)
			valSize := valDims.Size

			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					rect := image.Rectangle{Max: valSize}
					radius := gtx.Dp(th.Radius.RadiusSM)

					theme.DrawRRectBackground(gtx, rect, radius, th.Colors.Background)

					rr := clip.UniformRRect(rect, radius)
					theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, th.Colors.Border)

					return layout.Dimensions{Size: valSize}
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return renderVal(gtx)
				}),
			)
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

	// Reset active GPU paint color state
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

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

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderBtn := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeSM, label)
			lbl.Color = fgColor
			lbl.Font.Weight = font.Bold
			return lbl.Layout(gtx)
		})
	}

	btnDims := renderBtn(gtxContent)
	btnSize := btnDims.Size

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: btnSize}
			radius := gtx.Dp(th.Radius.RadiusSM)

			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			return layout.Dimensions{Size: btnSize}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderBtn(gtx)
		}),
	)
}
