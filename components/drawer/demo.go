package drawer

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	BottomDrawer *Drawer
	BtnTrigger   *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	s.BottomDrawer = New(Config{
		Title:       "System Telemetry",
		Description: "CPU: 2.1% | RAM: 189.5MB | Metal GPU 120 FPS",
		Open:        false,
	})

	s.BtnTrigger = button.New(button.Config{
		Text:    "Open Bottom Drawer",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.BottomDrawer.Open = true
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
			if s.BottomDrawer.Open {
				return s.BottomDrawer.Layout(gtx, th)
			}
			return layout.Dimensions{}
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
