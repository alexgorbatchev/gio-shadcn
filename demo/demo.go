/*
Package demo provides the full interactive 37-component gallery showcase application
for gio-shadcn following shadcn/ui design tokens and Gio immediate-mode rendering.
*/
package demo

import (
	"image"
	"log"
	"os"
	"sort"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/bnema/gio-shadcn/components/accordion"
	"github.com/bnema/gio-shadcn/components/alert"
	"github.com/bnema/gio-shadcn/components/aspectratio"
	"github.com/bnema/gio-shadcn/components/avatar"
	"github.com/bnema/gio-shadcn/components/badge"
	"github.com/bnema/gio-shadcn/components/breadcrumb"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/components/carousel"
	"github.com/bnema/gio-shadcn/components/checkbox"
	"github.com/bnema/gio-shadcn/components/collapsible"
	"github.com/bnema/gio-shadcn/components/command"
	"github.com/bnema/gio-shadcn/components/dialog"
	"github.com/bnema/gio-shadcn/components/drawer"
	"github.com/bnema/gio-shadcn/components/dropdownmenu"
	"github.com/bnema/gio-shadcn/components/empty"
	"github.com/bnema/gio-shadcn/components/hovercard"
	"github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/components/inputotp"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/components/numberinput"
	"github.com/bnema/gio-shadcn/components/pagination"
	"github.com/bnema/gio-shadcn/components/popover"
	"github.com/bnema/gio-shadcn/components/progress"
	"github.com/bnema/gio-shadcn/components/radio"
	"github.com/bnema/gio-shadcn/components/resizable"
	"github.com/bnema/gio-shadcn/components/scrollarea"
	selectcomp "github.com/bnema/gio-shadcn/components/select"
	"github.com/bnema/gio-shadcn/components/separator"
	"github.com/bnema/gio-shadcn/components/sheet"
	"github.com/bnema/gio-shadcn/components/skeleton"
	"github.com/bnema/gio-shadcn/components/slider"
	"github.com/bnema/gio-shadcn/components/spinner"
	switchcomp "github.com/bnema/gio-shadcn/components/switch"
	"github.com/bnema/gio-shadcn/components/table"
	"github.com/bnema/gio-shadcn/components/tabs"
	"github.com/bnema/gio-shadcn/components/textarea"
	"github.com/bnema/gio-shadcn/components/titlebar"
	"github.com/bnema/gio-shadcn/components/toast"
	"github.com/bnema/gio-shadcn/components/togglegroup"
	"github.com/bnema/gio-shadcn/components/tooltip"
	"github.com/bnema/gio-shadcn/components/tree"
	"github.com/bnema/gio-shadcn/theme"
)

type GalleryItem struct {
	ID        string
	Category  string
	Name      string
	clickable *widget.Clickable
}

var sidebarList = &widget.List{
	List: layout.List{
		Axis: layout.Vertical,
	},
}

var contentList = &widget.List{
	List: layout.List{
		Axis: layout.Vertical,
	},
}

