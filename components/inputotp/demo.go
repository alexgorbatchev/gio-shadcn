package inputotp

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	otpPin := New(Config{Length: 6})
	return otpPin.Layout(gtx, th)
}
