package empty_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/empty"
	"github.com/bnema/gio-shadcn/theme"
)

func TestEmptyStandard(t *testing.T) {
	th := theme.NewDark()
	e := empty.New(empty.Config{
		Title:       "No results",
		Description: "Try searching again",
		Icon:        lucide.Search,
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 200))}
	dims := e.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestEmptyWithAction(t *testing.T) {
	th := theme.NewDark()
	e := empty.New(empty.Config{
		Title:       "No tracks",
		Description: "Add your first track",
		Action: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 30)}
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 200))}
	dims := e.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
