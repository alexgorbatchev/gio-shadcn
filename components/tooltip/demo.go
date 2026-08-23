package tooltip

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	tipCallout := New(Config{Text: "ASIO Low Latency Buffer"})
	return tipCallout.Layout(gtx, th)
}
