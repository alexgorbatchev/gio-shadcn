/*
Package collapsible provides a simple expandable container component for gio-shadcn applications.

Collapsibles toggle content visibility following
shadcn/ui design principles.
*/
package collapsible

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Collapsible represents an open/close container component.
type Collapsible struct {
	clickable *widget.Clickable

	Open          bool
	Title         string
	Content       string
	Classes       string
	CustomHeader  layout.Widget
	ContentWidget layout.Widget
	Borderless    bool
}

// Config represents configuration for creating a Collapsible component.
type Config struct {
	Open          bool
	Title         string
	Content       string
	Classes       string
	CustomHeader  layout.Widget
	ContentWidget layout.Widget
	Borderless    bool
}

// New creates a new Collapsible component.
func New(config Config) *Collapsible {
	return &Collapsible{
		clickable:     new(widget.Clickable),
		Open:          config.Open,
		Title:         config.Title,
		Content:       config.Content,
		Classes:       config.Classes,
		CustomHeader:  config.CustomHeader,
		ContentWidget: config.ContentWidget,
		Borderless:    config.Borderless,
	}
}

// Layout renders the collapsible header trigger and content.
func (c *Collapsible) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if c.clickable.Clicked(gtx) {
		c.Open = !c.Open
	}

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(c.Classes)
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
				// Header Row
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							if c.CustomHeader != nil {
								return c.CustomHeader(gtx)
							}
							lbl := material.Label(mTheme, th.Typography.FontSizeBase, c.Title)
							lbl.Color = th.Colors.Foreground
							lbl.Font.Weight = font.Medium
							return lbl.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if c.Open {
								return lucide.ChevronDown.LayoutSize(gtx, unit.Dp(16), th.Colors.MutedFg)
							}
							return lucide.ChevronRight.LayoutSize(gtx, unit.Dp(16), th.Colors.MutedFg)
						}),
					)
				}),
				// Content Panel
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if !c.Open {
						return layout.Dimensions{}
					}
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if c.ContentWidget != nil {
								return c.ContentWidget(gtx)
							}
							lbl := material.Label(mTheme, th.Typography.FontSizeSM, c.Content)
							lbl.Color = th.Colors.MutedFg
							lbl.Alignment = text.Start
							return lbl.Layout(gtx)
						}),
					)
				}),
			)
		})
	}

	contentDims := renderContent(gtxContent)
	itemSize := contentDims.Size

	dims := c.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				if !c.Borderless {
					rect := image.Rectangle{Max: itemSize}
					radius := gtx.Dp(th.Radius.RadiusMD)
					theme.DrawRRectBackground(gtx, rect, radius, bgColor)
					rr := clip.UniformRRect(rect, radius)
					theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)
				}
				return layout.Dimensions{Size: itemSize}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return renderContent(gtx)
			}),
		)
	})

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
