package drawer_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/drawer"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestPanelInteriorAndOutsideClicks(t *testing.T) {
	th := theme.New()
	closed := 0
	calls := 0
	d := drawer.New(drawer.Config{Open: true, Height: 180, Title: "Panel", Classes: "bg-blue-500", OnClose: func() { closed++ }, Content: func(gtx layout.Context) layout.Dimensions { calls++; return layout.Dimensions{Size: image.Pt(20, 20)} }})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions { d.Layout(gtx, th); return th.RenderOverlays(gtx) }}
	h.Frame()
	if calls != 1 {
		t.Fatal("content rendered more than once")
	}
	h.Click(300, 350)
	if !d.Open || closed != 0 {
		t.Fatal("interior dismissed drawer")
	}
	h.Click(300, 50)
	if d.Open || closed != 1 {
		t.Fatalf("outside close=%v callbacks=%d", d.Open, closed)
	}
	for i := 0; i < 180; i++ {
		h.Frame()
	}
}
