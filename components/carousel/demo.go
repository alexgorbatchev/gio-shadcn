package carousel

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	slideCarousel := New(Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Slide 1: Audio Spectrum")
				lbl.Color = th.Colors.Foreground
				return lbl.Layout(gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Slide 2: Mixer Controls")
				lbl.Color = th.Colors.Foreground
				return lbl.Layout(gtx)
			},
		},
	})

	return slideCarousel.Layout(gtx, th)
}