// Run launches the full gio-shadcn component gallery window.
func Run() {
	go func() {
		w := &app.Window{}
		w.Option(app.Title("guipoc - Gio (gio-shadcn Gallery)"))
		w.Option(app.Size(1200, 800))
		w.Option(app.Maximized.Option())

		err := runWindow(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func runWindow(w *app.Window) error {
	th := theme.NewDark()

	tb := titlebar.NewTitleBar(
		titlebar.WithTitle("guipoc - Gio Component Gallery (37 Component Parity)"),
		titlebar.WithWindow(w),
		titlebar.WithVariant(theme.VariantSecondary),
	)

	// Gallery Component List (1 View Per Component - 37 Total)
	galleryItems := []*GalleryItem{
		// General
		{ID: "button", Category: "General", Name: "Button", clickable: new(widget.Clickable)},
		{ID: "badge", Category: "General", Name: "Badge", clickable: new(widget.Clickable)},
		{ID: "avatar", Category: "General", Name: "Avatar", clickable: new(widget.Clickable)},
		{ID: "label", Category: "General", Name: "Label & Typography", clickable: new(widget.Clickable)},

		// Form Controls
		{ID: "input", Category: "Form Controls", Name: "Text Input", clickable: new(widget.Clickable)},
		{ID: "textarea", Category: "Form Controls", Name: "Text Area", clickable: new(widget.Clickable)},
		{ID: "numberinput", Category: "Form Controls", Name: "Number Input", clickable: new(widget.Clickable)},
		{ID: "inputotp", Category: "Form Controls", Name: "Input OTP", clickable: new(widget.Clickable)},
		{ID: "checkbox", Category: "Form Controls", Name: "Checkbox", clickable: new(widget.Clickable)},
		{ID: "switch", Category: "Form Controls", Name: "Switch Toggle", clickable: new(widget.Clickable)},
		{ID: "radio", Category: "Form Controls", Name: "Radio Group", clickable: new(widget.Clickable)},
		{ID: "select", Category: "Form Controls", Name: "Select Dropdown", clickable: new(widget.Clickable)},
		{ID: "togglegroup", Category: "Form Controls", Name: "Toggle Group", clickable: new(widget.Clickable)},
		{ID: "slider", Category: "Form Controls", Name: "Range Slider", clickable: new(widget.Clickable)},

		// Feedback & Status
		{ID: "progress", Category: "Feedback", Name: "Progress Bar", clickable: new(widget.Clickable)},
		{ID: "skeleton", Category: "Feedback", Name: "Skeleton Loader", clickable: new(widget.Clickable)},
		{ID: "spinner", Category: "Feedback", Name: "Spinner Activity", clickable: new(widget.Clickable)},
		{ID: "alert", Category: "Feedback", Name: "Alert Callout", clickable: new(widget.Clickable)},
		{ID: "toast", Category: "Feedback", Name: "Toast Banner", clickable: new(widget.Clickable)},

		// Containers & Navigation
		{ID: "accordion", Category: "Navigation", Name: "Accordion", clickable: new(widget.Clickable)},
		{ID: "collapsible", Category: "Navigation", Name: "Collapsible", clickable: new(widget.Clickable)},
		{ID: "tree", Category: "Navigation", Name: "Tree View (DnD)", clickable: new(widget.Clickable)},
		{ID: "tabs", Category: "Navigation", Name: "Tabs", clickable: new(widget.Clickable)},
		{ID: "breadcrumb", Category: "Navigation", Name: "Breadcrumbs", clickable: new(widget.Clickable)},
		{ID: "pagination", Category: "Navigation", Name: "Pagination", clickable: new(widget.Clickable)},
		{ID: "table", Category: "Data Grid", Name: "Data Grid Table", clickable: new(widget.Clickable)},

		// Overlays & Popups
		{ID: "dialog", Category: "Overlays", Name: "Modal Dialog", clickable: new(widget.Clickable)},
		{ID: "sheet", Category: "Overlays", Name: "Side Sheet", clickable: new(widget.Clickable)},
		{ID: "drawer", Category: "Overlays", Name: "Bottom Drawer", clickable: new(widget.Clickable)},
		{ID: "popover", Category: "Overlays", Name: "Popover", clickable: new(widget.Clickable)},
		{ID: "dropdownmenu", Category: "Overlays", Name: "Dropdown Menu", clickable: new(widget.Clickable)},
		{ID: "tooltip", Category: "Overlays", Name: "Tooltip", clickable: new(widget.Clickable)},
		{ID: "hovercard", Category: "Overlays", Name: "Hover Card", clickable: new(widget.Clickable)},

		// Layout & Utilities
		{ID: "separator", Category: "Utilities", Name: "Separator", clickable: new(widget.Clickable)},
		{ID: "resizable", Category: "Utilities", Name: "Resizable Panels", clickable: new(widget.Clickable)},
		{ID: "carousel", Category: "Utilities", Name: "Carousel Slider", clickable: new(widget.Clickable)},
		{ID: "command", Category: "Utilities", Name: "Command Palette", clickable: new(widget.Clickable)},
		{ID: "empty", Category: "Utilities", Name: "Empty State", clickable: new(widget.Clickable)},
	}

	sort.Slice(galleryItems, func(i, j int) bool {
		return galleryItems[i].Name < galleryItems[j].Name
	})

	activeItemID := galleryItems[0].ID
	if testCat := os.Getenv("GIO_TEST_CATEGORY"); testCat != "" {
		activeItemID = testCat
	}

	// Theme Toggle
	var themeToggleBtn *button.Button
	themeToggleBtn = button.New(button.Config{
		Text:    "☀️ Light Mode",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			th.ToggleDark()
			if th.IsDark {
				themeToggleBtn.SetText("☀️ Light Mode")
			} else {
				themeToggleBtn.SetText("🌙 Dark Mode")
			}
			w.Invalidate()
		},
	})

	demoCard := card.New(card.Config{Variant: theme.VariantDefault})

	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			for {
				ev, ok := gtx.Event(
					key.Filter{Name: "1", Required: key.ModCtrl},
				)
				if !ok {
					break
				}
				_ = ev
			}

			// Background clip stack - fill FULL physical Retina framebuffer (e.Size)
			bgClip := clip.Rect{Max: e.Size}.Push(gtx.Ops)
			paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			bgClip.Pop()

			// Top Bar + Main Split View Layout
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// Window Title Bar
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return tb.Layout(gtx, th, w)
				}),

				// Header Toolbar
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    th.Spacing.Space4,
						Bottom: th.Spacing.Space4,
						Left:   th.Spacing.Space6,
						Right:  th.Spacing.Space6,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								lblTitle := label.NewTypography("guipoc - Gio Component Gallery", label.H2, "")
								return lblTitle.Layout(gtx, th)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return themeToggleBtn.Layout(gtx, th)
							}),
						)
					})
				}),

				// Main Content Split Area (Left Navigation Sidebar + Right Component Viewer)
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						// Left Sidebar Navigation List
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return renderSidebar(gtx, th, galleryItems, &activeItemID)
						}),

						// Right Component Gallery Viewport (Vertically Scrollable Auto)
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left:   th.Spacing.Space6,
								Right:  th.Spacing.Space6,
								Bottom: th.Spacing.Space6,
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return demoCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
									mTheme := th.MaterialTheme
									if mTheme == nil {
										mTheme = material.NewTheme()
									}
									return material.List(mTheme, contentList).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
										return renderComponentGalleryPage(gtx, th, activeItemID)
									})
								})
							})
						}),
					)
				}),
			)

			// RESET GPU PAINT COLOR TO BACKGROUND AT THE END OF THE FRAME LOOP BEFORE SUBMISSION
			paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

			e.Frame(&ops)
		}
	}
}

