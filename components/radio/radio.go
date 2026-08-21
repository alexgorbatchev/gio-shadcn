/*
Package radio provides an interactive radio button component for gio-shadcn applications.

Radio buttons allow users to select a single option from a set following
shadcn/ui design principles.
*/
package radio

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Radio represents an interactive radio button component.
type Radio struct {
	clickable *widget.Clickable

	Selected bool
	Disabled bool
	Classes  string
	OnChange func(bool)
}

// Config represents configuration for creating a Radio button.
type Config struct {
	Selected bool
	Disabled bool
	Classes  string
	OnChange func(bool)
}

// New creates a new Radio button component with the given configuration.
func New(config Config) *Radio {
	return &Radio{
		clickable: new(widget.Clickable),
		Selected:  config.Selected,
		Disabled:  config.Disabled,
		Classes:   config.Classes,
		OnChange:  config.OnChange,
	}
}

// Layout renders the radio circle and inner selected dot.
func (r *Radio) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if r.clickable.Clicked(gtx) && !r.Disabled {
		r.Selected = true
		if r.OnChange != nil {
			r.OnChange(true)
		}
	}

	boxSize := gtx.Dp(unit.Dp(16))
	size := image.Pt(boxSize, boxSize)
	gtx.Constraints = layout.Exact(size)

	borderColor := th.Colors.Border
	dotColor := th.Colors.Primary

	if r.Selected {
		borderColor = th.Colors.Primary
	}

	if r.Disabled {
		borderColor.A = 128
		dotColor.A = 128
	}

	styles := utils.ParseClasses(r.Classes)
	if styles.Background.A > 0 {
		dotColor = styles.Background
	}

	return r.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// Draw outer circle
		rect := image.Rectangle{Max: size}
		ellipse := clip.Ellipse(rect)

		// Outer stroke
		stroke := clip.Stroke{
			Path:  ellipse.Path(gtx.Ops),
			Width: 1.5,
		}
		paint.FillShape(gtx.Ops, borderColor, stroke.Op())

		// Draw inner dot if selected
		if r.Selected {
			dotSize := boxSize / 2
			dotMin := image.Pt(boxSize/4, boxSize/4)
			dotMax := image.Pt(boxSize/4+dotSize, boxSize/4+dotSize)
			dotRect := image.Rectangle{Min: dotMin, Max: dotMax}

			dotEllipse := clip.Ellipse(dotRect)
			paint.FillShape(gtx.Ops, dotColor, dotEllipse.Op(gtx.Ops))
		}

		return layout.Dimensions{Size: size}
	})
}
