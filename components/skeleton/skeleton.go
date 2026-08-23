/*
Package skeleton provides a shimmer placeholder component for gio-shadcn applications.

Skeletons display placeholder shapes while content is loading following
shadcn/ui design principles.
*/
package skeleton

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Skeleton represents a shimmer loading placeholder component.
type Skeleton struct {
	Width   unit.Dp
	Height  unit.Dp
	Classes string
}

// Config represents configuration for creating a Skeleton.
type Config struct {
	Width   unit.Dp
	Height  unit.Dp
	Classes string
}

// New creates a new Skeleton placeholder.
func New(config Config) *Skeleton {
	w := config.Width
	if w <= 0 {
		w = unit.Dp(100)
	}
	h := config.Height
	if h <= 0 {
		h = unit.Dp(20)
	}
	return &Skeleton{
		Width:   w,
		Height:  h,
		Classes: config.Classes,
	}
}

// Layout renders the placeholder box.
func (s *Skeleton) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	wPx := gtx.Dp(s.Width)
	hPx := gtx.Dp(s.Height)
	size := image.Pt(wPx, hPx)
	gtx.Constraints = layout.Exact(size)

	bgColor := th.Colors.Muted
	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	rect := image.Rectangle{Max: size}
	radius := gtx.Dp(th.Radius.RadiusMD)

	theme.DrawRRectBackground(gtx, rect, radius, bgColor)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return layout.Dimensions{Size: size}
}
