package tooltip_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/tooltip"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTooltipCreation(t *testing.T) {
	tt := tooltip.New(tooltip.Config{
		Text: "ASIO Low Latency Driver",
	})

	if tt.Text != "ASIO Low Latency Driver" {
		t.Errorf("expected Text to be 'ASIO Low Latency Driver', got %s", tt.Text)
	}
}

func TestTooltipLayout(t *testing.T) {
	th := theme.NewDark()
	tt := tooltip.New(tooltip.Config{
		Text: "ASIO Low Latency Driver",
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(200, 30)),
	}
	dims := tt.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Tooltip.Layout")
	}
}
