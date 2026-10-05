package theme

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/vibrantgio/effects/blur"
)

// Dimmer provides a reusable modal backdrop dimmer with optional backdrop blur that intercepts outside clicks to close overlays.
type Dimmer struct {
	clickable widget.Clickable
	backdrop  *blur.Backdrop
}

// NewDimmer creates a new reusable Dimmer.
func NewDimmer() *Dimmer {
	return &Dimmer{
		backdrop: new(blur.Backdrop),
	}
}

// UpdateBackdrop renders a blurred snapshot of the background layer behind the modal overlay.
func (d *Dimmer) UpdateBackdrop(layer func(ops *op.Ops), fullSize image.Point, sigma float64) error {
	if d.backdrop == nil {
		d.backdrop = new(blur.Backdrop)
	}
	return d.backdrop.Update(layer, fullSize, sigma, blur.WithDivisor(4))
}

// Release frees any offscreen GPU textures allocated by the backdrop.
func (d *Dimmer) Release() {
	if d.backdrop != nil {
		d.backdrop.Release()
	}
}

// Layout renders the dimmed backdrop with default alpha (160) filling the available constraints and invokes onClose when clicked.
func (d *Dimmer) Layout(gtx layout.Context, th *Theme, onClose func()) layout.Dimensions {
	return d.LayoutWithAlpha(gtx, th, 160, onClose)
}

// LayoutWithAlpha renders the dimmed backdrop with custom alpha (0-255) filling the available constraints and invokes onClose when clicked.
func (d *Dimmer) LayoutWithAlpha(gtx layout.Context, th *Theme, alpha uint8, onClose func()) layout.Dimensions {
	for d.clickable.Clicked(gtx) {
		if onClose != nil {
			onClose()
		}
	}

	return d.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		rect := image.Rectangle{Max: gtx.Constraints.Max}

		// 1. Render blurred background snapshot from theme or dimmer if available
		var imgOp paint.ImageOp
		var ok bool
		if th != nil {
			imgOp, ok = th.BackdropOp()
		}
		if !ok && d.backdrop != nil {
			imgOp, ok = d.backdrop.Op()
		}

		if ok {
			sz := imgOp.Size()
			if sz.X > 0 && sz.Y > 0 && rect.Max.X > 0 && rect.Max.Y > 0 {
				imgClip := clip.Rect(rect).Push(gtx.Ops)
				scaleX := float32(rect.Max.X) / float32(sz.X)
				scaleY := float32(rect.Max.Y) / float32(sz.Y)
				scale := f32.Affine2D{}.Scale(f32.Point{}, f32.Pt(scaleX, scaleY))
				tr := op.Affine(scale).Push(gtx.Ops)
				imgOp.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
				tr.Pop()
				imgClip.Pop()
			}
		}

		// 2. Render semi-transparent dark tint scrim over the background
		if alpha > 0 {
			dimAlpha := alpha
			if ok && dimAlpha > 110 {
				dimAlpha = 110 // ~43% dark tint when blur is active so the frosted glass effect is vivid
			}
			backdropColor := color.NRGBA{R: 0, G: 0, B: 0, A: dimAlpha}
			DrawRRectBackground(gtx, rect, 0, backdropColor)
		}

		return layout.Dimensions{Size: gtx.Constraints.Max}
	})
}
