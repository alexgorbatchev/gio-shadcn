package dropdownmenu_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/dropdownmenu"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDropdownMenuActionDropdownMenu(t *testing.T) {
	dm := dropdownmenu.New(dropdownmenu.Config{
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("Edit Track", "⌘E"),
			dropdownmenu.NewItem("Export FLAC", "⌘S"),
		},
		Open: true,
	})
	if len(dm.Items) != 2 {
		t.Fatalf("expected 2 items")
	}
}

func TestDropdownMenuTriggerButton(t *testing.T) {
	dm := dropdownmenu.New(dropdownmenu.Config{
		Open: false,
	})
	if dm.Open {
		t.Fatalf("expected Open false")
	}
}

func TestDropdownMenuItems(t *testing.T) {
	item := dropdownmenu.NewItem("Export FLAC", "⌘S")
	if item.Label != "Export FLAC" {
		t.Errorf("expected Label 'Export FLAC'")
	}
}

func TestDropdownMenuKeyboardShortcutBadges(t *testing.T) {
	item := dropdownmenu.NewItem("Edit Track", "⌘E")
	if item.Shortcut != "⌘E" {
		t.Errorf("expected Shortcut '⌘E'")
	}
}

func TestDropdownMenuOpenCloseState(t *testing.T) {
	th := theme.NewDark()
	dm := dropdownmenu.New(dropdownmenu.Config{
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("Edit Track", "⌘E"),
		},
		Open: true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(200, 100)),
	}
	dims := dm.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}
