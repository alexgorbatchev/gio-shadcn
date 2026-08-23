package collapsible

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	colContainer := New(Config{Title: "Advanced Mixer Settings", Content: "ASIO Direct hardware routing enabled.", Open: true})
	return colContainer.Layout(gtx, th)
}
