package progress

import (
	"fmt"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ProgDemo       *Progress
	ProgControlled *Progress
	BtnIncrement   *button.Button
	BtnReset       *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	s.ProgDemo = New(Config{Value: 0.66})
	s.ProgControlled = New(Config{Value: 0.33})

	s.BtnIncrement = button.New(button.Config{
		Text:    "+20%",
		Variant: theme.VariantOutline,
		OnClick: func() {
			s.ProgControlled.Value += 0.20
			if s.ProgControlled.Value > 1.0 {
				s.ProgControlled.Value = 1.0
			}
		},
	})

	s.BtnReset = button.New(button.Config{
		Text:    "Reset",
		Variant: theme.VariantGhost,
		OnClick: func() {
			s.ProgControlled.Value = 0.0
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
			return label.NewTypography("1. Default Progress (66%)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ProgDemo.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			pct := int(s.ProgControlled.Value * 100)
			return label.NewTypography(fmt.Sprintf("2. Controlled Progress (%d%%)", pct), label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ProgControlled.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnIncrement.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnReset.Layout(gtx, th) }),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
