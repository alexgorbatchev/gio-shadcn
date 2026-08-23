/*
Package pagination provides page navigation controls for gio-shadcn applications.

Pagination controls split large datasets into numbered pages following
shadcn/ui design principles.
*/
package pagination

import (
	"fmt"
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

// Pagination represents page navigation controls.
type Pagination struct {
	CurrentPage  int
	TotalPages   int
	OnSelectPage func(page int)

	prevBtn *widget.Clickable
	nextBtn *widget.Clickable
}

// Config represents configuration for creating Pagination controls.
type Config struct {
	CurrentPage  int
	TotalPages   int
	OnSelectPage func(page int)
}

// New creates a new Pagination component.
func New(config Config) *Pagination {
	curr := config.CurrentPage
	if curr < 1 {
		curr = 1
	}
	tot := config.TotalPages
	if tot < 1 {
		tot = 1
	}
	return &Pagination{
		CurrentPage:  curr,
		TotalPages:   tot,
		OnSelectPage: config.OnSelectPage,
		prevBtn:      new(widget.Clickable),
		nextBtn:      new(widget.Clickable),
	}
}

// Layout renders page number buttons and prev/next controls.
func (p *Pagination) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	if p.prevBtn.Clicked(gtx) && p.CurrentPage > 1 {
		p.CurrentPage--
		if p.OnSelectPage != nil {
			p.OnSelectPage(p.CurrentPage)
		}
	}

	if p.nextBtn.Clicked(gtx) && p.CurrentPage < p.TotalPages {
		p.CurrentPage++
		if p.OnSelectPage != nil {
			p.OnSelectPage(p.CurrentPage)
		}
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	children := make([]layout.FlexChild, 0, p.TotalPages+2)

	// Prev Button
	children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return p.prevBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return p.layoutPageButton(gtx, th, mTheme, "‹ Prev", false, p.CurrentPage <= 1)
		})
	}))

	// Page Number Buttons
	for i := 1; i <= p.TotalPages; i++ {
		pageNum := i
		isActive := pageNum == p.CurrentPage

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := new(widget.Clickable)
			if btn.Clicked(gtx) {
				p.CurrentPage = pageNum
				if p.OnSelectPage != nil {
					p.OnSelectPage(pageNum)
				}
			}
			return btn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return p.layoutPageButton(gtx, th, mTheme, fmt.Sprintf("%d", pageNum), isActive, false)
			})
		}))
	}

	// Next Button
	children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return p.nextBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return p.layoutPageButton(gtx, th, mTheme, "Next ›", false, p.CurrentPage >= p.TotalPages)
		})
	}))

	dims := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (p *Pagination) layoutPageButton(gtx layout.Context, th *theme.Theme, mTheme *material.Theme, labelText string, active, disabled bool) layout.Dimensions {
	bgColor := th.Colors.Muted
	fgColor := th.Colors.MutedFg

	if active {
		bgColor = th.Colors.Primary
		fgColor = th.Colors.PrimaryFg
	}

	if disabled {
		fgColor.A = 100
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space2,
		Bottom: th.Spacing.Space2,
		Left:   th.Spacing.Space3,
		Right:  th.Spacing.Space3,
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(mTheme, th.Typography.FontSizeSM, labelText)
			lbl.Color = fgColor
			if active {
				lbl.Font.Weight = font.Medium
			}
			return lbl.Layout(gtx)
		})
	}

	contentDims := renderContent(gtxContent)
	btnSize := contentDims.Size

	return layout.Stack{}.Layout(gtx,
		// Background drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: btnSize}
			radius := gtx.Dp(th.Radius.RadiusSM)
			theme.DrawRRectBackground(gtx, rect, radius, bgColor)
			return layout.Dimensions{Size: btnSize}
		}),

		// Label content drawn ON TOP of background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderContent(gtx)
		}),
	)
}
