package command

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	cmdPalette := New(Config{
		Placeholder: "Search command palette...",
		Items: []*Item{
			NewItem("Toggle Light/Dark Theme", "⌘T"),
			NewItem("Reset Master Audio Mixer", "⌘R"),
		},
	})

	return cmdPalette.Layout(gtx, th)
}
