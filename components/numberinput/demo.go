package numberinput

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	numStepper := New(Config{Value: 128.0, Step: 1.0, Min: 60.0, Max: 200.0})
	return numStepper.Layout(gtx, th)
}
