package radio

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	RadioDefault     *Radio
	RadioComfortable *Radio
	RadioCompact     *Radio

	RadioDisabledA   *Radio
	RadioDisabledB   *Radio
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	selectOption := func(target *Radio) {
		s.RadioDefault.Selected = (s.RadioDefault == target)
		s.RadioComfortable.Selected = (s.RadioComfortable == target)
		s.RadioCompact.Selected = (s.RadioCompact == target)
	}

	s.RadioDefault = New(Config{Selected: true, OnChange: func(b bool) {
		if b {
			selectOption(s.RadioDefault)
		}
	}})

	s.RadioComfortable = New(Config{Selected: false, OnChange: func(b bool) {
		if b {
			selectOption(s.RadioComfortable)
		}
	}})

	s.RadioCompact = New(Config{Selected: false, OnChange: func(b bool) {
		if b {
			selectOption(s.RadioCompact)
		}
	}})

	s.RadioDisabledA = New(Config{Selected: true, Disabled: true})
	s.RadioDisabledB = New(Config{Selected: false, Disabled: true})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("1. Density Radio Group (Official Demo)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.RadioDefault.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Default")
							lbl.Color = th.Colors.Foreground
							return lbl.Layout(gtx)
						}),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.RadioComfortable.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Comfortable")
							lbl.Color = th.Colors.Foreground
							return lbl.Layout(gtx)
						}),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.RadioCompact.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Compact")
							lbl.Color = th.Colors.Foreground
							return lbl.Layout(gtx)
						}),
					)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Disabled Radio Options", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.RadioDisabledA.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Disabled (Checked)")
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.RadioDisabledB.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Disabled (Unchecked)")
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
