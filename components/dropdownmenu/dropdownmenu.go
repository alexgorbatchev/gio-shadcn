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
	"github.com/bnema/gio-shadcn/components/button"
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
	Open          bool
	Items         []*Item
	Classes       string
	OnSelectItem  func(index int)
	TriggerButton *button.Button
	Trigger       layout.Widget

	backdropClick widget.Clickable
}

// Config represents configuration for creating a DropdownMenu.
type Config struct {
	TriggerText   string
	TriggerButton *button.Button
	Trigger       layout.Widget
	Open          bool
	Items         []*Item
	Classes       string
	OnSelectItem  func(index int)
}

// New creates a new DropdownMenu component.
func New(config Config) *DropdownMenu {
	dm := &DropdownMenu{
		Open:          config.Open,
		Items:         config.Items,
		Classes:       config.Classes,
		OnSelectItem:  config.OnSelectItem,
		TriggerButton: config.TriggerButton,
		Trigger:       config.Trigger,
	}

	if config.TriggerText != "" && dm.TriggerButton == nil {
		dm.TriggerButton = button.New(button.Config{
			Text:    config.TriggerText,
			Variant: theme.VariantOutline,
			OnClick: func() {
				dm.Open = !dm.Open
			},
		})
	}

	return dm
}

// Layout renders the dropdown menu panel or trigger with anchored menu when Open == true.
func (dm *DropdownMenu) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if dm.backdropClick.Clicked(gtx) {
		dm.Open = false
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	// 1. If trigger is present, render trigger and conditionally render anchored menu below it
	if dm.TriggerButton != nil || dm.Trigger != nil {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if dm.TriggerButton != nil {
					return dm.TriggerButton.Layout(gtx, th)
				}
				if dm.Trigger != nil {
					return dm.Trigger(gtx)
				}
				return layout.Dimensions{}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if !dm.Open {
					return layout.Dimensions{}
				}
				return layout.Inset{Top: th.Spacing.Space2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return dm.layoutMenuBox(gtx, th, mTheme)
				})
			}),
		)
	}

	// 2. Standalone menu rendering without trigger
	if !dm.Open {
		return layout.Dimensions{}
	}

	return dm.layoutMenuBox(gtx, th, mTheme)
}

func (dm *DropdownMenu) layoutMenuBox(gtx layout.Context, th *theme.Theme, mTheme *material.Theme) layout.Dimensions {
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

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (dm *DropdownMenu) layoutItem(gtx layout.Context, th *theme.Theme, mTheme *material.Theme, item *Item) layout.Dimensions {
	padding := layout.Inset{
		Top:    th.Spacing.Space2,
		Bottom: th.Spacing.Space2,
		Left:   th.Spacing.Space3,
		Right:  th.Spacing.Space3,
	}

	itemBg := th.Colors.Popover
	if item.clickable.Hovered() {
		itemBg = th.Colors.Secondary
	}

	return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtxContent := gtx
		gtxContent.Constraints.Min = image.Pt(0, 0)

		renderContent := func(gtx layout.Context) layout.Dimensions {
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
						lbl.Font.Weight = font.Medium
						return lbl.Layout(gtx)
					}),
				)
			})
		}

		itemDims := renderContent(gtxContent)
		itemSize := itemDims.Size

		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{Max: itemSize}
				radiusPx := gtx.Dp(th.Radius.RadiusSM)
				theme.DrawRRectBackground(gtx, rect, radiusPx, itemBg)
				return layout.Dimensions{Size: itemSize}
			}),

			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return renderContent(gtx)
			}),
		)
	})
}
