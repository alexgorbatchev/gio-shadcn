package sheet_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/sheet"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestPanelInteriorAndOutsideClicks(t *testing.T) {
	for _, side := range []sheet.Side{sheet.SideRight, sheet.SideLeft, sheet.SideTop, sheet.SideBottom} {
		th := theme.New()
		closed := 0
		calls := 0
		s := sheet.New(sheet.Config{Open: true, Side: side, Width: 200, Height: 180, Title: "Panel", Classes: "bg-blue-500", OnClose: func() { closed++ }, Content: func(gtx layout.Context) layout.Dimensions { calls++; return layout.Dimensions{Size: image.Pt(20, 20)} }})
		h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions { s.Layout(gtx, th); return th.RenderOverlays(gtx) }}
		h.Frame()
		if calls != 1 {
			t.Fatal("content rendered more than once")
		}
		x, y := float32(500), float32(300)
		outsideX, outsideY := float32(100), float32(100)
		switch side {
		case sheet.SideLeft:
			x = 100
			outsideX = 500
		case sheet.SideTop:
			x, y = 300, 140
			outsideX, outsideY = 300, 350
		case sheet.SideBottom:
			x, y = 300, 350
			outsideX, outsideY = 300, 50
		}
		h.Click(x, y)
		if !s.Open || closed != 0 {
			t.Fatalf("side %d interior dismissed", side)
		}
		h.Click(outsideX, outsideY)
		if s.Open || closed != 1 {
			t.Fatalf("side %d outside close=%v callbacks=%d", side, s.Open, closed)
		}
		for i := 0; i < 180; i++ {
			h.Frame()
		}
		if th.HasOverlays() {
			t.Fatal("overlay queue did not flush")
		}
	}
}
