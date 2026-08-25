package avatar

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	DemoBasic     *Avatar
	DemoBadge     *Avatar
	DemoBadgeIcon *Avatar
	DemoSizeSM    *Avatar
	DemoSizeMD    *Avatar
	DemoSizeLG    *Avatar
	GroupItem1    *Avatar
	GroupItem2    *Avatar
	GroupItem3    *Avatar
	GroupCount    *Avatar
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	// avatar-basic.tsx
	s.DemoBasic = New(Config{
		Initials: "CN",
		Size:     unit.Dp(40),
	})

	// avatar-badge.tsx (online green status dot)
	s.DemoBadge = New(Config{
		Initials:   "CN",
		ShowBadge:  true,
		BadgeColor: "green",
	})

	// avatar-badge-icon.tsx (plus icon badge)
	s.DemoBadgeIcon = New(Config{
		Initials:  "PP",
		ShowBadge: true,
	})

	// avatar-size.tsx (sm: 32dp, md: 40dp, lg: 56dp)
	s.DemoSizeSM = New(Config{Initials: "CN", Size: unit.Dp(32)})
	s.DemoSizeMD = New(Config{Initials: "CN", Size: unit.Dp(40)})
	s.DemoSizeLG = New(Config{Initials: "CN", Size: unit.Dp(56)})

	// avatar-group.tsx & avatar-group-count.tsx
	s.GroupItem1 = New(Config{Initials: "CN"})
	s.GroupItem2 = New(Config{Initials: "LR"})
	s.GroupItem3 = New(Config{Initials: "ER"})
	s.GroupCount = New(Config{Initials: "+3"})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Avatar (Basic)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoBasic.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Avatar (With Status Badge)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoBadge.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Avatar (With Action Badge)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoBadgeIcon.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Avatar (Sizes: SM, Default, LG)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoSizeSM.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoSizeMD.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoSizeLG.Layout(gtx, th) }),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Avatar (Group & Count)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.GroupItem1.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.GroupItem2.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.GroupItem3.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.GroupCount.Layout(gtx, th) }),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
