// Package testui drives actual Gio routers and layouts in component tests.
package testui

import (
	"image"
	"time"

	"gioui.org/f32"
	"gioui.org/io/input"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

type Harness struct {
	Router input.Router
	Ops    op.Ops
	Size   image.Point
	Widget layout.Widget
	Now    time.Time
}

func (h *Harness) Frame() layout.Dimensions {
	h.Ops.Reset()
	if h.Now.IsZero() {
		h.Now = time.Now()
	}
	dims := h.Widget(layout.Context{
		Ops: &h.Ops, Source: h.Router.Source(),
		Constraints: layout.Constraints{Max: h.Size},
		Metric:      unit.Metric{PxPerDp: 1, PxPerSp: 1}, Now: h.Now,
	})
	h.Router.Frame(&h.Ops)
	return dims
}

func (h *Harness) Pointer(kind pointer.Kind, x, y float32, buttons pointer.Buttons) {
	h.Router.Queue(pointer.Event{Kind: kind, Source: pointer.Mouse, Buttons: buttons, Position: f32.Pt(x, y)})
	h.Frame()
}

func (h *Harness) Click(x, y float32) {
	h.Pointer(pointer.Press, x, y, pointer.ButtonPrimary)
	h.Pointer(pointer.Release, x, y, 0)
	h.Frame()
}

func (h *Harness) Edit(text string, start, end int) {
	h.Router.Queue(key.EditEvent{Range: key.Range{Start: start, End: end}, Text: text})
	h.Frame()
	h.Frame()
}
