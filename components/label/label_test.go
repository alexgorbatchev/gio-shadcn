package label_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

func TestLabelTypographyLayout(t *testing.T) {
	th := theme.NewDark()

	h1 := label.NewTypography("Heading 1", label.H1, "")
	p := label.NewTypography("Body text", label.P, "")

	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 100)),
	}

	dimsH1 := h1.Layout(gtx, th)
	if dimsH1.Size.X <= 0 || dimsH1.Size.Y <= 0 {
		t.Errorf("H1 returned invalid dimensions %v", dimsH1.Size)
	}

	dimsP := p.Layout(gtx, th)
	if dimsP.Size.X <= 0 || dimsP.Size.Y <= 0 {
		t.Errorf("P returned invalid dimensions %v", dimsP.Size)
	}
}
