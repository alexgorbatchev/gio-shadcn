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
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Collapsible represents an open/close container component.
type Collapsible struct {
	clickable *widget.Clickable

	Open    bool
	Title   string
	Content string
	Classes string
}

// Config represents configuration for creating a Collapsible component.
type Config struct {
	Open    bool
	Title   string
	Content string
	Classes string
}

// New creates a new Collapsible component.
func New(config Config) *Collapsible {
	return &Collapsible{
		clickable: new(widget.Clickable),
		Open:      config.Open,
		Title:     config.Title,
		Content:   config.Content,
		Classes:   config.Classes,
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

	return c.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space3,
			Bottom: th.Spacing.Space3,
			Left:   th.Spacing.Space4,
			Right:  th.Spacing.Space4,
		}

		macro := op.Record(gtx.Ops)
		dims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(mTheme, th.Typography.FontSizeBase, c.Title)
							lbl.Color = th.Colors.Foreground
							lbl.Font.Weight = font.Medium
							return lbl.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							symbol := "▼"
							if !c.Open {
								symbol = "▶"
							}
							lbl := material.Label(mTheme, th.Typography.FontSizeSM, symbol)
							lbl.Color = th.Colors.MutedFg
							return lbl.Layout(gtx)
						}),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if !c.Open {
						return layout.Dimensions{}
					}
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(mTheme, th.Typography.FontSizeSM, c.Content)
							lbl.Color = th.Colors.MutedFg
							lbl.Alignment = text.Start
							return lbl.Layout(gtx)
						}),
					)
				}),
			)
		})
		callOp := macro.Stop()

		rect := image.Rectangle{Max: dims.Size}
		radius := gtx.Dp(th.Radius.RadiusMD)

		theme.DrawRRectBackground(gtx, rect, radius, bgColor)

		rr := clip.UniformRRect(rect, radius)
		theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

		callOp.Add(gtx.Ops)

		return dims
	})
}
