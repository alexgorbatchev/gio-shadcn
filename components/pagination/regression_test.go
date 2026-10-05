package pagination_test

import (
	"image"
	"testing"

	"gioui.org/io/input"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/pagination"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestPaginationNumberButtonSelectsPage(t *testing.T) {
	th := theme.New()
	p := pagination.New(pagination.Config{CurrentPage: 1, TotalPages: 5})
	h := testui.Harness{Size: image.Pt(600, 100), Widget: func(gtx layout.Context) layout.Dimensions { return p.Layout(gtx, th) }}
	h.Frame()
	var bounds []image.Rectangle
	for _, node := range h.Router.AppendSemantics(nil) {
		if node.Desc.Gestures&input.ClickGesture != 0 {
			bounds = append(bounds, node.Desc.Bounds)
		}
	}
	if len(bounds) != p.TotalPages+2 {
		t.Fatalf("unexpected clickable bounds: %v", bounds)
	}
	t.Logf("navigation hit bounds: %v", bounds)
	pageThree := bounds[3]
	h.Click(float32(pageThree.Min.X+pageThree.Max.X)/2, float32(pageThree.Min.Y+pageThree.Max.Y)/2)
	if p.CurrentPage != 3 {
		t.Fatalf("click on numbered page 3 leaves CurrentPage=%d", p.CurrentPage)
	}
}
