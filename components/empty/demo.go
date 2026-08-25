package empty

import (
	"gioui.org/layout"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	EmptyDemo     *Empty
	EmptyNoAction *Empty
	BtnAction     *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	btn := button.New(button.Config{Text: "Clear Filters", Variant: theme.VariantDefault})
	return &DemoState{
		BtnAction: btn,
		EmptyDemo: New(Config{
			Title:       "No results found",
			Description: "No tracks or playlists matched your search criteria.",
			Icon:        lucide.Search,
			Action: func(gtx layout.Context) layout.Dimensions {
				return btn.Layout(gtx, theme.NewDark())
			},
		}),
		EmptyNoAction: New(Config{
			Title:       "No Audio Inputs Connected",
			Description: "Connect an audio interface or ASIO device to begin.",
			Icon:        lucide.Radio,
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("1. Empty State with Action Button", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.EmptyDemo.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Simple Empty State", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.EmptyNoAction.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
