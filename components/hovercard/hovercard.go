/*
Package hovercard provides a hover preview card component for gio-shadcn applications.

HoverCards display rich preview content on hover following
shadcn/ui design principles.
*/
package hovercard

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

// HoverCard represents a hover preview card component.
type HoverCard struct {
	Hovered     bool
	Title       string
	Description string
	Classes     string
}

// Config represents configuration for creating a HoverCard.
type Config struct {
	Hovered     bool
	Title       string
	Description string
	Classes     string
}

// New creates a new HoverCard.
func New(config Config) *HoverCard {
	return &HoverCard{
		Hovered:     config.Hovered,
		Title:       config.Title,
		Description: config.Description,
		Classes:     config.Classes,
	}
}

// Layout renders the hover preview card with background drawn before text using layout.Stack.
func (h *HoverCard) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if !h.Hovered {
		return layout.Dimensions{}
	}

	if th == nil {
		th = theme.New()
	}

	bgColor := th.Colors.Card
	fgColor := th.Colors.CardFg
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(h.Classes)
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
					if h.Title == "" {
						return layout.Dimensions{}
					}
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, h.Title)
					lbl.Color = fgColor
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if h.Title != "" && h.Description != "" {
						return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx)
					}
					return layout.Dimensions{}
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if h.Description == "" {
						return layout.Dimensions{}
					}
					lbl := material.Label(mTheme, th.Typography.FontSizeXS, h.Description)
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
			)
		})
	}

	contentDims := renderContent(gtxContent)
	cardSize := contentDims.Size

	dims := layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: cardSize}
			radius := gtx.Dp(th.Radius.RadiusMD)

			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: cardSize}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)

	// Reset active GPU paint color state
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
