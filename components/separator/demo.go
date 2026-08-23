package separator

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	sepDivider := New(Config{Horizontal: true})
	return sepDivider.Layout(gtx, th)
}
