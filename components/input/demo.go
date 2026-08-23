package input

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	txtInput := Text("Enter track title...")
	maxWidth := gtx.Metric.Dp(400)
	gtx.Constraints.Max.X = maxWidth
	gtx.Constraints.Min.X = maxWidth

	return txtInput.Layout(gtx, th)
}
