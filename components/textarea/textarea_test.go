package textarea_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/textarea"
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

func TestTextAreaDisabled(t *testing.T) {
	th := theme.NewDark()
	ta := textarea.New(textarea.Config{Placeholder: "Disabled"})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims := ta.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
