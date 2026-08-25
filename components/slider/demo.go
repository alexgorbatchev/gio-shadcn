package slider

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	// 1. Controlled Slider
	ControlledSlider *Slider
	ControlledValue  float32

	// 2. Default Slider
	DemoSlider *Slider

	// 3. Disabled Slider
	DisabledSlider *Slider

	// 4. Multiple Slider
	MultiSlider1 *Slider
	MultiSlider2 *Slider

	// 5. Range Slider
	RangeSlider *Slider

	// 6. Vertical Sliders
	VerticalSlider1 *Slider
	VerticalSlider2 *Slider
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{
		ControlledValue: 0.7,
		DemoSlider:      New(Config{Value: 75.0, Min: 0.0, Max: 100.0}),
		DisabledSlider:  New(Config{Value: 50.0, Min: 0.0, Max: 100.0, Disabled: true}),
		MultiSlider1:    New(Config{Value: 20.0, Min: 0.0, Max: 100.0}),
		MultiSlider2:    New(Config{Value: 70.0, Min: 0.0, Max: 100.0}),
		RangeSlider:     New(Config{Value: 50.0, Min: 0.0, Max: 100.0}),
		VerticalSlider1: New(Config{Value: 50.0, Min: 0.0, Max: 100.0, Orientation: OrientationVertical}),
		VerticalSlider2: New(Config{Value: 25.0, Min: 0.0, Max: 100.0, Orientation: OrientationVertical}),
	}

	s.ControlledSlider = New(Config{
		Value: s.ControlledValue,
		Min:   0.0,
		Max:   1.0,
		OnChange: func(v float32) {
			s.ControlledValue = v
		},
	})

	return s
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
		// Header Title
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeXL, "Slider Showcase (6 Upstream Demos)")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		// Demo 1: Slider Controlled
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeBase, "1. Controlled Slider (slider-controlled.tsx)")
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, fmt.Sprintf("Value: %.2f", s.ControlledValue))
					lbl.Color = th.Colors.Primary
					return lbl.Layout(gtx)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ControlledSlider.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 2: Slider Demo
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "2. Default Slider (slider-demo.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoSlider.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 3: Slider Disabled
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "3. Disabled Slider (slider-disabled.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DisabledSlider.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 4: Slider Multiple
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "4. Multiple Sliders (slider-multiple.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.MultiSlider1.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.MultiSlider2.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 5: Slider Range
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "5. Range Slider (slider-range.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.RangeSlider.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 6: Slider Vertical
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "6. Vertical Sliders (slider-vertical.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtxH := gtx
			gtxH.Constraints.Max.Y = gtx.Dp(unit.Dp(160))
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtxH,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.VerticalSlider1.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.VerticalSlider2.Layout(gtx, th) }),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
