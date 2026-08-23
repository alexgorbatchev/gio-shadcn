package tabs

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	navTabs := New(Config{
		Tabs: []*Tab{
			NewTab("sink", "Kitchen Sink"),
			NewTab("deck", "Audio Deck"),
			NewTab("library", "Track Library"),
		},
		ActiveKey: "sink",
	})

	return navTabs.Layout(gtx, th)
}
