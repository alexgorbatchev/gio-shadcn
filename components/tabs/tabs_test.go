package tabs_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/tabs"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTabsCreation(t *testing.T) {
	tb := tabs.New(tabs.Config{
		Tabs: []*tabs.Tab{
			tabs.NewTab("sink", "Kitchen Sink"),
			tabs.NewTab("deck", "Audio Deck"),
		},
		ActiveKey: "sink",
	})

	if len(tb.Tabs) != 2 {
		t.Fatalf("expected 2 tabs, got %d", len(tb.Tabs))
	}

	if tb.ActiveKey != "sink" {
		t.Errorf("expected ActiveKey to be 'sink', got %s", tb.ActiveKey)
	}
}

func TestTabsLayout(t *testing.T) {
	th := theme.NewDark()
	tb := tabs.New(tabs.Config{
		Tabs: []*tabs.Tab{
			tabs.NewTab("sink", "Kitchen Sink"),
			tabs.NewTab("deck", "Audio Deck"),
		},
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 50)),
	}
	dims := tb.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Tabs.Layout")
	}
}
