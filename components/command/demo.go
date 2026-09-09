package command

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	BasicCmd     *Command
	DemoCmd      *Command
	DialogCmd    *Command
	GroupsCmd    *Command
	ShortcutsCmd *Command
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	// 1. command-basic.tsx (Calendar, Search Emoji, Calculator)
	s.BasicCmd = New(Config{
		TriggerText: "Open Basic Command Palette",
		Placeholder: "Type a command or search...",
		Items: []*Item{
			NewItem("Calendar", ""),
			NewItem("Search Emoji", ""),
			NewItem("Calculator", ""),
		},
	})

	// 2. command-demo.tsx (Suggestions + Settings with Icons & Shortcuts)
	s.DemoCmd = New(Config{
		TriggerText: "Open Full Palette (Suggestions & Settings)",
		Placeholder: "Type a command or search...",
		Items: []*Item{
			NewItemFull("Calendar", "", "Suggestions", lucide.Calendar, false),
			NewItemFull("Search Emoji", "", "Suggestions", lucide.FaceSlightlySmiling, false),
			NewItemFull("Calculator", "", "Suggestions", lucide.Calculator, true),
			NewItemFull("Profile", "⌘P", "Settings", lucide.User, false),
			NewItemFull("Billing", "⌘B", "Settings", lucide.CreditCard, false),
			NewItemFull("Settings", "⌘S", "Settings", lucide.Settings, false),
		},
	})

	// 3. command-dialog.tsx (Press ⌘J overlay palette)
	s.DialogCmd = New(Config{
		TriggerText: "Open Command Dialog (⌘J)",
		Placeholder: "Type a command or search...",
		Items: []*Item{
			NewItemFull("Find in Project", "⌘F", "Actions", lucide.Search, false),
			NewItemFull("Toggle Theme", "⌘T", "Actions", lucide.Sparkles, false),
			NewItemFull("Open Settings", "⌘,", "Actions", lucide.Settings, false),
		},
	})

	// 4. command-groups.tsx (Explicit Suggestions & Settings Groups)
	s.GroupsCmd = New(Config{
		TriggerText: "Open Grouped Command Palette",
		Placeholder: "Filter across groups...",
		Items: []*Item{
			NewItemFull("Calendar View", "", "Suggestions", lucide.Calendar, false),
			NewItemFull("Emoji Picker", "", "Suggestions", lucide.FaceSlightlySmiling, false),
			NewItemFull("Account Profile", "⌘P", "Settings", lucide.User, false),
		},
	})

	// 5. command-shortcuts.tsx (Profile ⌘P, Billing ⌘B, Settings ⌘S)
	s.ShortcutsCmd = New(Config{
		TriggerText: "Open Shortcuts Palette",
		Placeholder: "Search shortcuts...",
		Items: []*Item{
			NewItemFull("Profile", "⌘P", "Shortcuts", lucide.User, false),
			NewItemFull("Billing", "⌘B", "Shortcuts", lucide.CreditCard, false),
			NewItemFull("Settings", "⌘S", "Shortcuts", lucide.Settings, false),
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
			return label.NewTypography("1. Command Basic", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BasicCmd.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Command Demo (Groups, Icons & Shortcuts)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoCmd.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("3. Command Dialog (Shortcut: ⌘J)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DialogCmd.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("4. Command Groups", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.GroupsCmd.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("5. Command Shortcuts", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ShortcutsCmd.Layout(gtx, th) }),
	)
}

func (s *DemoState) RenderOverlay(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if s.BasicCmd.Open {
		return s.BasicCmd.Layout(gtx, th)
	}
	if s.DemoCmd.Open {
		return s.DemoCmd.Layout(gtx, th)
	}
	if s.DialogCmd.Open {
		return s.DialogCmd.Layout(gtx, th)
	}
	if s.GroupsCmd.Open {
		return s.GroupsCmd.Layout(gtx, th)
	}
	if s.ShortcutsCmd.Open {
		return s.ShortcutsCmd.Layout(gtx, th)
	}
	return layout.Dimensions{}
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}

func RenderOverlay(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.RenderOverlay(gtx, th)
}
