package slider_test

import (
	"image"
	"testing"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/slider"
	"github.com/bnema/gio-shadcn/internal/testui"
)

func TestHorizontalPointerDragAndClamping(t *testing.T) {
	s := slider.New(slider.Config{Min: 10, Max: 110, Classes: "bg-red-500"})
	changes := 0
	s.OnChange = func(v float32) {
		changes++
		if v < 10 || v > 110 {
			t.Errorf("unbounded callback %v", v)
		}
	}
	h := testui.Harness{Size: image.Pt(200, 100), Widget: func(gtx layout.Context) layout.Dimensions { return s.Layout(gtx, nil) }}
	h.Frame()
	h.Pointer(pointer.Press, 100, 8, pointer.ButtonPrimary)
	if s.Value != 60 {
		t.Fatalf("midpoint value %v", s.Value)
	}
	h.Pointer(pointer.Move, 300, 8, pointer.ButtonPrimary)
	if s.Value != 110 {
		t.Fatal("drag beyond maximum not clamped")
	}
	h.Pointer(pointer.Move, -20, 8, pointer.ButtonPrimary)
	if s.Value != 10 {
		t.Fatal("drag below minimum not clamped")
	}
	h.Pointer(pointer.Release, -20, 8, 0)
	if changes != 3 {
		t.Fatalf("callbacks=%d", changes)
	}
	s.Disabled = true
	h.Frame()
	h.Click(150, 8)
	if s.Value != 10 || changes != 3 {
		t.Fatal("disabled slider changed")
	}
	s.Value = 200
	h.Frame()
	s.Value = -100
	h.Frame()
}

func TestVerticalPointerDragAndClamping(t *testing.T) {
	s := slider.New(slider.Config{Min: 0, Max: 100, Orientation: slider.OrientationVertical})
	h := testui.Harness{Size: image.Pt(100, 200), Widget: func(gtx layout.Context) layout.Dimensions { return s.Layout(gtx, nil) }}
	h.Frame()
	h.Pointer(pointer.Press, 8, 100, pointer.ButtonPrimary)
	if s.Value != 50 {
		t.Fatal("vertical midpoint incorrect")
	}
	h.Pointer(pointer.Move, 8, -20, pointer.ButtonPrimary)
	if s.Value != 100 {
		t.Fatal("top not clamped")
	}
	h.Pointer(pointer.Move, 8, 300, pointer.ButtonPrimary)
	if s.Value != 0 {
		t.Fatal("bottom not clamped")
	}
	h.Pointer(pointer.Release, 8, 300, 0)
	s.Disabled = true
	h.Frame()
	h.Click(8, 100)
	if s.Value != 0 {
		t.Fatal("disabled vertical slider changed")
	}
	s.Value = 200
	h.Frame()
	s.Value = -100
	h.Frame()
}
