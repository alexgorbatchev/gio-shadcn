package radio

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	radioA := New(Config{Selected: true})
	radioB := New(Config{Selected: false})

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return radioA.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Master Output")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return radioB.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Headphones Cue")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
	)
}
