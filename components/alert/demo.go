package alert

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	InfoAlert *Alert
	DestAlert *Alert
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		InfoAlert: New(Config{Title: "Engine Status", Description: "CoreAudio buffer set to 64 samples."}),
		DestAlert: New(Config{Title: "Audio Clip Warning", Description: "Output signal clipped +1.2dB on Deck A.", Variant: theme.VariantDestructive}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.InfoAlert.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DestAlert.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
