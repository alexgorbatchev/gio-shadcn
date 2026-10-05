package textarea_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/textarea"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTextAreaStandard(t *testing.T) {
	th := theme.NewDark()
	ta := textarea.New(textarea.Config{Placeholder: "Write..."})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims := ta.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestTextAreaEditsAndCallsChange(t *testing.T) {
	th := theme.NewDark()
	changes := 0
	ta := textarea.New(textarea.Config{Placeholder: "Write...", OnChange: func(value string) {
		changes++
		if value != "hello\nworld" {
			t.Errorf("change text=%q", value)
		}
	}})
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions { return ta.Layout(gtx, th) }}
	h.Frame()
	h.Click(20, 20)
	h.Edit("hello\nworld", 0, 0)
	if ta.Text != "hello\nworld" || changes != 1 {
		t.Fatalf("text=%q callbacks=%d", ta.Text, changes)
	}
}
