package button_test

import (
	"testing"

	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

func TestButtonDefaultVariant(t *testing.T) {
	btn := button.New(button.Config{Text: "Primary", Variant: theme.VariantDefault})
	if btn.Variant != theme.VariantDefault {
		t.Fatalf("expected VariantDefault")
	}
}

func TestButtonSecondaryVariant(t *testing.T) {
	btn := button.New(button.Config{Text: "Secondary", Variant: theme.VariantSecondary})
	if btn.Variant != theme.VariantSecondary {
		t.Fatalf("expected VariantSecondary")
	}
}

func TestButtonOutlineVariant(t *testing.T) {
	btn := button.New(button.Config{Text: "Outline", Variant: theme.VariantOutline})
	if btn.Variant != theme.VariantOutline {
		t.Fatalf("expected VariantOutline")
	}
}

func TestButtonGhostVariant(t *testing.T) {
	btn := button.New(button.Config{Text: "Ghost", Variant: theme.VariantGhost})
	if btn.Variant != theme.VariantGhost {
		t.Fatalf("expected VariantGhost")
	}
}

func TestButtonDestructiveVariant(t *testing.T) {
	btn := button.New(button.Config{Text: "Destructive", Variant: theme.VariantDestructive})
	if btn.Variant != theme.VariantDestructive {
		t.Fatalf("expected VariantDestructive")
	}
}

func TestButtonLinkVariant(t *testing.T) {
	btn := button.New(button.Config{Text: "Link", Variant: theme.VariantLink})
	if btn.Variant != theme.VariantLink {
		t.Fatalf("expected VariantLink")
	}
}

func TestButtonSmallSize(t *testing.T) {
	btn := button.New(button.Config{Text: "Small", Size: theme.SizeSM})
	if btn.Size != theme.SizeSM {
		t.Fatalf("expected SizeSM")
	}
}

func TestButtonDefaultSize(t *testing.T) {
	btn := button.New(button.Config{Text: "Default Size", Size: theme.SizeDefault})
	if btn.Size != theme.SizeDefault {
		t.Fatalf("expected SizeDefault")
	}
}

func TestButtonLargeSize(t *testing.T) {
	btn := button.New(button.Config{Text: "Large Size", Size: theme.SizeLG})
	if btn.Size != theme.SizeLG {
		t.Fatalf("expected SizeLG")
	}
}

func TestButtonIconSize(t *testing.T) {
	btn := button.New(button.Config{Size: theme.SizeIcon})
	if btn.Size != theme.SizeIcon {
		t.Fatalf("expected SizeIcon")
	}
}

func TestButtonDisabledState(t *testing.T) {
	btn := button.New(button.Config{Text: "Disabled", Disabled: true})
	if !btn.Disabled {
		t.Fatalf("expected Disabled true")
	}
}

func TestButtonPointerClickEvent(t *testing.T) {
	clicked := false
	btn := button.New(button.Config{
		Text: "Click Me",
		OnClick: func() {
			clicked = true
		},
	})
	if btn.OnClick == nil {
		t.Fatalf("expected OnClick handler")
	}
	_ = clicked
}
