/*
Package theme provides a comprehensive theming system for gio-shadcn applications.
*/
package theme

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/vibrantgio/effects/blur"
)

type Theme struct {
	Colors        ColorScheme
	DarkColors    ColorScheme
	Typography    Typography
	Spacing       SpacingScale
	Radius        RadiusScale
	IsDark        bool
	MaterialTheme *material.Theme
	overlays      []layout.Widget
	backdrop      *blur.Backdrop
	backdropValid bool
}

// HasOverlays reports whether there are active window-level overlays queued for rendering.
func (t *Theme) HasOverlays() bool {
	return len(t.overlays) > 0
}

// UpdateBackdropIfNeeded renders a blurred snapshot of the background layer if not already cached.
func (t *Theme) UpdateBackdropIfNeeded(layer func(ops *op.Ops), size image.Point, sigma float64) {
	if t.backdropValid || size.X <= 0 || size.Y <= 0 {
		return
	}
	if t.backdrop == nil {
		t.backdrop = new(blur.Backdrop)
	}
	if err := t.backdrop.Update(layer, size, sigma, blur.WithDivisor(4)); err == nil {
		t.backdropValid = true
	}
}

// InvalidateBackdrop marks the cached backdrop snapshot as invalid so the next overlay opening captures fresh background.
func (t *Theme) InvalidateBackdrop() {
	t.backdropValid = false
}

// BackdropOp returns the cached blurred backdrop ImageOp if valid.
func (t *Theme) BackdropOp() (paint.ImageOp, bool) {
	if t == nil || t.backdrop == nil || !t.backdropValid {
		return paint.ImageOp{}, false
	}
	return t.backdrop.Op()
}

// ReleaseBackdrop frees any allocated offscreen GPU resources held by the backdrop.
func (t *Theme) ReleaseBackdrop() {
	if t.backdrop != nil {
		t.backdrop.Release()
		t.backdropValid = false
	}
}

// AddOverlay registers a window-level overlay (drawer, sheet, modal dialog, command palette)
// to be rendered at the root window level on top of all application content.
func (t *Theme) AddOverlay(w layout.Widget) {
	if w != nil {
		t.overlays = append(t.overlays, w)
	}
}

