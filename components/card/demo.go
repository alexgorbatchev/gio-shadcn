package card

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	CardStandard   *Card
	CardSmall      *Card
	CardEdgeToEdge *Card
	CardImage      *Card
	CardSpacing    *Card

	BtnDeploy      *button.Button
	BtnCancel      *button.Button
	BtnSave        *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		CardStandard:   New(Config{Variant: theme.VariantDefault}),
		CardSmall:      New(Config{Variant: theme.VariantDefault, Padding: layout.UniformInset(unit.Dp(16))}),
		CardEdgeToEdge: New(Config{Variant: theme.VariantDefault}),
		CardImage:      New(Config{Variant: theme.VariantDefault}),
		CardSpacing:    New(Config{Variant: theme.VariantDefault, Padding: layout.UniformInset(unit.Dp(32))}),

		BtnDeploy:      button.New(button.Config{Text: "Deploy Project", Variant: theme.VariantDefault}),
		BtnCancel:      button.New(button.Config{Text: "Cancel", Variant: theme.VariantOutline}),
		BtnSave:        button.New(button.Config{Text: "Save Preferences", Variant: theme.VariantSecondary}),
	}
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
			return label.NewTypography("1. Standard Project Card", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.CardStandard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return label.NewTypography("Create project", label.H3, "").Layout(gtx, th)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return label.NewTypography("Deploy your new project in one-click.", label.Muted, "").Layout(gtx, th)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnCancel.Layout(gtx, th) }),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnDeploy.Layout(gtx, th) }),
						)
					}),
				)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Small Compact Card", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.CardSmall.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return label.NewTypography("Notifications", label.H4, "").Layout(gtx, th)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return label.NewTypography("You have 3 unread messages.", label.Small, "").Layout(gtx, th)
					}),
				)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("3. Relaxed Spacing Card", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.CardSpacing.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return label.NewTypography("Custom Padding Settings", label.H4, "").Layout(gtx, th)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return label.NewTypography("This card demonstrates custom layout insets and padding scaling.", label.P, "").Layout(gtx, th)
					}),
				)
			})
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
