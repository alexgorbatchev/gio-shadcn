package card_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCardOptionsAffectContentBounds(t *testing.T) {
	c := card.NewCard(card.WithCardVariant(theme.VariantDefault), card.WithCardClasses("p-2 bg-red-500 rounded-full"), card.WithCardPadding(layout.UniformInset(20)))
	calls := 0
	h := testui.Harness{Size: image.Pt(200, 100), Widget: func(gtx layout.Context) layout.Dimensions {
		s := c.Update(gtx)
		if s.IsActive() || s.IsHovered() || s.IsPressed() || s.IsDisabled() {
			t.Fatal("passive card reports interaction")
		}
		return c.Layout(gtx, nil, func(gtx layout.Context) layout.Dimensions {
			calls++
			if gtx.Constraints.Max != image.Pt(184, 84) {
				t.Errorf("class padding bounds=%v", gtx.Constraints.Max)
			}
			return layout.Dimensions{Size: image.Pt(30, 20)}
		})
	}}
	if d := h.Frame(); d.Size != image.Pt(46, 36) || calls != 1 {
		t.Fatalf("card size=%v calls=%d", d.Size, calls)
	}
}

func TestEmptyHeaderHasNoText(t *testing.T) {
	header := card.NewHeader(card.HeaderConfig{})
	h := testui.Harness{Size: image.Pt(200, 100), Widget: func(gtx layout.Context) layout.Dimensions { return header.Layout(gtx, nil) }}
	if h.Frame().Size.Y != 0 {
		t.Fatal("empty header created text")
	}
}
