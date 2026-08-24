package popover

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	AnchoredPop *Popover
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		AnchoredPop: New(Config{Title: "Popover Title", Description: "Anchored card popover content box.", Open: false}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.AnchoredPop.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
