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

func TestSkeletonTextLineVariant(t *testing.T) {
	sk := skeleton.New(skeleton.Config{
		Width:  unit.Dp(180),
		Height: unit.Dp(16),
	})

	if sk.Width != unit.Dp(180) || sk.Height != unit.Dp(16) {
		t.Errorf("expected text line skeleton 180x16dp")
	}
}

func TestSkeletonAvatarCircleVariant(t *testing.T) {
	sk := skeleton.New(skeleton.Config{
		Width:  unit.Dp(40),
		Height: unit.Dp(40),
	})

	if sk.Width != unit.Dp(40) || sk.Height != unit.Dp(40) {
		t.Errorf("expected avatar circle skeleton 40x40dp")
	}
}

func TestSkeletonCardContainerVariant(t *testing.T) {
	sk := skeleton.New(skeleton.Config{
		Width:  unit.Dp(300),
		Height: unit.Dp(120),
	})

	if sk.Width != unit.Dp(300) || sk.Height != unit.Dp(120) {
		t.Errorf("expected card container skeleton 300x120dp")
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
