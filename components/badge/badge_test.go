package badge_test

import (
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/badge"
	"github.com/bnema/gio-shadcn/theme"
)

func TestBadgeCreation(t *testing.T) {
	bg := badge.New(badge.Config{
		Text:    "Test Badge",
		Variant: theme.VariantSecondary,
	})

	if bg.Text != "Test Badge" {
		t.Errorf("expected Text to be 'Test Badge', got %s", bg.Text)
	}

	if bg.Variant != theme.VariantSecondary {
		t.Errorf("expected Variant to be secondary, got %s", bg.Variant)
	}
}

func TestBadgeLayout(t *testing.T) {
	th := theme.NewDark()
	bg := badge.New(badge.Config{
		Text: "Engine Online",
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
	}
	dims := bg.Layout(gtx, th)

	if dims.Size.X < 0 || dims.Size.Y < 0 {
		t.Errorf("invalid dimensions returned from Badge.Layout")
	}
}
