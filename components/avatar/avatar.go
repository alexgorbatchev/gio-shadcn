/*
Package avatar provides a user avatar profile component for gio-shadcn applications.

Avatars display user profile pictures or fallback initials with optional status badges
following shadcn/ui design principles.
*/
package avatar

import (
	"image"
	"strings"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Avatar represents a user profile avatar component.
type Avatar struct {
	Initials   string
	Size       unit.Dp
	ShowBadge  bool
	BadgeColor string
	Classes    string
}

// Config represents configuration for creating an Avatar.
type Config struct {
	Initials   string
	Size       unit.Dp
	ShowBadge  bool
	BadgeColor string
	Classes    string
}

// New creates a new Avatar component with the given configuration.
func New(config Config) *Avatar {
	sz := config.Size
	if sz <= 0 {
		sz = unit.Dp(40)
	}
	return &Avatar{
		Initials:   strings.ToUpper(config.Initials),
		Size:       sz,
		ShowBadge:  config.ShowBadge,
		BadgeColor: config.BadgeColor,
		Classes:    config.Classes,
	}
}

// Layout renders the avatar circle and fallback initials with theme colors.
func (a *Avatar) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	sizePx := gtx.Dp(a.Size)
	size := image.Pt(sizePx, sizePx)
	gtx.Constraints = layout.Exact(size)

	bgColor := th.Colors.Muted
	fgColor := th.Colors.MutedFg
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(a.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	// Draw outer circular background
	rect := image.Rectangle{Max: size}
	ellipse := clip.Ellipse(rect)
	paint.FillShape(gtx.Ops, bgColor, ellipse.Op(gtx.Ops))

	// Draw border
	stroke := clip.Stroke{
		Path:  ellipse.Path(gtx.Ops),
		Width: 1.0,
	}
	paint.FillShape(gtx.Ops, borderColor, stroke.Op())

	// Draw centered initials
	mTheme := material.NewTheme()
	lblFontSize := th.Typography.FontSizeSM
	if sizePx > 48 {
		lblFontSize = th.Typography.FontSizeBase
	}

	lbl := material.Label(mTheme, lblFontSize, a.Initials)
	lbl.Color = fgColor
	lbl.Font.Weight = font.Medium
	lbl.Alignment = text.Middle

	_ = layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return lbl.Layout(gtx)
	})

	// Draw status badge dot at bottom right if enabled
	if a.ShowBadge {
		badgeSize := sizePx / 4
		if badgeSize < 8 {
			badgeSize = 8
		}

		badgeMin := image.Pt(sizePx-badgeSize, sizePx-badgeSize)
		badgeMax := image.Pt(sizePx, sizePx)
		badgeRect := image.Rectangle{Min: badgeMin, Max: badgeMax}

		badgeEllipse := clip.Ellipse(badgeRect)
		bColor := th.Colors.Primary
		if a.BadgeColor == "green" {
			bColor = th.Colors.Accent
		}

		paint.FillShape(gtx.Ops, bColor, badgeEllipse.Op(gtx.Ops))

		// Border around badge
		bStroke := clip.Stroke{
			Path:  badgeEllipse.Path(gtx.Ops),
			Width: 1.0,
		}
		paint.FillShape(gtx.Ops, th.Colors.Background, bStroke.Op())
	}

	return layout.Dimensions{Size: size}
}
