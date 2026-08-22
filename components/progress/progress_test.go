package progress_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/progress"
	"github.com/bnema/gio-shadcn/theme"
)

func TestProgressZeroPercent(t *testing.T) {
	p := progress.New(progress.Config{Value: 0.0})
	if p.Value != 0.0 {
		t.Errorf("expected value 0.0")
	}
}

func TestProgressSixtyFivePercent(t *testing.T) {
	p := progress.New(progress.Config{Value: 0.65})
	if p.Value != 0.65 {
		t.Errorf("expected value 0.65")
	}
}

func TestProgressHundredPercent(t *testing.T) {
	p := progress.New(progress.Config{Value: 1.0})
	if p.Value != 1.0 {
		t.Errorf("expected value 1.0")
	}
}

func TestProgressClampedRadiusHalfHeight(t *testing.T) {
	th := theme.NewDark()
	p := progress.New(progress.Config{Value: 0.5})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(200, 8)),
	}
	dims := p.Layout(gtx, th)
	if dims.Size.Y != 8 {
		t.Errorf("expected height 8")
	}
}

func TestProgressLayoutDimensions(t *testing.T) {
	th := theme.NewDark()
	p := progress.New(progress.Config{Value: 0.8})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 10)),
	}
	dims := p.Layout(gtx, th)
	if dims.Size.X != 300 {
		t.Errorf("expected width 300, got %d", dims.Size.X)
	}
}
