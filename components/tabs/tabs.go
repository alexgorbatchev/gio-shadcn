/*
Package tabs provides a tabbed navigation component for gio-shadcn applications.

Tabs switch between active views or content sections following
shadcn/ui design principles.
*/
package tabs

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

// Tab represents a single tab selector.
type Tab struct {
	Key       string
	Label     string
	clickable *widget.Clickable
}

// NewTab creates a new Tab.
func NewTab(key, label string) *Tab {
	return &Tab{
		Key:       key,
		Label:     label,
		clickable: new(widget.Clickable),
	}
}

// Tabs represents a tab navigation container component.
type Tabs struct {
	Tabs      []*Tab
	ActiveKey string
	Classes   string
	OnChange  func(string)
}

// Config represents configuration for creating Tabs.
type Config struct {
	Tabs      []*Tab
	ActiveKey string
	Classes   string
	OnChange  func(string)
}

// New creates a new Tabs component with the given configuration.
func New(config Config) *Tabs {
	active := config.ActiveKey
	if active == "" && len(config.Tabs) > 0 {
		active = config.Tabs[0].Key
	}
	return &Tabs{
		Tabs:      config.Tabs,
		ActiveKey: active,
		Classes:   config.Classes,
		OnChange:  config.OnChange,
	}
}

// Layout renders the horizontal tab container and manages active tab switching.
func (t *Tabs) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	children := make([]layout.FlexChild, 0, len(t.Tabs)*2)

	for i, tab := range t.Tabs {
		tab := tab

		if tab.clickable.Clicked(gtx) {
			t.ActiveKey = tab.Key
			if t.OnChange != nil {
				t.OnChange(t.ActiveKey)
			}
		}

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.layoutTab(gtx, th, tab)
		}))

		if i < len(t.Tabs)-1 {
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
			}))
		}
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space1,
		Bottom: th.Spacing.Space1,
		Left:   th.Spacing.Space1,
		Right:  th.Spacing.Space1,
	}

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: gtx.Constraints.Min}
			radius := gtx.Dp(th.Radius.RadiusMD)
			rr := clip.UniformRRect(rect, radius)
			paint.FillShape(gtx.Ops, th.Colors.Muted, rr.Op(gtx.Ops))
			return layout.Dimensions{Size: gtx.Constraints.Min}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)
			})
		}),
	)
}

func (t *Tabs) layoutTab(gtx layout.Context, th *theme.Theme, tab *Tab) layout.Dimensions {
	isActive := tab.Key == t.ActiveKey

	bgColor := th.Colors.Muted
	fgColor := th.Colors.MutedFg

	if isActive {
		bgColor = th.Colors.Background
		fgColor = th.Colors.Foreground
	}

	styles := utils.ParseClasses(t.Classes)
	if styles.Background.A > 0 && isActive {
		bgColor = styles.Background
	}

	mTheme := material.NewTheme()

	return tab.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space4,
			Right:  th.Spacing.Space4,
		}

		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{Max: gtx.Constraints.Min}
				radius := gtx.Dp(th.Radius.RadiusSM)
				rr := clip.UniformRRect(rect, radius)
				paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))
				return layout.Dimensions{Size: gtx.Constraints.Min}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, tab.Label)
					lbl.Color = fgColor
					if isActive {
						lbl.Font.Weight = font.Medium
					}
					return lbl.Layout(gtx)
				})
			}),
		)
	})
}
