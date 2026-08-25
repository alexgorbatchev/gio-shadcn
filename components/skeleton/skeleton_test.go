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
		Circle: true,
	})

	if !sk.Circle || sk.Width != unit.Dp(40) {
		t.Errorf("expected circular avatar skeleton")
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

func TestSkeletonAvatarDemo(t *testing.T) {
	skCircle := skeleton.New(skeleton.Config{Width: unit.Dp(40), Height: unit.Dp(40), Circle: true})
	skLine := skeleton.New(skeleton.Config{Width: unit.Dp(150), Height: unit.Dp(16)})
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims1 := skCircle.Layout(gtx, th)
	dims2 := skLine.Layout(gtx, th)
	if dims1.Size.X <= 0 || dims2.Size.X <= 0 {
		t.Errorf("expected valid avatar demo layout")
	}
}

func TestSkeletonCardDemo(t *testing.T) {
	skHead := skeleton.New(skeleton.Config{Width: unit.Dp(180), Height: unit.Dp(16)})
	skBody := skeleton.New(skeleton.Config{Width: unit.Dp(240), Height: unit.Dp(135)})
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 200))}
	dims1 := skHead.Layout(gtx, th)
	dims2 := skBody.Layout(gtx, th)
	if dims1.Size.X <= 0 || dims2.Size.X <= 0 {
		t.Errorf("expected valid card demo layout")
	}
}

func TestSkeletonDemo(t *testing.T) {
	demo := skeleton.NewDemoState()
	if demo == nil {
		t.Fatalf("expected non-nil demo state")
	}
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(400, 600))}
	dims := demo.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected valid demo layout")
	}
}

func TestSkeletonFormDemo(t *testing.T) {
	lbl := skeleton.New(skeleton.Config{Width: unit.Dp(80), Height: unit.Dp(16)})
	inp := skeleton.New(skeleton.Config{Width: unit.Dp(260), Height: unit.Dp(36)})
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims1 := lbl.Layout(gtx, th)
	dims2 := inp.Layout(gtx, th)
	if dims1.Size.X <= 0 || dims2.Size.X <= 0 {
		t.Errorf("expected valid form demo layout")
	}
}

func TestSkeletonTableDemo(t *testing.T) {
	col1 := skeleton.New(skeleton.Config{Width: unit.Dp(120), Height: unit.Dp(16)})
	col2 := skeleton.New(skeleton.Config{Width: unit.Dp(96), Height: unit.Dp(16)})
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims1 := col1.Layout(gtx, th)
	dims2 := col2.Layout(gtx, th)
	if dims1.Size.X <= 0 || dims2.Size.X <= 0 {
		t.Errorf("expected valid table demo layout")
	}
}

func TestSkeletonTextDemo(t *testing.T) {
	line1 := skeleton.New(skeleton.Config{Width: unit.Dp(260), Height: unit.Dp(16)})
	line2 := skeleton.New(skeleton.Config{Width: unit.Dp(195), Height: unit.Dp(16)})
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims1 := line1.Layout(gtx, th)
	dims2 := line2.Layout(gtx, th)
	if dims1.Size.X <= 0 || dims2.Size.X <= 0 {
		t.Errorf("expected valid text demo layout")
	}
}
