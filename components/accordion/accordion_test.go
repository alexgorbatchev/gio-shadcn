package accordion_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/accordion"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAccordionCreation(t *testing.T) {
	acc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", false),
			accordion.NewItem("Section 2", "Content 2", true),
		},
	})

	if len(acc.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(acc.Items))
	}

	if acc.Items[0].Title != "Section 1" {
		t.Errorf("expected item 0 title to be 'Section 1', got %s", acc.Items[0].Title)
	}
}

func TestAccordionLayout(t *testing.T) {
	th := theme.NewDark()
	acc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Section 1", "Content 1", false),
			accordion.NewItem("Section 2", "Content 2", true),
		},
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 200)),
	}
	dims := acc.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Accordion.Layout")
	}
}
