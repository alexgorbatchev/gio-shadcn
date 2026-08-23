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

// Layout renders the empty state title and description with background drawn first.
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

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space8,
		Bottom: th.Spacing.Space8,
		Left:   th.Spacing.Space8,
		Right:  th.Spacing.Space8,
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
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
		})
	}

	contentDims := renderContent(gtxContent)
	emptySize := contentDims.Size

	dims := layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: emptySize}
			radius := gtx.Dp(th.Radius.RadiusMD)

			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: emptySize}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)

	// Reset active GPU paint color state
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
