package collapsible_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/collapsible"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCollapsibleCreation(t *testing.T) {
	col := collapsible.New(collapsible.Config{
		Title:   "Advanced Audio Routing",
		Content: "ASIO multi-channel mapping enabled.",
		Open:    false,
	})

	if col.Title != "Advanced Audio Routing" {
		t.Errorf("expected Title to be 'Advanced Audio Routing', got %s", col.Title)
	}
}

func TestCollapsibleLayout(t *testing.T) {
	th := theme.NewDark()
	col := collapsible.New(collapsible.Config{
		Title:   "Advanced Audio Routing",
		Content: "ASIO multi-channel mapping enabled.",
		Open:    true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	dims := col.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Collapsible.Layout")
	}
}
