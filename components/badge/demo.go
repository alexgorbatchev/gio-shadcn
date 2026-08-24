package badge

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	BadgeDef  *Badge
	BadgeSec  *Badge
	BadgeOut  *Badge
	BadgeDest *Badge
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		BadgeDef:  New(Config{Text: "Default Badge", Variant: theme.VariantDefault}),
		BadgeSec:  New(Config{Text: "Secondary Badge", Variant: theme.VariantSecondary}),
		BadgeOut:  New(Config{Text: "Outline Badge", Variant: theme.VariantOutline}),
		BadgeDest: New(Config{Text: "Destructive Badge", Variant: theme.VariantDestructive}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeDef.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeSec.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeOut.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeDest.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
