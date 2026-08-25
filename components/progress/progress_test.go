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
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 10))}
	dims := p.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestProgressControlled(t *testing.T) {
	th := theme.NewDark()
	p := progress.New(progress.Config{Value: 0.0})
	p.Value = 0.5
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 10))}
	dims := p.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
