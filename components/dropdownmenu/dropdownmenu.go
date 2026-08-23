/*
Package dropdownmenu provides an action dropdown menu component for gio-shadcn applications.

DropdownMenus display action menus following
shadcn/ui design principles.
*/
package dropdownmenu

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Item represents a single menu item in the dropdown menu.
type Item struct {
	Label     string
	Shortcut  string
	clickable *widget.Clickable
}

// NewItem creates a new DropdownMenu Item.
func NewItem(label, shortcut string) *Item {
	return &Item{
		Label:     label,
		Shortcut:  shortcut,
		clickable: new(widget.Clickable),
	}
}

// DropdownMenu represents an action dropdown menu component.
type DropdownMenu struct {
	Open         bool
	Items        []*Item
	Classes      string
	OnSelectItem func(index int)
}

// Config represents configuration for creating a DropdownMenu.
type Config struct {
	Open         bool
	Items        []*Item
	Classes      string
	OnSelectItem func(index int)
}

// New creates a new DropdownMenu component.
func New(config Config) *DropdownMenu {
	return &DropdownMenu{
		Open:         config.Open,
		Items:        config.Items,
		Classes:      config.Classes,
		OnSelectItem: config.OnSelectItem,
	}
}

// Layout renders the dropdown menu panel when Open == true with background drawn first.
func (dm *DropdownMenu) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if !dm.Open {
		return layout.Dimensions{}
	}

	if th == nil {
		th = theme.New()
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		children := make([]layout.FlexChild, 0, len(dm.Items))

		for idx, item := range dm.Items {
			idx, item := idx, item

			if item.clickable.Clicked(gtx) {
				dm.Open = false
				if dm.OnSelectItem != nil {
					dm.OnSelectItem(idx)
				}
			}

			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return dm.layoutItem(gtx, th, mTheme, item)
			}))
		}

		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	}

	contentDims := renderContent(gtxContent)
	menuSize := contentDims.Size

	bgColor := th.Colors.Popover
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(dm.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	dims := layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: menuSize}
			radius := gtx.Dp(th.Radius.RadiusMD)

			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: menuSize}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)

	// Reset active GPU paint color state
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (dm *DropdownMenu) layoutItem(gtx layout.Context, th *theme.Theme, mTheme *material.Theme, item *Item) layout.Dimensions {
	return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space3,
			Right:  th.Spacing.Space3,
		}

		gtxContent := gtx
		gtxContent.Constraints.Min = image.Pt(0, 0)

		renderItem := func(gtx layout.Context) layout.Dimensions {
			return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(mTheme, th.Typography.FontSizeSM, item.Label)
						lbl.Color = th.Colors.PopoverFg
						return lbl.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if item.Shortcut == "" {
							return layout.Dimensions{}
						}
						lbl := material.Label(mTheme, th.Typography.FontSizeXS, item.Shortcut)
						lbl.Color = th.Colors.MutedFg
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					}),
				)
			})
		}

		itemDims := renderItem(gtxContent)
		itemSize := itemDims.Size

		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				if item.clickable.Hovered() {
					rect := image.Rectangle{Max: itemSize}
					theme.DrawRRectBackground(gtx, rect, 0, th.Colors.Secondary)
				}
				return layout.Dimensions{Size: itemSize}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return renderItem(gtx)
			}),
		)
	})
}
