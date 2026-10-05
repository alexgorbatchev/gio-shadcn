package button_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestPointerHarnessControl(t *testing.T) {
	th := theme.New()
	clicks := 0
	b := button.New(button.Config{Text: "Control", OnClick: func() { clicks++ }})
	h := testui.Harness{Size: image.Pt(200, 100), Widget: func(gtx layout.Context) layout.Dimensions { return b.Layout(gtx, th) }}
	h.Frame()
	h.Click(10, 10)
	if clicks != 1 {
		t.Fatalf("control button clicks = %d, want 1", clicks)
	}
}

func TestButtonUpdatePreservesCallback(t *testing.T) {
	th := theme.New()
	clicks := 0
	b := button.New(button.Config{Text: "Save", OnClick: func() { clicks++ }})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Constraints{Max: image.Pt(200, 100)}}
	b.Click()
	b.Update(gtx)
	b.Layout(gtx, th)
	if clicks != 1 {
		t.Fatalf("Update then Layout consumes click without OnClick: calls=%d", clicks)
	}
}
