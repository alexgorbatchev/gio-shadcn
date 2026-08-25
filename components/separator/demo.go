package separator

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	SepHorizontal *Separator
	SepVertical1  *Separator
	SepVertical2  *Separator
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		SepHorizontal: New(Config{Horizontal: true}),
		SepVertical1:  New(Config{Horizontal: false}),
		SepVertical2:  New(Config{Horizontal: false}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Radix UI / shadcn", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("An open-source UI component library.", label.Muted, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SepHorizontal.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return label.NewTypography("Blog", label.Small, "").Layout(gtx, th)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SepVertical1.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return label.NewTypography("Docs", label.Small, "").Layout(gtx, th)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SepVertical2.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return label.NewTypography("Source", label.Small, "").Layout(gtx, th)
				}),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
