package dropdownmenu_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/dropdownmenu"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDisabledMenuItemCannotSelect(t *testing.T) {
	th := theme.New()
	calls := 0
	item := dropdownmenu.NewItem("Disabled", "")
	item.Disabled = true
	item.OnSelect = func() { calls++ }
	m := dropdownmenu.New(dropdownmenu.Config{Open: true, Items: []*dropdownmenu.Item{item}})
	h := testui.Harness{Size: image.Pt(200, 200), Widget: func(gtx layout.Context) layout.Dimensions { return m.Layout(gtx, th) }}
	h.Frame()
	h.Click(15, 15)
	if calls != 0 {
		t.Fatalf("disabled menu item's OnSelect ran %d times", calls)
	}
}
