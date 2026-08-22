package collapsible_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/collapsible"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCollapsibleExpandedState(t *testing.T) {
	col := collapsible.New(collapsible.Config{
		Title:   "Expanded Section",
		Content: "Expanded body content",
		Open:    true,
	})
	if !col.Open {
		t.Fatalf("expected Open true")
	}
}

func TestCollapsibleCollapsedState(t *testing.T) {
	col := collapsible.New(collapsible.Config{
		Title:   "Collapsed Section",
		Content: "Collapsed body content",
		Open:    false,
	})
	if col.Open {
		t.Fatalf("expected Open false")
	}
}

func TestCollapsibleTriggerButtonHeader(t *testing.T) {
	col := collapsible.New(collapsible.Config{
		Title: "Trigger Button Header",
	})
	if col.Title != "Trigger Button Header" {
		t.Errorf("expected Title 'Trigger Button Header', got %s", col.Title)
	}
}

func TestCollapsibleContentBodyVisibilityToggle(t *testing.T) {
	th := theme.NewDark()
	col := collapsible.New(collapsible.Config{
		Title:   "Title",
		Content: "Body Content",
		Open:    true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	dims := col.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}
