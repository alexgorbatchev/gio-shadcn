package numberinput_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/numberinput"
	"github.com/bnema/gio-shadcn/theme"
)

func TestNumberInputCreation(t *testing.T) {
	ni := numberinput.New(numberinput.Config{
		Value: 128.0,
		Step:  1.0,
		Min:   60.0,
		Max:   200.0,
	})

	if ni.Value != 128.0 {
		t.Errorf("expected Value to be 128.0, got %f", ni.Value)
	}
}

func TestNumberInputLayout(t *testing.T) {
	th := theme.NewDark()
	ni := numberinput.New(numberinput.Config{
		Value: 128.0,
		Step:  1.0,
		Min:   60.0,
		Max:   200.0,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(200, 40)),
	}
	dims := ni.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from NumberInput.Layout")
	}
}
