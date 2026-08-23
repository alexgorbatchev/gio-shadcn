package dropdownmenu

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	dropdownMenu := New(Config{
		Items: []*Item{
			NewItem("Edit Track", "⌘E"),
			NewItem("Export FLAC", "⌘S"),
		},
		Open: false,
	})

	return dropdownMenu.Layout(gtx, th)
}
