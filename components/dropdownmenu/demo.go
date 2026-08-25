package dropdownmenu

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	DropdownDemo   *DropdownMenu
	MenuAccount    *DropdownMenu
	MenuCheckboxes *DropdownMenu
	MenuShortcuts  *DropdownMenu
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	s.DropdownDemo = New(Config{
		TriggerText: "Open dropdown menu",
		Open:        false,
		Items: []*Item{
			NewItem("Profile", "⇧⌘P"),
			NewItem("Billing", "⌘B"),
			NewItem("Settings", "⌘S"),
			NewItem("Keyboard shortcuts", "⌘K"),
			NewItem("Team", ""),
			NewItem("Invite users", ""),
			NewItem("Log out", "⇧⌘Q"),
		},
	})

	s.MenuAccount = New(Config{
		Open: true,
		Items: []*Item{
			NewItem("My Account", ""),
			NewItem("Profile", "⇧⌘P"),
			NewItem("Billing", "⌘B"),
			NewItem("Settings", "⌘S"),
		},
	})

	s.MenuCheckboxes = New(Config{
		Open: true,
		Items: []*Item{
			NewItem("✓ Status Bar", ""),
			NewItem("✓ Activity Bar", ""),
			NewItem("  Panel", ""),
		},
	})

	s.MenuShortcuts = New(Config{
		Open: true,
		Items: []*Item{
			NewItem("New Tab", "⌘T"),
			NewItem("New Window", "⌘N"),
			NewItem("Open File...", "⌘O"),
		},
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("1. Interactive Dropdown Menu (Click button to toggle)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DropdownDemo.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Account Menu Panel", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.MenuAccount.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("3. Checkbox Menu Items", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.MenuCheckboxes.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("4. Shortcuts Menu", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.MenuShortcuts.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
