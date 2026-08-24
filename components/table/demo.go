package table

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	TrackTable *Table
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
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

	return s.TrackTable.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
