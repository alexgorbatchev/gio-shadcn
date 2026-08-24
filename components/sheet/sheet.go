/*
Package sheet provides a side drawer overlay panel component for gio-shadcn applications.

Sheets display slide-over side panels following
shadcn/ui design principles.
*/
package sheet

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Sheet represents a side drawer panel component.
type Sheet struct {
	Title       string
	Description string
	Open        bool
	Width       unit.Dp
	Classes     string
	Content     layout.Widget

	OnClose       func()
	closeBtn      *button.Button
	backdropClick widget.Clickable
}

// Config represents configuration for creating a Sheet.
type Config struct {
	Title       string
	Description string
	Open        bool
	Width       unit.Dp
	Classes     string
	Content     layout.Widget
	OnClose     func()
}

// New creates a new Sheet side drawer.
func New(config Config) *Sheet {
	w := config.Width
	if w <= 0 {
		w = unit.Dp(320)
	}

	s := &Sheet{
		Title:       config.Title,
		Description: config.Description,
		Open:        config.Open,
		Width:       w,
		Classes:     config.Classes,
		Content:     config.Content,
		OnClose:     config.OnClose,
	}

	s.closeBtn = button.New(button.Config{
		Text:    "✕",
		Variant: theme.VariantGhost,
		Size:    theme.SizeSM,
		OnClick: func() {
			s.Open = false
			if s.OnClose != nil {
				s.OnClose()
			}
		},
	})

	return s
}

// Layout renders the dark backdrop overlay and side drawer panel when Open == true.
func (s *Sheet) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if !s.Open {
		return layout.Dimensions{}
	}

	if th == nil {
		th = theme.New()
	}

	// Process backdrop click to close when clicking outside
	if s.backdropClick.Clicked(gtx) {
		s.Open = false
		if s.OnClose != nil {
			s.OnClose()
		}
	}

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	sheetWidthPx := gtx.Dp(s.Width)

	dims := layout.Stack{}.Layout(gtx,
		// Dark backdrop overlay across full window that intercepts outside clicks
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return s.backdropClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{Max: gtx.Constraints.Max}
				backdropColor := color.NRGBA{R: 0, G: 0, B: 0, A: 160}
				theme.DrawRRectBackground(gtx, rect, 0, backdropColor)
				return layout.Dimensions{Size: gtx.Constraints.Max}
			})
		}),

		// Side Sheet panel positioned on the right edge
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = sheetWidthPx
				gtx.Constraints.Max.X = sheetWidthPx
				gtx.Constraints.Min.Y = gtx.Constraints.Max.Y

				padding := layout.Inset{
					Top:    th.Spacing.Space6,
					Bottom: th.Spacing.Space6,
					Left:   th.Spacing.Space6,
					Right:  th.Spacing.Space6,
				}

				sheetSize := image.Pt(sheetWidthPx, gtx.Constraints.Max.Y)

				return layout.Stack{}.Layout(gtx,
					// Sheet background drawn FIRST
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						rect := image.Rectangle{Max: sheetSize}
						theme.DrawRRectBackground(gtx, rect, 0, bgColor)

						// Left border line
						borderRect := image.Rectangle{Max: image.Pt(1, sheetSize.Y)}
						theme.DrawRRectBackground(gtx, borderRect, 0, borderColor)

						return layout.Dimensions{Size: sheetSize}
					}),

					// Sheet content drawn ON TOP of background
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
									return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
								}),
								// Custom or Illustrated Content Body
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									if s.Content != nil {
										return s.Content(gtx)
									}
									return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											lbl := material.Label(mTheme, th.Typography.FontSizeBase, "Track Audio Metadata & Details")
											lbl.Color = th.Colors.Foreground
											lbl.Font.Weight = font.SemiBold
											return lbl.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											lbl := material.Label(mTheme, th.Typography.FontSizeSM, "Format: FLAC 24-bit / 96kHz\nBPM: 128.00\nHarmonic Key: 8A\nChannels: Stereo\nDuration: 06:42")
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
