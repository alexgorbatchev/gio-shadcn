package aspectratio

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	Demo16x9     *AspectRatio
	DemoPortrait *AspectRatio
	DemoRTL      *AspectRatio
	DemoSquare   *AspectRatio
}

var defaultDemo = NewDemoState()

func renderPlaceholder(labelStr string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		th := theme.NewDark()
		rect := image.Rectangle{Max: gtx.Constraints.Min}
		theme.DrawRRectBackground(gtx, rect, gtx.Dp(th.Radius.RadiusMD), th.Colors.Muted)
		return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography(labelStr, label.Small, "").Layout(gtx, th)
		})
	}
}

func NewDemoState() *DemoState {
	s := &DemoState{}

	// aspect-ratio-demo.tsx (16:9)
	s.Demo16x9 = New(Config{
		Ratio:  16.0 / 9.0,
		Widget: renderPlaceholder("16:9 Landscape Image"),
	})

	// aspect-ratio-portrait.tsx (9:16)
	s.DemoPortrait = New(Config{
		Ratio:  9.0 / 16.0,
		Widget: renderPlaceholder("9:16 Portrait"),
	})

	// aspect-ratio-rtl.tsx (16:9 with RTL caption)
	s.DemoRTL = New(Config{
		Ratio:  16.0 / 9.0,
		Widget: renderPlaceholder("منظر طبيعي جميل (16:9)"),
	})

	// aspect-ratio-square.tsx (1:1)
	s.DemoSquare = New(Config{
		Ratio:  1.0,
		Widget: renderPlaceholder("1:1 Square Card"),
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Aspect Ratio (16:9 Landscape)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtxBounded := gtx
			gtxBounded.Constraints.Max.X = gtx.Dp(unit.Dp(360))
			return s.Demo16x9.Layout(gtxBounded, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Aspect Ratio (9:16 Portrait)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtxBounded := gtx
			gtxBounded.Constraints.Max.X = gtx.Dp(unit.Dp(160))
			return s.DemoPortrait.Layout(gtxBounded, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Aspect Ratio (1:1 Square)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtxBounded := gtx
			gtxBounded.Constraints.Max.X = gtx.Dp(unit.Dp(180))
			return s.DemoSquare.Layout(gtxBounded, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Aspect Ratio (RTL)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtxBounded := gtx
			gtxBounded.Constraints.Max.X = gtx.Dp(unit.Dp(360))
			return s.DemoRTL.Layout(gtxBounded, th)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
