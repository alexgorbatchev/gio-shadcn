package dialog

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	modalDialog := New(Config{
		Title:       "Reset Audio Mixer",
		Description: "Are you sure you want to reset all channel EQ gain levels?",
		Open:        false,
	})

	btnTrigger := button.New(button.Config{
		Text:    "Open Modal Dialog",
		Variant: theme.VariantOutline,
		OnClick: func() {
			modalDialog.Open = true
		},
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnTrigger.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return modalDialog.Layout(gtx, th) }),
	)
}
