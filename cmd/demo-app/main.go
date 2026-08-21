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

	"github.com/bnema/gio-shadcn/components/alert"
	"github.com/bnema/gio-shadcn/components/badge"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/components/checkbox"
	"github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/components/progress"
	"github.com/bnema/gio-shadcn/components/separator"
	switchcomp "github.com/bnema/gio-shadcn/components/switch"
	"github.com/bnema/gio-shadcn/components/titlebar"
	"github.com/bnema/gio-shadcn/theme"
)

func main() {
	go func() {
		w := &app.Window{}
		w.Option(app.Title("Gio-shadcn Demo"))
		w.Option(app.Size(900, 700))
		w.Option(app.Decorated(false)) // Disable system title bar

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
		titlebar.WithTitle("Gio-shadcn Demo (Flutter Parity Source of Truth)"),
		titlebar.WithWindow(w),
		titlebar.WithVariant(theme.VariantSecondary),
	)

	var zoomScale float32 = 1.0
	const minZoom = 0.5
	const maxZoom = 3.0
	const zoomStep = 0.1

	zoomLabel := label.NewTypography("Zoom: 100%", label.Small, "")

	// Initialize Button components
	primaryBtn := button.New(button.Config{
		Text:    "Primary Button",
		Variant: theme.VariantDefault,
		Size:    theme.SizeDefault,
		OnClick: func() { log.Println("Primary button clicked!") },
	})

	destructiveBtn := button.New(button.Config{
		Text:    "Destructive",
		Variant: theme.VariantDestructive,
		Size:    theme.SizeDefault,
		OnClick: func() { log.Println("Destructive button clicked!") },
	})

	outlineBtn := button.New(button.Config{
		Text:    "Outline",
		Variant: theme.VariantOutline,
		Size:    theme.SizeDefault,
		OnClick: func() { log.Println("Outline button clicked!") },
	})

	secondaryBtn := button.New(button.Config{
		Text:    "Secondary",
		Variant: theme.VariantSecondary,
		Size:    theme.SizeDefault,
		OnClick: func() { log.Println("Secondary button clicked!") },
	})

	ghostBtn := button.New(button.Config{
		Text:    "Ghost",
		Variant: theme.VariantGhost,
		Size:    theme.SizeDefault,
		OnClick: func() { log.Println("Ghost button clicked!") },
	})

	linkBtn := button.New(button.Config{
		Text:    "Link",
		Variant: theme.VariantLink,
		Size:    theme.SizeDefault,
		OnClick: func() { log.Println("Link button clicked!") },
	})

	// Theme toggle
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

	// Zoom controls
	zoomInBtn := button.New(button.Config{
		Text:    "+",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			if zoomScale < maxZoom {
				zoomScale += zoomStep
				if zoomScale > maxZoom {
					zoomScale = maxZoom
				}
				zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
				w.Invalidate()
			}
		},
	})

	zoomOutBtn := button.New(button.Config{
		Text:    "-",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			if zoomScale > minZoom {
				zoomScale -= zoomStep
				if zoomScale < minZoom {
					zoomScale = minZoom
				}
				zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
				w.Invalidate()
			}
		},
	})

	zoomResetBtn := button.New(button.Config{
		Text:    "Reset",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			zoomScale = 1.0
			zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
			w.Invalidate()
		},
	})

	// Containers & Inputs
	demoCard := card.New(card.Config{Variant: theme.VariantDefault})
	textInput := input.Text("Enter track name...")

	// Badges
	defaultBadge := badge.New(badge.Config{Text: "Default", Variant: theme.VariantDefault})
	secBadge := badge.New(badge.Config{Text: "Secondary", Variant: theme.VariantSecondary})
	outBadge := badge.New(badge.Config{Text: "Outline", Variant: theme.VariantOutline})
	destBadge := badge.New(badge.Config{Text: "Destructive", Variant: theme.VariantDestructive})

	// Progress & Toggles
	progBar := progress.New(progress.Config{Value: 0.65})
	sepLine := separator.New(separator.Config{Horizontal: true})
	audioSwitch := switchcomp.New(switchcomp.Config{Value: true, OnChange: func(v bool) { log.Printf("Audio switch: %v", v) }})
	gpuCheckbox := checkbox.New(checkbox.Config{Value: true, OnChange: func(v bool) { log.Printf("GPU checkbox: %v", v) }})

	// Alert
	systemAlert := alert.New(alert.Config{
		Title:       "ASIO Driver Active",
		Description: "Buffer size set to 64 samples (1.2ms latency).",
	})

	titleLabel := label.NewTypography("Gio-shadcn Showcase", label.H1, "")
	subtitleLabel := label.NewTypography("Ported from Flutter shadcn_flutter Source of Truth", label.P, "")

	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Process global zoom key events
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
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return tb.Layout(gtx, th, w)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    th.Spacing.Space6,
						Bottom: th.Spacing.Space6,
						Left:   th.Spacing.Space6,
						Right:  th.Spacing.Space6,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return titleLabel.Layout(gtx, th) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions { return subtitleLabel.Layout(gtx, th) }),
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
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    th.Spacing.Space2,
						Bottom: th.Spacing.Space6,
						Left:   th.Spacing.Space6,
						Right:  th.Spacing.Space6,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return demoCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								// Section 1: Buttons
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									sectionTitle := label.NewTypography("Buttons & Badges", label.H4, "")
									return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sectionTitle.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layoutButtonRow(gtx, th, primaryBtn, destructiveBtn, outlineBtn, secondaryBtn, ghostBtn, linkBtn)
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
											)
										}),
									)
								}),

								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sepLine.Layout(gtx, th) }),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

								// Section 2: Inputs & Toggles
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									sectionTitle := label.NewTypography("Inputs & Toggles", label.H4, "")
									return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sectionTitle.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											maxWidth := gtx.Metric.Dp(400)
											if gtx.Constraints.Max.X < maxWidth {
												maxWidth = gtx.Constraints.Max.X
											}
											gtx.Constraints.Max.X = maxWidth
											gtx.Constraints.Min.X = maxWidth
											return textInput.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return audioSwitch.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return label.NewTypography("High Quality Audio (96kHz)", label.Small, "").Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return gpuCheckbox.Layout(gtx, th) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return label.NewTypography("Hardware GPU Acceleration", label.Small, "").Layout(gtx, th)
												}),
											)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return label.NewTypography("Audio Buffer Progress (65%)", label.Small, "").Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions { return progBar.Layout(gtx, th) }),
											)
										}),
									)
								}),

								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sepLine.Layout(gtx, th) }),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

								// Section 3: Alerts
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									sectionTitle := label.NewTypography("Alert Callout", label.H4, "")
									return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return sectionTitle.Layout(gtx, th) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions { return systemAlert.Layout(gtx, th) }),
									)
								}),
							)
						})
					})
				}),
			)

			e.Frame(&ops)
		}
	}
}
