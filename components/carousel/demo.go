package carousel

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	SlideCarousel *Carousel
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		SlideCarousel: New(Config{
			Items: []layout.Widget{
				func(gtx layout.Context) layout.Dimensions {
					th := theme.NewDark()
					lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Slide 1: Audio Spectrum")
					lbl.Color = th.Colors.Foreground
					return lbl.Layout(gtx)
				},
				func(gtx layout.Context) layout.Dimensions {
					th := theme.NewDark()
					lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Slide 2: Mixer Controls")
					lbl.Color = th.Colors.Foreground
					return lbl.Layout(gtx)
				},
			},
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.SlideCarousel.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
