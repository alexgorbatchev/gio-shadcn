package drawer

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	DeliveryDrawer *Drawer
	BtnDelivery    *button.Button
	ProfileDrawer  *Drawer
	BtnProfile     *button.Button
	GoalDrawer     *Drawer
	BtnGoal        *button.Button
	HandleDrawer   *Drawer
	BtnHandle      *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	// 1. drawer-demo.tsx (Pick a delivery time)
	s.DeliveryDrawer = New(Config{
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
	s.BtnDelivery = button.New(button.Config{
		Text:    "Pick Delivery Time Drawer",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.DeliveryDrawer.Open = true
		},
	})

	// 2. drawer-dialog.tsx (Edit profile)
	s.ProfileDrawer = New(Config{
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
	s.BtnProfile = button.New(button.Config{
		Text:    "Edit Profile Drawer",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.ProfileDrawer.Open = true
		},
	})

	// 3. drawer-sides.tsx (Move Goal)
	s.GoalDrawer = New(Config{
		Title:       "Move Goal",
		Description: "Set your daily activity goal: 350 kcal / day.",
		Height:      unit.Dp(240),
		Open:        false,
	})
	s.BtnGoal = button.New(button.Config{
		Text:    "Move Goal Drawer",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.GoalDrawer.Open = true
		},
	})

	// 4. drawer-swipe-handle.tsx
	s.HandleDrawer = New(Config{
		Title:       "System Telemetry",
		Description: "CPU: 2.1% | RAM: 189.5 MB | Metal GPU Frame Rate: 120 FPS",
		Height:      unit.Dp(260),
		Open:        false,
	})
	s.BtnHandle = button.New(button.Config{
		Text:    "Telemetry Drawer (Swipe Handle)",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.HandleDrawer.Open = true
		},
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnDelivery.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnProfile.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnGoal.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnHandle.Layout(gtx, th) }),
			)
		}),

		// Overlays
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if s.DeliveryDrawer.Open {
				return s.DeliveryDrawer.Layout(gtx, th)
			}
			if s.ProfileDrawer.Open {
				return s.ProfileDrawer.Layout(gtx, th)
			}
			if s.GoalDrawer.Open {
				return s.GoalDrawer.Layout(gtx, th)
			}
			if s.HandleDrawer.Open {
				return s.HandleDrawer.Layout(gtx, th)
			}
			return layout.Dimensions{}
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
