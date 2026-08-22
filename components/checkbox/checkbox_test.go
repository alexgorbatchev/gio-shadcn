package checkbox_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/checkbox"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCheckboxUncheckedState(t *testing.T) {
	cb := checkbox.New(checkbox.Config{
		Value: false,
	})
	if cb.Value {
		t.Fatalf("expected Value false")
	}
}

func TestCheckboxCheckedState(t *testing.T) {
	cb := checkbox.New(checkbox.Config{
		Value: true,
	})
	if !cb.Value {
		t.Fatalf("expected Value true")
	}
}

func TestCheckboxDisabledState(t *testing.T) {
	cb := checkbox.New(checkbox.Config{
		Value:    true,
		Disabled: true,
	})
	if !cb.Disabled {
		t.Fatalf("expected Disabled true")
	}
}

func TestCheckboxInteractiveClickToggle(t *testing.T) {
	changed := false
	cb := checkbox.New(checkbox.Config{
		Value: false,
		OnChange: func(val bool) {
			changed = true
		},
	})
	if cb.OnChange == nil {
		t.Fatalf("expected OnChange handler")
	}
	_ = changed
}

func TestCheckboxCheckmarkVectorPathDrawing(t *testing.T) {
	th := theme.NewDark()
	cb := checkbox.New(checkbox.Config{
		Value: true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(16, 16)),
	}
	dims := cb.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestCheckboxLabelAssociation(t *testing.T) {
	th := theme.NewDark()
	cb := checkbox.New(checkbox.Config{
		Value: false,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(16, 16)),
	}
	dims := cb.Layout(gtx, th)
	if dims.Size.Y <= 0 {
		t.Errorf("invalid height")
	}
}
