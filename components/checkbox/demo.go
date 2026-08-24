package checkbox

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	ChkBox *Checkbox
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		ChkBox: New(Config{Value: true}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ChkBox.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			mTheme := th.MaterialTheme
			if mTheme == nil {
				mTheme = material.NewTheme()
			}
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Enable GPU Vector Acceleration")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
