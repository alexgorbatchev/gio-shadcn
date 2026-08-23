/*
Package scrollarea provides a custom styled scrollbar viewport component for gio-shadcn applications.

ScrollAreas render scrollable viewports following
shadcn/ui design principles.
*/
package scrollarea

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/bnema/gio-shadcn/theme"
)

// ScrollArea represents a custom scrollbar viewport component.
type ScrollArea struct {
	List   *widget.List
	Widget layout.Widget
}

// Config represents configuration for creating a ScrollArea.
type Config struct {
	List   *widget.List
	Widget layout.Widget
}

// New creates a new ScrollArea viewport wrapper.
func New(config Config) *ScrollArea {
	lst := config.List
	if lst == nil {
		lst = &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		}
	}
	return &ScrollArea{
		List:   lst,
		Widget: config.Widget,
	}
}

// Layout renders the scrollable area.
func (sa *ScrollArea) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if sa.Widget == nil {
		return layout.Dimensions{}
	}

	if th == nil {
		th = theme.New()
	}

	dims := sa.List.Layout(gtx, 1, func(gtx layout.Context, index int) layout.Dimensions {
		wDims := sa.Widget(gtx)

		// Draw subtle scrollbar track indicator
		trackWidth := gtx.Dp(th.Spacing.Space1)
		trackY := wDims.Size.Y
		if trackY > 0 {
			trackRect := image.Rectangle{
				Min: image.Pt(wDims.Size.X-trackWidth, 0),
				Max: image.Pt(wDims.Size.X, trackY),
			}
			theme.DrawRRectBackground(gtx, trackRect, 0, th.Colors.Border)
		}

		return wDims
	})

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
