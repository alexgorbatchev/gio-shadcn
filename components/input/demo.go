package input

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	TxtInput *Input
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		TxtInput: Text("Enter track title..."),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	maxWidth := gtx.Dp(unit.Dp(400))
	gtx.Constraints.Max.X = maxWidth
	gtx.Constraints.Min.X = maxWidth

	return s.TxtInput.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
