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

// Run launches the full gio-shadcn component gallery window.
func Run() {
	go func() {
		w := &app.Window{}
		w.Option(app.Title("guipoc - Gio (gio-shadcn Gallery)"))
		w.Option(app.Size(1200, 800))
		w.Option(app.Maximized.Option())
		w.Option(app.Decorated(false))

		err := runWindow(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func layoutButtonRow(gtx layout.Context, th *theme.Theme, buttons ...*button.Button) layout.Dimensions {
	children := make([]layout.FlexChild, 0, len(buttons)*2)
	for i, btn := range buttons {
		btn := btn
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return btn.Layout(gtx, th)
		}))
		if i < len(buttons)-1 {
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
			}))
		}
	}
	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)
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

	// Initialize gallery components
	btnPrimary := button.New(button.Config{Text: "Primary", Variant: theme.VariantDefault})
	btnSecondary := button.New(button.Config{Text: "Secondary", Variant: theme.VariantSecondary})
	btnOutline := button.New(button.Config{Text: "Outline", Variant: theme.VariantOutline})
	btnGhost := button.New(button.Config{Text: "Ghost", Variant: theme.VariantGhost})
	btnDestructive := button.New(button.Config{Text: "Destructive", Variant: theme.VariantDestructive})
	btnLink := button.New(button.Config{Text: "Link Button", Variant: theme.VariantLink})

	btnSM := button.New(button.Config{Text: "Small", Size: theme.SizeSM})
	btnDefaultSize := button.New(button.Config{Text: "Default Size", Size: theme.SizeDefault})
	btnLG := button.New(button.Config{Text: "Large Size", Size: theme.SizeLG})

	badgeDef := badge.New(badge.Config{Text: "Default Badge", Variant: theme.VariantDefault})
	badgeSec := badge.New(badge.Config{Text: "Secondary Badge", Variant: theme.VariantSecondary})
	badgeOut := badge.New(badge.Config{Text: "Outline Badge", Variant: theme.VariantOutline})
	badgeDest := badge.New(badge.Config{Text: "Destructive Badge", Variant: theme.VariantDestructive})

	avDJ := avatar.New(avatar.Config{Initials: "DJ", ShowBadge: true})
	avAG := avatar.New(avatar.Config{Initials: "AG", ShowBadge: false})

	txtInput := input.Text("Enter track title...")
	txtArea := textarea.New(textarea.Config{Placeholder: "Enter track descriptions and metadata notes..."})

	numStepper := numberinput.New(numberinput.Config{Value: 128.0, Step: 1.0, Min: 60.0, Max: 200.0})
	otpPin := inputotp.New(inputotp.Config{Length: 6})

	chkBox := checkbox.New(checkbox.Config{Value: true})
	swToggle := switchcomp.New(switchcomp.Config{Value: true})
	radioA := radio.New(radio.Config{Selected: true})
	radioB := radio.New(radio.Config{Selected: false})

	genreSel := selectcomp.New(selectcomp.Config{
		Options: []*selectcomp.Item{
			selectcomp.NewItem("house", "Progressive House"),
			selectcomp.NewItem("techno", "Techno"),
			selectcomp.NewItem("trance", "Trance"),
		},
		SelectedValue: "house",
	})

	tglGrp := togglegroup.New(togglegroup.Config{
		Items: []*togglegroup.Item{
			togglegroup.NewItem("mono", "Mono"),
			togglegroup.NewItem("stereo", "Stereo"),
			togglegroup.NewItem("5.1", "5.1 Surround"),
		},
		SelectedKey: "stereo",
	})

	rngSlider := slider.New(slider.Config{Value: 65.0, Min: 0.0, Max: 100.0})
	progBar := progress.New(progress.Config{Value: 0.65})

	shimmerSk := skeleton.New(skeleton.Config{Width: unit.Dp(180), Height: unit.Dp(24)})
	spinLoad := spinner.New(spinner.Config{})

	alertDefault := alert.New(alert.Config{Title: "Engine Status", Description: "CoreAudio buffer set to 64 samples."})
	alertDestruct := alert.New(alert.Config{Title: "Audio Clip Warning", Description: "Output signal clipped +1.2dB on Deck A.", Variant: theme.VariantDestructive})
	toastItem := toast.New(toast.Config{Title: "Track Exported", Description: "Exported to Starlight_Symphony.flac", Visible: false})

	// Accordion Demos (All 8 Official shadcn Demos)
	singleAcc := accordion.New(accordion.Config{
		Type: accordion.TypeSingle,
		Items: []*accordion.Item{
			accordion.NewItem("Is it accessible?", "Yes. It adheres to the WAI-ARIA design pattern.", true),
			accordion.NewItem("Is it styled?", "Yes. It comes with default styles.", false),
		},
	})
	multiAcc := accordion.New(accordion.Config{
		Type: accordion.TypeMultiple,
		Items: []*accordion.Item{
			accordion.NewItem("Audio Engine Specs", "Runs at 96kHz 24-bit precision.", true),
			accordion.NewItem("GPU Vector Pipeline", "Gio immediate-mode engine.", true),
		},
	})
	disabledAcc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItemConfig(accordion.ItemConfig{Title: "Disabled Section (Locked)", Content: "Locked content panel.", Disabled: true}),
		},
	})
	chevronAcc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItemConfig(accordion.ItemConfig{Title: "Chevron Indicator Item", Content: "Uses custom chevron symbol.", Icon: "v", Expanded: true}),
		},
	})
	customHeaderAcc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItemConfig(accordion.ItemConfig{Title: "Custom Header Badge Section", Content: "Custom header badge.", Expanded: true}),
		},
	})
	borderlessAcc := accordion.New(accordion.Config{
		Borderless: true,
		Items: []*accordion.Item{
			accordion.NewItem("Borderless Item 1", "Flush borderless container style.", true),
		},
	})
	nestedAcc := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItemConfig(accordion.ItemConfig{
				Title:    "Parent Section (Contains Inner Accordion)",
				Expanded: true,
				ContentWidget: func(gtx layout.Context) layout.Dimensions {
					return singleAcc.Layout(gtx, th)
				},
			}),
		},
	})
	controlledAcc := accordion.New(accordion.Config{
		Type: accordion.TypeMultiple,
		Items: []*accordion.Item{
			accordion.NewItem("Controlled Panel 1", "State toggled externally by buttons.", true),
		},
	})
	btnExpandAll := button.New(button.Config{
		Text:    "Expand All",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			for _, item := range controlledAcc.Items {
				item.Expanded = true
			}
			w.Invalidate()
		},
	})
	colContainer := collapsible.New(collapsible.Config{Title: "Advanced Mixer Settings", Content: "ASIO Direct hardware routing enabled.", Open: true})

	navTabs := tabs.New(tabs.Config{
		Tabs: []*tabs.Tab{
			tabs.NewTab("sink", "Kitchen Sink"),
			tabs.NewTab("deck", "Audio Deck"),
			tabs.NewTab("library", "Track Library"),
		},
		ActiveKey: "sink",
	})
	bCrumb := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Mixer", false),
			breadcrumb.NewItem("Deck A", true),
		},
	})
	pageControls := pagination.New(pagination.Config{CurrentPage: 1, TotalPages: 5})

	trackTable := table.New(table.Config{
		Headers: []string{"TITLE", "ARTIST", "BPM", "KEY", "GENRE"},
		Rows: []*table.Row{
			table.NewRow("Starlight Symphony", "Aethelgard", "128", "8A", "Progressive House"),
			table.NewRow("Quantum Drift", "Cyberpulse", "132", "11B", "Techno"),
			table.NewRow("Solar Flare", "Helios", "126", "4A", "Melodic Techno"),
		},
	})

	modalDialog := dialog.New(dialog.Config{
		Title:       "Reset Audio Mixer",
		Description: "Are you sure you want to reset all channel EQ gain levels?",
		Open:        false,
	})
	btnTriggerDialog := button.New(button.Config{
		Text:    "Open Modal Dialog",
		Variant: theme.VariantOutline,
		OnClick: func() {
			modalDialog.Open = true
			w.Invalidate()
		},
	})

	sheetDrawer := sheet.New(sheet.Config{
		Title:       "Track Inspector",
		Description: "Detailed FLAC metadata and harmonic key analysis.",
		Open:        false,
	})
	btnTriggerSheet := button.New(button.Config{
		Text:    "Open Side Sheet",
		Variant: theme.VariantOutline,
		OnClick: func() {
			sheetDrawer.Open = true
			w.Invalidate()
		},
	})

	bottomDrawer := drawer.New(drawer.Config{
		Title:       "System Telemetry",
		Description: "CPU: 2.1% | RAM: 189MB | Metal GPU 120 FPS",
		Open:        false,
	})
	btnTriggerDrawer := button.New(button.Config{
		Text:    "Open Bottom Drawer",
		Variant: theme.VariantOutline,
		OnClick: func() {
			bottomDrawer.Open = true
			w.Invalidate()
		},
	})

	anchoredPop := popover.New(popover.Config{Title: "Popover Title", Description: "Anchored card popover content box.", Open: false})
	dropdownMenu := dropdownmenu.New(dropdownmenu.Config{
		Items: []*dropdownmenu.Item{
			dropdownmenu.NewItem("Edit Track", "⌘E"),
			dropdownmenu.NewItem("Export FLAC", "⌘S"),
		},
		Open: false,
	})
	tipCallout := tooltip.New(tooltip.Config{Text: "ASIO Low Latency Buffer"})
	hoverCardItem := hovercard.New(hovercard.Config{Title: "Artist Profile", Description: "Aethelgard - Progressive House", Hovered: false})

	sepDivider := separator.New(separator.Config{Horizontal: true})
	cmdPalette := command.New(command.Config{
		Placeholder: "Search command palette...",
		Items: []*command.Item{
			command.NewItem("Toggle Light/Dark Theme", "⌘T"),
			command.NewItem("Reset Master Audio Mixer", "⌘R"),
		},
	})
	emptyBox := empty.New(empty.Config{})

	aspectWrapper := aspectratio.New(aspectratio.Config{
		Ratio: 16.0 / 9.0,
		Widget: func(gtx layout.Context) layout.Dimensions {
			return trackTable.Layout(gtx, th)
		},
	})

	resPanel := resizable.New(resizable.Config{
		Ratio: 0.5,
		LeftWidget: func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Left Split Panel", label.P, "").Layout(gtx, th)
		},
		RightWidget: func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Right Split Panel", label.P, "").Layout(gtx, th)
		},
	})

	slideCarousel := carousel.New(carousel.Config{
		Items: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Slide 1: Audio Spectrum", label.H4, "").Layout(gtx, th) },
			func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Slide 2: Mixer Controls", label.H4, "").Layout(gtx, th) },
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

						// Right Component Gallery Viewport
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left:   th.Spacing.Space6,
								Right:  th.Spacing.Space6,
								Bottom: th.Spacing.Space6,
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return demoCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
									return renderComponentGalleryPage(
										gtx, th, activeItemID,
										btnPrimary, btnSecondary, btnOutline, btnGhost, btnDestructive, btnLink, btnSM, btnDefaultSize, btnLG,
										badgeDef, badgeSec, badgeOut, badgeDest, avDJ, avAG,
										txtInput, txtArea, numStepper, otpPin,
										chkBox, swToggle, radioA, radioB, genreSel, tglGrp,
										rngSlider, progBar, shimmerSk, spinLoad,
										alertDefault, alertDestruct, toastItem,
										singleAcc, multiAcc, disabledAcc, chevronAcc, customHeaderAcc, borderlessAcc, nestedAcc, controlledAcc, btnExpandAll, colContainer,
										navTabs, bCrumb, pageControls,
										trackTable, aspectWrapper, resPanel, slideCarousel,
										modalDialog, btnTriggerDialog, sheetDrawer, btnTriggerSheet, bottomDrawer, btnTriggerDrawer,
										anchoredPop, dropdownMenu, tipCallout, hoverCardItem,
										sepDivider, cmdPalette, emptyBox,
									)
								})
							})
						}),
					)
				}),
			)

			if modalDialog.Open {
				modalDialog.Layout(gtx, th)
			}
			if sheetDrawer.Open {
				sheetDrawer.Layout(gtx, th)
			}
			if bottomDrawer.Open {
				bottomDrawer.Layout(gtx, th)
			}

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

