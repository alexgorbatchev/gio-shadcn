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
	s := slider.New(slider.Config{
		Value: 50.0,
		Min:   0.0,
		Max:   100.0,
	})

	if s.Value != 50.0 || s.Min != 0.0 || s.Max != 100.0 {
		t.Errorf("unexpected slider initialization")
	}
}

func TestSliderDisabledState(t *testing.T) {
	s := slider.New(slider.Config{
		Value:    25.0,
		Disabled: true,
	})

	if !s.Disabled {
		t.Errorf("expected slider to be disabled")
	}
}

func TestSliderControlledDemo(t *testing.T) {
	var val float32 = 0.5
	s := slider.New(slider.Config{
		Value: val,
		Min:   0.0,
		Max:   1.0,
		OnChange: func(v float32) {
			val = v
		},
	})
	if s.Value != 0.5 {
		t.Errorf("expected controlled slider value 0.5")
	}
}

func TestSliderMultipleDemo(t *testing.T) {
	s1 := slider.New(slider.Config{Value: 10.0, Min: 0.0, Max: 100.0})
	s2 := slider.New(slider.Config{Value: 70.0, Min: 0.0, Max: 100.0})
	if s1.Value != 10.0 || s2.Value != 70.0 {
		t.Errorf("expected multiple sliders")
	}
}

func TestSliderRangeDemo(t *testing.T) {
	s := slider.New(slider.Config{Value: 25.0, Min: 0.0, Max: 100.0})
	if s.Value != 25.0 {
		t.Errorf("expected range slider")
	}
}

func TestSliderVerticalDemo(t *testing.T) {
	th := theme.NewDark()
	s := slider.New(slider.Config{
		Value:       50.0,
		Min:         0.0,
		Max:         100.0,
		Orientation: slider.OrientationVertical,
	})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(30, 200)),
	}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid vertical slider layout dimensions")
	}
}

func TestSliderDemoLayout(t *testing.T) {
	demo := slider.NewDemoState()
	if demo == nil {
		t.Fatalf("expected non-nil demo state")
	}
	th := theme.NewDark()
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(400, 600)),
	}
	dims := demo.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected valid demo layout")
	}
}
