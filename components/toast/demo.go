package toast

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ToastItem *Toast
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		ToastItem: New(Config{Title: "Track Exported", Description: "Exported to Starlight_Symphony.flac", Visible: true}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.ToastItem.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
