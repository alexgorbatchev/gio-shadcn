/*
Package breadcrumb provides a breadcrumb trail navigation component for gio-shadcn applications.

Breadcrumbs indicate page hierarchy and location following
shadcn/ui design principles.
*/
package breadcrumb

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
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
	Items     []*Item
	Separator string
	OnSelect  func(index int)
}

// Config represents configuration for creating a Breadcrumb.
type Config struct {
	Items     []*Item
	Separator string
	OnSelect  func(index int)
}

// New creates a new Breadcrumb component.
func New(config Config) *Breadcrumb {
	return &Breadcrumb{
		Items:     config.Items,
		Separator: config.Separator,
		OnSelect:  config.OnSelect,
	}
}

// Layout renders the breadcrumb item trail with separators.
func (b *Breadcrumb) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	children := make([]layout.FlexChild, 0, len(b.Items)*2)

	for idx, item := range b.Items {
		idx, item := idx, item

		if item.clickable.Clicked(gtx) && b.OnSelect != nil {
			b.OnSelect(idx)
		}

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				fgColor := th.Colors.MutedFg
				if item.Active {
					fgColor = th.Colors.Foreground
				}
				lbl := material.Label(mTheme, th.Typography.FontSizeSM, item.Label)
				lbl.Color = fgColor
				if item.Active {
					lbl.Font.Weight = font.SemiBold
				}
				return lbl.Layout(gtx)
			})
		}))

		if idx < len(b.Items)-1 {
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Left: th.Spacing.Space2, Right: th.Spacing.Space2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					if b.Separator != "" {
						lbl := material.Label(mTheme, th.Typography.FontSizeSM, b.Separator)
						lbl.Color = th.Colors.MutedFg
						return lbl.Layout(gtx)
					}
					return lucide.ChevronRight.LayoutSize(gtx, unit.Dp(14), th.Colors.MutedFg)
				})
			}))
		}
	}

	dims := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
