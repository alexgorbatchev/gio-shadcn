package card

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	DemoCard *Card
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		DemoCard: New(Config{Variant: theme.VariantDefault}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.DemoCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Card Content Area")
		lbl.Color = th.Colors.Foreground
		return lbl.Layout(gtx)
	})
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
