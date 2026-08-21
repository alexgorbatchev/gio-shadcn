package table_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/table"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTableCreation(t *testing.T) {
	tbl := table.New(table.Config{
		Headers: []string{"TITLE", "ARTIST", "BPM"},
		Rows: []*table.Row{
			table.NewRow("Starlight Symphony", "Aethelgard", "128"),
			table.NewRow("Quantum Drift", "Cyberpulse", "132"),
		},
	})

	if len(tbl.Headers) != 3 {
		t.Fatalf("expected 3 headers, got %d", len(tbl.Headers))
	}

	if len(tbl.Rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(tbl.Rows))
	}
}

func TestTableLayout(t *testing.T) {
	th := theme.NewDark()
	tbl := table.New(table.Config{
		Headers: []string{"TITLE", "ARTIST", "BPM"},
		Rows: []*table.Row{
			table.NewRow("Starlight Symphony", "Aethelgard", "128"),
			table.NewRow("Quantum Drift", "Cyberpulse", "132"),
		},
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(400, 200)),
	}
	dims := tbl.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Table.Layout")
	}
}
