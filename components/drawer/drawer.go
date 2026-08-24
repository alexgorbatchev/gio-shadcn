/*
Package drawer provides a bottom sheet overlay panel component for gio-shadcn applications.

Drawers display slide-up bottom sheets following
shadcn/ui design principles.
*/
package drawer

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Drawer represents a bottom sheet overlay component.
type Drawer struct {
	Title       string
	Description string
	Open        bool
	Height      unit.Dp
	Classes     string
	Content     layout.Widget

	OnClose       func()
	closeBtn      *button.Button
	backdropClick widget.Clickable
}

// Config represents configuration for creating a Drawer.
type Config struct {
	Title       string
	Description string
	Open        bool
	Height      unit.Dp
	Classes     string
	Content     layout.Widget
	OnClose     func()
}

// New creates a new Drawer bottom sheet.
func New(config Config) *Drawer {
	h := config.Height
	if h <= 0 {
		h = unit.Dp(260)
	}

	d := &Drawer{
		Title:       config.Title,
		Description: config.Description,
		Open:        config.Open,
		Height:      h,
		Classes:     config.Classes,
		Content:     config.Content,
		OnClose:     config.OnClose,
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

	return d
}

// Layout renders the dark backdrop overlay and bottom drawer panel when Open == true.
func (d *Drawer) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if !d.Open {
		return layout.Dimensions{}
	}

	if th == nil {
		th = theme.New()
	}

	// Process backdrop click to close when clicking outside
	if d.backdropClick.Clicked(gtx) {
		d.Open = false
		if d.OnClose != nil {
			d.OnClose()
		}
	}

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(d.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	drawerHeightPx := gtx.Dp(d.Height)

	dims := layout.Stack{}.Layout(gtx,
		// Dark backdrop overlay across full window that intercepts outside clicks
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return d.backdropClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{Max: gtx.Constraints.Max}
				backdropColor := color.NRGBA{R: 0, G: 0, B: 0, A: 160}
				theme.DrawRRectBackground(gtx, rect, 0, backdropColor)
				return layout.Dimensions{Size: gtx.Constraints.Max}
			})
		}),

		// Drawer panel aligned to the SOUTH (bottom edge) of the viewport
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtxDrawer := gtx
			gtxDrawer.Constraints.Min.Y = 0

			return layout.S.Layout(gtxDrawer, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				gtx.Constraints.Min.Y = drawerHeightPx
				gtx.Constraints.Max.Y = drawerHeightPx

				padding := layout.Inset{
					Top:    th.Spacing.Space4,
					Bottom: th.Spacing.Space6,
					Left:   th.Spacing.Space6,
					Right:  th.Spacing.Space6,
				}

				drawerSize := image.Pt(gtx.Constraints.Max.X, drawerHeightPx)

				return layout.Stack{}.Layout(gtx,
					// Drawer background drawn FIRST
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						rect := image.Rectangle{Max: drawerSize}
						radiusPx := gtx.Dp(th.Radius.RadiusLG)

						theme.DrawRRectBackground(gtx, rect, radiusPx, bgColor)

						rr := clip.UniformRRect(rect, radiusPx)
						theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

						return layout.Dimensions{Size: drawerSize}
					}),

					// Drawer content drawn ON TOP of background
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
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
									return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Real-Time System & Hardware Metrics")
											lbl.Color = th.Colors.Foreground
											lbl.Font.Weight = font.SemiBold
											return lbl.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											lbl := material.Label(mTheme, th.Typography.FontSizeSM, "CPU Usage: 2.1% | Physical RAM: 189.5 MB | Metal GPU Frame Rate: 120 FPS | Buffer Latency: 0.7 ms")
											lbl.Color = th.Colors.MutedFg
											return lbl.Layout(gtx)
										}),
									)
								}),
							)
						})
					}),
				)
			})
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
