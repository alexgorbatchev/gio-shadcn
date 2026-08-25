package dialog

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ProfileDialog     *Dialog
	BtnProfile        *button.Button
	AlertDialog       *Dialog
	BtnAlert          *button.Button
	DestructiveDialog *Dialog
	BtnDestructive    *button.Button
	ShareDialog       *Dialog
	BtnShare          *button.Button
	ScrollDialog      *Dialog
	BtnScroll         *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	// 1. dialog-demo.tsx (Edit profile, Pedro Duarte)
	s.ProfileDialog = New(Config{
		Title:       "Edit profile",
		Description: "Make changes to your profile here. Click save when you're done.",
		ConfirmText: "Save changes",
		CancelText:  "Cancel",
		Content: func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(theme.NewDark().MaterialTheme, unit.Sp(13), "Name: Pedro Duarte\nUsername: @peduarte")
					lbl.Color = theme.DarkColorScheme().MutedFg
					return lbl.Layout(gtx)
				}),
			)
		},
		Open: false,
	})
	s.BtnProfile = button.New(button.Config{
		Text:    "Open Edit Profile Dialog",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.ProfileDialog.Open = true
		},
	})

	// 2. alert-dialog-demo.tsx (Are you absolutely sure?)
	s.AlertDialog = New(Config{
		Title:       "Are you absolutely sure?",
		Description: "This action cannot be undone. This will permanently delete your account from our servers.",
		ConfirmText: "Continue",
		CancelText:  "Cancel",
		Open:        false,
	})
	s.BtnAlert = button.New(button.Config{
		Text:    "Show Confirmation Alert Dialog",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.AlertDialog.Open = true
		},
	})

	// 3. alert-dialog-destructive.tsx (Delete chat? Trash2Icon)
	s.DestructiveDialog = New(Config{
		Title:       "Delete chat?",
		Description: "This will permanently delete this chat conversation and all memories.",
		ConfirmText: "Delete",
		CancelText:  "Cancel",
		Open:        false,
	})
	s.BtnDestructive = button.New(button.Config{
		Text:    "Delete Chat (Destructive)",
		Variant: theme.VariantDestructive,
		Icon:    lucide.Trash2,
		OnClick: func() {
			s.DestructiveDialog.Open = true
		},
	})

	// 4. dialog-close-button.tsx (Share link)
	s.ShareDialog = New(Config{
		Title:       "Share link",
		Description: "Anyone who has this link will be able to view this.",
		ConfirmText: "Copy Link",
		CancelText:  "Close",
		Content: func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(theme.NewDark().MaterialTheme, unit.Sp(12), "https://ui.shadcn.com/docs/installation")
			lbl.Color = theme.DarkColorScheme().Foreground
			return lbl.Layout(gtx)
		},
		Open: false,
	})
	s.BtnShare = button.New(button.Config{
		Text:    "Share Link Dialog",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.ShareDialog.Open = true
		},
	})

	// 5. dialog-scrollable-content.tsx
	s.ScrollDialog = New(Config{
		Title:       "Scrollable Content",
		Description: "This is a dialog with multi-paragraph scrollable content.",
		ConfirmText: "Accept",
		CancelText:  "Dismiss",
		Content: func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(theme.NewDark().MaterialTheme, unit.Sp(13), "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.")
			lbl.Color = theme.DarkColorScheme().MutedFg
			return lbl.Layout(gtx)
		},
		Open: false,
	})
	s.BtnScroll = button.New(button.Config{
		Text:    "Scrollable Content Dialog",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.ScrollDialog.Open = true
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
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnProfile.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnAlert.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnDestructive.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnShare.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnScroll.Layout(gtx, th) }),
			)
		}),

		// Overlays
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if s.ProfileDialog.Open {
				return s.ProfileDialog.Layout(gtx, th)
			}
			if s.AlertDialog.Open {
				return s.AlertDialog.Layout(gtx, th)
			}
			if s.DestructiveDialog.Open {
				return s.DestructiveDialog.Layout(gtx, th)
			}
			if s.ShareDialog.Open {
				return s.ShareDialog.Layout(gtx, th)
			}
			if s.ScrollDialog.Open {
				return s.ScrollDialog.Layout(gtx, th)
			}
			return layout.Dimensions{}
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
