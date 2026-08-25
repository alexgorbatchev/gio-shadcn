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
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Empty represents an empty state container component.
type Empty struct {
	Title       string
	Description string
	Classes     string
	Icon        *lucide.Icon
	Action      layout.Widget
}

// Config represents configuration for creating an Empty state component.
type Config struct {
	Title       string
	Description string
	Classes     string
	Icon        *lucide.Icon
	Action      layout.Widget
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
	ic := config.Icon
	if ic == nil {
		ic = lucide.Search
	}
	return &Empty{
		Title:       t,
		Description: d,
		Classes:     config.Classes,
		Icon:        ic,
		Action:      config.Action,
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
		Left:   th.Spacing.Space6,
		Right:  th.Spacing.Space6,
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if e.Icon != nil {
						return e.Icon.LayoutSize(gtx, unit.Dp(36), th.Colors.MutedFg)
					}
					return layout.Dimensions{}
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeLG, e.Title)
					lbl.Color = th.Colors.Foreground
					lbl.Font.Weight = font.Bold
					lbl.Alignment = text.Middle
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, e.Description)
					lbl.Color = th.Colors.MutedFg
					lbl.Alignment = text.Middle
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if e.Action != nil {
						return layout.Inset{Top: th.Spacing.Space4}.Layout(gtx, e.Action)
					}
					return layout.Dimensions{}
				}),
			)
		})
	}

	contentDims := renderContent(gtxContent)
	emptySize := contentDims.Size

	dims := layout.Stack{}.Layout(gtx,
		// Background drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: emptySize}
			radius := gtx.Dp(th.Radius.RadiusLG)
			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: emptySize}
		}),

		// Content drawn ON TOP of background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
