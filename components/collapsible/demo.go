package collapsible

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	BasicCol    *Collapsible
	DemoCol     *Collapsible
	FileTreeCol *Collapsible
	SettingsCol *Collapsible
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	// 1. collapsible-basic.tsx
	s.BasicCol = New(Config{
		Title: "Product details",
		Content: "This panel can be expanded or collapsed to reveal additional content.\n" +
			"Learn more about the product specifications and warranty details.",
		Open: true,
	})

	// 2. collapsible-demo.tsx (Order #4189, Status Shipped, 100 Market St)
	s.DemoCol = New(Config{
		Title: "Order #4189 — Shipped",
		Content: "Shipping Address: 100 Market St, San Francisco\n" +
			"Items: 2x Studio Headphones",
		Open: true,
	})

	// 3. collapsible-file-tree.tsx (components/ui/button.tsx, lib/utils.ts)
	s.FileTreeCol = New(Config{
		Title: "components/ui (File Tree)",
		ContentWidget: func(gtx layout.Context) layout.Dimensions {
			files := []string{"button.tsx", "card.tsx", "dialog.tsx", "input.tsx", "select.tsx", "table.tsx"}
			children := make([]layout.FlexChild, len(files))
			for i, f := range files {
				fileName := f
				children[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return lucide.File.LayoutSize(gtx, unit.Dp(14), theme.DarkColorScheme().MutedFg)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Width: unit.Dp(6)}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(theme.NewDark().MaterialTheme, unit.Sp(13), fileName)
							lbl.Color = theme.DarkColorScheme().Foreground
							return lbl.Layout(gtx)
						}),
					)
				})
			}
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
		},
		Open: true,
	})

	// 4. collapsible-settings.tsx (Radius X, Radius Y)
	s.SettingsCol = New(Config{
		Title:   "Radius Settings",
		Content: "Corner Radius: X = 8px, Y = 8px\nElement border radius configured.",
		Open:    false,
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BasicCol.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoCol.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.FileTreeCol.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SettingsCol.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
