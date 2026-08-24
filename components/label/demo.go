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
	PTypography     *Typography
	MutedTypography *Typography
	SmallTypography *Typography
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		H1Typography:    NewTypography("Typography Heading 1", H1, ""),
		H2Typography:    NewTypography("Typography Heading 2", H2, ""),
		H3Typography:    NewTypography("Typography Heading 3", H3, ""),
		H4Typography:    NewTypography("Typography Heading 4", H4, ""),
		PTypography:     NewTypography("Body paragraph demonstrating standard typography scaling.", P, ""),
		MutedTypography: NewTypography("Muted secondary text style.", Muted, ""),
		SmallTypography: NewTypography("Small caption text for fine print.", Small, ""),
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
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.PTypography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.MutedTypography.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SmallTypography.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
