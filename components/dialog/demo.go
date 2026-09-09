package dialog

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ProfileDialog     *Dialog
	AlertDialog       *Dialog
	DestructiveDialog *Dialog
	ShareDialog       *Dialog
	ScrollDialog      *Dialog
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	// 1. dialog-demo.tsx (Edit profile)
	s.ProfileDialog = New(Config{
		TriggerText: "Edit Profile Dialog",
		Title:       "Edit profile",
		Description: "Make changes to your profile here. Click save when you're done.",
		ConfirmText: "Save changes",
		CancelText:  "Cancel",
		Content: func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(theme.NewDark().MaterialTheme, unit.Sp(13), "Name: Pedro Duarte\nUsername: @peduarte")
					lbl.Color = theme.DarkColorScheme().Foreground
					return lbl.Layout(gtx)
				}),
			)
		},
		Open: false,
	})

	// 2. alert-dialog.tsx (Are you absolutely sure?)
	s.AlertDialog = New(Config{
		TriggerText: "Show Confirmation Alert Dialog",
		Title:       "Are you absolutely sure?",
		Description: "This action cannot be undone. This will permanently delete your account and remove your data from our servers.",
		ConfirmText: "Continue",
		CancelText:  "Cancel",
		Open:        false,
	})

	// 3. alert-dialog-destructive.tsx
	s.DestructiveDialog = New(Config{
		TriggerText: "Delete Project Dialog",
		Title:       "Delete Project",
		Description: "This will permanently delete the project repository and all associated artifacts.",
		ConfirmText: "Delete",
		CancelText:  "Cancel",
		Open:        false,
	})

	// 4. dialog-share.tsx
	s.ShareDialog = New(Config{
		TriggerText: "Share Link Dialog",
		Title:       "Share link",
		Description: "Anyone who has this link will be able to view this project.",
		ConfirmText: "Copy Link",
		CancelText:  "Close",
		Open:        false,
	})

	// 5. dialog-scrollable.tsx
	s.ScrollDialog = New(Config{
		TriggerText: "Terms & Conditions Dialog",
		Title:       "Terms and Conditions",
		Description: "Please review and accept our platform terms and privacy guidelines to continue.",
		ConfirmText: "I Accept",
		CancelText:  "Decline",
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
			return label.NewTypography("1. Profile Dialog", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ProfileDialog.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Alert Confirmation Dialog", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.AlertDialog.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("3. Destructive Action Dialog", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DestructiveDialog.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("4. Share Link Dialog", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ShareDialog.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("5. Scrollable Terms Dialog", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ScrollDialog.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
