package titlebar

import (
	"image"
	"testing"

	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTitlebarRegistersNativeMoveRegion(t *testing.T) {
	tb := NewTitleBar(WithTitle("Window"), WithControls(false))
	h := testui.Harness{Size: image.Pt(400, 40), Widget: func(gtx layout.Context) layout.Dimensions { return tb.Layout(gtx, theme.New(), nil) }}
	h.Frame()
	if action, ok := h.Router.ActionAt(f32.Pt(30, 20)); !ok || action != system.ActionMove {
		t.Fatalf("native move region=%v, %v", action, ok)
	}
}
