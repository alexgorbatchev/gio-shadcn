/*
Package resizable provides a split-view resizable panel component for gio-shadcn applications.

Resizables provide split panels with draggable divider handles following
shadcn/ui design principles.
*/
package resizable

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
)

// Resizable represents a split view panel container with a resizable handle.
type Resizable struct {
	Ratio       float32 // Left width ratio (0.1 to 0.9)
	LeftWidget  layout.Widget
	RightWidget layout.Widget
}

// Config represents configuration for creating a Resizable component.
type Config struct {
	Ratio       float32
	LeftWidget  layout.Widget
	RightWidget layout.Widget
}

// New creates a new Resizable panel wrapper.
func New(config Config) *Resizable {
	r := config.Ratio
	if r <= 0.1 || r >= 0.9 {
		r = 0.5
	}
	return &Resizable{
		Ratio:       r,
		LeftWidget:  config.LeftWidget,
		RightWidget: config.RightWidget,
	}
}

// Layout renders the split view panels and divider handle.
func (r *Resizable) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	totalWidth := gtx.Constraints.Max.X
	handleWidth := gtx.Dp(unit.Dp(4))
	leftWidth := int(float32(totalWidth-handleWidth) * r.Ratio)
	rightWidth := totalWidth - leftWidth - handleWidth

	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		// Left Panel
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints = layout.Exact(image.Pt(leftWidth, gtx.Constraints.Max.Y))
			if r.LeftWidget != nil {
				return r.LeftWidget(gtx)
			}
			return layout.Dimensions{Size: gtx.Constraints.Max}
		}),
		// Resizable Handle Divider
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			size := image.Pt(handleWidth, gtx.Constraints.Max.Y)
			rect := image.Rectangle{Max: size}
			paint.FillShape(gtx.Ops, th.Colors.Border, clip.Rect(rect).Op())
			return layout.Dimensions{Size: size}
		}),
		// Right Panel
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints = layout.Exact(image.Pt(rightWidth, gtx.Constraints.Max.Y))
			if r.RightWidget != nil {
				return r.RightWidget(gtx)
			}
			return layout.Dimensions{Size: gtx.Constraints.Max}
		}),
	)
}
