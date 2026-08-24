package inputotp

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	OtpPin *InputOTP
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		OtpPin: New(Config{Length: 6}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.OtpPin.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
