package skeleton

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	shimmerSk := New(Config{Width: unit.Dp(180), Height: unit.Dp(24)})
	return shimmerSk.Layout(gtx, th)
}
