package inputotp_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/inputotp"
	"github.com/bnema/gio-shadcn/theme"
)

func TestInputOTP6Digit(t *testing.T) {
	th := theme.NewDark()
	otp := inputotp.New(inputotp.Config{Length: 6})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 50))}
	dims := otp.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestInputOTP4Digit(t *testing.T) {
	th := theme.NewDark()
	otp := inputotp.New(inputotp.Config{Length: 4})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(200, 50))}
	dims := otp.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
