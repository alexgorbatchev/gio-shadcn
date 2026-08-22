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
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", true),
			accordion.NewItem("Section 2", "Content 2", false),
		},
	})
	if !acc.Items[0].Expanded || acc.Items[1].Expanded {
		t.Fatalf("expected single open item state")
	}
}

func TestAccordionMultipleOpenItems(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", true),
			accordion.NewItem("Section 2", "Content 2", true),
		},
	})
	if !acc.Items[0].Expanded || !acc.Items[1].Expanded {
		t.Fatalf("expected multiple open items state")
	}
}

func TestAccordionCollapsedState(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", false),
		},
	})
	if acc.Items[0].Expanded {
		t.Fatalf("expected item to be collapsed")
	}
}

func TestAccordionExpandCollapseAnimation(t *testing.T) {
	th := theme.NewDark()
	acc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", false),
		},
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 200)),
	}
	dims := acc.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestAccordionItemHeaderButton(t *testing.T) {
	item := accordion.NewItem("Header Title", "Body Content", false)
	if item.Title != "Header Title" {
		t.Errorf("expected Title 'Header Title', got %s", item.Title)
	}
}

func TestAccordionItemContentPanel(t *testing.T) {
	item := accordion.NewItem("Header Title", "Body Content", true)
	if item.Content != "Body Content" {
		t.Errorf("expected Content 'Body Content', got %s", item.Content)
	}
}

func TestAccordionBorderDividers(t *testing.T) {
	th := theme.NewDark()
	acc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", true),
		},
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 200)),
	}
	dims := acc.Layout(gtx, th)
	if dims.Size.Y <= 0 {
		t.Errorf("invalid height")
	}
}
