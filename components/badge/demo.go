package badge

import (
	"gioui.org/layout"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	BadgeDefault     *Badge
	BadgeSecondary   *Badge
	BadgeOutline     *Badge
	BadgeDestructive *Badge
	BadgeWithIcon    *Badge
	BadgeWithSpinner *Badge
	BadgeLink        *Badge
	BadgeCustomColor *Badge
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		BadgeDefault:     New(Config{Text: "Badge", Variant: theme.VariantDefault}),
		BadgeSecondary:   New(Config{Text: "Secondary", Variant: theme.VariantSecondary}),
		BadgeOutline:     New(Config{Text: "Outline", Variant: theme.VariantOutline}),
		BadgeDestructive: New(Config{Text: "Destructive", Variant: theme.VariantDestructive}),
		BadgeWithIcon:    New(Config{Text: "Verified", Variant: theme.VariantSecondary, Icon: lucide.Check}),
		BadgeWithSpinner: New(Config{Text: "Updating", Variant: theme.VariantOutline, Icon: lucide.LoaderCircle}),
		BadgeLink:        New(Config{Text: "Documentation ↗", Variant: theme.VariantDefault}),
		BadgeCustomColor: New(Config{Text: "Feature", Variant: theme.VariantSecondary}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Badge Variants", label.H3, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeDefault.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeSecondary.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeOutline.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeDestructive.Layout(gtx, th) }),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Badge with Icons & Link States", label.H3, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeWithIcon.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeWithSpinner.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeLink.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeCustomColor.Layout(gtx, th) }),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
