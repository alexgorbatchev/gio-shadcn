/*
Package togglegroup provides a grouped toggle selector component for gio-shadcn applications.

Toggle groups allow users to switch between exclusive option buttons following
shadcn/ui design principles.
*/
package togglegroup

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Item represents a single toggle item.
type Item struct {
	Key       string
	Label     string
	clickable *widget.Clickable
}

// NewItem creates a new Toggle Item.
func NewItem(key, label string) *Item {
	return &Item{
		Key:       key,
		Label:     label,
		clickable: new(widget.Clickable),
	}
}

// ToggleGroup represents a grouped toggle selector container.
type ToggleGroup struct {
	Items       []*Item
	SelectedKey string
	Classes     string
	OnChange    func(string)
}

// Config represents configuration for creating a ToggleGroup.
type Config struct {
	Items       []*Item
	SelectedKey string
	Classes     string
	OnChange    func(string)
}

// New creates a new ToggleGroup component.
func New(config Config) *ToggleGroup {
	sel := config.SelectedKey
	if sel == "" && len(config.Items) > 0 {
		sel = config.Items[0].Key
	}
	return &ToggleGroup{
		Items:       config.Items,
		SelectedKey: sel,
		Classes:     config.Classes,
		OnChange:    config.OnChange,
	}
}

// Layout renders the group of toggle items.
func (tg *ToggleGroup) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	children := make([]layout.FlexChild, 0, len(tg.Items)*2)

	for i, item := range tg.Items {
		item := item

		if item.clickable.Clicked(gtx) {
			tg.SelectedKey = item.Key
			if tg.OnChange != nil {
				tg.OnChange(tg.SelectedKey)
			}
		}

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return tg.layoutItem(gtx, th, item)
		}))

		if i < len(tg.Items)-1 {
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx)
			}))
		}
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space1,
		Bottom: th.Spacing.Space1,
		Left:   th.Spacing.Space1,
		Right:  th.Spacing.Space1,
	}

	dims := layout.Stack{}.Layout(gtx,
		// Toggle group container background drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: gtx.Constraints.Min}
			radius := gtx.Dp(th.Radius.RadiusMD)
			theme.DrawRRectBackground(gtx, rect, radius, th.Colors.Muted)
			return layout.Dimensions{Size: gtx.Constraints.Min}
		}),
		// Toggle items drawn ON TOP
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)
			})
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (tg *ToggleGroup) layoutItem(gtx layout.Context, th *theme.Theme, item *Item) layout.Dimensions {
	isSelected := item.Key == tg.SelectedKey

	bgColor := th.Colors.Muted
	fgColor := th.Colors.MutedFg

	if isSelected {
		bgColor = th.Colors.Background
		fgColor = th.Colors.Foreground
	}

	styles := utils.ParseClasses(tg.Classes)
	if styles.Background.A > 0 && isSelected {
		bgColor = styles.Background
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space3,
			Right:  th.Spacing.Space3,
		}

		return layout.Stack{}.Layout(gtx,
			// Item background drawn FIRST
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{Max: gtx.Constraints.Min}
				radius := gtx.Dp(th.Radius.RadiusSM)
				theme.DrawRRectBackground(gtx, rect, radius, bgColor)
				return layout.Dimensions{Size: gtx.Constraints.Min}
			}),
			// Item text drawn ON TOP
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, item.Label)
					lbl.Color = fgColor
					if isSelected {
						lbl.Font.Weight = font.Medium
					}
					return lbl.Layout(gtx)
				})
			}),
		)
	})
}
