# Pagination Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/pagination](https://ui.shadcn.com/docs/components/pagination)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `pagination` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Standard Page Bar Variant**
  - implementation:
    - pagination.go:41
  - tests:
    - pagination_test.go:13

### Capabilities & Features
- [x] **Previous & Next Navigation Buttons**
  - implementation:
    - pagination.go:61
  - tests:
    - pagination_test.go:34
- [x] **Active Page Button Highlight**
  - implementation:
    - pagination.go:81
  - tests:
    - pagination_test.go:24
- [x] **OnSelectPage Event Callback**
  - implementation:
    - pagination.go:65
  - tests:
    - pagination_test.go:48

### Demos
- [x] **1. Standard Pagination Bar**
  - implementation:
    - component: pagination.go:58
    - demo: demo.go:9
  - tests:
    - pagination_test.go:13
- [x] **2. Active Page Highlight**
  - implementation:
    - component: pagination.go:81
    - demo: demo.go:9
  - tests:
    - pagination_test.go:24
- [x] **3. Page Controls**
  - implementation:
    - component: pagination.go:61
    - demo: demo.go:9
  - tests:
    - pagination_test.go:34

---

## Code Structure & Entry Points
- `pagination.go`: Primary component widget layout and state logic.
- `pagination_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Modular component interactive demo layout (`components/pagination/demo.go`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
