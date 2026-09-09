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
		t.Errorf("expected Open to be true after trigger click")
	}
}

func TestDropdownMenuCheckboxToggle(t *testing.T) {
	toggled := false
	chkItem := dropdownmenu.NewCheckboxItem("Status Bar", true, func(checked bool) {
		toggled = true
	})

	if !chkItem.Checked {
		t.Errorf("expected initially checked")
	}
	if !chkItem.IsCheck {
		t.Errorf("expected IsCheck to be true")
	}

	// Simulate selecting checkbox item
	if chkItem.OnSelect != nil {
		chkItem.OnSelect()
	}
	if chkItem.Checked {
		t.Errorf("expected Checked to toggle to false")
	}
	if !toggled {
		t.Errorf("expected onToggle callback to have been called")
	}
}

func TestDropdownMenuFloatingOverlayNonStretching(t *testing.T) {
	th := theme.NewDark()
	dm := dropdownmenu.New(dropdownmenu.Config{
		TriggerText: "Actions",
		Open:        false,
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("One", ""),
			dropdownmenu.NewItem("Two", ""),
			dropdownmenu.NewItem("Three", ""),
		},
	})

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Constraints{Max: image.Pt(500, 500)}}
	dimsClosed := dm.Layout(gtx, th)

	// Now open the menu
	dm.Open = true
	dimsOpen := dm.Layout(gtx, th)

	// Floating overlay menu should report the same trigger height so host flex layout does NOT stretch!
	if dimsClosed.Size.Y != dimsOpen.Size.Y {
		t.Errorf("floating overlay menu changed host layout height: closed %d != open %d", dimsClosed.Size.Y, dimsOpen.Size.Y)
	}
}