func renderSidebar(gtx layout.Context, th *theme.Theme, items []*GalleryItem, activeID *string) layout.Dimensions {
	sidebarWidth := gtx.Dp(unit.Dp(240))
	gtx.Constraints.Min.X = sidebarWidth
	gtx.Constraints.Max.X = sidebarWidth

	padding := layout.Inset{
		Top:    th.Spacing.Space2,
		Bottom: th.Spacing.Space6,
		Left:   th.Spacing.Space6,
		Right:  th.Spacing.Space4,
	}

	return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.List(th.MaterialTheme, sidebarList).Layout(gtx, len(items)+2, func(gtx layout.Context, index int) layout.Dimensions {
			if index == 0 {
				lblCategory := label.NewTypography("COMPONENTS (37)", label.Small, "")
				return lblCategory.Layout(gtx, th)
			}
			if index == 1 {
				return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx)
			}

			item := items[index-2]
			isActive := item.ID == *activeID

			if item.clickable.Clicked(gtx) {
				if *activeID != item.ID {
					contentList.Position = layout.Position{}
				}
				*activeID = item.ID
			}

			return layout.Inset{Bottom: th.Spacing.Space1}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return item.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					padding := layout.Inset{
						Top:    th.Spacing.Space2,
						Bottom: th.Spacing.Space2,
						Left:   th.Spacing.Space3,
						Right:  th.Spacing.Space3,
					}

					itemBg := th.Colors.Background
					fgColor := th.Colors.Foreground

					if isActive {
						itemBg = th.Colors.Secondary
					} else if item.clickable.Hovered() {
						itemBg = th.Colors.Muted
					}

					return layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx layout.Context) layout.Dimensions {
							rect := image.Rectangle{Max: gtx.Constraints.Min}
							radiusPx := gtx.Dp(th.Radius.RadiusSM)
							theme.DrawRRectBackground(gtx, rect, radiusPx, itemBg)
							return layout.Dimensions{Size: gtx.Constraints.Min}
						}),
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := label.NewTypography(item.Name, label.Small, "")
								lbl.TextStyle.Color = &theme.ColorScheme{Foreground: fgColor}
								return lbl.Layout(gtx, th)
							})
						}),
					)
				})
			})
		})
	})
}

