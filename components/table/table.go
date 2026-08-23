/*
Package table provides a data grid table component for gio-shadcn applications.

Tables display structured data in rows and columns following
shadcn/ui design principles.
*/
package table

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

type Row struct {
	Cells     []string
	Selected  bool
	clickable *widget.Clickable
}

func NewRow(cells ...string) *Row {
	return &Row{
		Cells:     cells,
		clickable: new(widget.Clickable),
	}
}

type Table struct {
	Headers     []string
	Rows        []*Row
	Classes     string
	OnSelectRow func(index int)
}

type Config struct {
	Headers     []string
	Rows        []*Row
	Classes     string
	OnSelectRow func(index int)
}

func New(config Config) *Table {
	return &Table{
		Headers:     config.Headers,
		Rows:        config.Rows,
		Classes:     config.Classes,
		OnSelectRow: config.OnSelectRow,
	}
}

func (t *Table) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(t.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	children := make([]layout.FlexChild, 0, len(t.Rows)+1)

	children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return t.layoutHeader(gtx, th)
	}))

	for idx, row := range t.Rows {
		idx, row := idx, row

		if row.clickable.Clicked(gtx) {
			for i, r := range t.Rows {
				r.Selected = (i == idx)
			}
			if t.OnSelectRow != nil {
				t.OnSelectRow(idx)
			}
		}

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.layoutRow(gtx, th, row)
		}))
	}

	// Measure content dimensions with unconstrained Min
	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	}

	contentDims := renderContent(gtxContent)
	tableSize := contentDims.Size

	dims := layout.Stack{}.Layout(gtx,
		// Background & Border drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: tableSize}
			radius := gtx.Dp(th.Radius.RadiusMD)

			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: tableSize}
		}),

		// Table Rows drawn ON TOP
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (t *Table) layoutHeader(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space3,
		Bottom: th.Spacing.Space3,
		Left:   th.Spacing.Space4,
		Right:  th.Spacing.Space4,
	}

	return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		colChildren := make([]layout.FlexChild, 0, len(t.Headers))
		for _, h := range t.Headers {
			h := h
			colChildren = append(colChildren, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(mTheme, th.Typography.FontSizeXS, h)
				lbl.Color = th.Colors.MutedFg
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}))
		}
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, colChildren...)
	})
}

func (t *Table) layoutRow(gtx layout.Context, th *theme.Theme, row *Row) layout.Dimensions {
	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	bgColor := th.Colors.Background
	if row.Selected {
		bgColor = th.Colors.Secondary
	}

	return row.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		padding := layout.Inset{
			Top:    th.Spacing.Space3,
			Bottom: th.Spacing.Space3,
			Left:   th.Spacing.Space4,
			Right:  th.Spacing.Space4,
		}

		gtxContent := gtx
		gtxContent.Constraints.Min = image.Pt(0, 0)

		renderRowContent := func(gtx layout.Context) layout.Dimensions {
			return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				colChildren := make([]layout.FlexChild, 0, len(row.Cells))
				for _, cell := range row.Cells {
					cell := cell
					colChildren = append(colChildren, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(mTheme, th.Typography.FontSizeSM, cell)
						lbl.Color = th.Colors.Foreground
						return lbl.Layout(gtx)
					}))
				}
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, colChildren...)
			})
		}

		rowDims := renderRowContent(gtxContent)
		rowSize := rowDims.Size

		return layout.Stack{}.Layout(gtx,
			// Row Background drawn FIRST
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{Max: rowSize}
				theme.DrawRRectBackground(gtx, rect, 0, bgColor)

				lineRect := image.Rectangle{
					Min: image.Pt(0, rowSize.Y-1),
					Max: image.Pt(rowSize.X, rowSize.Y),
				}
				theme.DrawRRectBackground(gtx, lineRect, 0, th.Colors.Border)

				return layout.Dimensions{Size: rowSize}
			}),

			// Row Text drawn ON TOP
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return renderRowContent(gtx)
			}),
		)
	})
}
