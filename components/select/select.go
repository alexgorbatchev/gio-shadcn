/*
Package selectcomp provides a dropdown select menu component for gio-shadcn applications.

Selects allow users to choose an option from a collapsible menu following
shadcn/ui design principles.
*/
package selectcomp

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Item represents a single selectable option in the dropdown.
type Item struct {
	Value     string
	Label     string
	clickable *widget.Clickable
}

// NewItem creates a new Select Option Item.
func NewItem(value, label string) *Item {
	return &Item{
		Value:     value,
		Label:     label,
		clickable: new(widget.Clickable),
	}
}

// Select represents a dropdown select menu component.
type Select struct {
	Options       []*Item
	SelectedValue string
	Open          bool
	Classes       string
	OnChange      func(string)

	triggerBtn *widget.Clickable
}

// Config represents configuration for creating a Select component.
type Config struct {
	Options       []*Item
	SelectedValue string
	Open          bool
	Classes       string
	OnChange      func(string)
}

// New creates a new Select component.
func New(config Config) *Select {
	sel := config.SelectedValue
	if sel == "" && len(config.Options) > 0 {
		sel = config.Options[0].Value
	}
	return &Select{
		Options:       config.Options,
		SelectedValue: sel,
		Open:          config.Open,
		Classes:       config.Classes,
		OnChange:      config.OnChange,
		triggerBtn:    new(widget.Clickable),
	}
}

// Layout renders the trigger button and dropdown option list when open.
func (s *Select) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if s.triggerBtn.Clicked(gtx) {
		s.Open = !s.Open
	}

	mTheme := material.NewTheme()
	selectedLabel := s.SelectedValue

	for _, opt := range s.Options {
		if opt.Value == s.SelectedValue {
			selectedLabel = opt.Label
			break
		}
	}

	bgColor := th.Colors.Background
	borderColor := th.Colors.Input

	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	triggerDims := s.triggerBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space3,
			Right:  th.Spacing.Space3,
		}

		tDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, selectedLabel)
					lbl.Color = th.Colors.Foreground
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					arrow := "▼"
					if s.Open {
						arrow = "▲"
					}
					lbl := material.Label(mTheme, th.Typography.FontSizeXS, arrow)
					lbl.Color = th.Colors.MutedFg
					return lbl.Layout(gtx)
				}),
			)
		})

		rect := image.Rectangle{Max: tDims.Size}
		radius := gtx.Dp(th.Radius.RadiusMD)
		rr := clip.UniformRRect(rect, radius)

		paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

		stroke := clip.Stroke{
			Path:  rr.Path(gtx.Ops),
			Width: 1.0,
		}
		paint.FillShape(gtx.Ops, borderColor, stroke.Op())

		return tDims
	})

	if !s.Open {
		return triggerDims
	}

	// Render Dropdown Options List
	optionChildren := make([]layout.FlexChild, 0, len(s.Options))

	for _, opt := range s.Options {
		opt := opt // capture loop variable

		if opt.clickable.Clicked(gtx) {
			s.SelectedValue = opt.Value
			s.Open = false
			if s.OnChange != nil {
				s.OnChange(s.SelectedValue)
			}
		}

		optionChildren = append(optionChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return opt.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				padding := layout.Inset{
					Top:    th.Spacing.Space2,
					Bottom: th.Spacing.Space2,
					Left:   th.Spacing.Space3,
					Right:  th.Spacing.Space3,
				}
				itemDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, opt.Label)
					lbl.Color = th.Colors.PopoverFg
					return lbl.Layout(gtx)
				})

				itemBg := th.Colors.Popover
				if opt.Value == s.SelectedValue || opt.clickable.Hovered() {
					itemBg = th.Colors.Secondary
				}

				rect := image.Rectangle{Max: itemDims.Size}
				paint.FillShape(gtx.Ops, itemBg, clip.Rect(rect).Op())

				return itemDims
			})
		}))
	}

	optsDims := layout.Flex{Axis: layout.Vertical}.Layout(gtx, optionChildren...)

	rect := image.Rectangle{Max: optsDims.Size}
	radius := gtx.Dp(th.Radius.RadiusMD)
	rr := clip.UniformRRect(rect, radius)

	paint.FillShape(gtx.Ops, th.Colors.Popover, rr.Op(gtx.Ops))

	stroke := clip.Stroke{
		Path:  rr.Path(gtx.Ops),
		Width: 1.0,
	}
	paint.FillShape(gtx.Ops, th.Colors.Border, stroke.Op())

	return layout.Dimensions{
		Size: image.Pt(triggerDims.Size.X, triggerDims.Size.Y+optsDims.Size.Y),
	}
}
