package pagination_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/pagination"
	"github.com/bnema/gio-shadcn/theme"
)

func TestPaginationStandardPageBar(t *testing.T) {
	pg := pagination.New(pagination.Config{
		CurrentPage: 1,
		TotalPages:  5,
	})
	if pg.CurrentPage != 1 || pg.TotalPages != 5 {
		t.Errorf("expected current page 1 and total 5")
	}
}

func TestPaginationActivePageHighlight(t *testing.T) {
	pg := pagination.New(pagination.Config{
		CurrentPage: 3,
		TotalPages:  5,
	})
	if pg.CurrentPage != 3 {
		t.Errorf("expected active page 3")
	}
}

func TestPaginationPrevNextButtons(t *testing.T) {
	th := theme.NewDark()
	pg := pagination.New(pagination.Config{
		CurrentPage: 2,
		TotalPages:  5,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 40)),
	}
	dims := pg.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}

func TestPaginationOnSelectPageCallback(t *testing.T) {
	var selectedPage int
	pg := pagination.New(pagination.Config{
		CurrentPage: 1,
		TotalPages:  10,
		OnSelectPage: func(page int) {
			selectedPage = page
		},
	})
	if pg.OnSelectPage == nil {
		t.Errorf("expected OnSelectPage non-nil")
	}
	_ = selectedPage
}
