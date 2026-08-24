package breadcrumb

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	BreadcrumbPath *Breadcrumb
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		BreadcrumbPath: New(Config{
			Items: []*Item{
				NewItem("Home", false),
				NewItem("Mixer", false),
				NewItem("Deck A", true),
			},
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.BreadcrumbPath.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
