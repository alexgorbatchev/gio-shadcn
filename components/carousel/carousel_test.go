package carousel_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/carousel"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCarouselStandard(t *testing.T) {
	th := theme.NewDark()
	c := carousel.New(carousel.Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims := c.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestCarouselNavigation(t *testing.T) {
	th := theme.NewDark()
	c := carousel.New(carousel.Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	_ = c.Layout(gtx, th)
	if c.ActiveIndex != 0 {
		t.Errorf("expected active index 0, got %d", c.ActiveIndex)
	}
}
