package collapsible_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/collapsible"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCollapsibleBasic(t *testing.T) {
	c := collapsible.New(collapsible.Config{
		Title:   "Product details",
		Content: "This panel can be expanded or collapsed.",
		Open:    true,
	})
	if !c.Open {
		t.Errorf("expected Open to be true")
	}
}

func TestCollapsibleDemo(t *testing.T) {
	c := collapsible.New(collapsible.Config{
		Title:   "Order #4189",
		Content: "Shipping address: 100 Market St",
		Open:    true,
	})
	if c.Title != "Order #4189" {
		t.Errorf("unexpected title: %s", c.Title)
	}
}

func TestCollapsibleFileTree(t *testing.T) {
	c := collapsible.New(collapsible.Config{
		Title: "components/ui",
		ContentWidget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 50)}
		},
		Open: true,
	})
	if c.ContentWidget == nil {
		t.Errorf("expected non-nil ContentWidget")
	}
}

func TestCollapsibleSettings(t *testing.T) {
	c := collapsible.New(collapsible.Config{
		Title:   "Radius Settings",
		Content: "Radius X: 8px",
		Open:    false,
	})
	if c.Open {
		t.Errorf("expected Open to be false initially")
	}
}

func TestCollapsibleLayout(t *testing.T) {
	th := theme.NewDark()
	c := collapsible.New(collapsible.Config{
		Title:   "Test Collapsible",
		Content: "Content body",
		Open:    true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 200)),
	}
	dims := c.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Collapsible.Layout")
	}
}
