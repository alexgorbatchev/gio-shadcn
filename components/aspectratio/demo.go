package aspectratio

import (
	"image"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	ar := New(Config{
		Ratio: 16.0 / 9.0,
		Widget: func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: gtx.Constraints.Min}
			theme.DrawRRectBackground(gtx, rect, gtx.Dp(th.Radius.RadiusMD), th.Colors.Secondary)
			return layout.Dimensions{Size: gtx.Constraints.Min}
		},
	})

	return ar.Layout(gtx, th)
}
