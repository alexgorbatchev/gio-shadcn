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

// Row represents a single data row in the table.
type Row struct {
	Cells     []string
	Selected  bool
	clickable *widget.Clickable
}

// NewRow creates a new Table Row.
func NewRow(cells ...string) *Row {
	return &Row{
		Cells:     cells,
		clickable: new(widget.Clickable),
	}
}

// Table represents a data grid table component.
type Table struct {
	Headers     []string
	Rows        []*Row
	Classes     string
	OnSelectRow func(index int)
}

// Config represents configuration for creating a Table.
type Config struct {
	Headers     []string
	Rows        []*Row
	Classes     string
	OnSelectRow func(index int)
}

// New creates a new Table component with the given configuration.
func New(config Config) *Table {
	return &Table{
		Headers:     config.Headers,
		Rows:        config.Rows,
		Classes:     config.Classes,
		OnSelectRow: config.OnSelectRow,
	}
}

// Layout renders the table headers and data rows.
func (t *Table) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	children := make([]layout.FlexChild, 0, len(t.Rows)+1)

	// Render Header Row
	children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return t.layoutHeader(gtx, th)
	}))

	// Render Data Rows
	for idx, row := range t.Rows {
		idx, row := idx, row // capture loop variables

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

	tableDims := layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(t.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	rect := image.Rectangle{Max: tableDims.Size}
	radius := gtx.Dp(th.Radius.RadiusMD)
	rr := clip.UniformRRect(rect, radius)

	paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

	stroke := clip.Stroke{
		Path:  rr.Path(gtx.Ops),
		Width: 1.0,
	}
	paint.FillShape(gtx.Ops, borderColor, stroke.Op())

	return tableDims
}

func (t *Table) layoutHeader(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	mTheme := material.NewTheme()

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
	mTheme := material.NewTheme()

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

		rowDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
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

		rect := image.Rectangle{Max: rowDims.Size}
		paint.FillShape(gtx.Ops, bgColor, clip.Rect(rect).Op())

		// Bottom divider line
		lineRect := image.Rectangle{
			Min: image.Pt(0, rowDims.Size.Y-1),
			Max: image.Pt(rowDims.Size.X, rowDims.Size.Y),
		}
		paint.FillShape(gtx.Ops, th.Colors.Border, clip.Rect(lineRect).Op())

		return rowDims
	})
}
