package table

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	trackTable := New(Config{
		Headers: []string{"TITLE", "ARTIST", "BPM", "KEY", "GENRE"},
		Rows: []*Row{
			NewRow("Starlight Symphony", "Aethelgard", "128", "8A", "Progressive House"),
			NewRow("Quantum Drift", "Cyberpulse", "132", "11B", "Techno"),
			NewRow("Solar Flare", "Helios", "126", "4A", "Melodic Techno"),
		},
	})

	return trackTable.Layout(gtx, th)
}
