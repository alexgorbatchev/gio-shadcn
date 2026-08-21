package carousel_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/carousel"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCarouselCreation(t *testing.T) {
	car := carousel.New(carousel.Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 100)} },
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 100)} },
		},
		ActiveIndex: 0,
	})

	if len(car.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(car.Items))
	}
}

func TestCarouselLayout(t *testing.T) {
	th := theme.NewDark()
	car := carousel.New(carousel.Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 100)} },
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 100)} },
		},
		ActiveIndex: 0,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(200, 150)),
	}
	dims := car.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Carousel.Layout")
	}
}
