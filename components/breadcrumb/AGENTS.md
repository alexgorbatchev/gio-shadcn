# Breadcrumb Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/breadcrumb](https://ui.shadcn.com/docs/components/breadcrumb)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `breadcrumb` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Standard Path**
  - implementation:
    - breadcrumb.go:23
  - tests:
    - breadcrumb_test.go:13
- [x] **Active Page Link**
  - implementation:
    - breadcrumb.go:76
  - tests:
    - breadcrumb_test.go:24
- [x] **Custom Separators**
  - implementation:
    - breadcrumb.go:84
  - tests:
    - breadcrumb_test.go:35

### Capabilities & Features
- [x] **Horizontal Flex Layout**
  - implementation:
    - breadcrumb.go:91
  - tests:
    - breadcrumb_test.go:52
- [x] **Chevron Dividers**
  - implementation:
    - breadcrumb.go:85
  - tests:
    - breadcrumb_test.go:69
- [x] **Hover Highlight**
  - implementation:
    - breadcrumb.go:71
  - tests:
    - breadcrumb_test.go:87

---

## Code Structure & Entry Points
- `breadcrumb.go`: Primary component widget layout and state logic.
- `breadcrumb_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
