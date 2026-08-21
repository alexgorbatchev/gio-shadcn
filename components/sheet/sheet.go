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
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
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

	OnClose  func()
	closeBtn *button.Button
}

// Config represents configuration for creating a Sheet.
type Config struct {
	Title       string
	Description string
	Open        bool
	Width       unit.Dp
	Classes     string
	OnClose     func()
}

// New creates a new Sheet side drawer.
func New(config Config) *Sheet {
	w := config.Width
	if w <= 0 {
		w = unit.Dp(300)
	}

	s := &Sheet{
		Title:       config.Title,
		Description: config.Description,
		Open:        config.Open,
		Width:       w,
		Classes:     config.Classes,
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

	// Dark backdrop overlay
	backdropColor := color.NRGBA{R: 0, G: 0, B: 0, A: 160}
	paint.FillShape(gtx.Ops, backdropColor, clip.Rect{Max: gtx.Constraints.Max}.Op())

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := material.NewTheme()
	sheetWidthPx := gtx.Dp(s.Width)

	// Position side drawer on the right edge
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

		sheetDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
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
				// Description
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, s.Description)
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
			)
		})

		rect := image.Rectangle{Max: sheetDims.Size}
		paint.FillShape(gtx.Ops, bgColor, clip.Rect(rect).Op())

		// Left border line
		borderRect := image.Rectangle{Max: image.Pt(1, sheetDims.Size.Y)}
		paint.FillShape(gtx.Ops, borderColor, clip.Rect(borderRect).Op())

		return sheetDims
	})
}
