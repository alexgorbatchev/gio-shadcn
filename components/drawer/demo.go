package drawer

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	DeliveryDrawer *Drawer
	ProfileDrawer  *Drawer
	GoalDrawer     *Drawer
	HandleDrawer   *Drawer
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	// 1. drawer-demo.tsx (Pick a delivery time)
	s.DeliveryDrawer = New(Config{
		TriggerText: "Pick Delivery Time Drawer",
		Title:       "Pick a delivery time",
		Description: "We'll prepare your order as soon as possible.",
		Height:      unit.Dp(280),
		Content: func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(theme.NewDark().MaterialTheme, unit.Sp(13), "• Standard Delivery (25–35 min) [Fastest]\n• 5:00 PM – 5:15 PM\n• 6:00 PM – 6:15 PM (High Demand)")
					lbl.Color = theme.DarkColorScheme().MutedFg
					return lbl.Layout(gtx)
				}),
			)
		},
		Open: false,
	})

	// 2. drawer-dialog.tsx (Edit profile)
	s.ProfileDrawer = New(Config{
		TriggerText: "Edit Profile Drawer",
		Title:       "Edit profile",
		Description: "Make changes to your profile here. Click save when you're done.",
		Height:      unit.Dp(260),
		Content: func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(theme.NewDark().MaterialTheme, unit.Sp(13), "Email: shadcn@example.com\nUsername: @shadcn")
			lbl.Color = theme.DarkColorScheme().Foreground
			return lbl.Layout(gtx)
		},
		Open: false,
	})

	// 3. drawer-sides.tsx (Move Goal)
	s.GoalDrawer = New(Config{
		TriggerText: "Move Goal Drawer",
		Title:       "Move Goal",
		Description: "Set your daily activity goal: 350 kcal / day.",
		Height:      unit.Dp(240),
		Open:        false,
	})

	// 4. drawer-swipe-handle.tsx
	s.HandleDrawer = New(Config{
		TriggerText: "Telemetry Drawer (Swipe Handle)",
		Title:       "System Telemetry",
		Description: "CPU: 2.1% | RAM: 189.5 MB | Metal GPU Frame Rate: 120 FPS",
		Height:      unit.Dp(260),
		Open:        false,
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("1. Standard Bottom Drawer", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DeliveryDrawer.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Profile Dialog Drawer", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ProfileDrawer.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("3. Move Goal Drawer", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.GoalDrawer.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("4. Telemetry Drawer (Swipe Handle)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.HandleDrawer.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
