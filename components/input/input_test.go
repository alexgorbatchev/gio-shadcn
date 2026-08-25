package input_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/theme"
)

func TestInputStandard(t *testing.T) {
	th := theme.NewDark()
	inp := input.Text("Enter text...")
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 40))}
	dims := inp.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestInputDisabled(t *testing.T) {
	th := theme.NewDark()
	inp := input.NewInput(input.WithPlaceholder("Disabled"), input.WithInputDisabled(true))
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 40))}
	dims := inp.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
