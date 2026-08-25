package radio_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/radio"
	"github.com/bnema/gio-shadcn/theme"
)

func TestRadioSelected(t *testing.T) {
	th := theme.NewDark()
	r := radio.New(radio.Config{Selected: true})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(30, 30))}
	dims := r.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestRadioUnselected(t *testing.T) {
	th := theme.NewDark()
	r := radio.New(radio.Config{Selected: false})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(30, 30))}
	dims := r.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestRadioDisabled(t *testing.T) {
	th := theme.NewDark()
	r := radio.New(radio.Config{Selected: true, Disabled: true})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(30, 30))}
	dims := r.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
