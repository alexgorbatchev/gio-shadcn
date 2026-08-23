package progress

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	progBar := New(Config{Value: 0.65})
	return progBar.Layout(gtx, th)
}
