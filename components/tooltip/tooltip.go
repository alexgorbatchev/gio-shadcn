/*
Package tooltip provides a tooltip hover label component for gio-shadcn applications.

Tooltips display informative text popups following
shadcn/ui design principles.
*/
package tooltip

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Tooltip represents a tooltip label component.
type Tooltip struct {
	Text    string
	Classes string
}

// Config represents configuration for creating a Tooltip.
type Config struct {
	Text    string
	Classes string
}

// New creates a new Tooltip component with the given configuration.
func New(config Config) *Tooltip {
	return &Tooltip{
		Text:    config.Text,
		Classes: config.Classes,
	}
}

// Layout renders the tooltip popover box.
func (t *Tooltip) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	bgColor := th.Colors.Primary
	fgColor := th.Colors.PrimaryFg

	styles := utils.ParseClasses(t.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space1,
		Bottom: th.Spacing.Space1,
		Left:   th.Spacing.Space3,
		Right:  th.Spacing.Space3,
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeXS, t.Text)
			lbl.Color = fgColor
			lbl.Alignment = text.Middle
			return lbl.Layout(gtx)
		})
	}

	contentDims := renderContent(gtxContent)
	tipSize := contentDims.Size

	dims := layout.Stack{}.Layout(gtx,
		// Background drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: tipSize}
			radius := gtx.Dp(th.Radius.RadiusSM)
			theme.DrawRRectBackground(gtx, rect, radius, bgColor)
			return layout.Dimensions{Size: tipSize}
		}),

		// Text label drawn ON TOP
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
