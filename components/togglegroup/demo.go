package togglegroup

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	tglGrp := New(Config{
		Items: []*Item{
			NewItem("mono", "Mono"),
			NewItem("stereo", "Stereo"),
			NewItem("5.1", "5.1 Surround"),
		},
		SelectedKey: "stereo",
	})

	return tglGrp.Layout(gtx, th)
}
