package slider

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	rngSlider := New(Config{Value: 65.0, Min: 0.0, Max: 100.0})
	return rngSlider.Layout(gtx, th)
}
