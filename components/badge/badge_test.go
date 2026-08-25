package badge_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/badge"
	"github.com/bnema/gio-shadcn/theme"
)

func TestBadgeDefaultVariant(t *testing.T) {
	th := theme.NewDark()
	b := badge.New(badge.Config{Text: "Badge", Variant: theme.VariantDefault})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBadgeSecondaryVariant(t *testing.T) {
	th := theme.NewDark()
	b := badge.New(badge.Config{Text: "Secondary", Variant: theme.VariantSecondary})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBadgeOutlineVariant(t *testing.T) {
	th := theme.NewDark()
	b := badge.New(badge.Config{Text: "Outline", Variant: theme.VariantOutline})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBadgeDestructiveVariant(t *testing.T) {
	th := theme.NewDark()
	b := badge.New(badge.Config{Text: "Destructive", Variant: theme.VariantDestructive})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBadgeWithIcon(t *testing.T) {
	th := theme.NewDark()
	b := badge.New(badge.Config{Text: "Verified", Variant: theme.VariantSecondary, Icon: lucide.Check})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(120, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBadgeWithSpinner(t *testing.T) {
	th := theme.NewDark()
	b := badge.New(badge.Config{Text: "Updating", Variant: theme.VariantOutline, Icon: lucide.LoaderCircle})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(120, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBadgeLink(t *testing.T) {
	th := theme.NewDark()
	b := badge.New(badge.Config{Text: "Link ↗", Variant: theme.VariantDefault})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(120, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBadgeColors(t *testing.T) {
	th := theme.NewDark()
	b := badge.New(badge.Config{Text: "Custom Color", Variant: theme.VariantSecondary, Classes: "bg-blue-500"})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(120, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
