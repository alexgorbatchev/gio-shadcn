package progress_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/progress"
	"github.com/bnema/gio-shadcn/theme"
)

func TestProgressCreation(t *testing.T) {
	p := progress.New(progress.Config{
		Value: 0.65,
	})

	if p.Value != 0.65 {
		t.Errorf("expected Value to be 0.65, got %f", p.Value)
	}
}

func TestProgressLayout(t *testing.T) {
	th := theme.NewDark()
	p := progress.New(progress.Config{
		Value: 0.80,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 20)),
	}
	dims := p.Layout(gtx, th)

	if dims.Size.X != 300 {
		t.Errorf("expected width to be 300, got %d", dims.Size.X)
	}
}
