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
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Item represents a single command item in the palette.
type Item struct {
	Label     string
	Shortcut  string
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

// Layout renders the search bar and filterable command items.
func (c *Command) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	mTheme := material.NewTheme()
	query := strings.ToLower(c.searchEditor.Text())

	children := make([]layout.FlexChild, 0, len(c.Items)+2)

	// Search Input Row
	children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space3,
			Bottom: th.Spacing.Space3,
			Left:   th.Spacing.Space4,
			Right:  th.Spacing.Space4,
		}
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return material.Editor(mTheme, c.searchEditor, c.Placeholder).Layout(gtx)
		})
	}))

	// Bottom Separator Line
	children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		rect := image.Rectangle{Max: image.Pt(gtx.Constraints.Max.X, 1)}
		paint.FillShape(gtx.Ops, th.Colors.Border, clip.Rect(rect).Op())
		return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X, 1)}
	}))

	// Command Items List
	for idx, item := range c.Items {
		idx, item := idx, item

		if query != "" && !strings.Contains(strings.ToLower(item.Label), query) {
			continue
		}

		if item.clickable.Clicked(gtx) && c.OnSelectItem != nil {
			c.OnSelectItem(idx)
		}

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return c.layoutItem(gtx, th, mTheme, item)
		}))
	}

	cmdDims := layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(c.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	rect := image.Rectangle{Max: cmdDims.Size}
	radius := gtx.Dp(th.Radius.RadiusMD)
	rr := clip.UniformRRect(rect, radius)

	paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

	stroke := clip.Stroke{
		Path:  rr.Path(gtx.Ops),
		Width: 1.0,
	}
	paint.FillShape(gtx.Ops, borderColor, stroke.Op())

	return cmdDims
}

func (c *Command) layoutItem(gtx layout.Context, th *theme.Theme, mTheme *material.Theme, item *Item) layout.Dimensions {
	return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space4,
			Right:  th.Spacing.Space4,
		}

		itemDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, item.Label)
					lbl.Color = th.Colors.Foreground
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

		if item.clickable.Hovered() {
			rect := image.Rectangle{Max: itemDims.Size}
			paint.FillShape(gtx.Ops, th.Colors.Secondary, clip.Rect(rect).Op())
		}

		return itemDims
	})
}
