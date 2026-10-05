package numberinput

import (
	"image"
	"testing"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/internal/testui"
)

func TestStepperClampsAndCallsOnce(t *testing.T) {
	calls := 0
	ni := New(Config{Value: .5, Step: .3, Min: 0, Max: 1, OnChange: func(float32) { calls++ }})
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions { return ni.Layout(gtx, nil) }}
	h.Frame()
	ni.incBtn.Click()
	h.Frame()
	if ni.Value < .79 || ni.Value > .81 || calls != 1 {
		t.Fatalf("increment value=%v calls=%d", ni.Value, calls)
	}
	ni.incBtn.Click()
	h.Frame()
	ni.incBtn.Click()
	h.Frame()
	if ni.Value != 1 || calls != 2 {
		t.Fatal("maximum callback or clamp incorrect")
	}
	for i := 0; i < 5; i++ {
		ni.decBtn.Click()
		h.Frame()
	}
	if ni.Value != 0 || calls != 6 {
		t.Fatalf("minimum value=%v calls=%d", ni.Value, calls)
	}
}

func TestDecimalDisplayPreservesFraction(t *testing.T) {
	if got := formatValue(.75); got != "0.75" {
		t.Fatalf("fraction formatted %q", got)
	}
	if got := formatValue(128); got != "128" {
		t.Fatalf("integer formatted %q", got)
	}
}
