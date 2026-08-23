package switchcomp

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	swToggle := New(Config{Value: true})

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return swToggle.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "HQ Audio Engine")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
	)
}
