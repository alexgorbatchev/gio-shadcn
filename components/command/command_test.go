package command_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/command"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCommandCreation(t *testing.T) {
	cmd := command.New(command.Config{
		Items: []*command.Item{
			command.NewItem("Toggle Dark Mode", "⌘T"),
			command.NewItem("Reset Audio Engine", "⌘R"),
		},
	})

	if len(cmd.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(cmd.Items))
	}
}

func TestCommandLayout(t *testing.T) {
	th := theme.NewDark()
	cmd := command.New(command.Config{
		Items: []*command.Item{
			command.NewItem("Toggle Dark Mode", "⌘T"),
			command.NewItem("Reset Audio Engine", "⌘R"),
		},
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 150)),
	}
	dims := cmd.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Command.Layout")
	}
}
