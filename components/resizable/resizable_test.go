package resizable_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/resizable"
	"github.com/bnema/gio-shadcn/theme"
)

func TestResizableCreation(t *testing.T) {
	rz := resizable.New(resizable.Config{
		Ratio: 0.6,
	})

	if rz.Ratio != 0.6 {
		t.Errorf("expected Ratio to be 0.6, got %f", rz.Ratio)
	}
}

func TestResizableLayout(t *testing.T) {
	th := theme.NewDark()
	rz := resizable.New(resizable.Config{
		Ratio: 0.5,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(400, 200)),
	}
	dims := rz.Layout(gtx, th)

	if dims.Size.X != 400 || dims.Size.Y != 200 {
		t.Errorf("invalid dimensions returned from Resizable.Layout")
	}
}
