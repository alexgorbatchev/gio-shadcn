/*
Package empty provides an empty state placeholder component for gio-shadcn applications.

Empty states inform users when content or search results are unavailable following
shadcn/ui design principles.
*/
package empty

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Empty represents an empty state container component.
type Empty struct {
	Title       string
	Description string
	Classes     string
}

// Config represents configuration for creating an Empty state component.
type Config struct {
	Title       string
	Description string
	Classes     string
}

// New creates a new Empty state component.
func New(config Config) *Empty {
	t := config.Title
	if t == "" {
		t = "No Results Found"
	}
	d := config.Description
	if d == "" {
		d = "Try searching with a different keyword or filter."
	}
	return &Empty{
		Title:       t,
		Description: d,
		Classes:     config.Classes,
	}
}

// Layout renders the empty state title and description.
func (e *Empty) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	bgColor := th.Colors.Card
	borderColor := th.Colors.Border

	styles := utils.ParseClasses(e.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := material.NewTheme()

	padding := layout.Inset{
		Top:    th.Spacing.Space8,
		Bottom: th.Spacing.Space8,
		Left:   th.Spacing.Space8,
		Right:  th.Spacing.Space8,
	}

	return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		contentDims := layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(mTheme, th.Typography.FontSizeBase, e.Title)
				lbl.Color = th.Colors.Foreground
				lbl.Font.Weight = font.Bold
				lbl.Alignment = text.Middle
				return lbl.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(mTheme, th.Typography.FontSizeSM, e.Description)
				lbl.Color = th.Colors.MutedFg
				lbl.Alignment = text.Middle
				return lbl.Layout(gtx)
			}),
		)

		rect := image.Rectangle{Max: contentDims.Size}
		radius := gtx.Dp(th.Radius.RadiusMD)
		rr := clip.UniformRRect(rect, radius)

		paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

		stroke := clip.Stroke{
			Path:  rr.Path(gtx.Ops),
			Width: 1.0,
		}
		paint.FillShape(gtx.Ops, borderColor, stroke.Op())

		return contentDims
	})
}
