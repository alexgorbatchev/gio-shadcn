package tabs_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/tabs"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTabsBasic(t *testing.T) {
	th := theme.NewDark()
	tb := tabs.New(tabs.Config{
		Tabs: []*tabs.Tab{
			tabs.NewTab("account", "Account"),
			tabs.NewTab("password", "Password"),
		},
		ActiveKey: "account",
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 40))}
	dims := tb.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestTabsSelection(t *testing.T) {
	th := theme.NewDark()
	tb := tabs.New(tabs.Config{
		Tabs: []*tabs.Tab{
			tabs.NewTab("one", "One"),
			tabs.NewTab("two", "Two"),
		},
		ActiveKey: "one",
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 40))}
	dims := tb.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
