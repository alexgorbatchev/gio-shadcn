package dropdownmenu

import (
	"gioui.org/layout"
	"github.com/alexgorbatchev/gio-lucide"
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
		Icon:        lucide.ChevronDown,
		Open:        false,
		Items: []*Item{
			NewItemWithIcon("Profile", "⇧⌘P", lucide.User),
			NewItemWithIcon("Billing", "⌘B", lucide.CreditCard),
			NewItemWithIcon("Settings", "⌘S", lucide.Settings),
			NewItemWithIcon("Keyboard shortcuts", "⌘K", lucide.SlidersHorizontal),
			NewItemWithIcon("Team", "", lucide.Users),
			NewItemWithIcon("Invite users", "", lucide.UserPlus),
			NewItemWithIcon("Log out", "⇧⌘Q", lucide.LogOut),
		},
	})

	s.MenuAccount = New(Config{
		Open: true,
		Items: []*Item{
			NewItemWithIcon("My Account", "", lucide.User),
			NewItemWithIcon("Profile", "⇧⌘P", lucide.User),
			NewItemWithIcon("Billing", "⌘B", lucide.CreditCard),
			NewItemWithIcon("Settings", "⌘S", lucide.Settings),
		},
	})

	s.MenuCheckboxes = New(Config{
		Open: true,
		Items: []*Item{
			NewCheckboxItem("Status Bar", true, nil),
			NewCheckboxItem("Activity Bar", true, nil),
			NewCheckboxItem("Panel", false, nil),
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
			return label.NewTypography("1. Interactive Dropdown Menu (Click button to toggle floating overlay)", label.H4, "").Layout(gtx, th)
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
			return label.NewTypography("3. Checkbox Menu Items (Click to toggle checkmarks)", label.H4, "").Layout(gtx, th)
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
