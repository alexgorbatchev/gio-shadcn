/*
Package progress provides an animated progress bar component for gio-shadcn applications.

Progress bars indicate the completion status of a task or process following
shadcn/ui design principles with smooth spring physics transitions.
*/
package progress

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
	"github.com/vibrantgio/effects/spring"
)

// Progress represents an animated linear progress bar component.
type Progress struct {
	Value       float32 // Target value 0.0 to 1.0
	Height      unit.Dp
	Classes     string
	spring      *spring.Spring
	initialized bool
}

// Config represents configuration for creating a Progress component.
type Config struct {
	Value   float32
	Height  unit.Dp
	Classes string
}

// New creates a new Progress component with the given configuration.
func New(config Config) *Progress {
	h := config.Height
	if h <= 0 {
		h = unit.Dp(8)
	}
	val := config.Value
	if val < 0 {
		val = 0
	} else if val > 1 {
		val = 1
	}
	return &Progress{
		Value:   val,
		Height:  h,
		Classes: config.Classes,
	}
}

// Layout renders the progress bar with smooth spring physics animation.
func (p *Progress) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if !p.initialized {
		p.spring = spring.New(float64(p.Value), float64(p.Value), spring.Options{
			Stiffness: 120.0,
			Damping:   18.0,
		})
		p.initialized = true
	}

	// Update physics spring simulation
	p.spring.SetTarget(float64(p.Value))
	p.spring.Tick(60.0)
	currentVal := float32(p.spring.Value())
	if currentVal < 0 {
		currentVal = 0
	} else if currentVal > 1 {
		currentVal = 1
	}

	// Invalidate frame until animation settles
	if !p.spring.Settled(0.001) {
		gtx.Execute(op.InvalidateCmd{})
	}

	heightPx := gtx.Dp(p.Height)
	widthPx := gtx.Constraints.Max.X
	if widthPx <= 0 {
		widthPx = gtx.Dp(unit.Dp(200))
	}

	size := image.Pt(widthPx, heightPx)
	gtx.Constraints = layout.Exact(size)

	trackColor := th.Colors.Muted
	progressColor := th.Colors.Primary

	styles := utils.ParseClasses(p.Classes)
	if styles.Background.A > 0 {
		progressColor = styles.Background
	}

	radius := heightPx / 2
	if styles.Radius > 0 {
		radius = gtx.Dp(styles.Radius)
	}

	// 1. Draw track background
	trackRect := image.Rectangle{Max: size}
	trackRRect := theme.RRect(trackRect, radius)
	theme.DrawRRectBackground(gtx, trackRect, radius, trackColor)

	// 2. Draw animated filled progress bar clipped to current progress width
	filledWidth := int(float32(widthPx) * currentVal)
	if filledWidth > 0 {
		fillTrackClip := trackRRect.Op(gtx.Ops).Push(gtx.Ops)
		fillRect := image.Rect(0, 0, filledWidth, heightPx)
		fillRectClip := clip.Rect(fillRect).Push(gtx.Ops)
		paint.ColorOp{Color: progressColor}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
		fillRectClip.Pop()
		fillTrackClip.Pop()
	}

	// Reset active GPU paint color state
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return layout.Dimensions{Size: size}
}
