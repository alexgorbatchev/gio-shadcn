package input_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/theme"
)

func TestInputLayout(t *testing.T) {
	th := theme.NewDark()

	inp := input.Text("Enter text...")

	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 40)),
	}

	dims := inp.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("input returned invalid dimensions %v", dims.Size)
	}
}
