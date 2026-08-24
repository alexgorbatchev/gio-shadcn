package resizable

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ResPanel *Resizable
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		ResPanel: New(Config{
			Ratio: 0.5,
			LeftWidget: func(gtx layout.Context) layout.Dimensions {
				th := theme.NewDark()
				lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Left Split Panel")
				lbl.Color = th.Colors.Foreground
				return lbl.Layout(gtx)
			},
			RightWidget: func(gtx layout.Context) layout.Dimensions {
				th := theme.NewDark()
				lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Right Split Panel")
				lbl.Color = th.Colors.Foreground
				return lbl.Layout(gtx)
			},
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.ResPanel.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
