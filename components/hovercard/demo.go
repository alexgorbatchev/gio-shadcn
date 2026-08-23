package hovercard

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	hoverCardItem := New(Config{Title: "Artist Profile", Description: "Aethelgard - Progressive House", Hovered: false})
	return hoverCardItem.Layout(gtx, th)
}
