package dialog

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ModalDialog *Dialog
	BtnTrigger  *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	s.ModalDialog = New(Config{
		Title:       "Reset Audio Mixer",
		Description: "Are you sure you want to reset all channel EQ gain levels?",
		Open:        false,
	})

	s.BtnTrigger = button.New(button.Config{
		Text:    "Open Modal Dialog",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.ModalDialog.Open = true
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
			if s.ModalDialog.Open {
				return s.ModalDialog.Layout(gtx, th)
			}
			return layout.Dimensions{}
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
