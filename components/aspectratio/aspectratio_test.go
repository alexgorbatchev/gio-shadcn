package aspectratio_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/aspectratio"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAspectRatioCreation(t *testing.T) {
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 16.0 / 9.0,
	})

	if ar.Ratio <= 0 {
		t.Errorf("expected Ratio to be > 0, got %f", ar.Ratio)
	}
}

func TestAspectRatioLayout(t *testing.T) {
	th := theme.NewDark()
	ar := aspectratio.New(aspectratio.Config{
		Ratio: 16.0 / 9.0,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(160, 90)),
	}
	dims := ar.Layout(gtx, th)

	if dims.Size.X != 160 || dims.Size.Y != 90 {
		t.Errorf("expected size 160x90, got %dx%d", dims.Size.X, dims.Size.Y)
	}
}
