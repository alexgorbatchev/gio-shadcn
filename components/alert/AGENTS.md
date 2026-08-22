# Alert Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/alert](https://ui.shadcn.com/docs/components/alert)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `alert` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Default / Info Alert**
  - implementation:
    - alert.go:42
  - tests:
    - alert_test.go:13
- [x] **Destructive Warning Alert**
  - implementation:
    - alert.go:63
  - tests:
    - alert_test.go:23

### Capabilities & Features
- [x] **Title Header**
  - implementation:
    - alert.go:94
  - tests:
    - alert_test.go:33
- [x] **Description Body**
  - implementation:
    - alert.go:109
  - tests:
    - alert_test.go:42
- [x] **Variant Background Styling**
  - implementation:
    - alert.go:121
  - tests:
    - alert_test.go:51
- [x] **Border Stroke**
  - implementation:
    - alert.go:124
  - tests:
    - alert_test.go:66

### Demos (Official shadcn Demos)
- [x] **1. Default / Info Alert**
  - implementation:
    - alert.go:42
  - tests:
    - alert_test.go:13
- [x] **2. Destructive Warning Alert**
  - implementation:
    - alert.go:63
  - tests:
    - alert_test.go:23
- [x] **3. With Title & Description**
  - implementation:
    - alert.go:94
  - tests:
    - alert_test.go:33
- [x] **4. Custom Styling**
  - implementation:
    - alert.go:69
  - tests:
    - alert_test.go:51

---

## Code Structure & Entry Points
- `alert.go`: Primary component widget layout and state logic.
- `alert_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
