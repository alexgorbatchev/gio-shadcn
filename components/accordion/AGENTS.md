# Accordion Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/accordion](https://ui.shadcn.com/docs/components/accordion)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `accordion` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Single Open Item**
  - implementation:
    - accordion.go:55
  - tests:
    - accordion_test.go:13
- [x] **Multiple Open Items**
  - implementation:
    - accordion.go:55
  - tests:
    - accordion_test.go:13
- [x] **Collapsed State**
  - implementation:
    - accordion.go:55
  - tests:
    - accordion_test.go:13

### Capabilities & Features
- [x] **Expand/Collapse Animation State**
  - implementation:
    - accordion.go:55
  - tests:
    - accordion_test.go:13
- [x] **Item Header Button**
  - implementation:
    - accordion.go:22
  - tests:
    - accordion_test.go:15
- [x] **Item Content Panel**
  - implementation:
    - accordion.go:22
  - tests:
    - accordion_test.go:15
- [x] **Border Dividers**
  - implementation:
    - accordion.go:85
  - tests:
    - accordion_test.go:13

---

## Code Structure & Entry Points
- `accordion.go`: Primary component widget layout and state logic.
- `accordion_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
