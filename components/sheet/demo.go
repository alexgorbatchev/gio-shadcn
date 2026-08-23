package sheet

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	sheetDrawer := New(Config{
		Title:       "Track Inspector",
		Description: "Detailed FLAC metadata and harmonic key analysis.",
		Open:        false,
	})

	btnTrigger := button.New(button.Config{
		Text:    "Open Side Sheet",
		Variant: theme.VariantOutline,
		OnClick: func() {
			sheetDrawer.Open = true
		},
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnTrigger.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sheetDrawer.Layout(gtx, th) }),
	)
}
