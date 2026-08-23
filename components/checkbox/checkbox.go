/*
Package checkbox provides an interactive checkbox component for gio-shadcn applications.

Checkboxes allow users to select one or more items from a set following
shadcn/ui design principles.
*/
package checkbox

import (
	"image"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Checkbox represents an interactive checkbox component.
type Checkbox struct {
	clickable *widget.Clickable

	Value    bool
	Disabled bool
	Classes  string
	OnChange func(bool)
}

// Config represents configuration for creating a Checkbox.
type Config struct {
	Value    bool
	Disabled bool
	Classes  string
	OnChange func(bool)
}

// New creates a new Checkbox component with the given configuration.
func New(config Config) *Checkbox {
	return &Checkbox{
		clickable: new(widget.Clickable),
		Value:     config.Value,
		Disabled:  config.Disabled,
		Classes:   config.Classes,
		OnChange:  config.OnChange,
	}
}

// Layout renders the checkbox with the given graphics context and theme.
func (c *Checkbox) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if c.clickable.Clicked(gtx) && !c.Disabled {
		c.Value = !c.Value
		if c.OnChange != nil {
			c.OnChange(c.Value)
		}
	}

	boxSize := gtx.Dp(unit.Dp(16))
	size := image.Pt(boxSize, boxSize)
	gtx.Constraints = layout.Exact(size)

	bgColor := th.Colors.Background
	fgColor := th.Colors.PrimaryFg
	borderColor := th.Colors.Border

	if c.Value {
		bgColor = th.Colors.Primary
		borderColor = th.Colors.Primary
	}

	if c.Disabled {
		bgColor.A = 128
		borderColor.A = 128
		fgColor.A = 128
	}

	styles := utils.ParseClasses(c.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	dims := c.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// Draw box background FIRST
		rect := image.Rectangle{Max: size}
		radius := boxSize / 2
		if radius <= 0 {
			radius = 1
		}

		theme.DrawRRectBackground(gtx, rect, radius, bgColor)

		rr := clip.UniformRRect(rect, radius)
		theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

		// Draw checkmark vector path ON TOP when checked
		if c.Value {
			var p clip.Path
			p.Begin(gtx.Ops)

			// Scale checkmark relative to boxSize
			sz := float32(boxSize)
			p.MoveTo(f32.Pt(sz*0.25, sz*0.50))
			p.LineTo(f32.Pt(sz*0.45, sz*0.72))
			p.LineTo(f32.Pt(sz*0.75, sz*0.28))

			checkStroke := clip.Stroke{
				Path:  p.End(),
				Width: float32(gtx.Dp(unit.Dp(2))),
			}
			cl := checkStroke.Op().Push(gtx.Ops)
			paint.ColorOp{Color: fgColor}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			cl.Pop()
		}

		return layout.Dimensions{Size: size}
	})

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
