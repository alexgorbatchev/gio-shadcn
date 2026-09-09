package sheet

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	RightSheet  *Sheet
	LeftSheet   *Sheet
	TopSheet    *Sheet
	BottomSheet *Sheet
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	// 1. Right Side Sheet (Default)
	s.RightSheet = New(Config{
		TriggerText: "Open Right Side Sheet",
		Title:       "Right Side Sheet",
		Description: "Slides in from the right edge across the full application height.",
		Side:        SideRight,
		Width:       unit.Dp(360),
		Open:        false,
	})

	// 2. Left Side Sheet
	s.LeftSheet = New(Config{
		TriggerText: "Open Left Side Sheet",
		Title:       "Left Side Sheet",
		Description: "Slides in from the left edge across the full application height.",
		Side:        SideLeft,
		Width:       unit.Dp(360),
		Open:        false,
	})

	// 3. Top Sheet
	s.TopSheet = New(Config{
		TriggerText: "Open Top Sheet",
		Title:       "Top Sheet Banner",
		Description: "Slides down from the top edge across the full application width.",
		Side:        SideTop,
		Height:      unit.Dp(240),
		Open:        false,
	})

	// 4. Bottom Sheet
	s.BottomSheet = New(Config{
		TriggerText: "Open Bottom Sheet",
		Title:       "Bottom Sheet Panel",
		Description: "Slides up from the bottom edge across the full application width.",
		Side:        SideBottom,
		Height:      unit.Dp(260),
		Open:        false,
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("1. Right Side Sheet (Default)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.RightSheet.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Left Side Sheet", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.LeftSheet.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("3. Top Sheet", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.TopSheet.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("4. Bottom Sheet", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BottomSheet.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
