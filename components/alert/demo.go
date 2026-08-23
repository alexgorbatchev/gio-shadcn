package alert

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	infoAlert := New(Config{Title: "Engine Status", Description: "CoreAudio buffer set to 64 samples."})
	destAlert := New(Config{Title: "Audio Clip Warning", Description: "Output signal clipped +1.2dB on Deck A.", Variant: theme.VariantDestructive})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return infoAlert.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return destAlert.Layout(gtx, th) }),
	)
}
