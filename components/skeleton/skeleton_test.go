package skeleton_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/skeleton"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSkeletonCreation(t *testing.T) {
	sk := skeleton.New(skeleton.Config{
		Width:  unit.Dp(150),
		Height: unit.Dp(30),
	})

	if sk.Width != unit.Dp(150) {
		t.Errorf("expected Width to be 150dp, got %v", sk.Width)
	}
}

func TestSkeletonLayout(t *testing.T) {
	th := theme.NewDark()
	sk := skeleton.New(skeleton.Config{
		Width:  unit.Dp(150),
		Height: unit.Dp(30),
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(200, 50)),
	}
	dims := sk.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Skeleton.Layout")
	}
}
