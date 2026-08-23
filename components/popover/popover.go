/*
Package popover provides an anchored popover card component for gio-shadcn applications.

Popovers display floating content panels anchored to trigger elements following
shadcn/ui design principles.
*/
package popover

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

// Popover represents an anchored popover card.
type Popover struct {
	Open        bool
	Title       string
	Description string
	Classes     string
}

// Config represents configuration for creating a Popover.
type Config struct {
	Open        bool
	Title       string
	Description string
	Classes     string
}

// New creates a new Popover card component.
func New(config Config) *Popover {
	return &Popover{
		Open:        config.Open,
		Title:       config.Title,
		Description: config.Description,
		Classes:     config.Classes,
	}
}

// Layout renders the popover panel when Open == true.
func (p *Popover) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if !p.Open {
		return layout.Dimensions{}
	}

	if th == nil {
		th = theme.New()
	}

	bgColor := th.Colors.Popover
	fgColor := th.Colors.PopoverFg
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(p.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space4,
		Bottom: th.Spacing.Space4,
		Left:   th.Spacing.Space4,
		Right:  th.Spacing.Space4,
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if p.Title == "" {
						return layout.Dimensions{}
					}
					lbl := material.Label(mTheme, th.Typography.FontSizeBase, p.Title)
					lbl.Color = fgColor
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if p.Title != "" && p.Description != "" {
						return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
					}
					return layout.Dimensions{}
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if p.Description == "" {
						return layout.Dimensions{}
					}
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, p.Description)
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
			)
		})
	}

	contentDims := renderContent(gtxContent)
	popoverSize := contentDims.Size

	dims := layout.Stack{}.Layout(gtx,
		// Popover background drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: popoverSize}
			radius := gtx.Dp(th.Radius.RadiusMD)
			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: popoverSize}
		}),

		// Content drawn ON TOP of background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
