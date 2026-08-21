package pagination_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/pagination"
	"github.com/bnema/gio-shadcn/theme"
)

func TestPaginationCreation(t *testing.T) {
	pg := pagination.New(pagination.Config{
		CurrentPage: 1,
		TotalPages:  5,
	})

	if pg.CurrentPage != 1 {
		t.Errorf("expected CurrentPage to be 1, got %d", pg.CurrentPage)
	}

	if pg.TotalPages != 5 {
		t.Errorf("expected TotalPages to be 5, got %d", pg.TotalPages)
	}
}

func TestPaginationLayout(t *testing.T) {
	th := theme.NewDark()
	pg := pagination.New(pagination.Config{
		CurrentPage: 1,
		TotalPages:  5,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(400, 40)),
	}
	dims := pg.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Pagination.Layout")
	}
}
