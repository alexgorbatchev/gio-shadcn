package input_test

import (
	"image"
	"testing"

	"gioui.org/io/key"
	"gioui.org/layout"
	inputcomp "github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestInputFocusCallback(t *testing.T) {
	th := theme.New()
	focuses := 0
	i := inputcomp.NewInput(inputcomp.WithPlaceholder("Input"))
	i.OnFocus = func() { focuses++ }
	h := testui.Harness{Size: image.Pt(200, 100), Widget: func(gtx layout.Context) layout.Dimensions { return i.Layout(gtx, th) }}
	h.Frame()
	h.Click(20, 20)
	h.Router.Queue(key.EditEvent{Text: "hello"})
	h.Frame()
	h.Frame()
	if i.Text() != "hello" {
		t.Fatalf("editor did not receive edit: %q; probe inconclusive", i.Text())
	}
	if focuses != 1 {
		t.Fatalf("focused editor accepts typing, but OnFocus calls=%d", focuses)
	}
}
