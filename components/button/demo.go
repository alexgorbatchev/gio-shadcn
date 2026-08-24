package button

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	BtnPrimary     *Button
	BtnSecondary   *Button
	BtnOutline     *Button
	BtnGhost       *Button
	BtnDestructive *Button
	BtnLink        *Button

	BtnSM          *Button
	BtnDefaultSize *Button
	BtnLG          *Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		BtnPrimary:     New(Config{Text: "Primary", Variant: theme.VariantDefault}),
		BtnSecondary:   New(Config{Text: "Secondary", Variant: theme.VariantSecondary}),
		BtnOutline:     New(Config{Text: "Outline", Variant: theme.VariantOutline}),
		BtnGhost:       New(Config{Text: "Ghost", Variant: theme.VariantGhost}),
		BtnDestructive: New(Config{Text: "Destructive", Variant: theme.VariantDestructive}),
		BtnLink:        New(Config{Text: "Link Button", Variant: theme.VariantLink}),

		BtnSM:          New(Config{Text: "Small", Size: theme.SizeSM}),
		BtnDefaultSize: New(Config{Text: "Default Size", Size: theme.SizeDefault}),
		BtnLG:          New(Config{Text: "Large Size", Size: theme.SizeLG}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnPrimary.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnSecondary.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnOutline.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnGhost.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnDestructive.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnLink.Layout(gtx, th) }),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnSM.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnDefaultSize.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnLG.Layout(gtx, th) }),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
