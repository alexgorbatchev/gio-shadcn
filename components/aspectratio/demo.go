package aspectratio

import (
	"image"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	AspectRatio169 *AspectRatio
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		AspectRatio169: New(Config{
			Ratio: 16.0 / 9.0,
			Widget: func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{Max: gtx.Constraints.Min}
				theme.DrawRRectBackground(gtx, rect, 8, theme.NewDark().Colors.Secondary)
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.AspectRatio169.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
