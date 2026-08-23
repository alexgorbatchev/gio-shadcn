package selectcomp

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	genreSel := New(Config{
		Options: []*Item{
			NewItem("house", "Progressive House"),
			NewItem("techno", "Techno"),
			NewItem("trance", "Trance"),
		},
		SelectedValue: "house",
	})

	return genreSel.Layout(gtx, th)
}
