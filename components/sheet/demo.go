package sheet

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	SheetDrawer *Sheet
	BtnTrigger  *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	s.SheetDrawer = New(Config{
		Title:       "Track Inspector",
		Description: "Detailed FLAC metadata and harmonic key analysis.",
		Open:        false,
	})

	s.BtnTrigger = button.New(button.Config{
		Text:    "Open Side Sheet",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.SheetDrawer.Open = true
		},
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return s.BtnTrigger.Layout(gtx, th)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if s.SheetDrawer.Open {
				return s.SheetDrawer.Layout(gtx, th)
			}
			return layout.Dimensions{}
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
