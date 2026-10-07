package dialog_test

import (
	"image"
	"testing"

	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/dialog"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDialogTrapsNativeContentFocus(t *testing.T) {
	th := theme.New()
	var background, content widget.Editor
	var action widget.Clickable
	d := dialog.New(dialog.Config{Open: true, Content: func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(material.Editor(th.MaterialTheme, &content, "Content").Layout),
			layout.Rigid(material.Button(th.MaterialTheme, &action, "Action").Layout),
		)
	}})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
		return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions {
			material.Editor(th.MaterialTheme, &background, "Background").Layout(gtx)
			return d.Layout(gtx, th)
		})
	}}
	h.Frame()
	h.Frame()
	seenEditor, seenAction := false, false
	for _, mods := range []key.Modifiers{0, key.ModShift} {
		for range 12 {
			h.Key(key.NameTab, mods)
			if h.Router.Source().Focused(&background) {
				t.Fatal("Tab escaped to the background editor")
			}
			seenEditor = seenEditor || h.Router.Source().Focused(&content)
			seenAction = seenAction || h.Router.Source().Focused(&action)
		}
	}
	if !seenEditor || !seenAction {
		t.Fatalf("native custom content skipped: editor=%v, action=%v", seenEditor, seenAction)
	}
}

func TestDialogEscapeRestoresTriggerFocus(t *testing.T) {
	th := theme.New()
	cancels := 0
	d := dialog.New(dialog.Config{TriggerText: "Open", OnCancel: func() { cancels++ }})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
		return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions { return d.Layout(gtx, th) })
	}}
	h.Frame()
	d.TriggerButton.Click()
	h.Frame()
	h.Frame()
	h.Key(key.NameEscape, 0)
	if d.Open || cancels != 1 {
		t.Fatalf("Escape: open=%v, cancel calls=%d", d.Open, cancels)
	}
	h.Key(key.NameReturn, 0)
	if !d.Open {
		t.Fatal("Enter did not reopen the dialog from its restored trigger")
	}
}

func TestNestedDialogEscapeAndFocus(t *testing.T) {
	th := theme.New()
	innerCancels, outerCancels := 0, 0
	inner := dialog.New(dialog.Config{TriggerText: "Inner", Open: true, OnCancel: func() { innerCancels++ }})
	outer := dialog.New(dialog.Config{Open: true, OnCancel: func() { outerCancels++ }, Content: func(gtx layout.Context) layout.Dimensions {
		return inner.Layout(gtx, th)
	}})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
		return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions { return outer.Layout(gtx, th) })
	}}
	h.Frame()
	h.Frame()
	h.Key(key.NameEscape, 0)
	if inner.Open || !outer.Open || innerCancels != 1 || outerCancels != 0 {
		t.Fatalf("nested Escape: inner=%v, outer=%v, cancellations=%d/%d", inner.Open, outer.Open, innerCancels, outerCancels)
	}
	h.Key(key.NameReturn, 0)
	if !inner.Open || !outer.Open {
		t.Fatal("nested modal did not restore focus to its trigger inside the parent")
	}
}

func TestDialogCustomReturnFocusAndRemoval(t *testing.T) {
	for _, remove := range []bool{false, true} {
		t.Run(map[bool]string{false: "Escape", true: "unmount"}[remove], func(t *testing.T) {
			th := theme.New()
			var background widget.Editor
			d := dialog.New(dialog.Config{Open: true})
			d.Modal.ReturnFocus = func(gtx layout.Context) { gtx.Execute(key.FocusCmd{Tag: &background}) }
			present := true
			h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
				return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions {
					material.Editor(th.MaterialTheme, &background, "Background").Layout(gtx)
					if present {
						return d.Layout(gtx, th)
					}
					return layout.Dimensions{}
				})
			}}
			h.Frame()
			h.Frame()
			if remove {
				present = false
			} else {
				h.Key(key.NameEscape, 0)
			}
			h.Frame()
			h.Frame()
			h.Frame()
			h.Edit("restored", 0, 0)
			if background.Text() != "restored" {
				t.Fatalf("return target did not receive typing: %q", background.Text())
			}
		})
	}
}
