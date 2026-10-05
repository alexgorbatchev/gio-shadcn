package sheet_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/sheet"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSheetContentRenderedOnce(t *testing.T) {
	th := theme.New()
	calls := 0
	s := sheet.New(sheet.Config{Open: true, Content: func(gtx layout.Context) layout.Dimensions {
		calls++
		return layout.Dimensions{Size: image.Pt(100, 100)}
	}})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(600, 400))}
	s.Layout(gtx, th)
	th.RenderOverlays(gtx)
	if calls != 1 {
		t.Fatalf("Sheet without trigger renders Content %d times per frame", calls)
	}
}
