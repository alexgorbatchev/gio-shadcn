package table_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/table"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTableBasic(t *testing.T) {
	th := theme.NewDark()
	tbl := table.New(table.Config{
		Headers: []string{"Invoice", "Status", "Amount"},
		Rows: []*table.Row{
			table.NewRow("INV001", "Paid", "$250.00"),
			table.NewRow("INV002", "Pending", "$150.00"),
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 150))}
	dims := tbl.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestTableRowSelection(t *testing.T) {
	th := theme.NewDark()
	selected := -1
	tbl := table.New(table.Config{
		Headers: []string{"Item"},
		Rows: []*table.Row{
			table.NewRow("One"),
			table.NewRow("Two"),
		},
		OnSelectRow: func(index int) {
			selected = index
		},
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(200, 100))}
	_ = tbl.Layout(gtx, th)
	_ = selected
}