func renderComponentGalleryPage(gtx layout.Context, th *theme.Theme, activeID string) layout.Dimensions {
	switch activeID {
	case "accordion":
		return accordion.Demo(gtx, th)
	case "alert":
		return alert.Demo(gtx, th)
	case "aspectratio":
		return aspectratio.Demo(gtx, th)
	case "avatar":
		return avatar.Demo(gtx, th)
	case "badge":
		return badge.Demo(gtx, th)
	case "breadcrumb":
		return breadcrumb.Demo(gtx, th)
	case "button":
		return button.Demo(gtx, th)
	case "card":
		return card.Demo(gtx, th)
	case "carousel":
		return carousel.Demo(gtx, th)
	case "checkbox":
		return checkbox.Demo(gtx, th)
	case "collapsible":
		return collapsible.Demo(gtx, th)
	case "command":
		return command.Demo(gtx, th)
	case "dialog":
		return dialog.Demo(gtx, th)
	case "drawer":
		return drawer.Demo(gtx, th)
	case "dropdownmenu":
		return dropdownmenu.Demo(gtx, th)
	case "empty":
		return empty.Demo(gtx, th)
	case "hovercard":
		return hovercard.Demo(gtx, th)
	case "input":
		return input.Demo(gtx, th)
	case "inputotp":
		return inputotp.Demo(gtx, th)
	case "label":
		return label.Demo(gtx, th)
	case "numberinput":
		return numberinput.Demo(gtx, th)
	case "pagination":
		return pagination.Demo(gtx, th)
	case "popover":
		return popover.Demo(gtx, th)
	case "progress":
		return progress.Demo(gtx, th)
	case "radio":
		return radio.Demo(gtx, th)
	case "resizable":
		return resizable.Demo(gtx, th)
	case "scrollarea":
		return scrollarea.Demo(gtx, th)
	case "select":
		return selectcomp.Demo(gtx, th)
	case "separator":
		return separator.Demo(gtx, th)
	case "sheet":
		return sheet.Demo(gtx, th)
	case "skeleton":
		return skeleton.Demo(gtx, th)
	case "slider":
		return slider.Demo(gtx, th)
	case "spinner":
		return spinner.Demo(gtx, th)
	case "switch":
		return switchcomp.Demo(gtx, th)
	case "table":
		return table.Demo(gtx, th)
	case "tabs":
		return tabs.Demo(gtx, th)
	case "textarea":
		return textarea.Demo(gtx, th)
	case "titlebar":
		return titlebar.Demo(gtx, th, nil)
	case "toast":
		return toast.Demo(gtx, th)
	case "togglegroup":
		return togglegroup.Demo(gtx, th)
	case "tooltip":
		return tooltip.Demo(gtx, th)
	case "tree":
		return tree.Demo(gtx, th)
	default:
		return layout.Dimensions{}
	}
}
