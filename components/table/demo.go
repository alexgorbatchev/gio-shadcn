package table

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	InvoiceTable *Table
	TrackTable   *Table
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		InvoiceTable: New(Config{
			Headers: []string{"INVOICE", "STATUS", "METHOD", "AMOUNT"},
			Rows: []*Row{
				NewRow("INV001", "Paid", "Credit Card", "$250.00"),
				NewRow("INV002", "Pending", "PayPal", "$150.00"),
				NewRow("INV003", "Unpaid", "Bank Transfer", "$350.00"),
				NewRow("INV004", "Paid", "Credit Card", "$450.00"),
				NewRow("INV005", "Paid", "PayPal", "$550.00"),
			},
		}),
		TrackTable: New(Config{
			Headers: []string{"TITLE", "ARTIST", "BPM", "KEY", "GENRE"},
			Rows: []*Row{
				NewRow("Starlight Symphony", "Aethelgard", "128", "8A", "Progressive House"),
				NewRow("Quantum Drift", "Cyberpulse", "132", "11B", "Techno"),
				NewRow("Solar Flare", "Helios", "126", "4A", "Melodic Techno"),
			},
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("1. Invoices Table (Official Demo)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.InvoiceTable.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Audio Track Data Grid", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.TrackTable.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
