package hovercard_test

import (
	"image"
	"testing"

	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/hovercard"
	"github.com/bnema/gio-shadcn/internal/testui"
)

func TestHoverAndKeyboardPreview(t *testing.T) {
	c := hovercard.New(hovercard.Config{Title: "Preview", Description: "Details", Classes: "bg-blue-500"})
	focus := false
	h := testui.Harness{Size: image.Pt(400, 200), Widget: func(gtx layout.Context) layout.Dimensions {
		if focus {
			c.Focus(gtx)
			focus = false
		}
		return c.LayoutTrigger(gtx, nil, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(80, 30)} })
	}}
	h.Frame()
	h.Pointer(pointer.Move, 10, 10, 0)
	if !c.Hovered {
		t.Fatal("hover did not expose preview")
	}
	h.Pointer(pointer.Move, 350, 150, 0)
	if c.Hovered {
		t.Fatal("preview did not close")
	}
	focus = true
	h.Frame()
	h.Frame()
	if !c.Hovered {
		t.Fatal("keyboard focus did not expose preview")
	}
	h.Router.Queue(key.Event{Name: key.NameEscape, State: key.Press})
	h.Frame()
	if c.Hovered {
		t.Fatal("escape did not dismiss")
	}
	h.Click(10, 10)
	if h.Frame().Size != image.Pt(80, 30) {
		t.Fatal("preview changed trigger bounds")
	}
}
