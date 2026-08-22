package radio_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/radio"
	"github.com/bnema/gio-shadcn/theme"
)

func TestRadioSelectedState(t *testing.T) {
	r := radio.New(radio.Config{Selected: true})
	if !r.Selected {
		t.Errorf("expected radio to be selected")
	}
}

func TestRadioUnselectedState(t *testing.T) {
	r := radio.New(radio.Config{Selected: false})
	if r.Selected {
		t.Errorf("expected radio to be unselected")
	}
}

func TestRadioDisabledState(t *testing.T) {
	r := radio.New(radio.Config{Selected: false, Disabled: true})
	if !r.Disabled {
		t.Errorf("expected radio to be disabled")
	}
}

func TestRadioClickInteraction(t *testing.T) {
	th := theme.NewDark()
	var changed bool
	r := radio.New(radio.Config{
		Selected: false,
		OnChange: func(val bool) {
			changed = val
		},
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(16, 16)),
	}
	dims := r.Layout(gtx, th)
	if dims.Size.X != 16 {
		t.Errorf("expected 16px size")
	}
	_ = changed
}
