package main

import (
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/bnema/gio-shadcn/components/accordion"
	"github.com/bnema/gio-shadcn/components/alert"
	"github.com/bnema/gio-shadcn/components/aspectratio"
	"github.com/bnema/gio-shadcn/components/avatar"
	"github.com/bnema/gio-shadcn/components/badge"
	"github.com/bnema/gio-shadcn/components/breadcrumb"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/components/checkbox"
	"github.com/bnema/gio-shadcn/components/command"
	"github.com/bnema/gio-shadcn/components/dialog"
	"github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/components/inputotp"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/components/numberinput"
	"github.com/bnema/gio-shadcn/components/pagination"
	"github.com/bnema/gio-shadcn/components/progress"
	"github.com/bnema/gio-shadcn/components/radio"
	selectcomp "github.com/bnema/gio-shadcn/components/select"
	"github.com/bnema/gio-shadcn/components/separator"
	"github.com/bnema/gio-shadcn/components/skeleton"
	"github.com/bnema/gio-shadcn/components/slider"
	"github.com/bnema/gio-shadcn/components/spinner"
	switchcomp "github.com/bnema/gio-shadcn/components/switch"
	"github.com/bnema/gio-shadcn/components/table"
	"github.com/bnema/gio-shadcn/components/tabs"
	"github.com/bnema/gio-shadcn/components/textarea"
	"github.com/bnema/gio-shadcn/components/titlebar"
	"github.com/bnema/gio-shadcn/components/togglegroup"
	"github.com/bnema/gio-shadcn/components/tooltip"
	"github.com/bnema/gio-shadcn/theme"
)

