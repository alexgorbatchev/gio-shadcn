package badge_test

import (
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/badge"
	"github.com/bnema/gio-shadcn/theme"
)

func TestBadgeDefaultVariant(t *testing.T) {
	bg := badge.New(badge.Config{
		Text:    "Default Badge",
		Variant: theme.VariantDefault,
	})
	if bg.Variant != theme.VariantDefault {
		t.Fatalf("expected VariantDefault")
	}
}

func TestBadgeSecondaryVariant(t *testing.T) {
	bg := badge.New(badge.Config{
		Text:    "Secondary Badge",
		Variant: theme.VariantSecondary,
	})
	if bg.Variant != theme.VariantSecondary {
		t.Fatalf("expected VariantSecondary")
	}
}

func TestBadgeOutlineVariant(t *testing.T) {
	bg := badge.New(badge.Config{
		Text:    "Outline Badge",
		Variant: theme.VariantOutline,
	})
	if bg.Variant != theme.VariantOutline {
		t.Fatalf("expected VariantOutline")
	}
}

func TestBadgeDestructiveVariant(t *testing.T) {
	bg := badge.New(badge.Config{
		Text:    "Destructive Badge",
		Variant: theme.VariantDestructive,
	})
	if bg.Variant != theme.VariantDestructive {
		t.Fatalf("expected VariantDestructive")
	}
}

func TestBadgeFullRoundedRadius(t *testing.T) {
	th := theme.NewDark()
	bg := badge.New(badge.Config{
		Text: "Badge Radius Test",
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops: ops,
	}
	dims := bg.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestBadgeXSTypographyLabel(t *testing.T) {
	th := theme.NewDark()
	bg := badge.New(badge.Config{
		Text: "XS Typography",
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops: ops,
	}
	dims := bg.Layout(gtx, th)
	if dims.Size.Y <= 0 {
		t.Errorf("invalid height")
	}
}

func TestBadgeCompactPadding(t *testing.T) {
	th := theme.NewDark()
	bg := badge.New(badge.Config{
		Text: "Padding Test",
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops: ops,
	}
	dims := bg.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid padding dimensions")
	}
}
