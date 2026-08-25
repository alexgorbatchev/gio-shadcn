package separator_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/separator"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSeparatorHorizontal(t *testing.T) {
	th := theme.NewDark()
	s := separator.New(separator.Config{Horizontal: true})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(200, 1))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestSeparatorVertical(t *testing.T) {
	th := theme.NewDark()
	s := separator.New(separator.Config{Horizontal: false})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(1, 20))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
