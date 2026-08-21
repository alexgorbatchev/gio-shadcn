/*
Package carousel provides an image/card slider carousel component for gio-shadcn applications.

Carousels display slidable media or card items following
shadcn/ui design principles.
*/
package carousel

import (
	"fmt"
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

// Carousel represents an image/card slider carousel component.
type Carousel struct {
	Items       []layout.Widget
	ActiveIndex int

	prevBtn *widget.Clickable
	nextBtn *widget.Clickable
}

// Config represents configuration for creating a Carousel.
type Config struct {
	Items       []layout.Widget
	ActiveIndex int
}

// New creates a new Carousel slider.
func New(config Config) *Carousel {
	idx := config.ActiveIndex
	if idx < 0 || idx >= len(config.Items) {
		idx = 0
	}
	return &Carousel{
		Items:       config.Items,
		ActiveIndex: idx,
		prevBtn:     new(widget.Clickable),
		nextBtn:     new(widget.Clickable),
	}
}

// Layout renders the current slide and navigation controls.
func (c *Carousel) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if len(c.Items) == 0 {
		return layout.Dimensions{}
	}

	if th == nil {
		th = theme.New()
	}

	if c.prevBtn.Clicked(gtx) && c.ActiveIndex > 0 {
		c.ActiveIndex--
	}

	if c.nextBtn.Clicked(gtx) && c.ActiveIndex < len(c.Items)-1 {
		c.ActiveIndex++
	}

	mTheme := material.NewTheme()

	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		// Slide Content
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if c.ActiveIndex >= 0 && c.ActiveIndex < len(c.Items) {
				return c.Items[c.ActiveIndex](gtx)
			}
			return layout.Dimensions{}
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
		}),
		// Navigation Bar Row
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return c.prevBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(mTheme, th.Typography.FontSizeSM, "‹ Prev")
						lbl.Color = th.Colors.Foreground
						if c.ActiveIndex <= 0 {
							lbl.Color = th.Colors.MutedFg
						}
						return lbl.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeXS, fmt.Sprintf("%d / %d", c.ActiveIndex+1, len(c.Items)))
					lbl.Color = th.Colors.MutedFg
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return c.nextBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(mTheme, th.Typography.FontSizeSM, "Next ›")
						lbl.Color = th.Colors.Foreground
						if c.ActiveIndex >= len(c.Items)-1 {
							lbl.Color = th.Colors.MutedFg
						}
						return lbl.Layout(gtx)
					})
				}),
			)
		}),
	)
}

func _(i image.Point) {} // unused image guard
