# Select Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/select](https://ui.shadcn.com/docs/components/select)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `select` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Select Dropdown Field**
  - implementation:
    - select.go:2
  - tests:
    - select_test.go:1

### Capabilities & Features
- [x] **Option Items List**
  - implementation:
    - select.go:4
  - tests:
    - select_test.go:15
- [x] **Active Selected Value**
  - implementation:
    - select.go:74
  - tests:
    - select_test.go:13
- [x] **Open/Close Overlay Dropdown**
  - implementation:
    - select.go:74
  - tests:
    - select_test.go:13

---

## Code Structure & Entry Points
- `select.go`: Primary component widget layout and state logic.
- `select_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
