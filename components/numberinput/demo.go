package numberinput

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	BpmStepper      *NumberInput
	GainStepper     *NumberInput
	QuantityStepper *NumberInput
	DecimalStepper  *NumberInput
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		BpmStepper:      New(Config{Value: 128.0, Step: 1.0, Min: 60.0, Max: 200.0}),
		GainStepper:     New(Config{Value: 0.0, Step: 0.5, Min: -12.0, Max: 12.0}),
		QuantityStepper: New(Config{Value: 1.0, Step: 1.0, Min: 1.0, Max: 10.0}),
		DecimalStepper:  New(Config{Value: 0.75, Step: 0.05, Min: 0.0, Max: 1.0}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Track Tempo (BPM Stepper)")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BpmStepper.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Channel Gain Level (-12dB to +12dB)")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.GainStepper.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Deck Quantize Grid (1 to 10)")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.QuantityStepper.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Crossfader Curve Precision (0.0 to 1.0)")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DecimalStepper.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
