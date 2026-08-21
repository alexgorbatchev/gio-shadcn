/*
Package breadcrumb provides a breadcrumb trail navigation component for gio-shadcn applications.

Breadcrumbs indicate page hierarchy and location following
shadcn/ui design principles.
*/
package breadcrumb

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

// Item represents a single breadcrumb location item.
type Item struct {
	Label     string
	Active    bool
	clickable *widget.Clickable
}

// NewItem creates a new Breadcrumb Item.
func NewItem(label string, active bool) *Item {
	return &Item{
		Label:     label,
		Active:    active,
		clickable: new(widget.Clickable),
	}
}

// Breadcrumb represents a breadcrumb trail container component.
type Breadcrumb struct {
	Items    []*Item
	OnSelect func(index int)
}

// Config represents configuration for creating a Breadcrumb.
type Config struct {
	Items    []*Item
	OnSelect func(index int)
}

// New creates a new Breadcrumb component.
func New(config Config) *Breadcrumb {
	return &Breadcrumb{
		Items:    config.Items,
		OnSelect: config.OnSelect,
	}
}

// Layout renders the breadcrumb item trail with slashes.
func (b *Breadcrumb) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	mTheme := material.NewTheme()
	children := make([]layout.FlexChild, 0, len(b.Items)*2)

	for idx, item := range b.Items {
		idx, item := idx, item // capture loop variables

		if item.clickable.Clicked(gtx) && b.OnSelect != nil {
			b.OnSelect(idx)
		}

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(mTheme, th.Typography.FontSizeSM, item.Label)
				lbl.Color = th.Colors.MutedFg
				if item.Active {
					lbl.Color = th.Colors.Foreground
					lbl.Font.Weight = font.Medium
				}
				return lbl.Layout(gtx)
			})
		}))

		if idx < len(b.Items)-1 {
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Left: th.Spacing.Space2, Right: th.Spacing.Space2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, "/")
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				})
			}))
		}
	}

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)
}
