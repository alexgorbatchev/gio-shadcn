package alert_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/alert"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAlertActionDemo(t *testing.T) {
	al := alert.New(alert.Config{
		Title:       "Dark mode is now available",
		Description: "Enable it under your profile settings to get started.",
		Action: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(60, 24)}
		},
	})
	if al.Action == nil {
		t.Errorf("expected Action widget to be non-nil")
	}
}

func TestAlertBasicDemo(t *testing.T) {
	al := alert.New(alert.Config{
		Icon:        lucide.CircleCheck,
		Title:       "Account updated successfully",
		Description: "Your profile information has been saved.",
	})
	if al.Icon == nil {
		t.Errorf("expected Icon to be non-nil")
	}
}

func TestAlertColorsDemo(t *testing.T) {
	al := alert.New(alert.Config{
		Icon:        lucide.TriangleAlert,
		Title:       "Your subscription will expire in 3 days.",
		Description: "Renew now to avoid service interruption.",
	})
	if al.Title == "" {
		t.Errorf("expected Title to be set")
	}
}

func TestAlertStandardDemo(t *testing.T) {
	al1 := alert.New(alert.Config{
		Icon:        lucide.CircleCheck,
		Title:       "Payment successful",
		Description: "Your payment of $29.99 has been processed.",
	})
	al2 := alert.New(alert.Config{
		Icon:        lucide.Info,
		Title:       "New feature available",
		Description: "We've added dark mode support.",
	})
	if al1.Title == "" || al2.Title == "" {
		t.Errorf("expected titles to be set")
	}
}

func TestAlertDestructiveDemo(t *testing.T) {
	al := alert.New(alert.Config{
		Variant:     theme.VariantDestructive,
		Icon:        lucide.CircleAlert,
		Title:       "Payment failed",
		Description: "Your payment could not be processed.",
	})
	if al.Variant != theme.VariantDestructive {
		t.Errorf("expected VariantDestructive, got %v", al.Variant)
	}
}

func TestAlertRTLDemo(t *testing.T) {
	al := alert.New(alert.Config{
		Icon:        lucide.CircleCheck,
		Title:       "تم الدفع بنجاح",
		Description: "تمت معالجة دفعتك البالغة 29.99 دولارًا.",
	})
	if al.Title == "" {
		t.Errorf("expected RTL title to be set")
	}
}

func TestAlertLayout(t *testing.T) {
	th := theme.NewDark()
	al := alert.New(alert.Config{
		Title:       "System Alert",
		Description: "Buffer rate 96kHz.",
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 100)),
	}
	dims := al.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions from Alert.Layout: %v", dims.Size)
	}
}
