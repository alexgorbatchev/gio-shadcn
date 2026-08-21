/*
Package slider provides an interactive range slider component for gio-shadcn applications.

Sliders allow users to select a numeric value within a range following
shadcn/ui design principles.
*/
package slider

import (
	"image"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Slider represents an interactive range slider component.
type Slider struct {
	Value    float32
	Min      float32
	Max      float32
	Disabled bool
	Classes  string
	OnChange func(float32)

	dragID pointer.ID
}

// Config represents configuration for creating a Slider.
type Config struct {
	Value    float32
	Min      float32
	Max      float32
	Disabled bool
	Classes  string
	OnChange func(float32)
}

// New creates a new Slider component with the given configuration.
func New(config Config) *Slider {
	minVal := config.Min
	maxVal := config.Max
	if maxVal <= minVal {
		maxVal = 1.0
		minVal = 0.0
	}
	val := config.Value
	if val < minVal {
		val = minVal
	} else if val > maxVal {
		val = maxVal
	}
	return &Slider{
		Value:    val,
		Min:      minVal,
		Max:      maxVal,
		Disabled: config.Disabled,
		Classes:  config.Classes,
		OnChange: config.OnChange,
	}
}

// Layout renders the slider track and thumb with mouse/pointer drag interaction.
func (s *Slider) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	trackHeight := gtx.Dp(unit.Dp(6))
	thumbRadius := gtx.Dp(unit.Dp(8))
	widthPx := gtx.Constraints.Max.X
	if widthPx <= 0 {
		widthPx = gtx.Dp(unit.Dp(200))
	}
	heightPx := thumbRadius * 2

	size := image.Pt(widthPx, heightPx)
	gtx.Constraints = layout.Exact(size)

	// Process pointer drag/click events
	if !s.Disabled {
		area := clip.Rect(image.Rectangle{Max: size}).Push(gtx.Ops)
		event.Op(gtx.Ops, s)
		area.Pop()

		for {
			ev, ok := gtx.Event(pointer.Filter{
				Target: s,
				Kinds:  pointer.Press | pointer.Drag | pointer.Release,
			})
			if !ok {
				break
			}
			if pEv, ok := ev.(pointer.Event); ok {
				if pEv.Kind == pointer.Press || pEv.Kind == pointer.Drag {
					ratio := pEv.Position.X / float32(widthPx)
					if ratio < 0 {
						ratio = 0
					} else if ratio > 1 {
						ratio = 1
					}
					s.Value = s.Min + ratio*(s.Max-s.Min)
					if s.OnChange != nil {
						s.OnChange(s.Value)
					}
				}
			}
		}
	}

	// Calculate filled progress fraction
	ratio := (s.Value - s.Min) / (s.Max - s.Min)
	if ratio < 0 {
		ratio = 0
	} else if ratio > 1 {
		ratio = 1
	}

	trackColor := th.Colors.Muted
	progressColor := th.Colors.Primary
	thumbColor := th.Colors.PrimaryFg

	if s.Disabled {
		trackColor.A = 128
		progressColor.A = 128
		thumbColor.A = 128
	}

	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		progressColor = styles.Background
	}

	// Draw track background (centered vertically)
	trackY := (heightPx - trackHeight) / 2
	trackRect := image.Rectangle{
		Min: image.Pt(0, trackY),
		Max: image.Pt(widthPx, trackY+trackHeight),
	}
	rrTrack := clip.UniformRRect(trackRect, trackHeight/2)
	paint.FillShape(gtx.Ops, trackColor, rrTrack.Op(gtx.Ops))

	// Draw filled portion
	progWidth := int(float32(widthPx) * ratio)
	if progWidth > 0 {
		progRect := image.Rectangle{
			Min: image.Pt(0, trackY),
			Max: image.Pt(progWidth, trackY+trackHeight),
		}
		rrProg := clip.UniformRRect(progRect, trackHeight/2)
		paint.FillShape(gtx.Ops, progressColor, rrProg.Op(gtx.Ops))
	}

	// Draw thumb circle
	thumbX := int(float32(widthPx) * ratio)
	thumbMin := image.Pt(thumbX-thumbRadius, 0)
	thumbMax := image.Pt(thumbX+thumbRadius, heightPx)
	thumbRect := image.Rectangle{Min: thumbMin, Max: thumbMax}

	ellipseThumb := clip.Ellipse(thumbRect)
	paint.FillShape(gtx.Ops, thumbColor, ellipseThumb.Op(gtx.Ops))

	// Draw thumb border
	stroke := clip.Stroke{
		Path:  ellipseThumb.Path(gtx.Ops),
		Width: 1.5,
	}
	paint.FillShape(gtx.Ops, th.Colors.Primary, stroke.Op())

	return layout.Dimensions{Size: size}
}
