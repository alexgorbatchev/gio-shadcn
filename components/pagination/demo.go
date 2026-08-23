package pagination

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	pageControls := New(Config{CurrentPage: 1, TotalPages: 5})
	return pageControls.Layout(gtx, th)
}
