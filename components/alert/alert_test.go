package alert_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/alert"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAlertDefaultVariant(t *testing.T) {
	al := alert.New(alert.Config{
		Title:       "Default Info Alert",
		Description: "Audio buffer set to 64 samples.",
		Variant:     theme.VariantDefault,
	})
	if al.Variant != theme.VariantDefault {
		t.Fatalf("expected VariantDefault")
	}
}

func TestAlertDestructiveVariant(t *testing.T) {
	al := alert.New(alert.Config{
		Title:       "Audio Clip Warning",
		Description: "Output signal clipped +1.2dB on Deck A.",
		Variant:     theme.VariantDestructive,
	})
	if al.Variant != theme.VariantDestructive {
		t.Fatalf("expected VariantDestructive")
	}
}

func TestAlertTitleHeader(t *testing.T) {
	al := alert.New(alert.Config{
		Title: "Header Title",
	})
	if al.Title != "Header Title" {
		t.Errorf("expected Title 'Header Title', got %s", al.Title)
	}
}

func TestAlertDescriptionBody(t *testing.T) {
	al := alert.New(alert.Config{
		Description: "Body description text",
	})
	if al.Description != "Body description text" {
		t.Errorf("expected Description 'Body description text', got %s", al.Description)
	}
}

func TestAlertVariantBackgroundStyling(t *testing.T) {
	th := theme.NewDark()
	al := alert.New(alert.Config{
		Title:   "Alert Title",
		Variant: theme.VariantDestructive,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	dims := al.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestAlertBorderStroke(t *testing.T) {
	th := theme.NewDark()
	al := alert.New(alert.Config{
		Title: "Alert Title",
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	dims := al.Layout(gtx, th)
	if dims.Size.Y <= 0 {
		t.Errorf("invalid height")
	}
}
