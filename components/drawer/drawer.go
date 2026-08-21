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
	"gioui.org/widget/material"
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

	OnClose  func()
	closeBtn *button.Button
}

// Config represents configuration for creating a Drawer.
type Config struct {
	Title       string
	Description string
	Open        bool
	Height      unit.Dp
	Classes     string
	OnClose     func()
}

// New creates a new Drawer bottom sheet.
func New(config Config) *Drawer {
	h := config.Height
	if h <= 0 {
		h = unit.Dp(250)
	}

	d := &Drawer{
		Title:       config.Title,
		Description: config.Description,
		Open:        config.Open,
		Height:      h,
		Classes:     config.Classes,
		OnClose:     config.OnClose,
	}

	d.closeBtn = button.New(button.Config{
		Text:    "✕",
		Variant: theme.VariantGhost,
		Size:    theme.SizeSM,
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

	// Dark backdrop overlay
	backdropColor := color.NRGBA{R: 0, G: 0, B: 0, A: 160}
	paint.FillShape(gtx.Ops, backdropColor, clip.Rect{Max: gtx.Constraints.Max}.Op())

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(d.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := material.NewTheme()
	drawerHeightPx := gtx.Dp(d.Height)

	// Position bottom drawer on the bottom edge
	return layout.S.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		gtx.Constraints.Min.Y = drawerHeightPx
		gtx.Constraints.Max.Y = drawerHeightPx

		padding := layout.Inset{
			Top:    th.Spacing.Space4,
			Bottom: th.Spacing.Space6,
			Left:   th.Spacing.Space6,
			Right:  th.Spacing.Space6,
		}

		drawerDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// Handle indicator bar
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						handleSize := image.Pt(gtx.Dp(unit.Dp(36)), gtx.Dp(unit.Dp(4)))
						handleRect := image.Rectangle{Max: handleSize}
						rrHandle := clip.UniformRRect(handleRect, gtx.Dp(unit.Dp(2)))
						paint.FillShape(gtx.Ops, th.Colors.MutedFg, rrHandle.Op(gtx.Ops))
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
				// Description
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, d.Description)
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
			)
		})

		rect := image.Rectangle{Max: drawerDims.Size}
		radius := gtx.Dp(th.Radius.RadiusLG)
		rr := clip.UniformRRect(rect, radius)

		paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

		stroke := clip.Stroke{
			Path:  rr.Path(gtx.Ops),
			Width: 1.0,
		}
		paint.FillShape(gtx.Ops, borderColor, stroke.Op())

		return drawerDims
	})
}
