package carousel_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/carousel"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCarouselHorizontalSlide(t *testing.T) {
	car := carousel.New(carousel.Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
		},
		ActiveIndex: 0,
	})
	if len(car.Items) != 2 {
		t.Fatalf("expected 2 items")
	}
}

func TestCarouselActiveIndexTracking(t *testing.T) {
	car := carousel.New(carousel.Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
		},
		ActiveIndex: 1,
	})
	if car.ActiveIndex != 1 {
		t.Fatalf("expected ActiveIndex 1")
	}
}

func TestCarouselNextPrevControls(t *testing.T) {
	th := theme.NewDark()
	car := carousel.New(carousel.Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
			func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{Size: image.Pt(100, 50)} },
		},
		ActiveIndex: 0,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	dims := car.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestCarouselSlideItemContainer(t *testing.T) {
	th := theme.NewDark()
	itemRan := false
	car := carousel.New(carousel.Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions {
				itemRan = true
				return layout.Dimensions{Size: image.Pt(50, 50)}
			},
		},
		ActiveIndex: 0,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	_ = car.Layout(gtx, th)
	if !itemRan {
		t.Errorf("expected slide item to execute")
	}
}
