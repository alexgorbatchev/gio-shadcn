package spinner_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/spinner"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSpinnerCreation(t *testing.T) {
	sp := spinner.New(spinner.Config{})

	if sp.Size <= 0 {
		t.Errorf("expected Size to be > 0, got %v", sp.Size)
	}
}

func TestSpinnerLayout(t *testing.T) {
	th := theme.NewDark()
	sp := spinner.New(spinner.Config{})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(50, 50)),
	}
	dims := sp.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Spinner.Layout")
	}
}
