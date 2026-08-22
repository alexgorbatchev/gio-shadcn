# Checkbox Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/checkbox](https://ui.shadcn.com/docs/components/checkbox)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `checkbox` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Unchecked State**
  - implementation:
    - checkbox.go:42
  - tests:
    - checkbox_test.go:13
- [x] **Checked State**
  - implementation:
    - checkbox.go:73
  - tests:
    - checkbox_test.go:22
- [x] **Disabled State**
  - implementation:
    - checkbox.go:78
  - tests:
    - checkbox_test.go:31

### Capabilities & Features
- [x] **Interactive Click Toggle**
  - implementation:
    - checkbox.go:60
  - tests:
    - checkbox_test.go:41
- [x] **Checkmark Vector Path Drawing**
  - implementation:
    - checkbox.go:104
  - tests:
    - checkbox_test.go:54
- [x] **Label Association**
  - implementation:
    - checkbox.go:88
  - tests:
    - checkbox_test.go:69

---

### Demos
- [x] **Default Checkbox**
  - implementation:
    - checkbox.go:53
  - tests:
    - checkbox_test.go:23
- [x] **With Label**
  - implementation:
    - checkbox.go:42
  - tests:
    - checkbox_test.go:33
- [x] **Disabled Checkbox**
  - implementation:
    - checkbox.go:27
  - tests:
    - checkbox_test.go:32
- [x] **Form Control Checkbox**
  - implementation:
    - checkbox.go:93
  - tests:
    - checkbox_test.go:55

## Code Structure & Entry Points
- `checkbox.go`: Primary component widget layout and state logic.
- `checkbox_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
