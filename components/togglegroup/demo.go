package togglegroup

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	TglGrp *ToggleGroup
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		TglGrp: New(Config{
			Items: []*Item{
				NewItem("mono", "Mono"),
				NewItem("stereo", "Stereo"),
				NewItem("5.1", "5.1 Surround"),
			},
			SelectedKey: "stereo",
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.TglGrp.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
