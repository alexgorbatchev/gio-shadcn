package titlebar_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/titlebar"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTitleBarLayout(t *testing.T) {
	th := theme.NewDark()

	tb := titlebar.NewTitleBar(
		titlebar.WithTitle("Test TitleBar"),
		titlebar.WithVariant(theme.VariantSecondary),
	)

	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(800, 36)),
	}

	dims := tb.Layout(gtx, th, nil)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("titlebar returned invalid dimensions %v", dims.Size)
	}
}
