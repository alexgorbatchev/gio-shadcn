package checkbox_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/checkbox"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCheckboxUnchecked(t *testing.T) {
	th := theme.NewDark()
	c := checkbox.New(checkbox.Config{Value: false})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(30, 30))}
	dims := c.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestCheckboxChecked(t *testing.T) {
	th := theme.NewDark()
	c := checkbox.New(checkbox.Config{Value: true})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(30, 30))}
	dims := c.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestCheckboxDisabled(t *testing.T) {
	th := theme.NewDark()
	c := checkbox.New(checkbox.Config{Value: true, Disabled: true})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(30, 30))}
	dims := c.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestCheckboxClickToggle(t *testing.T) {
	th := theme.NewDark()
	toggled := false
	c := checkbox.New(checkbox.Config{
		Value: false,
		OnChange: func(val bool) {
			toggled = val
		},
	})
	h := testui.Harness{Size: image.Pt(100, 40), Widget: func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, th) }}
	h.Frame()
	h.Click(10, 10)
	if !toggled || !c.Value {
		t.Fatal("click did not check the checkbox or invoke OnChange")
	}
	h.Click(10, 10)
	if toggled || c.Value {
		t.Fatal("second click did not uncheck the checkbox")
	}
}
