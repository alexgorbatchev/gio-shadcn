# Pagination Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/pagination](https://ui.shadcn.com/docs/components/pagination)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `pagination` component implementation or unit tests in this directory, developers and AI agents MUST:
1. **Maintain Parity:** Preserve strict 100% visual token and feature parity with `shadcn/ui` and the source-of-truth Flutter POC implementation.
2. **Keep Checklist Updated:** Update the feature checklist below if new variants, subcomponents, configuration options, or interactive states are added.
3. **GPU State Safety:** Ensure all rounded rectangle drawings pass through `theme.DrawRRectBackground` / `theme.DrawStroke` with radius clamped to half the component height (`heightPx / 2`).
4. **Prevent Color State Leaking:** End all `Layout` functions with an explicit GPU color state reset:
   ```go
   paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)
   ```
5. **Unit Test Coverage:** Verify that `go test ./...` passes cleanly with zero compilation or `go vet` errors.

---

## Component Variants & Features Checklist

### Variants
- [x] **Standard Page Bar (Pages 1-5)**
  - implementation:
    - pagination.go:39
  - tests:
    - pagination_test.go:13
- [x] **Active Page Highlight Variant**
  - implementation:
    - pagination.go:91
  - tests:
    - pagination_test.go:23

### Capabilities & Features
- [x] **Previous / Next Navigation Buttons**
  - implementation:
    - pagination.go:83
  - tests:
    - pagination_test.go:33
- [x] **OnSelectPage Event Callback**
  - implementation:
    - pagination.go:67
  - tests:
    - pagination_test.go:50
- [x] **Page Button Hover & Active Styling**
  - implementation:
    - pagination.go:116
  - tests:
    - pagination_test.go:33

### Demos
- [x] **Standard Pagination**
  - implementation:
    - pagination.go:39
  - tests:
    - pagination_test.go:13
- [x] **Previous / Next Buttons**
  - implementation:
    - pagination.go:83
  - tests:
    - pagination_test.go:33
- [x] **Page Controls**
  - implementation:
    - pagination.go:90
  - tests:
    - pagination_test.go:23

---

## Code Structure & Entry Points
- `pagination.go`: Primary component widget layout and state logic.
- `pagination_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
