package button_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

func TestButtonVariantsAndSizes(t *testing.T) {
	th := theme.NewDark()

	btnPrimary := button.New(button.Config{Text: "Primary", Variant: theme.VariantDefault})
	btnSecondary := button.New(button.Config{Text: "Secondary", Variant: theme.VariantSecondary})
	btnOutline := button.New(button.Config{Text: "Outline", Variant: theme.VariantOutline})
	btnGhost := button.New(button.Config{Text: "Ghost", Variant: theme.VariantGhost})
	btnDestructive := button.New(button.Config{Text: "Destructive", Variant: theme.VariantDestructive})
	btnLink := button.New(button.Config{Text: "Link", Variant: theme.VariantLink})

	btnSM := button.New(button.Config{Text: "Small", Size: theme.SizeSM})
	btnDefaultSize := button.New(button.Config{Text: "Default Size", Size: theme.SizeDefault})
	btnLG := button.New(button.Config{Text: "Large Size", Size: theme.SizeLG})

	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(800, 600)),
	}

	buttons := []*button.Button{
		btnPrimary, btnSecondary, btnOutline, btnGhost, btnDestructive, btnLink,
		btnSM, btnDefaultSize, btnLG,
	}

	for _, btn := range buttons {
		dims := btn.Layout(gtx, th)
		if dims.Size.X <= 0 || dims.Size.Y <= 0 {
			t.Errorf("button %s returned invalid dimensions %v", btn.Text, dims.Size)
		}
	}
}
