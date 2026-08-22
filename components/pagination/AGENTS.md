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
- [x] **Pagination Bar**
  - implementation:
    - pagination.go:38
  - tests:
    - pagination_test.go:13

### Capabilities & Features
- [x] **Previous Page Button**
  - implementation:
    - pagination.go:64
  - tests:
    - pagination_test.go:13
- [x] **Next Page Button**
  - implementation:
    - pagination.go:71
  - tests:
    - pagination_test.go:13
- [x] **Page Number Buttons**
  - implementation:
    - pagination.go:94
  - tests:
    - pagination_test.go:23
- [x] **Active Page Highlight**
  - implementation:
    - pagination.go:124
  - tests:
    - pagination_test.go:18

---

## Code Structure & Entry Points
- `pagination.go`: Primary component widget layout and state logic.
- `pagination_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
