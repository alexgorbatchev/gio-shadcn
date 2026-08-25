# Drawer Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/drawer](https://ui.shadcn.com/docs/components/drawer)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `drawer` component implementation or unit tests in this directory, developers and AI agents MUST:
1. **Maintain Parity:** Preserve strict 100% visual token and feature parity with `shadcn/ui` and the source-of-truth Flutter POC implementation.
2. **Keep Checklist Updated:** Update the feature checklist below if new variants, subcomponents, configuration options, or interactive states are added.
3. **GPU State Safety:** Ensure all rounded rectangle drawings pass through `theme.DrawRRectBackground` / `theme.DrawStroke` with radius clamped to half the component height (`heightPx / 2`).
4. **Prevent Color State Leaks:** End all `Layout` functions with an explicit GPU color state reset:
   ```go
   paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)
   ```
5. **Unit Test Coverage:** Verify that `go test ./...` passes cleanly with zero compilation or `go vet` errors.

---

## Component Variants & Features Checklist

### Variants
- [x] **Bottom Sheet Drawer Panel**
  - implementation:
    - drawer.go:49
  - tests:
    - drawer_test.go:13

### Capabilities & Features
- [x] **Dark Backdrop Overlay & Outside Click-To-Close**
  - implementation:
    - drawer.go:107
  - tests:
    - drawer_test.go:13
- [x] **South Viewport Edge Alignment**
  - implementation:
    - drawer.go:121
  - tests:
    - drawer_test.go:47
- [x] **Drag Handle Indicator**
  - implementation:
    - drawer.go:153
  - tests:
    - drawer_test.go:35
- [x] **Header Title & Close Button**
  - implementation:
    - drawer.go:164
  - tests:
    - drawer_test.go:24
- [x] **Custom Content Widget**
  - implementation:
    - drawer.go:193
  - tests:
    - drawer_test.go:13

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/drawer-demo.tsx
  - implementation:
    - component: drawer.go:49
    - demo: demo.go:26
  - tests:
    - drawer_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/drawer-dialog.tsx
  - implementation:
    - component: drawer.go:49
    - demo: demo.go:48
  - tests:
    - drawer_test.go:24
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/drawer-sides.tsx
  - implementation:
    - component: drawer.go:49
    - demo: demo.go:68
  - tests:
    - drawer_test.go:35
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/drawer-swipe-handle.tsx
  - implementation:
    - component: drawer.go:153
    - demo: demo.go:83
  - tests:
    - drawer_test.go:47

---

## Code Structure & Entry Points
- `drawer.go`: Primary component widget layout and state logic.
- `drawer_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Exported interactive demo widget (`Demo`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
