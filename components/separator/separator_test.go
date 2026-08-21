package separator_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/separator"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSeparatorCreation(t *testing.T) {
	sep := separator.New(separator.Config{
		Horizontal: true,
	})

	if !sep.Horizontal {
		t.Errorf("expected Horizontal to be true")
	}
}

func TestSeparatorLayout(t *testing.T) {
	th := theme.NewDark()
	sep := separator.New(separator.Config{
		Horizontal: true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(400, 20)),
	}
	dims := sep.Layout(gtx, th)

	if dims.Size.X != 400 {
		t.Errorf("expected width to be 400, got %d", dims.Size.X)
	}
}
