// Package hover manages native pointer hover and keyboard focus for popups.
package hover

import (
	"image"

	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
)

type Trigger struct {
	click     gesture.Click
	dismissed bool
}

func (t *Trigger) Focus(gtx layout.Context) { gtx.Execute(key.FocusCmd{Tag: t}) }

func (t *Trigger) Layout(gtx layout.Context, w layout.Widget) (layout.Dimensions, bool) {
	for {
		e, ok := t.click.Update(gtx.Source)
		if !ok {
			break
		}
		if e.Kind == gesture.KindPress {
			t.Focus(gtx)
		}
	}
	for {
		e, ok := gtx.Event(key.FocusFilter{Target: t}, key.Filter{Focus: t, Name: key.NameEscape})
		if !ok {
			break
		}
		if k, ok := e.(key.Event); ok && k.State == key.Press {
			t.dismissed = true
		}
	}
	active := t.click.Hovered() || gtx.Focused(t)
	if !active {
		t.dismissed = false
	}
	macro := op.Record(gtx.Ops)
	dims := w(gtx)
	content := macro.Stop()
	defer clip.Rect(image.Rectangle{Max: dims.Size}).Push(gtx.Ops).Pop()
	t.click.Add(gtx.Ops)
	event.Op(gtx.Ops, t)
	content.Add(gtx.Ops)
	return dims, active && !t.dismissed
}
