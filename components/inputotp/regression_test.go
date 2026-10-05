package inputotp_test

import (
	"image"
	"testing"

	"gioui.org/io/key"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/inputotp"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestOTPRejectsOverlongCellText(t *testing.T) {
	th := theme.New()
	otp := inputotp.New(inputotp.Config{Length: 4, Value: "1234"})
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions { return otp.Layout(gtx, th) }}
	h.Frame()
	h.Click(15, 15)
	h.Router.Queue(key.EditEvent{Range: key.Range{}, Text: "5678901"})
	h.Frame()
	h.Frame()
	if len(otp.Value) > otp.Length {
		t.Fatalf("Length=%d but one cell accepts Value=%q", otp.Length, otp.Value)
	}
	if otp.Value == "1234" {
		t.Fatal("edit did not reach editor; probe inconclusive")
	}
}
