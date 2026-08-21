package dropdownmenu_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/dropdownmenu"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDropdownMenuCreation(t *testing.T) {
	dm := dropdownmenu.New(dropdownmenu.Config{
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("Edit Track", "⌘E"),
			dropdownmenu.NewItem("Export FLAC", "⌘S"),
		},
		Open: true,
	})

	if len(dm.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(dm.Items))
	}
}

func TestDropdownMenuLayout(t *testing.T) {
	th := theme.NewDark()
	dm := dropdownmenu.New(dropdownmenu.Config{
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("Edit Track", "⌘E"),
			dropdownmenu.NewItem("Export FLAC", "⌘S"),
		},
		Open: true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(200, 100)),
	}
	dims := dm.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from DropdownMenu.Layout")
	}
}
