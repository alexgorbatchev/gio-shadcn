package textarea

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	txtArea := New(Config{Placeholder: "Enter track descriptions and metadata notes..."})
	return txtArea.Layout(gtx, th)
}
