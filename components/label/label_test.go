package label_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

func TestLabelHeading1(t *testing.T) {
	h1 := label.NewTypography("Heading 1", label.H1, "")
	if h1.Element != label.H1 {
		t.Fatalf("expected Element H1")
	}
}

func TestLabelHeading2(t *testing.T) {
	h2 := label.NewTypography("Heading 2", label.H2, "")
	if h2.Element != label.H2 {
		t.Fatalf("expected Element H2")
	}
}

func TestLabelHeading3(t *testing.T) {
	h3 := label.NewTypography("Heading 3", label.H3, "")
	if h3.Element != label.H3 {
		t.Fatalf("expected Element H3")
	}
}

func TestLabelHeading4(t *testing.T) {
	h4 := label.NewTypography("Heading 4", label.H4, "")
	if h4.Element != label.H4 {
		t.Fatalf("expected Element H4")
	}
}

func TestLabelBodyParagraph(t *testing.T) {
	p := label.NewTypography("Body text", label.P, "")
	if p.Element != label.P {
		t.Fatalf("expected Element P")
	}
}

func TestLabelMutedAndSmall(t *testing.T) {
	muted := label.NewTypography("Muted text", label.Muted, "")
	small := label.NewTypography("Small text", label.Small, "")
	if muted.Element != label.Muted || small.Element != label.Small {
		t.Fatalf("expected Muted and Small elements")
	}
}

func TestLabelTypographyFontScale(t *testing.T) {
	th := theme.NewDark()
	h1 := label.NewTypography("Heading 1", label.H1, "")
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 100)),
	}
	dims := h1.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestLabelThemeColorIntegration(t *testing.T) {
	th := theme.NewDark()
	lbl := label.NewLabel(label.WithLabelText("Test Label"))
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 50)),
	}
	dims := lbl.Layout(gtx, th)
	if dims.Size.Y <= 0 {
		t.Errorf("invalid height")
	}
}
