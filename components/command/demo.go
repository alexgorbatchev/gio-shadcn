package command

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	CmdPalette *Command
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		CmdPalette: New(Config{
			Placeholder: "Search command palette...",
			Items: []*Item{
				NewItem("Toggle Light/Dark Theme", "⌘T"),
				NewItem("Reset Master Audio Mixer", "⌘R"),
			},
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.CmdPalette.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
