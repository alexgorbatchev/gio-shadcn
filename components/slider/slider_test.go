package slider_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/slider"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSliderCreation(t *testing.T) {
	sl := slider.New(slider.Config{
		Value: 65.0,
		Min:   0.0,
		Max:   100.0,
	})

	if sl.Value != 65.0 {
		t.Errorf("expected Value to be 65.0, got %f", sl.Value)
	}
}

func TestSliderLayout(t *testing.T) {
	th := theme.NewDark()
	sl := slider.New(slider.Config{
		Value: 50.0,
		Min:   0.0,
		Max:   100.0,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 30)),
	}
	dims := sl.Layout(gtx, th)

	if dims.Size.X != 300 {
		t.Errorf("expected width to be 300, got %d", dims.Size.X)
	}
}
