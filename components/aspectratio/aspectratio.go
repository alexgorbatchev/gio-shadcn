/*
Package aspectratio provides a fixed ratio container layout component for gio-shadcn applications.

AspectRatios maintain proportional width-to-height constraints following
shadcn/ui design principles.
*/
package aspectratio

import (
	"image"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

// AspectRatio represents a container enforcing a fixed width-to-height ratio.
type AspectRatio struct {
	Ratio  float32 // Width / Height (e.g., 16.0 / 9.0 = 1.777)
	Widget layout.Widget
}

// Config represents configuration for creating an AspectRatio wrapper.
type Config struct {
	Ratio  float32
	Widget layout.Widget
}

// New creates a new AspectRatio wrapper.
func New(config Config) *AspectRatio {
	r := config.Ratio
	if r <= 0 {
		r = 16.0 / 9.0
	}
	return &AspectRatio{
		Ratio:  r,
		Widget: config.Widget,
	}
}

// Layout applies the aspect ratio constraints and renders the child widget.
func (ar *AspectRatio) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	w := gtx.Constraints.Max.X
	if w <= 0 {
		w = gtx.Constraints.Min.X
	}

	h := int(float32(w) / ar.Ratio)
	size := image.Pt(w, h)
	gtx.Constraints = layout.Exact(size)

	if ar.Widget != nil {
		return ar.Widget(gtx)
	}

	return layout.Dimensions{Size: size}
}
