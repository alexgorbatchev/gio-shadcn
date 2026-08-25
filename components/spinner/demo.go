package spinner

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/badge"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/empty"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	// 1. Badge Spinners
	SpinBadge1 *Spinner
	BadgeSync  *badge.Badge
	SpinBadge2 *Spinner
	BadgeUpd   *badge.Badge
	SpinBadge3 *Spinner
	BadgeProc  *badge.Badge

	// 2. Button Spinners
	SpinBtn1 *Spinner
	BtnLoad  *button.Button
	SpinBtn2 *Spinner
	BtnWait  *button.Button
	SpinBtn3 *Spinner
	BtnProc  *button.Button

	// 3. Custom Spinner
	CustomSpin *Spinner

	// 4. Item Demo Spinner
	DemoSpin *Spinner

	// 5. Empty State Spinner
	EmptySpin *Spinner
	EmptyCard *empty.Empty
	EmptyBtn  *button.Button

	// 6. Sizes
	SpinSizeSM *Spinner
	SpinSizeMD *Spinner
	SpinSizeLG *Spinner
	SpinSizeXL *Spinner
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{
		SpinBadge1: New(Config{Size: unit.Dp(14)}),
		BadgeSync:  badge.New(badge.Config{Text: "Syncing", Variant: theme.VariantDefault}),
		SpinBadge2: New(Config{Size: unit.Dp(14)}),
		BadgeUpd:   badge.New(badge.Config{Text: "Updating", Variant: theme.VariantSecondary}),
		SpinBadge3: New(Config{Size: unit.Dp(14)}),
		BadgeProc:  badge.New(badge.Config{Text: "Processing", Variant: theme.VariantOutline}),

		SpinBtn1: New(Config{Size: unit.Dp(16)}),
		BtnLoad:  button.New(button.Config{Text: "Loading...", Size: theme.SizeSM, Disabled: true}),
		SpinBtn2: New(Config{Size: unit.Dp(16)}),
		BtnWait:  button.New(button.Config{Text: "Please wait", Variant: theme.VariantOutline, Size: theme.SizeSM, Disabled: true}),
		SpinBtn3: New(Config{Size: unit.Dp(16)}),
		BtnProc:  button.New(button.Config{Text: "Processing", Variant: theme.VariantSecondary, Size: theme.SizeSM, Disabled: true}),

		CustomSpin: New(Config{Size: unit.Dp(16)}),
		DemoSpin:   New(Config{Size: unit.Dp(20)}),

		EmptySpin: New(Config{Size: unit.Dp(32)}),
		EmptyCard: empty.New(empty.Config{}),
		EmptyBtn:  button.New(button.Config{Text: "Cancel", Variant: theme.VariantOutline, Size: theme.SizeSM}),

		SpinSizeSM: New(Config{Size: unit.Dp(12)}),
		SpinSizeMD: New(Config{Size: unit.Dp(16)}),
		SpinSizeLG: New(Config{Size: unit.Dp(24)}),
		SpinSizeXL: New(Config{Size: unit.Dp(32)}),
	}
	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}
	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Title
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeXL, "Spinner Showcase (6 Upstream Demos)")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		// Demo 1: Spinner Badge
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "1. Spinner in Badges (spinner-badge.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinBadge1.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeSync.Layout(gtx, th) }),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinBadge2.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeUpd.Layout(gtx, th) }),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinBadge3.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BadgeProc.Layout(gtx, th) }),
					)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 2: Spinner Button
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "2. Spinner in Buttons (spinner-button.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinBtn1.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnLoad.Layout(gtx, th) }),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinBtn2.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnWait.Layout(gtx, th) }),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinBtn3.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnProc.Layout(gtx, th) }),
					)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 3: Spinner Custom
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "3. Custom Spinner (spinner-custom.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.CustomSpin.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 4: Spinner Demo
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "4. Item Spinner Row (spinner-demo.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoSpin.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, "Processing payment...")
					lbl.Color = th.Colors.Foreground
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, "$100.00")
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 5: Spinner Empty
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "5. Empty State Spinner (spinner-empty.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.EmptySpin.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Processing your request")
					lbl.Color = th.Colors.Foreground
					lbl.Font.Weight = font.SemiBold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, "Please wait while we process your request. Do not refresh the page.")
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.EmptyBtn.Layout(gtx, th) }),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 6: Spinner Sizes
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "6. Spinner Sizes (spinner-size.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinSizeSM.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinSizeMD.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinSizeLG.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SpinSizeXL.Layout(gtx, th) }),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
