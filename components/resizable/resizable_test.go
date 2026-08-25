package resizable_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/resizable"
	"github.com/bnema/gio-shadcn/theme"
)

func TestResizableBasic(t *testing.T) {
	th := theme.NewDark()
	r := resizable.New(resizable.Config{
		Ratio: 0.5,
		LeftWidget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 50)}
		},
		RightWidget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 50)}
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims := r.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestResizableCustomRatio(t *testing.T) {
	th := theme.NewDark()
	r := resizable.New(resizable.Config{
		Ratio: 0.35,
		LeftWidget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 50)}
		},
		RightWidget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 50)}
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims := r.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