func renderComponentGalleryPage(
	gtx layout.Context, th *theme.Theme, activeID string,
	btnPrimary, btnSecondary, btnOutline, btnGhost, btnDestructive, btnLink, btnSM, btnDefaultSize, btnLG *button.Button,
	badgeDef, badgeSec, badgeOut, badgeDest *badge.Badge, avDJ, avAG *avatar.Avatar,
	txtInput *input.Input, txtArea *textarea.TextArea, numStepper *numberinput.NumberInput, otpPin *inputotp.InputOTP,
	chkBox *checkbox.Checkbox, swToggle *switchcomp.Switch, radioA, radioB *radio.Radio, genreSel *selectcomp.Select, tglGrp *togglegroup.ToggleGroup,
	rngSlider *slider.Slider, progBar *progress.Progress, shimmerSk *skeleton.Skeleton, spinLoad *spinner.Spinner,
	alertDefault, alertDestruct *alert.Alert, toastItem *toast.Toast,
	singleAcc, multiAcc, disabledAcc, chevronAcc, customHeaderAcc, borderlessAcc, nestedAcc, controlledAcc *accordion.Accordion, btnExpandAll *button.Button, colContainer *collapsible.Collapsible,
	navTabs *tabs.Tabs, bCrumb *breadcrumb.Breadcrumb, pageControls *pagination.Pagination,
	trackTable *table.Table, aspectWrapper *aspectratio.AspectRatio, resPanel *resizable.Resizable, slideCarousel *carousel.Carousel,
	modalDialog *dialog.Dialog, btnTriggerDialog *button.Button, sheetDrawer *sheet.Sheet, btnTriggerSheet *button.Button, bottomDrawer *drawer.Drawer, btnTriggerDrawer *button.Button,
	anchoredPop *popover.Popover, dropdownMenu *dropdownmenu.DropdownMenu, tipCallout *tooltip.Tooltip, hoverCardItem *hovercard.HoverCard,
	sepDivider *separator.Separator, cmdPalette *command.Command, emptyBox *empty.Empty,
) layout.Dimensions {

	switch activeID {
	// 1. Button
	case "button":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Button Variants", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layoutButtonRow(gtx, th, btnPrimary, btnSecondary, btnOutline, btnGhost, btnDestructive, btnLink) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Button Sizes", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layoutButtonRow(gtx, th, btnSM, btnDefaultSize, btnLG) }),
		)

	// 2. Badge
	case "badge":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Badge Variants", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return badgeDef.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return badgeSec.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return badgeOut.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return badgeDest.Layout(gtx, th) }),
				)
			}),
		)

	// 3. Avatar
	case "avatar":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("User Avatars & Status Badges", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return avDJ.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return avAG.Layout(gtx, th) }),
				)
			}),
		)

	// 4. Label & Typography
	case "label":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Typography Heading 1", label.H1, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Typography Heading 2", label.H2, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Typography Heading 3", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Typography Heading 4", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Body paragraph demonstrating standard typography scaling.", label.P, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Muted secondary text style.", label.Muted, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Small caption text for fine print.", label.Small, "").Layout(gtx, th) }),
		)

	// 5. Input
	case "input":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Text Input", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				maxWidth := gtx.Metric.Dp(400)
				gtx.Constraints.Max.X = maxWidth
				gtx.Constraints.Min.X = maxWidth
				return txtInput.Layout(gtx, th)
			}),
		)

	// 6. Textarea
	case "textarea":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Multi-line Text Area", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return txtArea.Layout(gtx, th) }),
		)

	// 7. Numberinput
	case "numberinput":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Number Input Stepper", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return numStepper.Layout(gtx, th) }),
		)

	// 8. Inputotp
	case "inputotp":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Input OTP PIN Code", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return otpPin.Layout(gtx, th) }),
		)

	// 9. Checkbox
	case "checkbox":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Checkbox Controls", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return chkBox.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Enable GPU Vector Acceleration", label.P, "").Layout(gtx, th) }),
				)
			}),
		)

	// 10. Switch
	case "switch":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Switch Toggle", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return swToggle.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space3}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("HQ Audio Engine", label.P, "").Layout(gtx, th) }),
				)
			}),
		)

	// 11. Radio
	case "radio":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Radio Options", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return radioA.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Master Output", label.P, "").Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return radioB.Layout(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Headphones Cue", label.P, "").Layout(gtx, th) }),
				)
			}),
		)

	// 12. Select
	case "select":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Select Dropdown Menu", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return genreSel.Layout(gtx, th) }),
		)

	// 13. Togglegroup
	case "togglegroup":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Toggle Group Buttons", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return tglGrp.Layout(gtx, th) }),
		)

	// 14. Slider
	case "slider":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Range Slider Fader", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return rngSlider.Layout(gtx, th) }),
		)

	// 15. Progress
	case "progress":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Progress Bar Indicator", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return progBar.Layout(gtx, th) }),
		)

	// 16. Skeleton
	case "skeleton":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Skeleton Loader Bar", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return shimmerSk.Layout(gtx, th) }),
		)

	// 17. Spinner
	case "spinner":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Activity Loading Spinner", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return spinLoad.Layout(gtx, th) }),
		)

	// 18. Alert
	case "alert":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Alert Callout Banners", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return alertDefault.Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return alertDestruct.Layout(gtx, th) }),
		)

	// 19. Toast
	case "toast":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Toast Notification Banner", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return toastItem.Layout(gtx, th) }),
		)

	// 20. Accordion
	case "accordion":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Accordion Showcase (All 8 Official Demos)", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("1. Single Open Accordion (Default)", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return singleAcc.Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("2. Multiple Open Accordion", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return multiAcc.Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("3. Disabled Item Accordion", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return disabledAcc.Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("4. Chevron Icon Accordion", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return chevronAcc.Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("5. Custom Header Badge Section", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return customHeaderAcc.Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("6. Borderless Variant Accordion", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return borderlessAcc.Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("7. Nested Accordion", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return nestedAcc.Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("8. Controlled Accordion State", label.H4, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnExpandAll.Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return controlledAcc.Layout(gtx, th) }),
		)

	// 21. Collapsible
	case "collapsible":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Collapsible Container", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return colContainer.Layout(gtx, th) }),
		)

	// 22. Tabs
	case "tabs":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Tabs Navigation Header", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return navTabs.Layout(gtx, th) }),
		)

	// 23. Breadcrumb
	case "breadcrumb":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Breadcrumb Navigation Path", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return bCrumb.Layout(gtx, th) }),
		)

	// 24. Pagination
	case "pagination":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Pagination Page Controls", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return pageControls.Layout(gtx, th) }),
		)

	// 25. Table
	case "table":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Master Data Grid Table", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return trackTable.Layout(gtx, th) }),
		)

	// 26. Dialog
	case "dialog":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Modal Dialog Window", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnTriggerDialog.Layout(gtx, th) }),
		)

	// 27. Sheet
	case "sheet":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Side Sheet Drawer", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnTriggerSheet.Layout(gtx, th) }),
		)

	// 28. Drawer
	case "drawer":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Bottom Drawer Panel", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return btnTriggerDrawer.Layout(gtx, th) }),
		)

	// 29. Popover
	case "popover":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Anchored Popover Card", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return anchoredPop.Layout(gtx, th) }),
		)

	// 30. Dropdownmenu
	case "dropdownmenu":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Dropdown Action Menu", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return dropdownMenu.Layout(gtx, th) }),
		)

	// 31. Tooltip
	case "tooltip":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Tooltip Callout Text", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return tipCallout.Layout(gtx, th) }),
		)

	// 32. Hovercard
	case "hovercard":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Hover Card User Profile", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return hoverCardItem.Layout(gtx, th) }),
		)

	// 33. Separator
	case "separator":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Separator Divider Lines", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sepDivider.Layout(gtx, th) }),
		)

	// 34. Resizable
	case "resizable":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Resizable Split Panels", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return resPanel.Layout(gtx, th) }),
		)

	// 35. Carousel
	case "carousel":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Carousel Slide Viewer", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return slideCarousel.Layout(gtx, th) }),
		)

	// 36. Command
	case "command":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Command Palette", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return cmdPalette.Layout(gtx, th) }),
		)

	// 37. Empty
	case "empty":
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Empty State Box", label.H3, "").Layout(gtx, th) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions { return emptyBox.Layout(gtx, th) }),
		)

	default:
		return layout.Dimensions{}
	}
}
