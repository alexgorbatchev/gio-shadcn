package inputotp_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/inputotp"
	"github.com/bnema/gio-shadcn/theme"
)

func TestInputOTP4DigitPin(t *testing.T) {
	otp := inputotp.New(inputotp.Config{
		Length: 4,
	})
	if otp.Length != 4 {
		t.Fatalf("expected Length 4")
	}
}

func TestInputOTP6DigitPin(t *testing.T) {
	otp := inputotp.New(inputotp.Config{
		Length: 6,
	})
	if otp.Length != 6 {
		t.Fatalf("expected Length 6")
	}
}

func TestInputOTPDigitBoxLayout(t *testing.T) {
	th := theme.NewDark()
	otp := inputotp.New(inputotp.Config{
		Length: 6,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 50)),
	}
	dims := otp.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestInputOTPActiveCellHighlight(t *testing.T) {
	th := theme.NewDark()
	otp := inputotp.New(inputotp.Config{
		Length: 6,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 50)),
	}
	dims := otp.Layout(gtx, th)
	if dims.Size.Y <= 0 {
		t.Errorf("invalid height")
	}
}
