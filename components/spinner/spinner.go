/*
Package spinner provides a loading activity indicator component for gio-shadcn applications.

Spinners indicate background processing following
shadcn/ui design principles.
*/
package spinner

import (
	"image"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Spinner represents a circular loading indicator.
type Spinner struct {
	Size    unit.Dp
	Classes string
}

// Config represents configuration for creating a Spinner.
type Config struct {
	Size    unit.Dp
	Classes string
}

// New creates a new Spinner component.
func New(config Config) *Spinner {
	sz := config.Size
	if sz <= 0 {
		sz = unit.Dp(24)
	}
	return &Spinner{
		Size:    sz,
		Classes: config.Classes,
	}
}

// Layout renders the circular arc spinner.
func (s *Spinner) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	sizePx := gtx.Dp(s.Size)
	size := image.Pt(sizePx, sizePx)
	gtx.Constraints = layout.Exact(size)

	spinColor := th.Colors.Primary
	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		spinColor = styles.Background
	}

	// Draw arc stroke path
	center := float32(sizePx) / 2.0
	radius := center - float32(gtx.Dp(unit.Dp(2)))
	strokeWidth := float32(gtx.Dp(unit.Dp(2.5)))

	var p clip.Path
	p.Begin(gtx.Ops)

	// Draw 270 degree arc
	startAngle := float64(0)
	endAngle := float64(270) * (math.Pi / 180.0)

	first := true
	for a := startAngle; a <= endAngle; a += 0.1 {
		x := center + radius*float32(math.Cos(a))
		y := center + radius*float32(math.Sin(a))
		if first {
			p.MoveTo(f32.Pt(x, y))
			first = false
		} else {
			p.LineTo(f32.Pt(x, y))
		}
	}

	stroke := clip.Stroke{
		Path:  p.End(),
		Width: strokeWidth,
	}
	paint.FillShape(gtx.Ops, spinColor, stroke.Op())

	return layout.Dimensions{Size: size}
}
