package switchcomp_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	switchcomp "github.com/bnema/gio-shadcn/components/switch"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSwitchCreation(t *testing.T) {
	sw := switchcomp.New(switchcomp.Config{
		Value: true,
	})

	if !sw.Value {
		t.Errorf("expected Value to be true")
	}
}

func TestSwitchLayout(t *testing.T) {
	th := theme.NewDark()
	sw := switchcomp.New(switchcomp.Config{
		Value: false,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(100, 50)),
	}
	dims := sw.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Switch.Layout")
	}
}
