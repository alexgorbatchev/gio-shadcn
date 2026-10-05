# gio-shadcn

Native Go components for [Gio](https://gioui.org/), using design tokens inspired by [shadcn/ui](https://ui.shadcn.com/). Components keep their input state across frames and share light or dark themes with embedded Geist fonts.

The library is under development. Component specifications in `components/*/AGENTS.md` distinguish implemented behavior from upstream examples that have not been verified. This project does not claim complete shadcn/ui parity.

## Install

The module requires Go 1.26.2 or later and Gio's [platform prerequisites](https://gioui.org/doc/install).

```sh
go get github.com/bnema/gio-shadcn
```

## Use a component

Construct stateful components once, outside the frame loop. Most packages use `New(Config{...})`; input, label, and titlebar also offer functional options. Consult the package source for each constructor.

```go
package main

import (
    "log"
    "os"

    "gioui.org/app"
    "gioui.org/layout"
    "gioui.org/op"
    "gioui.org/op/paint"
    "github.com/bnema/gio-shadcn/components/button"
    "github.com/bnema/gio-shadcn/theme"
)

func main() {
    go func() {
        w := new(app.Window)
        w.Option(app.Title("gio-shadcn"))
        th := theme.New()
        defer th.ReleaseBackdrop()
        btn := button.New(button.Config{
            Text: "Click me",
            Variant: theme.VariantDefault,
            OnClick: func() { log.Print("clicked") },
        })
        var ops op.Ops
        for {
            switch e := w.Event().(type) {
            case app.DestroyEvent:
                if e.Err != nil { log.Print(e.Err) }
                os.Exit(0)
            case app.FrameEvent:
                gtx := app.NewContext(&ops, e)
                paint.Fill(&ops, th.Colors.Background)
                layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
                    return btn.Layout(gtx, th)
                })
                th.RenderOverlays(gtx)
                paint.ColorOp{Color: th.Colors.Background}.Add(&ops)
                e.Frame(&ops)
            }
        }
    }()
    app.Main()
}
```

## Components

The repository contains 42 component packages:

| Area | Packages |
| --- | --- |
| Actions and text | button, badge, label, titlebar |
| Forms | input, textarea, numberinput, inputotp, checkbox, switch, radio, select, togglegroup, slider |
| Containers and layout | card, aspectratio, scrollarea, separator, resizable, carousel |
| Navigation and data | accordion, collapsible, tree, tabs, breadcrumb, pagination, table |
| Overlays | dialog, sheet, drawer, popover, dropdownmenu, tooltip, hovercard, command |
| Feedback | avatar, progress, skeleton, spinner, alert, toast, empty |

## Themes and overlays

Use `theme.New()` for light mode or `theme.NewDark()` for dark mode. `th.ToggleDark()` switches the active color scheme. Customize spacing and typography through the theme fields; corner tokens include `th.Radius.RadiusMD`. `theme.NewThemeFromJSON(path)` loads light and dark colors and initializes the default fonts, spacing, typography and radii. Other JSON configuration fields are currently metadata; customize those theme fields in Go.

Render queued overlays after the main layout, using the full window constraints. Release the theme's backdrop resources when the window closes. The gallery in [demo/demo.go](demo/demo.go) shows backdrop capture and root overlay rendering.

Tooltip and hover card expose `LayoutTrigger(gtx, th, widget)` for pointer hover and keyboard focus. A tooltip's `Layout` draws only while `Open` is true; hover card's `Layout` uses `Hovered` for externally controlled previews.

## Develop

```sh
go run ./demo/cmd              # Interactive component gallery
go build -o bin/demo-app ./demo/cmd
go test ./...
go vet ./...
go fmt ./...
```

Equivalent tasks are available through the [justfile](justfile). Read [AGENTS.md](AGENTS.md) and the relevant component specification before changing code. Behavioral changes require regression tests, GPU drawing must use the theme helpers, and documentation mappings must be checked manually.

## License

[MIT](LICENSE).
