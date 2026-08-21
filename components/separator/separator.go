/*
Package separator provides a divider line component for gio-shadcn applications.

Separators visually separate content in lists and layouts following
shadcn/ui design principles.
*/
package separator

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Separator represents a divider line component.
type Separator struct {
	Horizontal bool
	Thickness  unit.Dp
	Classes    string
}

// Config represents configuration for creating a Separator component.
type Config struct {
	Horizontal bool
	Thickness  unit.Dp
	Classes    string
}

// New creates a new Separator component with the given configuration.
func New(config Config) *Separator {
	t := config.Thickness
	if t <= 0 {
		t = unit.Dp(1)
	}
	return &Separator{
		Horizontal: config.Horizontal,
		Thickness:  t,
		Classes:    config.Classes,
	}
}

// Layout renders the separator line with the given graphics context and theme.
func (s *Separator) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	thickPx := gtx.Dp(s.Thickness)
	if thickPx < 1 {
		thickPx = 1
	}

	var size image.Point
	if s.Horizontal {
		widthPx := gtx.Constraints.Max.X
		if widthPx <= 0 {
			widthPx = gtx.Dp(unit.Dp(100))
		}
		size = image.Pt(widthPx, thickPx)
	} else {
		heightPx := gtx.Constraints.Max.Y
		if heightPx <= 0 {
			heightPx = gtx.Dp(unit.Dp(100))
		}
		size = image.Pt(thickPx, heightPx)
	}

	gtx.Constraints = layout.Exact(size)

	lineColor := th.Colors.Border
	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		lineColor = styles.Background
	}

	rect := image.Rectangle{Max: size}
	paint.FillShape(gtx.Ops, lineColor, clip.Rect(rect).Op())

	return layout.Dimensions{Size: size}
}
