package breadcrumb_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/breadcrumb"
	"github.com/bnema/gio-shadcn/theme"
)

func TestBreadcrumbCreation(t *testing.T) {
	bc := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Audio Engine", false),
			breadcrumb.NewItem("Mixer", true),
		},
	})

	if len(bc.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(bc.Items))
	}
}

func TestBreadcrumbLayout(t *testing.T) {
	th := theme.NewDark()
	bc := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Audio Engine", false),
			breadcrumb.NewItem("Mixer", true),
		},
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 30)),
	}
	dims := bc.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Breadcrumb.Layout")
	}
}
