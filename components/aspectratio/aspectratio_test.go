package aspectratio_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/aspectratio"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAspectRatioDemo16x9(t *testing.T) {
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 16.0 / 9.0,
		Widget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: gtx.Constraints.Min}
		},
	})
	if ar.Ratio != 16.0/9.0 {
		t.Errorf("expected ratio 16/9, got %v", ar.Ratio)
	}
}

func TestAspectRatioPortraitDemo(t *testing.T) {
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 9.0 / 16.0,
		Widget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: gtx.Constraints.Min}
		},
	})
	if ar.Ratio != 9.0/16.0 {
		t.Errorf("expected ratio 9/16, got %v", ar.Ratio)
	}
}

func TestAspectRatioSquareDemo(t *testing.T) {
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 1.0,
		Widget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: gtx.Constraints.Min}
		},
	})
	if ar.Ratio != 1.0 {
		t.Errorf("expected ratio 1.0, got %v", ar.Ratio)
	}
}

func TestAspectRatioRTLDemo(t *testing.T) {
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 16.0 / 9.0,
		Widget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: gtx.Constraints.Min}
		},
	})
	th := theme.NewDark()
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(360, 200)),
	}
	dims := ar.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from AspectRatio.Layout: %v", dims.Size)
	}
}
