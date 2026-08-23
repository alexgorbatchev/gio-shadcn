package titlebar

import (
	"gioui.org/app"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme, w *app.Window) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	tb := NewTitleBar(
		WithTitle("Window TitleBar"),
		WithWindow(w),
		WithVariant(theme.VariantSecondary),
	)

	return tb.Layout(gtx, th, w)
}