func main() {
	go func() {
		w := &app.Window{}
		w.Option(app.Title("Gio-shadcn Full Parity Showcase"))
		w.Option(app.Size(1000, 800))
		w.Option(app.Decorated(false))

		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func updateWindowColors(w *app.Window, th *theme.Theme) {
	w.Option(app.NavigationColor(th.Colors.Background))
	w.Option(app.StatusColor(th.Colors.Background))
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
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
}

func run(w *app.Window) error {
	th := theme.NewDark()

	updateWindowColors(w, th)

	tb := titlebar.NewTitleBar(
		titlebar.WithTitle("Gio-shadcn Full Parity Showcase (Flutter Source of Truth)"),
		titlebar.WithWindow(w),
		titlebar.WithVariant(theme.VariantSecondary),
	)

	var zoomScale float32 = 1.0
	const minZoom = 0.5
	const maxZoom = 3.0
	const zoomStep = 0.1

	zoomLabel := label.NewTypography("Zoom: 100%", label.Small, "")

	// Buttons
	primaryBtn := button.New(button.Config{Text: "Primary", Variant: theme.VariantDefault, OnClick: func() { log.Println("Primary clicked!") }})
	destructiveBtn := button.New(button.Config{Text: "Destructive", Variant: theme.VariantDestructive, OnClick: func() { log.Println("Destructive clicked!") }})
	outlineBtn := button.New(button.Config{Text: "Outline", Variant: theme.VariantOutline, OnClick: func() { log.Println("Outline clicked!") }})
	secondaryBtn := button.New(button.Config{Text: "Secondary", Variant: theme.VariantSecondary, OnClick: func() { log.Println("Secondary clicked!") }})
	ghostBtn := button.New(button.Config{Text: "Ghost", Variant: theme.VariantGhost, OnClick: func() { log.Println("Ghost clicked!") }})
	linkBtn := button.New(button.Config{Text: "Link", Variant: theme.VariantLink, OnClick: func() { log.Println("Link clicked!") }})

	// Theme toggle & zoom controls
	var themeToggleBtn *button.Button
	themeToggleBtn = button.New(button.Config{
		Text:    "☀️ Light Mode",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			th.ToggleDark()
			updateWindowColors(w, th)
			if th.IsDark {
				themeToggleBtn.SetText("☀️ Light Mode")
			} else {
				themeToggleBtn.SetText("🌙 Dark Mode")
			}
			w.Invalidate()
		},
	})

	zoomInBtn := button.New(button.Config{Text: "+", Variant: theme.VariantOutline, Size: theme.SizeSM, OnClick: func() {
		if zoomScale < maxZoom {
			zoomScale += zoomStep
			zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
			w.Invalidate()
		}
	}})

	zoomOutBtn := button.New(button.Config{Text: "-", Variant: theme.VariantOutline, Size: theme.SizeSM, OnClick: func() {
		if zoomScale > minZoom {
			zoomScale -= zoomStep
			zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
			w.Invalidate()
		}
	}})

	zoomResetBtn := button.New(button.Config{Text: "Reset", Variant: theme.VariantOutline, Size: theme.SizeSM, OnClick: func() {
		zoomScale = 1.0
		zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
		w.Invalidate()
	}})

	// Breadcrumb & Tabs
	bCrumb := breadcrumb.New(breadcrumb.Config{
		Items: []*breadcrumb.Item{
			breadcrumb.NewItem("Home", false),
			breadcrumb.NewItem("Audio Engine", false),
			breadcrumb.NewItem("Mixer Controls", true),
		},
	})

	navTabs := tabs.New(tabs.Config{
		Tabs: []*tabs.Tab{
			tabs.NewTab("sink", "Kitchen Sink"),
			tabs.NewTab("deck", "Audio Deck"),
			tabs.NewTab("library", "Track Library"),
		},
		ActiveKey: "sink",
	})

	// Containers & Inputs
	demoCard := card.New(card.Config{Variant: theme.VariantDefault})
	textInput := input.Text("Enter track name...")
	textAreaInput := textarea.New(textarea.Config{Placeholder: "Enter track notes..."})
	bpmStepper := numberinput.New(numberinput.Config{Value: 128.0, Step: 1.0, Min: 60.0, Max: 200.0})
	pinOTP := inputotp.New(inputotp.Config{Length: 6})

	// Select Dropdown
	genreSelect := selectcomp.New(selectcomp.Config{
		Options: []*selectcomp.Item{
			selectcomp.NewItem("house", "Progressive House"),
			selectcomp.NewItem("techno", "Techno"),
			selectcomp.NewItem("trance", "Trance"),
		},
		SelectedValue: "house",
	})

	// Badges & Avatars
	defaultBadge := badge.New(badge.Config{Text: "Default", Variant: theme.VariantDefault})
	secBadge := badge.New(badge.Config{Text: "Secondary", Variant: theme.VariantSecondary})
	outBadge := badge.New(badge.Config{Text: "Outline", Variant: theme.VariantOutline})
	destBadge := badge.New(badge.Config{Text: "Destructive", Variant: theme.VariantDestructive})

	av1 := avatar.New(avatar.Config{Initials: "DJ", ShowBadge: true})
	av2 := avatar.New(avatar.Config{Initials: "AG", ShowBadge: false})

	// Interactive Form Controls
	progBar := progress.New(progress.Config{Value: 0.65})
	sepLine := separator.New(separator.Config{Horizontal: true})
	audioSwitch := switchcomp.New(switchcomp.Config{Value: true})
	gpuCheckbox := checkbox.New(checkbox.Config{Value: true})
	radioMaster := radio.New(radio.Config{Selected: true})

	gainSlider := slider.New(slider.Config{Value: 65.0, Min: 0.0, Max: 100.0})

	// Loaders
	shimmerSk := skeleton.New(skeleton.Config{Width: unit.Dp(120), Height: unit.Dp(20)})
	spinLoad := spinner.New(spinner.Config{})

	// Accordions & Tooltip Callout
	infoAccordion := accordion.New(accordion.Config{
		Items: []*accordion.Item{
			accordion.NewItem("Audio Engine Specifications", "Runs at 96kHz 24-bit floating point precision with low-jitter clocking.", false),
			accordion.NewItem("GPU Vector Pipeline", "Gio immediate-mode engine renders vector paths directly on GPU with Metal shaders.", true),
		},
	})

	// Toggle Group
	modeToggles := togglegroup.New(togglegroup.Config{
		Items: []*togglegroup.Item{
			togglegroup.NewItem("mono", "Mono"),
			togglegroup.NewItem("stereo", "Stereo"),
			togglegroup.NewItem("surround", "5.1 Surround"),
		},
		SelectedKey: "stereo",
	})

	// Data Grid Table
	trackTable := table.New(table.Config{
		Headers: []string{"TITLE", "ARTIST", "BPM", "KEY", "GENRE"},
		Rows: []*table.Row{
			table.NewRow("Starlight Symphony", "Aethelgard", "128", "8A", "Progressive House"),
			table.NewRow("Quantum Drift", "Cyberpulse", "132", "11B", "Techno"),
			table.NewRow("Solar Flare", "Helios", "126", "4A", "Melodic Techno"),
		},
	})

	// Pagination
	pageControls := pagination.New(pagination.Config{CurrentPage: 1, TotalPages: 5})

	// Modal Dialog (Triggered by button)
	modalDialog := dialog.New(dialog.Config{
		Title:       "Confirm Engine Reset",
		Description: "Are you sure you want to reset all audio buffers and channel gain levels?",
		Open:        false,
	})

	resetDialogBtn := button.New(button.Config{
		Text:    "Open Reset Dialog",
		Variant: theme.VariantOutline,
		OnClick: func() {
			modalDialog.Open = true
			w.Invalidate()
		},
	})

	// Command Palette
	cmdPalette := command.New(command.Config{
		Placeholder: "Type command or filter (e.g. 'Reset')...",
		Items: []*command.Item{
			command.NewItem("Toggle Light/Dark Theme", "⌘T"),
			command.NewItem("Reset Master Audio Mixer", "⌘R"),
			command.NewItem("Export Audio Track to FLAC", "⌘E"),
		},
	})

	tipCallout := tooltip.New(tooltip.Config{Text: "ASIO Low Latency Buffer: 64 samples"})
	systemAlert := alert.New(alert.Config{Title: "ASIO Driver Active", Description: "Buffer size set to 64 samples (1.2ms latency)."})

	aspectWrapper := aspectratio.New(aspectratio.Config{
		Ratio: 16.0 / 9.0,
		Widget: func(gtx layout.Context) layout.Dimensions {
			return trackTable.Layout(gtx, th)
		},
	})

	titleLabel := label.NewTypography("Gio-shadcn Showcase", label.H1, "")
	subtitleLabel := label.NewTypography("100% Component Parity Ported from Flutter Source of Truth", label.P, "")

	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			for {
				ev, ok := gtx.Event(
					key.Filter{Name: "+", Required: key.ModCtrl},
					key.Filter{Name: "=", Required: key.ModCtrl},
					key.Filter{Name: "-", Required: key.ModCtrl},
					key.Filter{Name: "0", Required: key.ModCtrl},
				)
				if !ok {
					break
				}
				if e, ok := ev.(key.Event); ok && e.State == key.Press {
					switch {
					case (e.Name == "+" || e.Name == "=") && e.Modifiers == key.ModCtrl:
						if zoomScale < maxZoom {
							zoomScale += zoomStep
							zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
							w.Invalidate()
						}
					case e.Name == "-" && e.Modifiers == key.ModCtrl:
						if zoomScale > minZoom {
							zoomScale -= zoomStep
							zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
							w.Invalidate()
						}
					case e.Name == "0" && e.Modifiers == key.ModCtrl:
						zoomScale = 1.0
						zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
						w.Invalidate()
					}
				}
			}

			if zoomScale != 1.0 {
				gtx.Metric = unit.Metric{
					PxPerDp: e.Metric.PxPerDp * zoomScale,
					PxPerSp: e.Metric.PxPerSp * zoomScale,
				}
			}

			background := clip.Rect{Max: gtx.Constraints.Max}.Op()
			paint.FillShape(gtx.Ops, th.Colors.Background, background)

			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return tb.Layout(gtx, th, w) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    th.Spacing.Space4,
						Bottom: th.Spacing.Space4,
						Left:   th.Spacing.Space6,
						Right:  th.Spacing.Space6,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return titleLabel.Layout(gtx, th) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return subtitleLabel.Layout(gtx, th) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return bCrumb.Layout(gtx, th) }),
								)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return themeToggleBtn.Layout(gtx, th) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return zoomOutBtn.Layout(gtx, th) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return zoomInBtn.Layout(gtx, th) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return zoomResetBtn.Layout(gtx, th) }),
								)
							}),
						)
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Left: th.Spacing.Space6, Right: th.Spacing.Space6, Bottom: th.Spacing.Space3}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return navTabs.Layout(gtx, th)
					})
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    th.Spacing.Space2,
						Bottom: th.Spacing.Space6,
						Left:   th.Spacing.Space6,
						Right:  th.Spacing.Space6,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return demoCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								// Section 1: Buttons, Badges, Avatars & Loaders
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									sectionTitle := label.NewTypography("Buttons, Badges, Avatars & Loaders", label.H4, "")
									return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sectionTitle.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layoutButtonRow(gtx, th, primaryBtn, destructiveBtn, outlineBtn, secondaryBtn, ghostBtn, linkBtn, resetDialogBtn)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return defaultBadge.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return secBadge.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return outBadge.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return destBadge.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return av1.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return av2.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return shimmerSk.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return spinLoad.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return tipCallout.Layout(gtx, th) }),
											)
										}),
									)
								}),

								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sepLine.Layout(gtx, th) }),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

								// Section 2: Form Inputs, Sliders, OTP, Select & Steppers
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									sectionTitle := label.NewTypography("Form Controls, Sliders, Select & Steppers", label.H4, "")
									return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sectionTitle.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													maxWidth := gtx.Metric.Dp(250)
													gtx.Constraints.Max.X = maxWidth
													gtx.Constraints.Min.X = maxWidth
													return textInput.Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return genreSelect.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return bpmStepper.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return pinOTP.Layout(gtx, th) }),
											)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return textAreaInput.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return audioSwitch.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return label.NewTypography("HQ Audio", label.Small, "").Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return gpuCheckbox.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return label.NewTypography("GPU Accel", label.Small, "").Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return radioMaster.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return label.NewTypography("Master Out", label.Small, "").Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return modeToggles.Layout(gtx, th) }),
											)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return label.NewTypography("Gain Slider (65%)", label.Small, "").Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return gainSlider.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return progBar.Layout(gtx, th) }),
											)
										}),
									)
								}),

								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sepLine.Layout(gtx, th) }),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

								// Section 3: Data Grid Table, Command Palette & Accordion
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									sectionTitle := label.NewTypography("Data Grid Table, Command Palette & Accordion", label.H4, "")
									return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sectionTitle.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return aspectWrapper.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return pageControls.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return cmdPalette.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return infoAccordion.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return systemAlert.Layout(gtx, th) }),
									)
								}),
							)
						})
					})
				}),

				// Modal Dialog Overlay
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return modalDialog.Layout(gtx, th)
				}),
			)

			e.Frame(&ops)
		}
	}
}
