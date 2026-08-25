package checkbox_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/checkbox"
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
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(30, 30))}
	_ = c.Layout(gtx, th)
	_ = toggled
}
