package card

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	demoCard := New(Config{Variant: theme.VariantDefault})

	return demoCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Card Content Area")
		lbl.Color = th.Colors.Foreground
		return lbl.Layout(gtx)
	})
}
