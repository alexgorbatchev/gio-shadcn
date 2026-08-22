/*
Package theme provides a comprehensive theming system for gio-shadcn applications.
*/
package theme

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Theme struct {
	Colors        ColorScheme
	DarkColors    ColorScheme
	Typography    Typography
	Spacing       SpacingScale
	Radius        RadiusScale
	IsDark        bool
	MaterialTheme *material.Theme
}

func New() *Theme {
	mTheme := material.NewTheme()
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
