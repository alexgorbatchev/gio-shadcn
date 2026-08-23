package breadcrumb

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	bCrumb := New(Config{
		Items: []*Item{
			NewItem("Home", false),
			NewItem("Mixer", false),
			NewItem("Deck A", true),
		},
	})

	return bCrumb.Layout(gtx, th)
}
