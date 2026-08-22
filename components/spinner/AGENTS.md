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
- [x] **Small Spinner (16dp)**
  - implementation:
    - spinner.go:32
  - tests:
    - spinner_test.go:13
- [x] **Default Activity Spinner (24dp)**
  - implementation:
    - spinner.go:32
  - tests:
    - spinner_test.go:22
- [x] **Large Loader Spinner (48dp)**
  - implementation:
    - spinner.go:32
  - tests:
    - spinner_test.go:31

### Capabilities & Features
- [x] **Circular Arc Stroke Path**
  - implementation:
    - spinner.go:75
  - tests:
    - spinner_test.go:40
- [x] **Spin Animation Indicator**
  - implementation:
    - spinner.go:47
  - tests:
    - spinner_test.go:40

### Demos
- [x] **Small Spinner Demo**
  - implementation:
    - spinner.go:32
  - tests:
    - spinner_test.go:13
- [x] **Default Activity Spinner Demo**
  - implementation:
    - spinner.go:32
  - tests:
    - spinner_test.go:22
- [x] **Large Loader Spinner Demo**
  - implementation:
    - spinner.go:32
  - tests:
    - spinner_test.go:31

---

## Code Structure & Entry Points
- `spinner.go`: Primary component widget layout and state logic.
- `spinner_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
