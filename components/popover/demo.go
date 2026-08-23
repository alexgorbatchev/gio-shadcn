package popover

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	anchoredPop := New(Config{Title: "Popover Title", Description: "Anchored card popover content box.", Open: false})
	return anchoredPop.Layout(gtx, th)
}
