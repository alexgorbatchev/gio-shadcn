package numberinput

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	NumStepper *NumberInput
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		NumStepper: New(Config{Value: 128.0, Step: 1.0, Min: 60.0, Max: 200.0}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.NumStepper.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
