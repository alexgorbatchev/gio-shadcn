package empty_test

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/empty"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
	"image"
	"testing"
)

func TestActionRunsOnce(t *testing.T) {
	calls := 0
	e := empty.New(empty.Config{Action: func(gtx layout.Context) layout.Dimensions {
		calls++
		return layout.Dimensions{Size: image.Pt(100, 30)}
	}})
	h := testui.Harness{Size: image.Pt(300, 300), Widget: func(gtx layout.Context) layout.Dimensions { return e.Layout(gtx, theme.New()) }}
	h.Frame()
	if calls != 1 {
		t.Fatalf("action calls=%d, want 1", calls)
	}
}
