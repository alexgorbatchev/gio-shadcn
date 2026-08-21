package inputotp_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/inputotp"
	"github.com/bnema/gio-shadcn/theme"
)

func TestInputOTPCreation(t *testing.T) {
	otp := inputotp.New(inputotp.Config{
		Length: 6,
	})

	if otp.Length != 6 {
		t.Errorf("expected Length to be 6, got %d", otp.Length)
	}
}

func TestInputOTPLayout(t *testing.T) {
	th := theme.NewDark()
	otp := inputotp.New(inputotp.Config{
		Length: 6,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 50)),
	}
	dims := otp.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from InputOTP.Layout")
	}
}
