package command_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/command"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCommandBasic(t *testing.T) {
	cmd := command.New(command.Config{
		Placeholder: "Type a command...",
		Items: []*command.Item{
			command.NewItem("Calendar", ""),
			command.NewItem("Calculator", ""),
		},
	})
	if len(cmd.Items) != 2 {
		t.Errorf("expected 2 items, got %d", len(cmd.Items))
	}
}

func TestCommandDemoWithIcons(t *testing.T) {
	cmd := command.New(command.Config{
		Items: []*command.Item{
			command.NewItemFull("Profile", "⌘P", "Settings", lucide.User, false),
			command.NewItemFull("Billing", "⌘B", "Settings", lucide.CreditCard, false),
		},
	})
	if cmd.Items[0].Icon == nil {
		t.Errorf("expected non-nil icon")
	}
}

func TestCommandGroupsAndShortcuts(t *testing.T) {
	cmd := command.New(command.Config{
		Items: []*command.Item{
			command.NewItemFull("Search Emoji", "", "Suggestions", lucide.FaceSlightlySmiling, false),
			command.NewItemFull("Settings", "⌘S", "Settings", lucide.Settings, false),
		},
	})
	if cmd.Items[0].Group != "Suggestions" || cmd.Items[1].Group != "Settings" {
		t.Errorf("unexpected groups")
	}
}

func TestCommandLayout(t *testing.T) {
	th := theme.NewDark()
	cmd := command.New(command.Config{
		Items: []*command.Item{
			command.NewItem("Item 1", "⌘1"),
		},
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 300)),
	}
	dims := cmd.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Command.Layout")
	}
}
