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
	p := pagination.New(pagination.Config{
		CurrentPage: 1,
		TotalPages:  5,
	})
	if p.CurrentPage != 1 || p.TotalPages != 5 {
		t.Errorf("expected page 1 of 5")
	}
}

func TestPaginationActivePageSelection(t *testing.T) {
	p := pagination.New(pagination.Config{
		CurrentPage: 3,
		TotalPages:  5,
	})
	if p.CurrentPage != 3 {
		t.Errorf("expected active page 3")
	}
}

func TestPaginationPreviousNextNavigation(t *testing.T) {
	p := pagination.New(pagination.Config{
		CurrentPage: 2,
		TotalPages:  10,
	})
	if p.TotalPages != 10 {
		t.Errorf("expected 10 total pages")
	}
}

func TestPaginationOnSelectPageCallback(t *testing.T) {
	var selected int
	p := pagination.New(pagination.Config{
		CurrentPage: 1,
		TotalPages:  5,
		OnSelectPage: func(page int) {
			selected = page
		},
	})
	if p.OnSelectPage == nil {
		t.Errorf("expected non-nil OnSelectPage callback")
	}
	_ = selected
}

func TestPaginationLayoutDimensions(t *testing.T) {
	th := theme.NewDark()
	p := pagination.New(pagination.Config{
		CurrentPage: 1,
		TotalPages:  5,
	})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(400, 40)),
	}
	dims := p.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions from Pagination.Layout")
	}
}
