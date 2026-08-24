# Slider Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/slider](https://ui.shadcn.com/docs/components/slider)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `slider` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Range Slider Fader**
  - implementation:
    - slider.go:44
  - tests:
    - slider_test.go:13

### Capabilities & Features
- [x] **Interactive Pointer Drag**
  - implementation:
    - slider.go:88
  - tests:
    - slider_test.go:36
- [x] **Min / Max Value Range**
  - implementation:
    - slider.go:45
  - tests:
    - slider_test.go:13
- [x] **Track Fill Portion**
  - implementation:
    - slider.go:151
  - tests:
    - slider_test.go:51
- [x] **Circular Thumb Knob**
  - implementation:
    - slider.go:160
  - tests:
    - slider_test.go:51
- [x] **Clamped Radius Half-Height**
  - implementation:
    - slider.go:145
  - tests:
    - slider_test.go:51

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/slider-controlled.tsx
  - implementation:
    - component: slider.go:69
    - demo: demo.go:16
  - tests:
    - slider_test.go:37
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/slider-demo.tsx
  - implementation:
    - component: slider.go:69
    - demo: demo.go:16
  - tests:
    - slider_test.go:50
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/slider-disabled.tsx
  - implementation:
    - component: slider.go:69
    - demo: demo.go:16
  - tests:
    - slider_test.go:24
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/slider-multiple.tsx
  - implementation:
    - component: slider.go:69
    - demo: demo.go:16
  - tests:
    - slider_test.go:50
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/slider-range.tsx
  - implementation:
    - component: slider.go:69
    - demo: demo.go:16
  - tests:
    - slider_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/slider-vertical.tsx
  - implementation:
    - component: slider.go:69
    - demo: demo.go:16
  - tests:
    - slider_test.go:50

---

## Code Structure & Entry Points
- `slider.go`: Primary component widget layout and state logic.
- `slider_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Interactive component gallery demo.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
