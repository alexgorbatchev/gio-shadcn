package carousel

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	CarouselDemo     *Carousel
	CarouselSize     *Carousel
	CarouselMultiple *Carousel
	DemoCards        []*card.Card
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	cards := make([]*card.Card, 5)
	for i := range cards {
		cards[i] = card.New(card.Config{Variant: theme.VariantDefault, Padding: layout.UniformInset(unit.Dp(24))})
	}

	items := make([]layout.Widget, 5)
	for i := range items {
		idx := i + 1
		c := cards[i]
		items[i] = func(gtx layout.Context) layout.Dimensions {
			th := theme.NewDark()
			return c.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.MaterialTheme, th.Typography.FontSize2XL, fmt.Sprintf("%d", idx))
					lbl.Color = th.Colors.Foreground
					return lbl.Layout(gtx)
				})
			})
		}
	}

	return &DemoState{
		CarouselDemo: New(Config{
			Items: items,
		}),
		CarouselSize: New(Config{
			Items: items[:3],
		}),
		CarouselMultiple: New(Config{
			Items: items,
		}),
		DemoCards: cards,
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Carousel Standard Slider (1 to 5)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.CarouselDemo.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Carousel Sized Cards", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.CarouselSize.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
