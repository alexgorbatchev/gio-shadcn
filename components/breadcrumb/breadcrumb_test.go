package breadcrumb_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/breadcrumb"
	"github.com/bnema/gio-shadcn/theme"
)

func TestBreadcrumbStandard(t *testing.T) {
	th := theme.NewDark()
	b := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Components", false),
			breadcrumb.NewItem("Breadcrumb", true),
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(400, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBreadcrumbCustomSeparator(t *testing.T) {
	th := theme.NewDark()
	b := breadcrumb.New(breadcrumb.Config{
		Separator: "/",
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Docs", true),
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBreadcrumbEllipsis(t *testing.T) {
	th := theme.NewDark()
	b := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("...", false),
			breadcrumb.NewItem("Components", true),
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 30))}
	dims := b.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestBreadcrumbInteractiveClick(t *testing.T) {
	th := theme.NewDark()
	clicked := -1
	b := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Page", true),
		},
		OnSelect: func(index int) {
			clicked = index
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 30))}
	_ = b.Layout(gtx, th)
	_ = clicked
}
