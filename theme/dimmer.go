package theme

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget"
)

// Dimmer provides a reusable modal backdrop dimmer that intercepts outside clicks to close overlays.
type Dimmer struct {
	clickable widget.Clickable
}

// NewDimmer creates a new reusable Dimmer.
func NewDimmer() *Dimmer {
	return &Dimmer{}
}

// Layout renders the dimmed backdrop filling the available constraints and invokes onClose when clicked.
func (d *Dimmer) Layout(gtx layout.Context, th *Theme, onClose func()) layout.Dimensions {
	for d.clickable.Clicked(gtx) {
		if onClose != nil {
			onClose()
		}
	}

	return d.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		rect := image.Rectangle{Max: gtx.Constraints.Max}
		backdropColor := color.NRGBA{R: 0, G: 0, B: 0, A: 160}
		DrawRRectBackground(gtx, rect, 0, backdropColor)
		return layout.Dimensions{Size: gtx.Constraints.Max}
	})
}
