package alert_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/alert"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAlertCreation(t *testing.T) {
	al := alert.New(alert.Config{
		Title:       "System Alert",
		Description: "Audio buffer size set to 64 samples.",
	})

	if al.Title != "System Alert" {
		t.Errorf("expected Title to be 'System Alert', got %s", al.Title)
	}
}

func TestAlertLayout(t *testing.T) {
	th := theme.NewDark()
	al := alert.New(alert.Config{
		Title:       "System Alert",
		Description: "Audio buffer size set to 64 samples.",
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	dims := al.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Alert.Layout")
	}
}