// RenderOverlays renders all registered window-level overlays across the full window constraints.
func (t *Theme) RenderOverlays(gtx layout.Context) layout.Dimensions {
	defer func() { paint.ColorOp{Color: t.Colors.Background}.Add(gtx.Ops) }()
	if len(t.overlays) == 0 {
		return layout.Dimensions{}
	}
	current := t.overlays
	t.overlays = nil

	for _, overlay := range current {
		overlay(gtx)
	}
	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func New() *Theme {
	mTheme := material.NewTheme()
	mTheme.Shaper = NewGeistShaper()
	mTheme.Face = TypefaceGeist
	colors := LightColorScheme()
	mTheme.Palette.Fg = colors.Foreground
	mTheme.Palette.Bg = colors.Background

	return &Theme{
		Colors:        colors,
		DarkColors:    DarkColorScheme(),
		Typography:    DefaultTypography(),
		Spacing:       DefaultSpacing(),
		Radius:        DefaultRadius(),
		IsDark:        false,
		MaterialTheme: mTheme,
	}
}

func NewDark() *Theme {
	mTheme := material.NewTheme()
	mTheme.Shaper = NewGeistShaper()
	mTheme.Face = TypefaceGeist
	colors := DarkColorScheme()
	mTheme.Palette.Fg = colors.Foreground
	mTheme.Palette.Bg = colors.Background

	return &Theme{
		Colors:        colors,
		DarkColors:    LightColorScheme(),
		Typography:    DefaultTypography(),
		Spacing:       DefaultSpacing(),
		Radius:        DefaultRadius(),
		IsDark:        true,
		MaterialTheme: mTheme,
	}
}

func (t *Theme) ToggleDark() {
	if t.IsDark {
		t.Colors, t.DarkColors = t.DarkColors, t.Colors
		t.IsDark = false
	} else {
		t.Colors, t.DarkColors = t.DarkColors, t.Colors
		t.IsDark = true
	}
	if t.MaterialTheme != nil {
		t.MaterialTheme.Palette.Fg = t.Colors.Foreground
		t.MaterialTheme.Palette.Bg = t.Colors.Background
	}
}

// DrawRRectBackground safely draws a rounded rectangle fill with isolated push/pop clips.
func DrawRRectBackground(gtx layout.Context, rect image.Rectangle, radius int, c color.NRGBA) {
	if c.A == 0 || rect.Dx() <= 0 || rect.Dy() <= 0 {
		return
	}
	rr := RRect(rect, radius)
	cl := rr.Push(gtx.Ops)
	paint.ColorOp{Color: c}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	cl.Pop()
}

// RRect creates a rounded rectangle with radii safe for narrow and short bounds.
func RRect(rect image.Rectangle, radius int) clip.RRect {
	radius = max(0, min(radius, rect.Dx()/2, rect.Dy()/2))
	return clip.UniformRRect(rect, radius)
}

// DrawCornerBackground fills a rectangle with independently rounded corners.
func DrawCornerBackground(gtx layout.Context, rr clip.RRect, c color.NRGBA) {
	if rr.Rect.Dx() <= 0 || rr.Rect.Dy() <= 0 {
		return
	}
	limit := min(rr.Rect.Dx(), rr.Rect.Dy()) / 2
	rr.NW = max(0, min(rr.NW, limit))
	rr.NE = max(0, min(rr.NE, limit))
	rr.SW = max(0, min(rr.SW, limit))
	rr.SE = max(0, min(rr.SE, limit))
	defer rr.Push(gtx.Ops).Pop()
	DrawRRectBackground(gtx, rr.Rect, 0, c)
}

// DrawStroke safely draws a stroke path with isolated push/pop clips.
func DrawStroke(gtx layout.Context, path clip.PathSpec, width float32, c color.NRGBA) {
	if width <= 0 || c.A == 0 {
		return
	}
	stroke := clip.Stroke{
		Path:  path,
		Width: width,
	}
	cl := stroke.Op().Push(gtx.Ops)
	paint.ColorOp{Color: c}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	cl.Pop()
}

// RecordLayout measures a widget while retaining its operations for one replay.
// This prevents measurement from drawing or processing child input twice.
func RecordLayout(gtx layout.Context, w layout.Widget) (layout.Dimensions, op.CallOp) {
	macro := op.Record(gtx.Ops)
	dims := w(gtx)
	return dims, macro.Stop()
}

func ValidateTheme(t *Theme) error {
	if t == nil {
		return fmt.Errorf("theme cannot be nil")
	}

	if err := validateColorScheme(&t.Colors); err != nil {
		return fmt.Errorf("invalid light colors: %w", err)
	}

	if err := validateColorScheme(&t.DarkColors); err != nil {
		return fmt.Errorf("invalid dark colors: %w", err)
	}

	return nil
}

func validateColorScheme(cs *ColorScheme) error {
	if cs == nil {
		return fmt.Errorf("color scheme cannot be nil")
	}

	colors := []struct {
		name  string
		color color.NRGBA
	}{
		{"background", cs.Background},
		{"foreground", cs.Foreground},
		{"primary", cs.Primary},
		{"primary-foreground", cs.PrimaryFg},
		{"secondary", cs.Secondary},
		{"secondary-foreground", cs.SecondaryFg},
		{"border", cs.Border},
	}

	for _, c := range colors {
		if c.color.A == 0 {
			return fmt.Errorf("color %s has zero alpha", c.name)
		}
	}

	return nil
}

type Component interface {
	Layout(gtx layout.Context, theme *Theme) layout.Dimensions
	Update(gtx layout.Context) ComponentState
}

type ComponentState interface {
	IsActive() bool
	IsHovered() bool
	IsPressed() bool
	IsDisabled() bool
}

type Variant string
type Size string

const (
	VariantDefault     Variant = "default"
	VariantDestructive Variant = "destructive"
	VariantOutline     Variant = "outline"
	VariantSecondary   Variant = "secondary"
	VariantGhost       Variant = "ghost"
	VariantLink        Variant = "link"
)

const (
	SizeDefault Size = "default"
	SizeSM      Size = "sm"
	SizeLG      Size = "lg"
	SizeIcon    Size = "icon"
)
