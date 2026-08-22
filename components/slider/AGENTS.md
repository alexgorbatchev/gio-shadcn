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
    - slider_test.go:25
- [x] **Min / Max Value Range**
  - implementation:
    - slider.go:45
  - tests:
    - slider_test.go:16
- [x] **Track Fill Portion**
  - implementation:
    - slider.go:151
  - tests:
    - slider_test.go:33
- [x] **Circular Thumb Knob**
  - implementation:
    - slider.go:160
  - tests:
    - slider_test.go:33
- [x] **Clamped Radius Half-Height**
  - implementation:
    - slider.go:142
  - tests:
    - slider_test.go:33

---

### Demos
- [x] **Single Thumb Range Slider**
  - implementation:
    - slider.go:70
  - tests:
    - slider_test.go:26
- [x] **Volume Fader (65%)**
  - implementation:
    - slider.go:71
  - tests:
    - slider_test.go:15
- [x] **Pitch Bend Slider**
  - implementation:
    - slider.go:72
  - tests:
    - slider_test.go:28
- [x] **Disabled Slider**
  - implementation:
    - slider.go:28
  - tests:
    - slider_test.go:13

## Code Structure & Entry Points
- `slider.go`: Primary component widget layout and state logic.
- `slider_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
