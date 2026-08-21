package togglegroup_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/togglegroup"
	"github.com/bnema/gio-shadcn/theme"
)

func TestToggleGroupCreation(t *testing.T) {
	tg := togglegroup.New(togglegroup.Config{
		Items: []*togglegroup.Item{
			togglegroup.NewItem("grid", "Grid"),
			togglegroup.NewItem("list", "List"),
		},
		SelectedKey: "grid",
	})

	if len(tg.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(tg.Items))
	}

	if tg.SelectedKey != "grid" {
		t.Errorf("expected SelectedKey to be 'grid', got %s", tg.SelectedKey)
	}
}

func TestToggleGroupLayout(t *testing.T) {
	th := theme.NewDark()
	tg := togglegroup.New(togglegroup.Config{
		Items: []*togglegroup.Item{
			togglegroup.NewItem("grid", "Grid"),
			togglegroup.NewItem("list", "List"),
		},
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(200, 40)),
	}
	dims := tg.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from ToggleGroup.Layout")
	}
}
