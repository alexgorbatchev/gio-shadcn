/*
Package toast provides a notification banner component for gio-shadcn applications.

Toasts display brief notification banners following
shadcn/ui design principles.
*/
package toast

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Toast represents a floating toast notification item.
type Toast struct {
	Title       string
	Description string
	Visible     bool
	Variant     theme.Variant
	Classes     string
}

// Config represents configuration for creating a Toast.
type Config struct {
	Title       string
	Description string
	Visible     bool
	Variant     theme.Variant
	Classes     string
}

// New creates a new Toast notification.
func New(config Config) *Toast {
	v := config.Variant
	if v == "" {
		v = theme.VariantDefault
	}
	return &Toast{
		Title:       config.Title,
		Description: config.Description,
		Visible:     config.Visible,
		Variant:     v,
		Classes:     config.Classes,
	}
}

// Layout renders the toast notification banner if Visible == true.
func (t *Toast) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if !t.Visible {
		return layout.Dimensions{}
	}

	if th == nil {
		th = theme.New()
	}

	bgColor := th.Colors.Card
	fgColor := th.Colors.CardFg
	borderColor := th.Colors.Border

	if t.Variant == theme.VariantDestructive {
		bgColor = th.Colors.Destructive
		fgColor = th.Colors.DestructiveFg
		borderColor = th.Colors.Destructive
	}

	styles := utils.ParseClasses(t.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space3,
		Bottom: th.Spacing.Space3,
		Left:   th.Spacing.Space4,
		Right:  th.Spacing.Space4,
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if t.Title == "" {
						return layout.Dimensions{}
					}
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, t.Title)
					lbl.Color = fgColor
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if t.Title != "" && t.Description != "" {
						return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx)
					}
					return layout.Dimensions{}
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if t.Description == "" {
						return layout.Dimensions{}
					}
					lbl := material.Label(mTheme, th.Typography.FontSizeXS, t.Description)
					lbl.Color = th.Colors.MutedFg
					if t.Variant == theme.VariantDestructive {
						lbl.Color = fgColor
					}
					return lbl.Layout(gtx)
				}),
			)
		})
	}

	contentDims := renderContent(gtxContent)
	toastSize := contentDims.Size

	dims := layout.Stack{}.Layout(gtx,
		// Background & Border drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: toastSize}
			radius := gtx.Dp(th.Radius.RadiusMD)

			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: toastSize}
		}),

		// Toast text content drawn ON TOP
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
