package selectcomp_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	selectcomp "github.com/bnema/gio-shadcn/components/select"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSelectCreation(t *testing.T) {
	sel := selectcomp.New(selectcomp.Config{
		Options: []*selectcomp.Item{
			selectcomp.NewItem("house", "Progressive House"),
			selectcomp.NewItem("techno", "Techno"),
		},
		SelectedValue: "house",
	})

	if sel.SelectedValue != "house" {
		t.Errorf("expected SelectedValue to be 'house', got %s", sel.SelectedValue)
	}
}

func TestSelectLayout(t *testing.T) {
	th := theme.NewDark()
	sel := selectcomp.New(selectcomp.Config{
		Options: []*selectcomp.Item{
			selectcomp.NewItem("house", "Progressive House"),
			selectcomp.NewItem("techno", "Techno"),
		},
		SelectedValue: "house",
		Open:          true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(200, 100)),
	}
	dims := sel.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Select.Layout")
	}
}
