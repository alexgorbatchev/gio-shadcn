package card_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCardLayout(t *testing.T) {
	th := theme.NewDark()

	c := card.New(card.Config{Variant: theme.VariantDefault})

	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(800, 600)),
	}

	dims := c.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: image.Pt(200, 100)}
	})

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("card returned invalid dimensions %v", dims.Size)
	}
}
