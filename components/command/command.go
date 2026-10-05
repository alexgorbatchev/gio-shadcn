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
	"gioui.org/gesture"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/button"
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
	Open          bool
	Placeholder   string
	Items         []*Item
	Classes       string
	OnSelectItem  func(index int)
	TriggerButton *button.Button
	Trigger       layout.Widget

	searchEditor *widget.Editor
	dimmer       *theme.Dimmer
	panel        gesture.Click
	wasOpen      bool
}

// Config represents configuration for creating a Command palette.
type Config struct {
	Open          bool
	Placeholder   string
	Items         []*Item
	Classes       string
	TriggerText   string
	TriggerButton *button.Button
	Trigger       layout.Widget
	OnSelectItem  func(index int)
}

// New creates a new Command palette component.
func New(config Config) *Command {
	ph := config.Placeholder
	if ph == "" {
		ph = "Type a command or search..."
	}
	ed := new(widget.Editor)
	ed.SingleLine = true

	cmd := &Command{
		Open:          config.Open,
		Placeholder:   ph,
		Items:         config.Items,
		Classes:       config.Classes,
		OnSelectItem:  config.OnSelectItem,
		TriggerButton: config.TriggerButton,
		Trigger:       config.Trigger,
		searchEditor:  ed,
		dimmer:        theme.NewDimmer(),
	}

	if config.TriggerText != "" && cmd.TriggerButton == nil {
		cmd.TriggerButton = button.New(button.Config{
			Text:    config.TriggerText,
			Variant: theme.VariantOutline,
			OnClick: func() {
				cmd.Open = true
			},
		})
	}

	return cmd
}

// Layout renders the trigger or palette, floating centered above all elements with a dimmed backdrop when Open == true.
func (c *Command) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}
	defer func() { paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops) }()

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	// 1. If trigger is configured, render trigger and modal floating palette with dimmed backdrop when Open == true
	if c.TriggerButton != nil || c.Trigger != nil {
		var triggerDims layout.Dimensions
		if c.TriggerButton != nil {
			triggerDims = c.TriggerButton.Layout(gtx, th)
		} else if c.Trigger != nil {
			triggerDims = c.Trigger(gtx)
		}

		if c.Open {
			th.AddOverlay(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(gtx.Constraints.Max)
				if !c.wasOpen {
					gtx.Execute(key.FocusCmd{Tag: c.searchEditor})
					c.wasOpen = true
				}
				for {
					if _, ok := c.panel.Update(gtx.Source); !ok {
						break
					}
				}
				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
					// Dimmed backdrop filling entire window that closes on outside click
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						return c.dimmer.Layout(gtx, th, func() {
							c.Open = false
						})
					}),
					// Centered modal floating command palette
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							maxW := gtx.Dp(unit.Dp(540))
							if gtx.Constraints.Max.X > maxW {
								gtx.Constraints.Max.X = maxW
								gtx.Constraints.Min.X = maxW
							}
							return c.layoutPaletteBox(gtx, th, mTheme, true)
						})
					}),
				)
			})
		} else {
			c.wasOpen = false
		}

		return triggerDims
	}

	// 2. Inline/embedded command box rendering
	return c.layoutPaletteBox(gtx, th, mTheme, false)
}

func (c *Command) layoutPaletteBox(gtx layout.Context, th *theme.Theme, mTheme *material.Theme, isModal bool) layout.Dimensions {
	for {
		if _, ok := c.searchEditor.Update(gtx); !ok {
			break
		}
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
						return layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							ed := material.Editor(mTheme, c.searchEditor, c.Placeholder)
							ed.TextSize = th.Typography.FontSizeSM
							ed.Color = th.Colors.Foreground
							ed.HintColor = th.Colors.MutedFg
							return ed.Layout(gtx)
						})
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

			if item.clickable.Clicked(gtx) && !item.Disabled {
				if isModal {
					c.Open = false
				}
				if c.OnSelectItem != nil {
					c.OnSelectItem(idx)
				}
			}

			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return c.layoutItem(gtx, th, mTheme, item)
			}))
		}

		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	}

	macro := op.Record(gtx.Ops)
	contentDims := renderContent(gtxContent)
	callOp := macro.Stop()

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

			rr := theme.RRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: cmdSize}
		}),

		// Recorded content played EXACTLY ONCE
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if isModal {
				defer clip.Rect(image.Rectangle{Max: cmdSize}).Push(gtx.Ops).Pop()
				c.panel.Add(gtx.Ops)
			}
			callOp.Add(gtx.Ops)
			return contentDims
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

	fgColor := th.Colors.Foreground
	if item.Disabled {
		fgColor = th.Colors.MutedFg
	}

	macro := op.Record(gtx.Ops)
	contentDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
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
				return layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, item.Label)
					lbl.Color = fgColor
					return lbl.Layout(gtx)
				})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if item.Shortcut == "" {
					return layout.Dimensions{}
				}
				kbdPadding := layout.Inset{
					Top:    unit.Dp(2),
					Bottom: unit.Dp(2),
					Left:   unit.Dp(6),
					Right:  unit.Dp(6),
				}
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					macroKbd := op.Record(gtx.Ops)
					kbdDims := kbdPadding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(mTheme, th.Typography.FontSizeXS, item.Shortcut)
						lbl.Color = th.Colors.MutedFg
						lbl.Font.Weight = font.Medium
						return lbl.Layout(gtx)
					})
					callKbd := macroKbd.Stop()

					kbdSize := kbdDims.Size
					minH := gtx.Dp(unit.Dp(20))
					if kbdSize.Y < minH {
						kbdSize.Y = minH
					}

					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx layout.Context) layout.Dimensions {
							rect := image.Rectangle{Max: kbdSize}
							radius := gtx.Dp(th.Radius.RadiusSM)
							theme.DrawRRectBackground(gtx, rect, radius, th.Colors.Muted)
							rr := theme.RRect(rect, radius)
							theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, th.Colors.Border)
							return layout.Dimensions{Size: kbdSize}
						}),
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								callKbd.Add(gtx.Ops)
								return kbdDims
							})
						}),
					)
				})
			}),
		)
	})
	callOp := macro.Stop()

	itemSize := contentDims.Size

	return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				if item.clickable.Hovered() && !item.Disabled {
					rect := image.Rectangle{Max: itemSize}
					radius := gtx.Dp(th.Radius.RadiusSM)
					theme.DrawRRectBackground(gtx, rect, radius, th.Colors.Secondary)
				}
				return layout.Dimensions{Size: itemSize}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				callOp.Add(gtx.Ops)
				return contentDims
			}),
		)
	})
}
