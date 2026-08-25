package accordion_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/accordion"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAccordionBasic(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Type: accordion.TypeSingle,
		Items: []*accordion.Item{
			accordion.NewItem("How do I reset my password?", "Click on 'Forgot Password' on the login page.", true),
			accordion.NewItem("Can I change my subscription plan?", "Yes, you can upgrade or downgrade.", false),
			accordion.NewItem("What payment methods do you accept?", "We accept credit cards, PayPal, bank transfers.", false),
		},
	})
	if !acc.Items[0].Expanded || acc.Items[1].Expanded {
		t.Errorf("expected item 0 expanded and item 1 collapsed")
	}
}

func TestAccordionBorders(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Type: accordion.TypeSingle,
		Items: []*accordion.Item{
			accordion.NewItem("How does billing work?", "We offer monthly and annual subscription plans.", true),
			accordion.NewItem("Is my data secure?", "Yes. We use end-to-end encryption.", false),
		},
	})
	if len(acc.Items) != 2 {
		t.Errorf("expected 2 items, got %d", len(acc.Items))
	}
}

func TestAccordionCard(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Type: accordion.TypeSingle,
		Items: []*accordion.Item{
			accordion.NewItem("What subscription plans do you offer?", "Starter ($9/mo), Pro ($29/mo), Enterprise ($99/mo).", true),
			accordion.NewItem("How do I cancel my subscription?", "You can cancel anytime from account settings.", false),
		},
	})
	if !acc.Items[0].Expanded {
		t.Errorf("expected first item expanded")
	}
}

func TestAccordionStandardDemo(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Type: accordion.TypeSingle,
		Items: []*accordion.Item{
			accordion.NewItem("What are your shipping options?", "Standard, express, overnight shipping.", true),
			accordion.NewItem("What is your return policy?", "Returns accepted within 30 days.", false),
		},
	})
	if !acc.Items[0].Expanded {
		t.Errorf("expected shipping item expanded")
	}
}

func TestAccordionDisabledDemo(t *testing.T) {
	item := accordion.NewItemConfig(accordion.ItemConfig{
		Title:    "Premium feature information",
		Content:  "This section contains information about premium features.",
		Expanded: false,
		Disabled: true,
	})
	if !item.Disabled {
		t.Errorf("expected item to be disabled")
	}
}

func TestAccordionMultipleDemo(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Type: accordion.TypeMultiple,
		Items: []*accordion.Item{
			accordion.NewItem("Notification Settings", "Manage how you receive notifications.", true),
			accordion.NewItem("Privacy & Security", "Control privacy settings.", true),
		},
	})
	if !acc.Items[0].Expanded || !acc.Items[1].Expanded {
		t.Errorf("expected both items expanded in multiple mode")
	}
}

func TestAccordionRTLDemo(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Type: accordion.TypeSingle,
		Items: []*accordion.Item{
			accordion.NewItem("كيف يمكنني إعادة تعيين كلمة المرور؟", "انقر على 'نسيت كلمة المرور'.", true),
		},
	})
	if len(acc.Items) != 1 || !acc.Items[0].Expanded {
		t.Errorf("expected 1 RTL item expanded")
	}
}

func TestAccordionLayout(t *testing.T) {
	th := theme.NewDark()
	acc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", true),
		},
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 300)),
	}
	dims := acc.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Accordion.Layout")
	}
}
