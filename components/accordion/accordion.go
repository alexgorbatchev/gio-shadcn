/*
Package accordion provides an expandable accordion component for gio-shadcn applications.

Accordions display collapsible content panels following
shadcn/ui design principles.
*/
package accordion

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Item represents a single expandable item in an accordion.
type Item struct {
	Title     string
	Content   string
	Expanded  bool
	clickable *widget.Clickable
}

// NewItem creates a new Accordion Item.
func NewItem(title, content string, expanded bool) *Item {
	return &Item{
		Title:     title,
		Content:   content,
		Expanded:  expanded,
		clickable: new(widget.Clickable),
	}
}

// Accordion represents an accordion container component.
type Accordion struct {
	Items   []*Item
	Classes string
}

// Config represents configuration for creating an Accordion.
type Config struct {
	Items   []*Item
	Classes string
}

// New creates a new Accordion component with the given configuration.
func New(config Config) *Accordion {
	return &Accordion{
		Items:   config.Items,
		Classes: config.Classes,
	}
}

// Layout renders the list of accordion items and manages expanded/collapsed states.
func (a *Accordion) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	children := make([]layout.FlexChild, 0, len(a.Items)*2)

	for i, item := range a.Items {
		item := item // capture loop variable

		// Handle trigger click
		if item.clickable.Clicked(gtx) {
			item.Expanded = !item.Expanded
		}

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return a.layoutItem(gtx, th, item)
		}))

		if i < len(a.Items)-1 {
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
			}))
		}
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

func (a *Accordion) layoutItem(gtx layout.Context, th *theme.Theme, item *Item) layout.Dimensions {
	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(a.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := material.NewTheme()

	return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space3,
			Bottom: th.Spacing.Space3,
			Left:   th.Spacing.Space4,
			Right:  th.Spacing.Space4,
		}

		itemDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// Header Trigger Row
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(mTheme, th.Typography.FontSizeBase, item.Title)
							lbl.Color = th.Colors.Foreground
							lbl.Font.Weight = font.Medium
							return lbl.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							symbol := "+"
							if item.Expanded {
								symbol = "-"
							}
							lbl := material.Label(mTheme, th.Typography.FontSizeBase, symbol)
							lbl.Color = th.Colors.MutedFg
							lbl.Font.Weight = font.Bold
							return lbl.Layout(gtx)
						}),
					)
				}),

				// Collapsible Content
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if !item.Expanded {
						return layout.Dimensions{}
					}
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(mTheme, th.Typography.FontSizeSM, item.Content)
							lbl.Color = th.Colors.MutedFg
							lbl.Alignment = text.Start
							return lbl.Layout(gtx)
						}),
					)
				}),
			)
		})

		rect := image.Rectangle{Max: itemDims.Size}
		radius := gtx.Dp(th.Radius.RadiusMD)
		rr := clip.UniformRRect(rect, radius)

		paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

		stroke := clip.Stroke{
			Path:  rr.Path(gtx.Ops),
			Width: 1.0,
		}
		paint.FillShape(gtx.Ops, borderColor, stroke.Op())

		return itemDims
	})
}
