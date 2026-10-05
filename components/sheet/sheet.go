/*
Package sheet provides a side modal sheet/drawer component for gio-shadcn applications.

Sheets extend dialog and drawer semantics by sliding in from the screen edges
following shadcn/ui design principles with smooth physics animations.
*/
package sheet

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
	"github.com/vibrantgio/effects/spring"
)

type Side int

const (
	SideRight Side = iota
	SideLeft
	SideTop
	SideBottom
)

// Sheet represents a side drawer panel component.
type Sheet struct {
	Title         string
	Description   string
	Open          bool
	Width         unit.Dp
	Height        unit.Dp
	Side          Side
	Classes       string
	Content       layout.Widget
	TriggerButton *button.Button
	Trigger       layout.Widget

	OnClose  func()
	closeBtn *button.Button
	dimmer   *theme.Dimmer
	spring   *spring.Spring
}

// Config represents configuration for creating a Sheet.
type Config struct {
	TriggerText   string
	TriggerButton *button.Button
	Trigger       layout.Widget
	Title         string
	Description   string
	Open          bool
	Width         unit.Dp
	Height        unit.Dp
	Side          Side
	Classes       string
	Content       layout.Widget
	OnClose       func()
}

// New creates a new Sheet side drawer.
func New(config Config) *Sheet {
	w := config.Width
	if w <= 0 {
		w = unit.Dp(360)
	}
	h := config.Height
	if h <= 0 {
		h = unit.Dp(300)
	}

	initVal := 0.0
	if config.Open {
		initVal = 1.0
	}

	s := &Sheet{
		Title:         config.Title,
		Description:   config.Description,
		Open:          config.Open,
		Width:         w,
		Height:        h,
		Side:          config.Side,
		Classes:       config.Classes,
		Content:       config.Content,
		TriggerButton: config.TriggerButton,
		Trigger:       config.Trigger,
		OnClose:       config.OnClose,
		dimmer:        theme.NewDimmer(),
		spring: spring.New(initVal, initVal, spring.Options{
			Stiffness: 1200.0,
			Damping:   69.0,
		}),
	}

	if config.TriggerText != "" && s.TriggerButton == nil {
		s.TriggerButton = button.New(button.Config{
			Text:    config.TriggerText,
			Variant: theme.VariantOutline,
			OnClick: func() {
				s.Open = true
			},
		})
	}

	s.closeBtn = button.New(button.Config{
		Variant: theme.VariantGhost,
		Size:    theme.SizeIcon,
		Icon:    lucide.X,
		OnClick: func() {
			s.Open = false
			if s.OnClose != nil {
				s.OnClose()
			}
		},
	})

	return s
}

// Layout renders the trigger, and when Open == true or animating, queues the full-screen side sheet overlay.
func (s *Sheet) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	var triggerDims layout.Dimensions
	if s.TriggerButton != nil {
		triggerDims = s.TriggerButton.Layout(gtx, th)
	} else if s.Trigger != nil {
		triggerDims = s.Trigger(gtx)
	}

	if s.spring == nil {
		initVal := 0.0
		if s.Open {
			initVal = 1.0
		}
		s.spring = spring.New(initVal, initVal, spring.Options{
			Stiffness: 1200.0,
			Damping:   69.0,
		})
	}

	if !s.Open && s.spring.Settled(0.001) && s.spring.Value() <= 0.001 {
		return triggerDims
	}

	th.AddOverlay(func(gtx layout.Context) layout.Dimensions {
		return s.renderOverlay(gtx, th)
	})

	if s.TriggerButton != nil || s.Trigger != nil {
		return triggerDims
	}
	return s.renderOverlay(gtx, th)
}

func (s *Sheet) renderOverlay(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	if s.spring == nil {
		initVal := 0.0
		if s.Open {
			initVal = 1.0
		}
		s.spring = spring.New(initVal, initVal, spring.Options{
			Stiffness: 1200.0,
			Damping:   69.0,
		})
	}

	target := 0.0
	if s.Open {
		target = 1.0
	}
	s.spring.SetTarget(target)
	s.spring.Tick(60.0)
	progress := float32(s.spring.Value())
	if progress < 0 {
		progress = 0
	} else if progress > 1 {
		progress = 1
	}

	if !s.spring.Settled(0.001) {
		gtx.Execute(op.InvalidateCmd{})
	}

	if progress <= 0.001 && !s.Open {
		return layout.Dimensions{}
	}

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	sheetWidthPx := gtx.Dp(s.Width)
	sheetHeightPx := gtx.Dp(s.Height)
	windowSize := gtx.Constraints.Max

	macro := op.Record(gtx.Ops)
	layout.Stack{}.Layout(gtx,
		// 1. Dark backdrop dimmer across full window with animated fade & blur
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			alpha := uint8(float32(160) * progress)
			return s.dimmer.LayoutWithAlpha(gtx, th, alpha, func() {
				s.Open = false
				if s.OnClose != nil {
					s.OnClose()
				}
			})
		}),

		// 2. Side drawer panel aligned to the configured edge of viewport with slide animation
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtxSheet := gtx
			gtxSheet.Constraints = layout.Exact(windowSize)

			var align layout.Direction
			var sheetSize image.Point
			var slideOffset image.Point

			switch s.Side {
			case SideLeft:
				align = layout.W
				sheetSize = image.Pt(sheetWidthPx, windowSize.Y)
				slideOffset = image.Pt(-int(float32(sheetWidthPx)*(1.0-progress)), 0)
			case SideTop:
				align = layout.N
				sheetSize = image.Pt(windowSize.X, sheetHeightPx)
				slideOffset = image.Pt(0, -int(float32(sheetHeightPx)*(1.0-progress)))
			case SideBottom:
				align = layout.S
				sheetSize = image.Pt(windowSize.X, sheetHeightPx)
				slideOffset = image.Pt(0, int(float32(sheetHeightPx)*(1.0-progress)))
			default: // SideRight
				align = layout.E
				sheetSize = image.Pt(sheetWidthPx, windowSize.Y)
				slideOffset = image.Pt(int(float32(sheetWidthPx)*(1.0-progress)), 0)
			}

			return align.Layout(gtxSheet, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(sheetSize)

				// Apply physics slide offset
				op.Offset(slideOffset).Add(gtx.Ops)

				padding := layout.Inset{
					Top:    th.Spacing.Space6,
					Bottom: th.Spacing.Space6,
					Left:   th.Spacing.Space6,
					Right:  th.Spacing.Space6,
				}

				return layout.Stack{}.Layout(gtx,
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						rect := image.Rectangle{Max: sheetSize}
						theme.DrawRRectBackground(gtx, rect, 0, bgColor)

						rr := clip.RRect{Rect: rect}
						theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

						return layout.Dimensions{Size: sheetSize}
					}),

					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								// Header Row with Title and Close Button
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
										layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
											lbl := material.Label(mTheme, th.Typography.FontSizeXL, s.Title)
											lbl.Color = th.Colors.Foreground
											lbl.Font.Weight = font.Bold
											return lbl.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return s.closeBtn.Layout(gtx, th)
										}),
									)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
								}),
								// Description Body
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									lbl := material.Label(mTheme, th.Typography.FontSizeSM, s.Description)
									lbl.Color = th.Colors.MutedFg
									return lbl.Layout(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx)
								}),
								// Custom or Illustrated Content Body
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									if s.Content != nil {
										return s.Content(gtx)
									}
									return layout.Dimensions{}
								}),
							)
						})
					}),
				)
			})
		}),
	)
	callOp := macro.Stop()
	callOp.Add(gtx.Ops)

	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return layout.Dimensions{Size: windowSize}
}
