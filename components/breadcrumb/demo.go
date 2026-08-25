package breadcrumb

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	BreadcrumbDemo      *Breadcrumb
	BreadcrumbSlash     *Breadcrumb
	BreadcrumbEllipsis  *Breadcrumb
	BreadcrumbDropdown  *Breadcrumb
	BreadcrumbLinks     *Breadcrumb
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		BreadcrumbDemo: New(Config{
			Items: []*Item{
				NewItem("Home", false),
				NewItem("Components", false),
				NewItem("Breadcrumb", true),
			},
		}),
		BreadcrumbSlash: New(Config{
			Separator: "/",
			Items: []*Item{
				NewItem("Home", false),
				NewItem("Docs", false),
				NewItem("Components", false),
				NewItem("Breadcrumb", true),
			},
		}),
		BreadcrumbEllipsis: New(Config{
			Items: []*Item{
				NewItem("Home", false),
				NewItem("...", false),
				NewItem("Components", false),
				NewItem("Breadcrumb", true),
			},
		}),
		BreadcrumbDropdown: New(Config{
			Items: []*Item{
				NewItem("Home", false),
				NewItem("Templates ▾", false),
				NewItem("Dashboard", true),
			},
		}),
		BreadcrumbLinks: New(Config{
			Items: []*Item{
				NewItem("Workspace", false),
				NewItem("Audio Decks", false),
				NewItem("Deck A", true),
			},
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Breadcrumb Chevron (Default)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BreadcrumbDemo.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Breadcrumb Slash Separator", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BreadcrumbSlash.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Breadcrumb with Ellipsis", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BreadcrumbEllipsis.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Breadcrumb Dropdown & Links", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BreadcrumbDropdown.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BreadcrumbLinks.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
