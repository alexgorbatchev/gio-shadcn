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

func TestSpinnerSmallSize(t *testing.T) {
	sp := spinner.New(spinner.Config{
		Size: unit.Dp(16),
	})

	if sp.Size != unit.Dp(16) {
		t.Errorf("expected Size 16dp, got %v", sp.Size)
	}
}

func TestSpinnerDefaultSize(t *testing.T) {
	sp := spinner.New(spinner.Config{})

	if sp.Size != unit.Dp(24) {
		t.Errorf("expected default Size 24dp, got %v", sp.Size)
	}
}

func TestSpinnerLargeSize(t *testing.T) {
	sp := spinner.New(spinner.Config{
		Size: unit.Dp(48),
	})

	if sp.Size != unit.Dp(48) {
		t.Errorf("expected Size 48dp, got %v", sp.Size)
	}
}

func TestSpinnerArcStrokeLayout(t *testing.T) {
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
