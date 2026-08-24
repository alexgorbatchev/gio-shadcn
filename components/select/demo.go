package selectcomp

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	GenreSel *Select
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		GenreSel: New(Config{
			Options: []*Item{
				NewItem("house", "Progressive House"),
				NewItem("techno", "Techno"),
				NewItem("trance", "Trance"),
			},
			SelectedValue: "house",
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.GenreSel.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
