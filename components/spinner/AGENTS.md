# Spinner Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/spinner](https://ui.shadcn.com/docs/components/spinner)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `spinner` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Activity Loading Spinner**
  - implementation:
    - spinner.go:34
  - tests:
    - spinner_test.go:13

### Capabilities & Features
- [x] **Circular Arc Stroke**
  - implementation:
    - spinner.go:68
  - tests:
    - spinner_test.go:31
- [x] **Spin Animation Indicator**
  - implementation:
    - spinner.go:88
  - tests:
    - spinner_test.go:33

---

### Demos
- [x] **Small Spinner**
  - implementation:
    - spinner.go:48
  - tests:
    - spinner_test.go:22
- [x] **Default Activity Spinner**
  - implementation:
    - spinner.go:49
  - tests:
    - spinner_test.go:15
- [x] **Large Loader Spinner**
  - implementation:
    - spinner.go:50
  - tests:
    - spinner_test.go:24

## Code Structure & Entry Points
- `spinner.go`: Primary component widget layout and state logic.
- `spinner_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
