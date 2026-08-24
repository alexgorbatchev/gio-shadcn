package switchcomp

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	SwToggle *Switch
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		SwToggle: New(Config{Value: true}),
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

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SwToggle.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "HQ Audio Engine")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
