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
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

type Type string

const (
	TypeSingle   Type = "single"
	TypeMultiple Type = "multiple"
)

type Item struct {
	Title         string
	Content       string
	Expanded      bool
	Disabled      bool
	Icon          string
	CustomHeader  layout.Widget
	ContentWidget layout.Widget
	clickable     *widget.Clickable
}

type ItemConfig struct {
	Title         string
	Content       string
	Expanded      bool
	Disabled      bool
	Icon          string
	CustomHeader  layout.Widget
	ContentWidget layout.Widget
}

func NewItem(title, content string, expanded bool) *Item {
	return &Item{
		Title:     title,
		Content:   content,
		Expanded:  expanded,
		clickable: new(widget.Clickable),
	}
}

func NewItemConfig(config ItemConfig) *Item {
	return &Item{
		Title:         config.Title,
		Content:       config.Content,
		Expanded:      config.Expanded,
		Disabled:      config.Disabled,
		Icon:          config.Icon,
		CustomHeader:  config.CustomHeader,
		ContentWidget: config.ContentWidget,
		clickable:     new(widget.Clickable),
	}
}

type Accordion struct {
	Type       Type
	Items      []*Item
	Borderless bool
	Classes    string
}

type Config struct {
	Type       Type
	Items      []*Item
	Borderless bool
	Classes    string
}

func New(config Config) *Accordion {
	t := config.Type
	if t == "" {
		t = TypeSingle
	}
	return &Accordion{
		Type:       t,
		Items:      config.Items,
		Borderless: config.Borderless,
		Classes:    config.Classes,
	}
}

func (a *Accordion) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	children := make([]layout.FlexChild, 0, len(a.Items)*2)

	for i, item := range a.Items {
		item := item

		if item.clickable.Clicked(gtx) && !item.Disabled {
			if a.Type == TypeSingle {
				targetState := !item.Expanded
				for _, other := range a.Items {
					other.Expanded = false
				}
				item.Expanded = targetState
			} else {
				item.Expanded = !item.Expanded
			}
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

	dims := layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (a *Accordion) layoutItem(gtx layout.Context, th *theme.Theme, item *Item) layout.Dimensions {
	bgColor := th.Colors.Card
	borderColor := th.Colors.Border
	fgColor := th.Colors.Foreground
	mutedColor := th.Colors.MutedFg

	if item.Disabled {
		fgColor.A = 128
		mutedColor.A = 128
	}

	styles := utils.ParseClasses(a.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
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
								if item.CustomHeader != nil {
									return item.CustomHeader(gtx)
								}
								lbl := material.Label(mTheme, th.Typography.FontSizeBase, item.Title)
								lbl.Color = fgColor
								lbl.Font.Weight = font.Medium
								return lbl.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								if item.Expanded {
									return lucide.ChevronUp.LayoutSize(gtx, unit.Dp(16), mutedColor)
								}
								return lucide.ChevronDown.LayoutSize(gtx, unit.Dp(16), mutedColor)
							}),
						)
					}),

					// Expandable Content Body
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if !item.Expanded {
							return layout.Dimensions{}
						}
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								if item.ContentWidget != nil {
									return item.ContentWidget(gtx)
								}
								lbl := material.Label(mTheme, th.Typography.FontSizeSM, item.Content)
								lbl.Color = mutedColor
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

		return layout.Stack{}.Layout(gtx,
			// Background drawn FIRST
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				if !a.Borderless {
					rect := image.Rectangle{Max: itemSize}
					radius := gtx.Dp(th.Radius.RadiusMD)
					theme.DrawRRectBackground(gtx, rect, radius, bgColor)

					rr := clip.UniformRRect(rect, radius)
					theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)
				}
				return layout.Dimensions{Size: itemSize}
			}),

			// Text content drawn ON TOP of background
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return renderContent(gtx)
			}),
		)
	})
}
