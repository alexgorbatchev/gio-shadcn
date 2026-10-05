package theme_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/vibrantgio/effects/spring"
)

func TestDimmerLayoutAndClose(t *testing.T) {
	th := theme.NewDark()
	dimmer := theme.NewDimmer()
	defer dimmer.Release()

	closed := false
	onClose := func() {
		closed = true
	}

	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(800, 600)),
	}

	dims := dimmer.Layout(gtx, th, onClose)
	if dims.Size.X != 800 || dims.Size.Y != 600 {
		t.Errorf("expected dimmer to fill 800x600, got %v", dims.Size)
	}
	if closed {
		t.Errorf("expected closed to remain false before click event")
	}

	// Test LayoutWithAlpha
	dimsAlpha := dimmer.LayoutWithAlpha(gtx, th, 120, onClose)
	if dimsAlpha.Size.X != 800 || dimsAlpha.Size.Y != 600 {
		t.Errorf("expected dimmer with alpha to fill 800x600, got %v", dimsAlpha.Size)
	}
}

func TestSpringSpeed(t *testing.T) {
	s := spring.New(0, 0, spring.Options{
		Stiffness: 1200.0,
		Damping:   69.0,
	})
	s.SetTarget(1.0)
	settledIn := -1
	for f := 0; f < 60; f++ {
		s.Tick(60.0)
		if s.Settled(0.005) {
			settledIn = f
			break
		}
	}
	t.Logf("Spring with k=1200, c=69 settled in %d frames (approx %.1f ms)", settledIn, float64(settledIn)*16.67)
}
