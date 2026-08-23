package toast

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	toastItem := New(Config{Title: "Track Exported", Description: "Exported to Starlight_Symphony.flac", Visible: true})
	return toastItem.Layout(gtx, th)
}
