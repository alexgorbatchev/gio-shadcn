package aspectratio_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/aspectratio"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAspectRatio16By9(t *testing.T) {
	th := theme.NewDark()
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 16.0 / 9.0,
	})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(160, 90)),
	}
	dims := ar.Layout(gtx, th)
	if dims.Size.X != 160 || dims.Size.Y != 90 {
		t.Fatalf("expected size 160x90, got %dx%d", dims.Size.X, dims.Size.Y)
	}
}

func TestAspectRatio4By3(t *testing.T) {
	th := theme.NewDark()
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 4.0 / 3.0,
	})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(120, 90)),
	}
	dims := ar.Layout(gtx, th)
	if dims.Size.X != 120 || dims.Size.Y != 90 {
		t.Fatalf("expected size 120x90, got %dx%d", dims.Size.X, dims.Size.Y)
	}
}

func TestAspectRatio1By1Square(t *testing.T) {
	th := theme.NewDark()
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 1.0,
	})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(100, 100)),
	}
	dims := ar.Layout(gtx, th)
	if dims.Size.X != 100 || dims.Size.Y != 100 {
		t.Fatalf("expected size 100x100, got %dx%d", dims.Size.X, dims.Size.Y)
	}
}

func TestAspectRatioProportionalConstraintLayout(t *testing.T) {
	th := theme.NewDark()
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 2.0,
	})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(200, 100)),
	}
	dims := ar.Layout(gtx, th)
	if dims.Size.Y != 100 {
		t.Errorf("expected height 100, got %d", dims.Size.Y)
	}
}

func TestAspectRatioWrappedChildFitting(t *testing.T) {
	th := theme.NewDark()
	childRan := false
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 16.0 / 9.0,
		Widget: func(gtx layout.Context) layout.Dimensions {
			childRan = true
			return layout.Dimensions{Size: gtx.Constraints.Min}
		},
	})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(160, 90)),
	}
	_ = ar.Layout(gtx, th)
	if !childRan {
		t.Errorf("expected wrapped child widget to be executed")
	}
}
