/*
Package switchcomp provides an interactive toggle switch component for gio-shadcn applications.

Switches toggle between on and off states following shadcn/ui design principles.
*/
package switchcomp

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Switch represents a toggle switch component.
type Switch struct {
	clickable *widget.Clickable

	Value    bool
	Disabled bool
	Classes  string
	OnChange func(bool)
}

// Config represents configuration for creating a Switch.
type Config struct {
	Value    bool
	Disabled bool
	Classes  string
	OnChange func(bool)
}

// New creates a new Switch component with the given configuration.
func New(config Config) *Switch {
	return &Switch{
		clickable: new(widget.Clickable),
		Value:     config.Value,
		Disabled:  config.Disabled,
		Classes:   config.Classes,
		OnChange:  config.OnChange,
	}
}

// Layout renders the switch with the given graphics context and theme.
func (s *Switch) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if s.clickable.Clicked(gtx) && !s.Disabled {
		s.Value = !s.Value
		if s.OnChange != nil {
			s.OnChange(s.Value)
		}
	}

	trackWidth := gtx.Dp(unit.Dp(36))
	trackHeight := gtx.Dp(unit.Dp(20))
	thumbSize := gtx.Dp(unit.Dp(16))
	padding := gtx.Dp(unit.Dp(2))

	size := image.Pt(trackWidth, trackHeight)
	gtx.Constraints = layout.Exact(size)

	// Determine colors
	trackColor := th.Colors.Muted
	thumbColor := th.Colors.MutedFg

	if s.Value {
		trackColor = th.Colors.Primary
		thumbColor = th.Colors.PrimaryFg
	}

	if s.Disabled {
		trackColor.A = 128
		thumbColor.A = 128
	}

	styles := utils.ParseClasses(s.Classes)
	if styles.Background.A > 0 {
		trackColor = styles.Background
	}

	dims := s.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// Draw track background FIRST
		trackRect := image.Rectangle{Max: size}
		trackRadius := trackHeight / 2
		theme.DrawRRectBackground(gtx, trackRect, trackRadius, trackColor)

		// Draw border if unselected
		if !s.Value {
			rrTrack := clip.UniformRRect(trackRect, trackRadius)
			theme.DrawStroke(gtx, rrTrack.Path(gtx.Ops), 1.0, th.Colors.Border)
		}

		// Calculate thumb position
		thumbX := padding
		if s.Value {
			thumbX = trackWidth - thumbSize - padding
		}
		thumbY := (trackHeight - thumbSize) / 2

		thumbMin := image.Pt(thumbX, thumbY)
		thumbMax := image.Pt(thumbX+thumbSize, thumbY+thumbSize)
		thumbRect := image.Rectangle{Min: thumbMin, Max: thumbMax}
		thumbRadius := thumbSize / 2

		// Draw thumb knob ON TOP of track
		theme.DrawRRectBackground(gtx, thumbRect, thumbRadius, thumbColor)

		return layout.Dimensions{Size: size}
	})

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

// Clicked returns true if the switch was toggled in the current frame.
func (s *Switch) Clicked(gtx layout.Context) bool {
	return s.clickable.Clicked(gtx) && !s.Disabled
}

func _(c color.NRGBA) {} // unused color reference guard
