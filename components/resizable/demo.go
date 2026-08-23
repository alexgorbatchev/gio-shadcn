package resizable

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	resPanel := New(Config{
		Ratio: 0.5,
		LeftWidget: func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Left Split Panel")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		},
		RightWidget: func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Right Split Panel")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		},
	})

	return resPanel.Layout(gtx, th)
}
