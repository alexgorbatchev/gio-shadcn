package tabs

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	NavTabs *Tabs
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		NavTabs: New(Config{
			Tabs: []*Tab{
				NewTab("sink", "Kitchen Sink"),
				NewTab("deck", "Audio Deck"),
				NewTab("library", "Track Library"),
			},
			ActiveKey: "sink",
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.NavTabs.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
