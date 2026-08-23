package drawer

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	bottomDrawer := New(Config{
		Title:       "System Telemetry",
		Description: "CPU: 2.1% | RAM: 189.5MB | Metal GPU 120 FPS",
		Open:        false,
	})

	btnTrigger := button.New(button.Config{
		Text:    "Open Bottom Drawer",
		Variant: theme.VariantOutline,
		OnClick: func() {
			bottomDrawer.Open = true
		},
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnTrigger.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return bottomDrawer.Layout(gtx, th) }),
	)
}
