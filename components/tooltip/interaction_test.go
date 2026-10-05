package tooltip_test

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/tooltip"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
	"image"
	"testing"
)

func TestTooltipClosedWithoutHover(t *testing.T) {
	tp := tooltip.New(tooltip.Config{Text: "Hint"})
	h := testui.Harness{Size: image.Pt(200, 100), Widget: func(gtx layout.Context) layout.Dimensions { return tp.Layout(gtx, theme.New()) }}
	if dims := h.Frame(); dims.Size != (image.Point{}) {
		t.Fatalf("closed tooltip occupies %v", dims.Size)
	}
}

func TestTooltipHoverTrigger(t *testing.T) {
	tp := tooltip.New(tooltip.Config{Text: "Hint"})
	th := theme.New()
	h := testui.Harness{Size: image.Pt(200, 100), Widget: func(gtx layout.Context) layout.Dimensions {
		return tp.LayoutTrigger(gtx, th, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(80, 30)} })
	}}
	h.Frame()
	h.Pointer(pointer.Move, 10, 10, 0)
	if !tp.Open {
		t.Fatal("hover did not open tooltip")
	}
	h.Pointer(pointer.Move, 180, 90, 0)
	if tp.Open {
		t.Fatal("leaving trigger did not close tooltip")
	}
}
