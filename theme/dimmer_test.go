package theme_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDimmerLayoutAndClose(t *testing.T) {
	th := theme.NewDark()
	dimmer := theme.NewDimmer()

	closed := false
	onClose := func() {
		closed = true
	}

	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(800, 600)),
	}

	dims := dimmer.Layout(gtx, th, onClose)
	if dims.Size.X != 800 || dims.Size.Y != 600 {
		t.Errorf("expected dimmer to fill 800x600, got %v", dims.Size)
	}
	if closed {
		t.Errorf("expected closed to remain false before click event")
	}
}
