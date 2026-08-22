package command_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/command"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCommandPaletteSearchBox(t *testing.T) {
	cmd := command.New(command.Config{
		Placeholder: "Search actions...",
		Items: []*command.Item{
			command.NewItem("Toggle Dark Mode", "⌘T"),
		},
	})
	if cmd.Placeholder != "Search actions..." {
		t.Fatalf("expected Placeholder 'Search actions...'")
	}
}

func TestCommandSearchInputFilter(t *testing.T) {
	th := theme.NewDark()
	cmd := command.New(command.Config{
		Items: []*command.Item{
			command.NewItem("Toggle Dark Mode", "⌘T"),
			command.NewItem("Reset Audio Engine", "⌘R"),
		},
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 150)),
	}
	dims := cmd.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestCommandItemList(t *testing.T) {
	cmd := command.New(command.Config{
		Items: []*command.Item{
			command.NewItem("Item 1", "⌘1"),
			command.NewItem("Item 2", "⌘2"),
		},
	})
	if len(cmd.Items) != 2 {
		t.Errorf("expected 2 items")
	}
}

func TestCommandShortcutBadges(t *testing.T) {
	item := command.NewItem("Toggle Theme", "⌘T")
	if item.Shortcut != "⌘T" {
		t.Errorf("expected Shortcut '⌘T'")
	}
}
