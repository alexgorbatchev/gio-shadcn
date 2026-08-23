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

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

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

	// 1. Render Trigger Button
	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	padding := layout.Inset{
		Top:    th.Spacing.Space2,
		Bottom: th.Spacing.Space2,
		Left:   th.Spacing.Space3,
		Right:  th.Spacing.Space3,
	}

	renderTriggerContent := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
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
	}

	triggerDims := s.triggerBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		contentDims := renderTriggerContent(gtxContent)
		tSize := contentDims.Size

		return layout.Stack{}.Layout(gtx,
			// Trigger Background drawn FIRST
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{Max: tSize}
				radius := gtx.Dp(th.Radius.RadiusMD)
				theme.DrawRRectBackground(gtx, rect, radius, bgColor)

				rr := clip.UniformRRect(rect, radius)
				theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

				return layout.Dimensions{Size: tSize}
			}),

			// Trigger Text Content drawn ON TOP
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return renderTriggerContent(gtx)
			}),
		)
	})

	if !s.Open {
		paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)
		return triggerDims
	}

	// 2. Render Dropdown Options List
	optionChildren := make([]layout.FlexChild, 0, len(s.Options))

	for _, opt := range s.Options {
		opt := opt

		if opt.clickable.Clicked(gtx) {
			s.SelectedValue = opt.Value
			s.Open = false
			if s.OnChange != nil {
				s.OnChange(s.SelectedValue)
			}
		}

		optionChildren = append(optionChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return opt.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				itemBg := th.Colors.Popover
				if opt.Value == s.SelectedValue || opt.clickable.Hovered() {
					itemBg = th.Colors.Secondary
				}

				renderItemContent := func(gtx layout.Context) layout.Dimensions {
					return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(mTheme, th.Typography.FontSizeSM, opt.Label)
						lbl.Color = th.Colors.PopoverFg
						return lbl.Layout(gtx)
					})
				}

				itemContentDims := renderItemContent(gtxContent)
				itemSize := itemContentDims.Size

				return layout.Stack{}.Layout(gtx,
					// Item background drawn FIRST
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						rect := image.Rectangle{Max: itemSize}
						theme.DrawRRectBackground(gtx, rect, 0, itemBg)
						return layout.Dimensions{Size: itemSize}
					}),

					// Item text drawn ON TOP
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return renderItemContent(gtx)
					}),
				)
			})
		}))
	}

	optsContentDims := layout.Flex{Axis: layout.Vertical}.Layout(gtxContent, optionChildren...)
	optsSize := optsContentDims.Size

	optsDims := layout.Stack{}.Layout(gtx,
		// Popover list background drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: optsSize}
			radius := gtx.Dp(th.Radius.RadiusMD)
			theme.DrawRRectBackground(gtx, rect, radius, th.Colors.Popover)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, th.Colors.Border)

			return layout.Dimensions{Size: optsSize}
		}),

		// Options list drawn ON TOP
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, optionChildren...)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return layout.Dimensions{
		Size: image.Pt(triggerDims.Size.X, triggerDims.Size.Y+optsDims.Size.Y),
	}
}
