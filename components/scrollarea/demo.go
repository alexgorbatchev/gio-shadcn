package scrollarea

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ScrollAreaWidget *ScrollArea
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		ScrollAreaWidget: New(Config{
			Widget: func(gtx layout.Context) layout.Dimensions {
				th := theme.NewDark()
				lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Scrollable Container Body Area")
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

	return s.ScrollAreaWidget.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
