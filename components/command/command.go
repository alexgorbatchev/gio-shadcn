/*
Package command provides a command palette search component for gio-shadcn applications.

Commands display searchable action items and shortcuts following
shadcn/ui design principles.
*/
package command

import (
	"image"
	"strings"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Item represents a single command item in the palette.
type Item struct {
	Label     string
	Shortcut  string
	Group     string
	Icon      *lucide.Icon
	Disabled  bool
	clickable *widget.Clickable
}

// NewItem creates a new Command Item.
func NewItem(label, shortcut string) *Item {
	return &Item{
		Label:     label,
		Shortcut:  shortcut,
		clickable: new(widget.Clickable),
	}
}

// NewItemFull creates a new Command Item with icon and group.
func NewItemFull(label, shortcut, group string, icon *lucide.Icon, disabled bool) *Item {
	return &Item{
		Label:     label,
		Shortcut:  shortcut,
		Group:     group,
		Icon:      icon,
		Disabled:  disabled,
		clickable: new(widget.Clickable),
	}
}

// Command represents a command search palette component.
type Command struct {
	Placeholder  string
	Items        []*Item
	Classes      string
	OnSelectItem func(index int)

	searchEditor *widget.Editor
}

// Config represents configuration for creating a Command palette.
type Config struct {
	Placeholder  string
	Items        []*Item
	Classes      string
	OnSelectItem func(index int)
}

// New creates a new Command palette component.
func New(config Config) *Command {
	ph := config.Placeholder
	if ph == "" {
		ph = "Type a command or search..."
	}
	ed := new(widget.Editor)
	ed.SingleLine = true

	return &Command{
		Placeholder:  ph,
		Items:        config.Items,
		Classes:      config.Classes,
		OnSelectItem: config.OnSelectItem,
		searchEditor: ed,
	}
}

// Layout renders the search bar and filterable command items with background drawn first.
func (c *Command) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	query := strings.ToLower(c.searchEditor.Text())

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		children := make([]layout.FlexChild, 0, len(c.Items)*2+2)

		// Search Input Row
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			padding := layout.Inset{
				Top:    th.Spacing.Space3,
				Bottom: th.Spacing.Space3,
				Left:   th.Spacing.Space4,
				Right:  th.Spacing.Space4,
			}
			return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Right: th.Spacing.Space2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return lucide.Search.LayoutSize(gtx, unit.Dp(16), th.Colors.MutedFg)
						})
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(mTheme, c.searchEditor, c.Placeholder).Layout(gtx)
					}),
				)
			})
		}))

		// Bottom Separator Line
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: image.Pt(gtx.Constraints.Max.X, 1)}
			theme.DrawRRectBackground(gtx, rect, 0, th.Colors.Border)
			return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X, 1)}
		}))

		// Command Items List with optional group headers
		lastGroup := ""
		for idx, item := range c.Items {
			idx, item := idx, item

			if query != "" && !strings.Contains(strings.ToLower(item.Label), query) {
				continue
			}

			if item.Group != "" && item.Group != lastGroup {
				lastGroup = item.Group
				groupTitle := item.Group
				children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					padding := layout.Inset{
						Top:    th.Spacing.Space2,
						Bottom: th.Spacing.Space1,
						Left:   th.Spacing.Space4,
						Right:  th.Spacing.Space4,
					}
					return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(mTheme, unit.Sp(11), strings.ToUpper(groupTitle))
						lbl.Color = th.Colors.MutedFg
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					})
				}))
			}

			if item.clickable.Clicked(gtx) && !item.Disabled && c.OnSelectItem != nil {
				c.OnSelectItem(idx)
			}

			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return c.layoutItem(gtx, th, mTheme, item)
			}))
		}

		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	}

	contentDims := renderContent(gtxContent)
	cmdSize := contentDims.Size

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(c.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	dims := layout.Stack{}.Layout(gtx,
		// Background drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: cmdSize}
			radius := gtx.Dp(th.Radius.RadiusLG)
			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: cmdSize}
		}),

		// Content drawn ON TOP
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (c *Command) layoutItem(gtx layout.Context, th *theme.Theme, mTheme *material.Theme, item *Item) layout.Dimensions {
	padding := layout.Inset{
		Top:    th.Spacing.Space2,
		Bottom: th.Spacing.Space2,
		Left:   th.Spacing.Space4,
		Right:  th.Spacing.Space4,
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	fgColor := th.Colors.Foreground
	if item.Disabled {
		fgColor.A = 120
	}

	renderItemContent := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if item.Icon != nil {
						return layout.Inset{Right: th.Spacing.Space2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return item.Icon.LayoutSize(gtx, unit.Dp(16), fgColor)
						})
					}
					return layout.Dimensions{}
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, item.Label)
					lbl.Color = fgColor
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if item.Shortcut == "" {
						return layout.Dimensions{}
					}
					lbl := material.Label(mTheme, th.Typography.FontSizeXS, item.Shortcut)
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
			)
		})
	}

	contentDims := renderItemContent(gtxContent)
	itemSize := contentDims.Size

	return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				if item.clickable.Hovered() && !item.Disabled {
					rect := image.Rectangle{Max: itemSize}
					theme.DrawRRectBackground(gtx, rect, 0, th.Colors.Secondary)
				}
				return layout.Dimensions{Size: itemSize}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return renderItemContent(gtx)
			}),
		)
	})
}
