package pagination

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	PageControls *Pagination
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		PageControls: New(Config{CurrentPage: 1, TotalPages: 5}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.PageControls.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
