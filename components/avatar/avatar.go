/*
Package avatar provides a user avatar profile component for gio-shadcn applications.
*/
package avatar

import (
	"image"
	"strings"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

type Avatar struct {
	Initials   string
	Size       unit.Dp
	ShowBadge  bool
	BadgeColor string
	Classes    string
}

type Config struct {
	Initials   string
	Size       unit.Dp
	ShowBadge  bool
	BadgeColor string
	Classes    string
}

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

	rect := image.Rectangle{Max: size}
	ellipse := clip.Ellipse(rect)

	// Draw background and border safely
	bgClip := ellipse.Push(gtx.Ops)
	theme.DrawRRectBackground(gtx, rect, sizePx/2, bgColor)
	bgClip.Pop()

	theme.DrawStroke(gtx, ellipse.Path(gtx.Ops), 1.0, borderColor)

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}
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

		bClip := badgeEllipse.Push(gtx.Ops)
		theme.DrawRRectBackground(gtx, badgeRect, badgeSize/2, bColor)
		bClip.Pop()

		theme.DrawStroke(gtx, badgeEllipse.Path(gtx.Ops), 1.0, th.Colors.Background)
	}

	return layout.Dimensions{Size: size}
}
