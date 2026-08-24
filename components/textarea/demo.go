package textarea

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	TxtArea *TextArea
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		TxtArea: New(Config{Placeholder: "Enter track descriptions and metadata notes..."}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.TxtArea.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
