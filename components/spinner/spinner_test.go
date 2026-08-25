package spinner_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/spinner"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSpinnerCreation(t *testing.T) {
	s := spinner.New(spinner.Config{
		Size: unit.Dp(32),
	})

	if s.Size != unit.Dp(32) {
		t.Errorf("expected spinner size 32dp")
	}
}

func TestSpinnerDefaultSize(t *testing.T) {
	s := spinner.New(spinner.Config{})
	if s.Size != unit.Dp(24) {
		t.Errorf("expected default spinner size 24dp")
	}
}

func TestSpinnerBadgeDemo(t *testing.T) {
	s := spinner.New(spinner.Config{Size: unit.Dp(14)})
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(50, 50))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected valid spinner badge layout")
	}
}

func TestSpinnerButtonDemo(t *testing.T) {
	s := spinner.New(spinner.Config{Size: unit.Dp(16)})
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(50, 50))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected valid spinner button layout")
	}
}

func TestSpinnerCustomDemo(t *testing.T) {
	s := spinner.New(spinner.Config{Size: unit.Dp(16)})
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(50, 50))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected valid custom spinner layout")
	}
}

func TestSpinnerEmptyDemo(t *testing.T) {
	s := spinner.New(spinner.Config{Size: unit.Dp(32)})
	th := theme.NewDark()
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 100))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected valid empty spinner layout")
	}
}

func TestSpinnerDemoLayout(t *testing.T) {
	demo := spinner.NewDemoState()
	if demo == nil {
		t.Fatalf("expected non-nil demo state")
	}
	th := theme.NewDark()
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(400, 600)),
	}
	dims := demo.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected valid demo layout")
	}
}
