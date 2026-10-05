package progress_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/progress"
	"github.com/bnema/gio-shadcn/theme"
)

func TestProgressBasic(t *testing.T) {
	th := theme.NewDark()
	p := progress.New(progress.Config{Value: 0.66})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 8))}
	dims := p.Layout(gtx, th)
	if dims.Size.X != 300 || dims.Size.Y != 8 {
		t.Errorf("expected 300x8, got %v", dims.Size)
	}
}

func TestProgressControlled(t *testing.T) {
	th := theme.NewDark()
	p := progress.New(progress.Config{Value: 0.0})
	p.Value = 0.5
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 8))}
	dims := p.Layout(gtx, th)
	if dims.Size.X != 300 || dims.Size.Y != 8 {
		t.Errorf("expected 300x8, got %v", dims.Size)
	}
}

func TestProgressValueClamping(t *testing.T) {
	pLow := progress.New(progress.Config{Value: -0.5})
	if pLow.Value != 0.0 {
		t.Errorf("expected value clamped to 0.0, got %f", pLow.Value)
	}

	pHigh := progress.New(progress.Config{Value: 1.5})
	if pHigh.Value != 1.0 {
		t.Errorf("expected value clamped to 1.0, got %f", pHigh.Value)
	}
}

func TestProgressAnimationEasing(t *testing.T) {
	th := theme.NewDark()
	p := progress.New(progress.Config{Value: 0.0})

	// Layout at 0.0
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(200, 8))}
	p.Layout(gtx, th)

	// Update to 0.75
	p.Value = 0.75

	// Simulate frames
	for frame := 0; frame < 60; frame++ {
		gtx.Ops.Reset()
		p.Layout(gtx, th)
	}
}
