package accordion_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/accordion"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAccordionSingleOpenItem(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Type: accordion.TypeSingle,
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", true),
			accordion.NewItem("Section 2", "Content 2", false),
		},
	})
	if !acc.Items[0].Expanded || acc.Items[1].Expanded {
		t.Errorf("expected item 0 expanded and item 1 collapsed")
	}
}

func TestAccordionMultipleOpenItems(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Type: accordion.TypeMultiple,
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", true),
			accordion.NewItem("Section 2", "Content 2", true),
		},
	})
	if !acc.Items[0].Expanded || !acc.Items[1].Expanded {
		t.Errorf("expected both items expanded in multiple mode")
	}
}

func TestAccordionDisabledState(t *testing.T) {
	item := accordion.NewItemConfig(accordion.ItemConfig{
		Title:    "Disabled Item",
		Content:  "Content",
		Expanded: false,
		Disabled: true,
	})
	if !item.Disabled {
		t.Errorf("expected item to be disabled")
	}
}

func TestAccordionCustomHeaderAndIcon(t *testing.T) {
	item := accordion.NewItemConfig(accordion.ItemConfig{
		Title:    "Header",
		Content:  "Body",
		Icon:     "v",
		Expanded: true,
	})
	if item.Icon != "v" {
		t.Errorf("expected icon 'v', got %s", item.Icon)
	}
}

func TestAccordionBorderlessVariant(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Borderless: true,
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", true),
		},
	})
	if !acc.Borderless {
		t.Errorf("expected Borderless to be true")
	}
}

func TestAccordionNestedContentWidget(t *testing.T) {
	innerAcc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Inner Section", "Inner Content", true),
		},
	})
	outerItem := accordion.NewItemConfig(accordion.ItemConfig{
		Title:    "Outer Section",
		Expanded: true,
		ContentWidget: func(gtx layout.Context) layout.Dimensions {
			return innerAcc.Layout(gtx, theme.NewDark())
		},
	})
	if outerItem.ContentWidget == nil {
		t.Errorf("expected ContentWidget to be non-nil")
	}
}

func TestAccordionLayoutDimensions(t *testing.T) {
	th := theme.NewDark()
	acc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", true),
			accordion.NewItem("Section 2", "Content 2", false),
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
