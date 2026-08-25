package tooltip_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/tooltip"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTooltipBasic(t *testing.T) {
	th := theme.NewDark()
	tp := tooltip.New(tooltip.Config{Text: "Add to library"})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(200, 30))}
	dims := tp.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestTooltipCustomText(t *testing.T) {
	th := theme.NewDark()
	tp := tooltip.New(tooltip.Config{Text: "Keyboard shortcut: ⌘S"})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(250, 30))}
	dims := tp.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
