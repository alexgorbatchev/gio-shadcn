package skeleton

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	// 1. Avatar Demo
	AvatarCircle *Skeleton
	AvatarLine1  *Skeleton
	AvatarLine2  *Skeleton

	// 2. Card Demo
	CardWrapper  *card.Card
	CardHeader1  *Skeleton
	CardHeader2  *Skeleton
	CardBody     *Skeleton

	// 3. Main Demo
	DemoCircle *Skeleton
	DemoLine1  *Skeleton
	DemoLine2  *Skeleton

	// 4. Form Demo
	FormLabel1  *Skeleton
	FormInput1  *Skeleton
	FormLabel2  *Skeleton
	FormInput2  *Skeleton
	FormButton  *Skeleton

	// 5. Table Demo
	TableRows [][]*Skeleton

	// 6. Text Demo
	TextLine1 *Skeleton
	TextLine2 *Skeleton
	TextLine3 *Skeleton
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{
		// 1. Avatar Demo
		AvatarCircle: New(Config{Width: unit.Dp(40), Height: unit.Dp(40), Circle: true}),
		AvatarLine1:  New(Config{Width: unit.Dp(150), Height: unit.Dp(16)}),
		AvatarLine2:  New(Config{Width: unit.Dp(100), Height: unit.Dp(16)}),

		// 2. Card Demo
		CardWrapper: card.New(card.Config{Variant: theme.VariantDefault}),
		CardHeader1: New(Config{Width: unit.Dp(180), Height: unit.Dp(16)}),
		CardHeader2: New(Config{Width: unit.Dp(120), Height: unit.Dp(16)}),
		CardBody:    New(Config{Width: unit.Dp(240), Height: unit.Dp(135)}),

		// 3. Main Demo
		DemoCircle: New(Config{Width: unit.Dp(48), Height: unit.Dp(48), Circle: true}),
		DemoLine1:  New(Config{Width: unit.Dp(250), Height: unit.Dp(16)}),
		DemoLine2:  New(Config{Width: unit.Dp(200), Height: unit.Dp(16)}),

		// 4. Form Demo
		FormLabel1: New(Config{Width: unit.Dp(80), Height: unit.Dp(16)}),
		FormInput1: New(Config{Width: unit.Dp(260), Height: unit.Dp(36)}),
		FormLabel2: New(Config{Width: unit.Dp(96), Height: unit.Dp(16)}),
		FormInput2: New(Config{Width: unit.Dp(260), Height: unit.Dp(36)}),
		FormButton: New(Config{Width: unit.Dp(96), Height: unit.Dp(36)}),

		// 6. Text Demo
		TextLine1: New(Config{Width: unit.Dp(260), Height: unit.Dp(16)}),
		TextLine2: New(Config{Width: unit.Dp(260), Height: unit.Dp(16)}),
		TextLine3: New(Config{Width: unit.Dp(195), Height: unit.Dp(16)}),
	}

	// 5. Table Demo
	for i := 0; i < 5; i++ {
		row := []*Skeleton{
			New(Config{Width: unit.Dp(120), Height: unit.Dp(16)}),
			New(Config{Width: unit.Dp(96), Height: unit.Dp(16)}),
			New(Config{Width: unit.Dp(80), Height: unit.Dp(16)}),
		}
		s.TableRows = append(s.TableRows, row)
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
			lbl := material.Label(mTheme, th.Typography.FontSizeXL, "Skeleton Showcase (6 Upstream Demos)")
			lbl.Color = th.Colors.Foreground
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		// Demo 1: Skeleton Avatar
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "1. Skeleton Avatar (skeleton-avatar.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.AvatarCircle.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.AvatarLine1.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.AvatarLine2.Layout(gtx, th) }),
					)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 2: Skeleton Card
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "2. Skeleton Card (skeleton-card.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.CardWrapper.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.CardHeader1.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.CardHeader2.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.CardBody.Layout(gtx, th) }),
				)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 3: Skeleton Demo
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "3. Skeleton Demo (skeleton-demo.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoCircle.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoLine1.Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoLine2.Layout(gtx, th) }),
					)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 4: Skeleton Form
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "4. Skeleton Form (skeleton-form.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.FormLabel1.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.FormInput1.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.FormLabel2.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.FormInput2.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.FormButton.Layout(gtx, th) }),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 5: Skeleton Table
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "5. Skeleton Table (skeleton-table.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			rows := make([]layout.FlexChild, 0, len(s.TableRows)*2)
			for i, r := range s.TableRows {
				row := r
				rows = append(rows, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return row[0].Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return row[1].Layout(gtx, th) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return row[2].Layout(gtx, th) }),
					)
				}))
				if i < len(s.TableRows)-1 {
					rows = append(rows, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
					}))
				}
			}
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rows...)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		// Demo 6: Skeleton Text
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeBase, "6. Skeleton Text (skeleton-text.tsx)")
			lbl.Color = th.Colors.MutedFg
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.TextLine1.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.TextLine2.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.TextLine3.Layout(gtx, th) }),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
