package collapsible_test

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/collapsible"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
	"image"
	"testing"
)

func TestContentRunsOnce(t *testing.T) {
	calls := 0
	c := collapsible.New(collapsible.Config{Title: "Details", Open: true, ContentWidget: func(gtx layout.Context) layout.Dimensions {
		calls++
		return layout.Dimensions{Size: image.Pt(100, 30)}
	}})
	h := testui.Harness{Size: image.Pt(300, 200), Widget: func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, theme.New()) }}
	h.Frame()
	if calls != 1 {
		t.Fatalf("content calls=%d, want 1", calls)
	}
}
