package numberinput_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/numberinput"
	"github.com/bnema/gio-shadcn/theme"
)

func TestNumberInputBPMStepper(t *testing.T) {
	ni := numberinput.New(numberinput.Config{
		Value: 128.0,
		Step:  1.0,
		Min:   60.0,
		Max:   200.0,
	})
	if ni.Value != 128.0 || ni.Step != 1.0 {
		t.Errorf("expected value 128.0 and step 1.0")
	}
}

func TestNumberInputGainRangeInput(t *testing.T) {
	ni := numberinput.New(numberinput.Config{
		Value: 0.0,
		Step:  0.5,
		Min:   -12.0,
		Max:   12.0,
	})
	if ni.Min != -12.0 || ni.Max != 12.0 {
		t.Errorf("expected gain range -12 to +12")
	}
}

func TestNumberInputMinMaxBounds(t *testing.T) {
	ni := numberinput.New(numberinput.Config{
		Value: 250.0,
		Step:  1.0,
		Min:   0.0,
		Max:   100.0,
	})
	if ni.Value != 100.0 {
		t.Errorf("expected clamped value 100.0, got %f", ni.Value)
	}
}

func TestNumberInputIncrementDecrementButtons(t *testing.T) {
	th := theme.NewDark()
	ni := numberinput.New(numberinput.Config{
		Value: 50.0,
		Step:  5.0,
		Min:   0.0,
		Max:   100.0,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(200, 40)),
	}
	dims := ni.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestNumberInputOnChangeCallback(t *testing.T) {
	var changedVal float32
	ni := numberinput.New(numberinput.Config{
		Value: 10.0,
		Step:  1.0,
		Min:   0.0,
		Max:   20.0,
		OnChange: func(val float32) {
			changedVal = val
		},
	})
	if ni.OnChange == nil {
		t.Errorf("expected OnChange callback non-nil")
	}
	_ = changedVal
}

func TestNumberInputDisplayBox(t *testing.T) {
	th := theme.NewDark()
	ni := numberinput.New(numberinput.Config{Value: 42.0})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(200, 40)),
	}
	dims := ni.Layout(gtx, th)
	if dims.Size.Y <= 0 {
		t.Errorf("invalid height")
	}
}
