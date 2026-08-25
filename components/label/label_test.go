package label_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

func TestLabelH1(t *testing.T) {
	th := theme.NewDark()
	h1 := label.NewTypography("Heading 1", label.H1, "")
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 50))}
	dims := h1.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestLabelH2(t *testing.T) {
	th := theme.NewDark()
	h2 := label.NewTypography("Heading 2", label.H2, "")
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 40))}
	dims := h2.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestLabelH3(t *testing.T) {
	th := theme.NewDark()
	h3 := label.NewTypography("Heading 3", label.H3, "")
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 35))}
	dims := h3.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestLabelH4(t *testing.T) {
	th := theme.NewDark()
	h4 := label.NewTypography("Heading 4", label.H4, "")
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 30))}
	dims := h4.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestLabelParagraph(t *testing.T) {
	th := theme.NewDark()
	p := label.NewTypography("Body paragraph text", label.P, "")
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 30))}
	dims := p.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
