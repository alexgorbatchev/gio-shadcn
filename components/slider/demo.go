package slider

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	RngSlider *Slider
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		RngSlider: New(Config{Value: 65.0, Min: 0.0, Max: 100.0}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.RngSlider.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
