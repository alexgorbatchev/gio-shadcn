package titlebar

import (
	"gioui.org/app"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	TitleBar *TitleBar
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		TitleBar: NewTitleBar(
			WithTitle("Window TitleBar"),
			WithVariant(theme.VariantSecondary),
		),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme, w *app.Window) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.TitleBar.Layout(gtx, th, w)
}

func Demo(gtx layout.Context, th *theme.Theme, w *app.Window) layout.Dimensions {
	return defaultDemo.Layout(gtx, th, w)
}
