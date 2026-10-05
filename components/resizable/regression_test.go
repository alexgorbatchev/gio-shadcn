package resizable_test

import (
	"image"
	"testing"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/resizable"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestResizableDragChangesRatio(t *testing.T) {
	th := theme.New()
	r := resizable.New(resizable.Config{Ratio: .5})
	h := testui.Harness{Size: image.Pt(200, 100), Widget: func(gtx layout.Context) layout.Dimensions { return r.Layout(gtx, th) }}
	h.Frame()
	h.Router.Queue(pointer.Event{Kind: pointer.Press, Source: pointer.Mouse, Buttons: pointer.ButtonPrimary, Position: f32.Pt(100, 50)})
	h.Frame()
	h.Router.Queue(pointer.Event{Kind: pointer.Move, Source: pointer.Mouse, Buttons: pointer.ButtonPrimary, Position: f32.Pt(150, 50)})
	h.Frame()
	h.Router.Queue(pointer.Event{Kind: pointer.Release, Source: pointer.Mouse, Position: f32.Pt(150, 50)})
	h.Frame()
	if r.Ratio == .5 {
		t.Fatalf("divider drag from x=100 to x=150 leaves Ratio=%v", r.Ratio)
	}
}
