/*
Package drawer provides a bottom sheet overlay component for gio-shadcn applications.

Drawers display slide-up contextual panels anchored to the screen edge
following shadcn/ui design principles with smooth physics transitions.
*/
package drawer

import (
	"image"

	"gioui.org/font"
	"gioui.org/gesture"
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

// Drawer represents a bottom sheet overlay component.
type Drawer struct {
	Title         string
	Description   string
	Open          bool
	Height        unit.Dp
	Classes       string
	Content       layout.Widget
	TriggerButton *button.Button
	Trigger       layout.Widget
	Modal         theme.Modal

	OnClose  func()
	closeBtn *button.Button
	dimmer   *theme.Dimmer
	spring   *spring.Spring
	panel    gesture.Click
}

// Config represents configuration for creating a Drawer.
type Config struct {
	TriggerText   string
	TriggerButton *button.Button
	Trigger       layout.Widget
	Title         string
	Description   string
	Open          bool
	Height        unit.Dp
	Classes       string
	Content       layout.Widget
	OnClose       func()
}

// New creates a new Drawer bottom sheet.
func New(config Config) *Drawer {
	h := config.Height
	if h <= 0 {
		h = unit.Dp(280)
	}

	initVal := 0.0
	if config.Open {
		initVal = 1.0
	}

	d := &Drawer{
		Title:         config.Title,
		Description:   config.Description,
		Open:          config.Open,
		Height:        h,
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

	if config.TriggerText != "" && d.TriggerButton == nil {
		d.TriggerButton = button.New(button.Config{
			Text:    config.TriggerText,
			Variant: theme.VariantOutline,
			OnClick: func() {
				d.Open = true
			},
		})
	}

	d.closeBtn = button.New(button.Config{
		Variant: theme.VariantGhost,
		Size:    theme.SizeIcon,
		Icon:    lucide.X,
		OnClick: func() {
			d.Open = false
			if d.OnClose != nil {
				d.OnClose()
			}
		},
	})
	if d.TriggerButton != nil {
		d.Modal.ReturnFocus = d.TriggerButton.Focus
	}

	return d
}

// Layout renders the trigger, and when Open == true or animating, queues the full-screen bottom drawer overlay.
func (d *Drawer) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}
	defer func() { paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops) }()

	var triggerDims layout.Dimensions
	if d.TriggerButton != nil {
		triggerDims = d.TriggerButton.Layout(gtx, th)
	} else if d.Trigger != nil {
		triggerDims = d.Trigger(gtx)
	}

	if d.spring == nil {
		initVal := 0.0
		if d.Open {
			initVal = 1.0
		}
		d.spring = spring.New(initVal, initVal, spring.Options{
			Stiffness: 1200.0,
			Damping:   69.0,
		})
	}

	if !d.Open && d.spring.Settled(0.001) && d.spring.Value() <= 0.001 {
		return triggerDims
	}

	th.AddModal(&d.Modal, &d.Open, d.closeBtn.Focus, d.OnClose, func(gtx layout.Context) layout.Dimensions {
		return d.renderOverlay(gtx, th)
	})
	if d.TriggerButton != nil || d.Trigger != nil {
		return triggerDims
	}
	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func (d *Drawer) renderOverlay(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	gtx.Constraints = layout.Exact(gtx.Constraints.Max)
	defer func() { paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops) }()
	for {
		if _, ok := d.panel.Update(gtx.Source); !ok {
			break
		}
	}
	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	if d.spring == nil {
		initVal := 0.0
		if d.Open {
			initVal = 1.0
		}
		d.spring = spring.New(initVal, initVal, spring.Options{
			Stiffness: 1200.0,
			Damping:   69.0,
		})
	}

	target := 0.0
	if d.Open {
		target = 1.0
	}
	d.spring.SetTarget(target)
	d.spring.Tick(60.0)
	progress := float32(d.spring.Value())
	if progress < 0 {
		progress = 0
	} else if progress > 1 {
		progress = 1
	}

	if !d.spring.Settled(0.001) {
		d.Modal.Invalidate()
	}

	if progress <= 0.001 && !d.Open {
		return layout.Dimensions{}
	}

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(d.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	drawerHeightPx := min(gtx.Dp(d.Height), gtx.Constraints.Max.Y)
	windowSize := gtx.Constraints.Max

	slideOffset := int(float32(drawerHeightPx) * (1.0 - progress))

	macro := op.Record(gtx.Ops)
	layout.Stack{}.Layout(gtx,
		// Full window dimmer backdrop with animated fade & blur
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			alpha := uint8(float32(160) * progress)
			return d.dimmer.LayoutWithAlpha(gtx, th, alpha, func() {
				d.Open = false
				if d.OnClose != nil {
					d.OnClose()
				}
			})
		}),

		// Drawer panel anchored to the SOUTH (bottom edge) across the FULL window width with slide animation
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtxDrawer := gtx
			gtxDrawer.Constraints = layout.Exact(windowSize)

			return layout.S.Layout(gtxDrawer, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(image.Pt(windowSize.X, drawerHeightPx))

				// Apply physics slide offset
				op.Offset(image.Pt(0, slideOffset)).Add(gtx.Ops)

				padding := layout.Inset{
					Top:    th.Spacing.Space4,
					Bottom: th.Spacing.Space6,
					Left:   th.Spacing.Space6,
					Right:  th.Spacing.Space6,
				}

				drawerSize := image.Pt(windowSize.X, drawerHeightPx)

				return layout.Stack{}.Layout(gtx,
					// Drawer background drawn FIRST with top rounded corners
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						rect := image.Rectangle{Max: drawerSize}
						radiusPx := max(0, min(gtx.Dp(th.Radius.RadiusLG), rect.Dx()/2, rect.Dy()/2))

						var rr clip.RRect
						rr.Rect = rect
						rr.NW = radiusPx
						rr.NE = radiusPx
						rr.SW = 0
						rr.SE = 0

						theme.DrawCornerBackground(gtx, rr, bgColor)

						theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

						return layout.Dimensions{Size: drawerSize}
					}),

					// Drawer content
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						defer clip.Rect(image.Rectangle{Max: drawerSize}).Push(gtx.Ops).Pop()
						d.panel.Add(gtx.Ops)
						return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								// Handle indicator bar
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										handleSize := image.Pt(gtx.Dp(unit.Dp(36)), gtx.Dp(unit.Dp(4)))
										handleRect := image.Rectangle{Max: handleSize}
										theme.DrawRRectBackground(gtx, handleRect, gtx.Dp(unit.Dp(2)), th.Colors.MutedFg)
										return layout.Dimensions{Size: handleSize}
									})
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx)
								}),
								// Header Row with Title and Close Button
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
										layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
											lbl := material.Label(mTheme, th.Typography.FontSizeXL, d.Title)
											lbl.Color = th.Colors.Foreground
											lbl.Font.Weight = font.Bold
											return lbl.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return d.closeBtn.Layout(gtx, th)
										}),
									)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
								}),
								// Description Body
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									lbl := material.Label(mTheme, th.Typography.FontSizeSM, d.Description)
									lbl.Color = th.Colors.MutedFg
									return lbl.Layout(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
								}),
								// Custom or Illustrated Content Body
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									if d.Content != nil {
										return d.Content(gtx)
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
