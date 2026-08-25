package card_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCardStandard(t *testing.T) {
	th := theme.NewDark()
	c := card.New(card.Config{Variant: theme.VariantDefault})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 200))}
	dims := c.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: image.Pt(200, 100)}
	})
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestCardSmall(t *testing.T) {
	th := theme.NewDark()
	c := card.New(card.Config{Variant: theme.VariantDefault, Padding: layout.UniformInset(unit.Dp(16))})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 150))}
	dims := c.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: image.Pt(150, 80)}
	})
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestCardSpacing(t *testing.T) {
	th := theme.NewDark()
	c := card.New(card.Config{Variant: theme.VariantDefault, Padding: layout.UniformInset(unit.Dp(32))})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(400, 300))}
	dims := c.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: image.Pt(250, 150)}
	})
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
