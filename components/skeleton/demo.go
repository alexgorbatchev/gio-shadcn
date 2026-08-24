package skeleton

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ShimmerSk *Skeleton
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		ShimmerSk: New(Config{Width: unit.Dp(180), Height: unit.Dp(24)}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.ShimmerSk.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
