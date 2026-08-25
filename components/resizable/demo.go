package resizable

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ResDemo   *Resizable
	ResHandle *Resizable
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		ResDemo: New(Config{
			Ratio: 0.5,
			LeftWidget: func(gtx layout.Context) layout.Dimensions {
				th := theme.NewDark()
				lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Panel One")
				lbl.Color = th.Colors.Foreground
				return layout.Center.Layout(gtx, lbl.Layout)
			},
			RightWidget: func(gtx layout.Context) layout.Dimensions {
				th := theme.NewDark()
				lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Panel Two")
				lbl.Color = th.Colors.Foreground
				return layout.Center.Layout(gtx, lbl.Layout)
			},
		}),
		ResHandle: New(Config{
			Ratio: 0.35,
			LeftWidget: func(gtx layout.Context) layout.Dimensions {
				th := theme.NewDark()
				lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Sidebar (35%)")
				lbl.Color = th.Colors.MutedFg
				return layout.Center.Layout(gtx, lbl.Layout)
			},
			RightWidget: func(gtx layout.Context) layout.Dimensions {
				th := theme.NewDark()
				lbl := material.Label(th.MaterialTheme, th.Typography.FontSizeBase, "Main Content (65%)")
				lbl.Color = th.Colors.Foreground
				return layout.Center.Layout(gtx, lbl.Layout)
			},
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("1. Resizable 50/50 Split (Drag divider)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ResDemo.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Sidebar / Content Split (35/65)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ResHandle.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
