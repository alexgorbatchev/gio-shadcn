package dialog_test

import (
	"image"
	"testing"

	"gioui.org/io/key"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/dialog"
	inputcomp "github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDialogInteriorDoesNotDismiss(t *testing.T) {
	th := theme.New()
	cancels := 0
	d := dialog.New(dialog.Config{Open: true, Title: "Dialog", OnCancel: func() { cancels++ }, Content: func(gtx layout.Context) layout.Dimensions {
		theme.DrawRRectBackground(gtx, image.Rect(0, 0, 200, 100), 0, th.Colors.Card)
		return layout.Dimensions{Size: image.Pt(200, 100)}
	}})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions { d.Layout(gtx, th); return th.RenderOverlays(gtx) }}
	h.Frame()
	h.Click(300, 200)
	if !d.Open || cancels != 0 {
		t.Fatalf("click inside dialog: Open=%v, OnCancel calls=%d", d.Open, cancels)
	}
}

func TestDialogBlocksBackgroundTyping(t *testing.T) {
	th := theme.New()
	i := inputcomp.NewInput(inputcomp.WithPlaceholder("Background input"))
	d := dialog.New(dialog.Config{Title: "Modal"})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
		i.Layout(gtx, th)
		d.Layout(gtx, th)
		return th.RenderOverlays(gtx)
	}}
	h.Frame()
	h.Click(20, 20)
	h.Router.Queue(key.EditEvent{Text: "before"})
	h.Frame()
	h.Frame()
	if i.Text() != "before" {
		t.Fatalf("background focus control failed: %q", i.Text())
	}
	d.Open = true
	h.Frame()
	h.Router.Queue(key.EditEvent{Range: key.Range{Start: 6, End: 6}, Text: "after"})
	h.Frame()
	h.Frame()
	if i.Text() != "before" {
		t.Fatalf("typing while modal is open changes background input to %q", i.Text())
	}
}
