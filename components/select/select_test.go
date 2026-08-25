package selectcomp_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	selectcomp "github.com/bnema/gio-shadcn/components/select"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSelectBasic(t *testing.T) {
	th := theme.NewDark()
	s := selectcomp.New(selectcomp.Config{
		Options: []*selectcomp.Item{
			selectcomp.NewItem("apple", "Apple"),
			selectcomp.NewItem("banana", "Banana"),
		},
		SelectedValue: "apple",
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(200, 40))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestSelectOpen(t *testing.T) {
	th := theme.NewDark()
	s := selectcomp.New(selectcomp.Config{
		Options: []*selectcomp.Item{
			selectcomp.NewItem("apple", "Apple"),
			selectcomp.NewItem("banana", "Banana"),
		},
		SelectedValue: "apple",
		Open:          true,
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(200, 150))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
