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
    - drawer.go:21
  - tests:
    - drawer_test.go:13

### Capabilities & Features
- [x] **Dark Backdrop Overlay**
  - implementation:
    - drawer.go:121
  - tests:
    - drawer_test.go:23
- [x] **Backdrop Click-To-Close**
  - implementation:
    - drawer.go:101
  - tests:
    - drawer_test.go:39
- [x] **South Viewport Edge Alignment (`layout.S`)**
  - implementation:
    - drawer.go:126
  - tests:
    - drawer_test.go:53
- [x] **Drag Handle Indicator Bar**
  - implementation:
    - drawer.go:172
  - tests:
    - drawer_test.go:68
- [x] **Close Button (`closeBtn`)**
  - implementation:
    - drawer.go:181
  - tests:
    - drawer_test.go:77
- [x] **Custom / Illustrated Content Body**
  - implementation:
    - drawer.go:202
  - tests:
    - drawer_test.go:86

### Demos (Official shadcn Demos)
- [x] **1. Bottom Sheet Drawer**
  - implementation:
    - drawer.go:92
  - tests:
    - drawer_test.go:13
- [x] **2. Drag Handle Bar**
  - implementation:
    - drawer.go:172
  - tests:
    - drawer_test.go:68
- [x] **3. Telemetry Metrics Content**
  - implementation:
    - drawer.go:202
  - tests:
    - drawer_test.go:86
- [x] **4. Backdrop Click-To-Close**
  - implementation:
    - drawer.go:101
  - tests:
    - drawer_test.go:39

---

## Code Structure & Entry Points
- `drawer.go`: Primary component widget layout and state logic.
- `drawer_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
