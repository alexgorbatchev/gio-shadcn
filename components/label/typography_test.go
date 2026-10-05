package label_test

import (
	"image"
	"testing"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTypographyElementsAndOverrides(t *testing.T) {
	th := theme.New()
	for _, element := range []label.TypographyElement{label.H1, label.H2, label.H3, label.H4, label.P, label.Small, label.Lead, label.Large, label.Muted, "unknown"} {
		typ := label.NewTypography("Hello", element, "")
		h := testui.Harness{Size: image.Pt(400, 200), Widget: func(gtx layout.Context) layout.Dimensions { return typ.Layout(gtx, th) }}
		original := h.Frame()
		if original.Size.X <= 0 || original.Size.Y <= 0 {
			t.Fatalf("empty %q text", element)
		}
		typ.SetText("Longer text")
		typ.SetElement(label.H2)
		typ.SetTextStyle(theme.TextStyle{Size: unit.Sp(48), Weight: font.Bold, Alignment: text.Middle, Color: &th.DarkColors})
		typ.Classes = "bg-red-500"
		if d := h.Frame(); d.Size.Y < 48 {
			t.Fatalf("override ignored for %q: %v", element, d.Size)
		}
		th.MaterialTheme = nil
		h.Frame()
		h.Widget = func(gtx layout.Context) layout.Dimensions { return typ.Layout(gtx, nil) }
		h.Frame()
	}
}

func TestLabelOptionsAndPassiveState(t *testing.T) {
	th := theme.New()
	for _, size := range []theme.Size{theme.SizeSM, theme.SizeDefault, theme.SizeLG} {
		l := label.NewLabel(label.WithLabelText("Label"), label.WithLabelClasses("bg-blue-500"), label.WithLabelVariant(theme.VariantSecondary), label.WithLabelSize(size), label.WithTextStyle(theme.TextStyle{Weight: font.Bold, Color: &th.DarkColors}))
		h := testui.Harness{Size: image.Pt(400, 200), Widget: func(gtx layout.Context) layout.Dimensions {
			s := l.Update(gtx)
			if s.IsActive() || s.IsHovered() || s.IsPressed() || s.IsDisabled() {
				t.Fatal("passive label reports interaction")
			}
			return l.Layout(gtx, th)
		}}
		if h.Frame().Size.Y <= 0 {
			t.Fatal("label not rendered")
		}
		l.SetText("Other")
		l.SetTextStyle(theme.TextStyle{Size: 40})
		l.Variant = "unknown"
		th.MaterialTheme = nil
		if h.Frame().Size.Y < 40 {
			t.Fatal("label setter ignored")
		}
	}
}
