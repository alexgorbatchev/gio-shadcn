package dropdownmenu

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	DropdownMenu *DropdownMenu
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		DropdownMenu: New(Config{
			Items: []*Item{
				NewItem("Edit Track", "⌘E"),
				NewItem("Export FLAC", "⌘S"),
			},
			Open: false,
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.DropdownMenu.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
