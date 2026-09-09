/*
Package spinner provides a loading activity indicator component for gio-shadcn applications.

Spinners indicate background processing following
shadcn/ui design principles with continuous smooth rotation animations.
*/
package spinner

import (
	"image"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
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

// Layout renders the smoothly rotating circular arc spinner.
func (s *Spinner) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	// Invalidate frame for continuous fluid 60/120 FPS rotation
	gtx.Execute(op.InvalidateCmd{})

	sizePx := gtx.Dp(s.Size)
	size := image.Pt(sizePx, sizePx)
	gtx.Constraints = layout.Exact(size)

	spinColor := th.Colors.Primary
	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		spinColor = styles.Background
	}

	center := float32(sizePx) / 2.0
	radius := center - float32(gtx.Dp(unit.Dp(2.5)))
	strokeWidth := float32(gtx.Dp(unit.Dp(2.5)))

	var p clip.Path
	p.Begin(gtx.Ops)

	// Continuous time-driven rotation angle
	nanos := gtx.Now.UnixNano()
	spinPhase := float64(nanos%1_000_000_000) / 1_000_000_000.0 * 2 * math.Pi

	// Draw 270 degree rotating arc
	first := true
	for a := 0.0; a <= 1.5*math.Pi; a += 0.1 {
		angle := a + spinPhase
		x := center + radius*float32(math.Cos(angle))
		y := center + radius*float32(math.Sin(angle))
		if first {
			p.MoveTo(f32.Pt(x, y))
			first = false
		} else {
			p.LineTo(f32.Pt(x, y))
		}
	}

	theme.DrawStroke(gtx, p.End(), strokeWidth, spinColor)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return layout.Dimensions{Size: size}
}
