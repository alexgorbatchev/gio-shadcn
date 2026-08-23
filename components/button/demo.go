package button

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	btnPrimary := New(Config{Text: "Primary", Variant: theme.VariantDefault})
	btnSecondary := New(Config{Text: "Secondary", Variant: theme.VariantSecondary})
	btnOutline := New(Config{Text: "Outline", Variant: theme.VariantOutline})
	btnGhost := New(Config{Text: "Ghost", Variant: theme.VariantGhost})
	btnDestructive := New(Config{Text: "Destructive", Variant: theme.VariantDestructive})
	btnLink := New(Config{Text: "Link Button", Variant: theme.VariantLink})

	btnSM := New(Config{Text: "Small", Size: theme.SizeSM})
	btnDefaultSize := New(Config{Text: "Default Size", Size: theme.SizeDefault})
	btnLG := New(Config{Text: "Large Size", Size: theme.SizeLG})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnPrimary.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnSecondary.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnOutline.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnGhost.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnDestructive.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnLink.Layout(gtx, th) }),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnSM.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnDefaultSize.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnLG.Layout(gtx, th) }),
			)
		}),
	)
}
