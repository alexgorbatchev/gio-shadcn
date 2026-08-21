package radio_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/radio"
	"github.com/bnema/gio-shadcn/theme"
)

func TestRadioCreation(t *testing.T) {
	r := radio.New(radio.Config{
		Selected: true,
	})

	if !r.Selected {
		t.Errorf("expected Selected to be true")
	}
}

func TestRadioLayout(t *testing.T) {
	th := theme.NewDark()
	r := radio.New(radio.Config{
		Selected: true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(50, 50)),
	}
	dims := r.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Radio.Layout")
	}
}
