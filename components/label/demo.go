package label

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	H1Typography    *Typography
	H2Typography    *Typography
	H3Typography    *Typography
	H4Typography    *Typography
	LeadTypography  *Typography
	PTypography     *Typography
	LargeTypography *Typography
	SmallTypography *Typography
	MutedTypography *Typography
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		H1Typography:    NewTypography("The Joke Tax Chronicles", H1, ""),
		H2Typography:    NewTypography("The People of the Kingdom", H2, ""),
		H3Typography:    NewTypography("The Joke Tax", H3, ""),
		H4Typography:    NewTypography("People stopped telling jokes", H4, ""),
		LeadTypography:  NewTypography("A modal series of components designed for maximum accessibility.", P, ""),
		PTypography:     NewTypography("The king thought long and hard, and finally decided to tax all jokes in the kingdom.", P, ""),
		LargeTypography: NewTypography("Are you absolutely sure?", H4, ""),
		SmallTypography: NewTypography("Email address is required for verification.", Small, ""),
		MutedTypography: NewTypography("Enter your email address.", Muted, ""),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.H1Typography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.H2Typography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.H3Typography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.H4Typography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.LeadTypography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.PTypography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.LargeTypography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SmallTypography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.MutedTypography.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
