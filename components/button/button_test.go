package button_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

func TestButtonDefault(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Button", Variant: theme.VariantDefault})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 40))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonSecondary(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Secondary", Variant: theme.VariantSecondary})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 40))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonOutline(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Outline", Variant: theme.VariantOutline})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 40))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonGhost(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Ghost", Variant: theme.VariantGhost})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 40))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonDestructive(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Destructive", Variant: theme.VariantDestructive})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 40))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonLink(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Link", Variant: theme.VariantLink})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 40))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonWithIcon(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Login with Email", Variant: theme.VariantDefault, Icon: lucide.Mail})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(160, 40))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonIconOnly(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Variant: theme.VariantOutline, Size: theme.SizeIcon, Icon: lucide.ChevronRight})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(40, 40))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonSizes(t *testing.T) {
	th := theme.NewDark()
	btnSM := button.New(button.Config{Text: "Small", Size: theme.SizeSM})
	btnLG := button.New(button.Config{Text: "Large", Size: theme.SizeLG})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 40))}
	_ = btnSM.Layout(gtx, th)
	_ = btnLG.Layout(gtx, th)
}
