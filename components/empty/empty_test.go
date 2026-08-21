package empty_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/empty"
	"github.com/bnema/gio-shadcn/theme"
)

func TestEmptyCreation(t *testing.T) {
	emp := empty.New(empty.Config{})

	if emp.Title != "No Results Found" {
		t.Errorf("expected default Title, got %s", emp.Title)
	}
}

func TestEmptyLayout(t *testing.T) {
	th := theme.NewDark()
	emp := empty.New(empty.Config{})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 150)),
	}
	dims := emp.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Empty.Layout")
	}
}
