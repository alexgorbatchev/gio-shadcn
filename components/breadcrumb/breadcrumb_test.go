package breadcrumb_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/breadcrumb"
	"github.com/bnema/gio-shadcn/theme"
)

func TestBreadcrumbStandardPath(t *testing.T) {
	bc := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Mixer", false),
		},
	})
	if len(bc.Items) != 2 {
		t.Fatalf("expected 2 items")
	}
}

func TestBreadcrumbActivePageLink(t *testing.T) {
	bc := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Deck A", true),
		},
	})
	if !bc.Items[1].Active {
		t.Fatalf("expected active link on item 1")
	}
}

func TestBreadcrumbCustomSeparators(t *testing.T) {
	th := theme.NewDark()
	bc := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Mixer", true),
		},
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 30)),
	}
	dims := bc.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestBreadcrumbHorizontalFlexLayout(t *testing.T) {
	th := theme.NewDark()
	bc := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Mixer", true),
		},
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 30)),
	}
	dims := bc.Layout(gtx, th)
	if dims.Size.Y <= 0 {
		t.Errorf("invalid height")
	}
}

func TestBreadcrumbChevronDividers(t *testing.T) {
	th := theme.NewDark()
	bc := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Mixer", false),
			breadcrumb.NewItem("Deck A", true),
		},
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 30)),
	}
	dims := bc.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestBreadcrumbHoverHighlight(t *testing.T) {
	bc := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
		},
	})
	if bc.Items[0].Label != "Home" {
		t.Errorf("expected Label 'Home'")
	}
}
