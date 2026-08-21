package checkbox_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/checkbox"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCheckboxCreation(t *testing.T) {
	cb := checkbox.New(checkbox.Config{
		Value: true,
	})

	if !cb.Value {
		t.Errorf("expected Value to be true")
	}
}

func TestCheckboxLayout(t *testing.T) {
	th := theme.NewDark()
	cb := checkbox.New(checkbox.Config{
		Value: true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(50, 50)),
	}
	dims := cb.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Checkbox.Layout")
	}
}
