/*
Package progress provides a progress bar component for gio-shadcn applications.

Progress bars indicate the completion status of a task or process following
shadcn/ui design principles.
*/
package progress

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Progress represents a linear progress bar component.
type Progress struct {
	Value   float32 // 0.0 to 1.0
	Height  unit.Dp
	Classes string
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

// Layout renders the progress bar with the given graphics context and theme.
func (p *Progress) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	heightPx := gtx.Dp(p.Height)
	widthPx := gtx.Constraints.Max.X
	if widthPx <= 0 {
		widthPx = gtx.Dp(unit.Dp(200))
	}

	size := image.Pt(widthPx, heightPx)
	gtx.Constraints = layout.Exact(size)

	// Background track color and foreground progress color
	trackColor := th.Colors.Muted
	progressColor := th.Colors.Primary

	styles := utils.ParseClasses(p.Classes)
	if styles.Background.A > 0 {
		progressColor = styles.Background
	}

	radius := gtx.Dp(th.Radius.RadiusFull)

	// Draw background track
	trackRect := image.Rectangle{Max: size}
	theme.DrawRRectBackground(gtx, trackRect, radius, trackColor)

	// Draw filled progress bar
	if p.Value > 0 {
		progWidth := int(float32(widthPx) * p.Value)
		if progWidth > 0 {
			progRect := image.Rectangle{Max: image.Pt(progWidth, heightPx)}
			theme.DrawRRectBackground(gtx, progRect, radius, progressColor)
		}
	}

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return layout.Dimensions{Size: size}
}
