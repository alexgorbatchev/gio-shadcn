package scrollarea

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	sa := New(Config{
		Widget: func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Scrollable Container Body Area")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		},
	})
	return sa.Layout(gtx, th)
}
