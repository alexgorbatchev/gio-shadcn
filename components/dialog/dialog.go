/*
Package dialog provides a modal dialog box component for gio-shadcn applications.

Dialogs display modal confirmation windows following
shadcn/ui design principles.
*/
package dialog

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Dialog represents a modal window dialog component.
type Dialog struct {
	Title       string
	Description string
	Open        bool
	ConfirmText string
	CancelText  string
	Classes     string

	OnConfirm func()
	OnCancel  func()

	cancelBtn  *button.Button
	confirmBtn *button.Button
}

// Config represents configuration for creating a Dialog.
type Config struct {
	Title       string
	Description string
	Open        bool
	ConfirmText string
	CancelText  string
	Classes     string

	OnConfirm func()
	OnCancel  func()
}

// New creates a new Dialog component.
func New(config Config) *Dialog {
	confText := config.ConfirmText
	if confText == "" {
		confText = "Confirm"
	}
	cancText := config.CancelText
	if cancText == "" {
		cancText = "Cancel"
	}

	d := &Dialog{
		Title:       config.Title,
		Description: config.Description,
		Open:        config.Open,
		ConfirmText: confText,
		CancelText:  cancText,
		Classes:     config.Classes,
		OnConfirm:   config.OnConfirm,
		OnCancel:    config.OnCancel,
	}

	d.cancelBtn = button.New(button.Config{
		Text:    d.CancelText,
		Variant: theme.VariantOutline,
		OnClick: func() {
			d.Open = false
			if d.OnCancel != nil {
				d.OnCancel()
			}
		},
	})

	d.confirmBtn = button.New(button.Config{
		Text:    d.ConfirmText,
		Variant: theme.VariantDefault,
		OnClick: func() {
			d.Open = false
			if d.OnConfirm != nil {
				d.OnConfirm()
			}
		},
	})

	return d
}

// Layout renders the modal backdrop and dialog window if Open == true.
func (d *Dialog) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if !d.Open {
		return layout.Dimensions{}
	}

	if th == nil {
		th = theme.New()
	}

	// Backdrop dark overlay
	backdropColor := color.NRGBA{R: 0, G: 0, B: 0, A: 160}
	paint.FillShape(gtx.Ops, backdropColor, clip.Rect{Max: gtx.Constraints.Max}.Op())

	// Centered dialog card
	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(d.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := material.NewTheme()

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space6,
			Bottom: th.Spacing.Space6,
			Left:   th.Spacing.Space6,
			Right:  th.Spacing.Space6,
		}

		cardDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// Header Title
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSize2XL, d.Title)
					lbl.Color = th.Colors.Foreground
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
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
					return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx)
				}),
				// Action Buttons Row
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return d.cancelBtn.Layout(gtx, th)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return d.confirmBtn.Layout(gtx, th)
						}),
					)
				}),
			)
		})

		rect := image.Rectangle{Max: cardDims.Size}
		radius := gtx.Dp(th.Radius.RadiusLG)
		rr := clip.UniformRRect(rect, radius)

		paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

		stroke := clip.Stroke{
			Path:  rr.Path(gtx.Ops),
			Width: 1.0,
		}
		paint.FillShape(gtx.Ops, borderColor, stroke.Op())

		return cardDims
	})
}
