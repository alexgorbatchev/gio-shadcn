package togglegroup_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/togglegroup"
	"github.com/bnema/gio-shadcn/theme"
)

func TestToggleGroupBasic(t *testing.T) {
	th := theme.NewDark()
	tg := togglegroup.New(togglegroup.Config{
		Items: []*togglegroup.Item{
			togglegroup.NewItem("bold", "B"),
			togglegroup.NewItem("italic", "I"),
		},
		SelectedKey: "bold",
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(200, 40))}
	dims := tg.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestToggleGroupSelection(t *testing.T) {
	th := theme.NewDark()
	selected := ""
	tg := togglegroup.New(togglegroup.Config{
		Items: []*togglegroup.Item{
			togglegroup.NewItem("a", "A"),
			togglegroup.NewItem("b", "B"),
		},
		SelectedKey: "a",
		OnChange: func(key string) {
			selected = key
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(200, 40))}
	_ = tg.Layout(gtx, th)
	_ = selected
}
