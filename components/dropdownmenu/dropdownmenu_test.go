package dropdownmenu_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/dropdownmenu"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDropdownMenuBasic(t *testing.T) {
	th := theme.NewDark()
	dm := dropdownmenu.New(dropdownmenu.Config{
		Open: true,
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("Profile", "⇧⌘P"),
			dropdownmenu.NewItem("Settings", "⌘S"),
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(240, 100))}
	dims := dm.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestDropdownMenuClosed(t *testing.T) {
	th := theme.NewDark()
	dm := dropdownmenu.New(dropdownmenu.Config{
		Open: false,
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("Profile", ""),
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(240, 100))}
	dims := dm.Layout(gtx, th)
	if dims.Size.X != 0 || dims.Size.Y != 0 {
		t.Errorf("expected 0 dimensions for closed menu without trigger")
	}
}

func TestDropdownMenuTriggerToggle(t *testing.T) {
	th := theme.NewDark()
	dm := dropdownmenu.New(dropdownmenu.Config{
		TriggerText: "Open Menu",
		Open:        false,
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("Profile", "⇧⌘P"),
		},
	})

	if dm.Open {
		t.Fatalf("expected Open to be false initially")
	}

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 200))}
	dims := dm.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected trigger dimensions to be positive, got %v", dims.Size)
	}

	// Trigger button must exist and be toggleable
	if dm.TriggerButton == nil {
		t.Fatalf("expected TriggerButton to be initialized")
	}

	dm.TriggerButton.OnClick()
	if !dm.Open {
		t.Errorf("expected Open to become true after trigger button click")
	}
}

func TestDropdownMenuItemSelection(t *testing.T) {
	th := theme.NewDark()
	selected := -1
	dm := dropdownmenu.New(dropdownmenu.Config{
		Open: true,
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("First Item", "⌘1"),
			dropdownmenu.NewItem("Second Item", "⌘2"),
		},
		OnSelectItem: func(index int) {
			selected = index
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(240, 100))}
	_ = dm.Layout(gtx, th)
	_ = selected
}
