# Text Area Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/textarea](https://ui.shadcn.com/docs/components/textarea)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `textarea` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Standard Text Area**
  - implementation:
    - textarea.go:42
  - tests:
    - textarea_test.go:13
- [x] **Prefilled Text Area**
  - implementation:
    - textarea.go:45
  - tests:
    - textarea_test.go:14
- [x] **Disabled Text Area**
  - implementation:
    - textarea.go:73
  - tests:
    - textarea_test.go:24

### Capabilities & Features
- [x] **Multi-Line Text Editing**
  - implementation:
    - textarea.go:44
  - tests:
    - textarea_test.go:13
- [x] **Placeholder Text**
  - implementation:
    - textarea.go:107
  - tests:
    - textarea_test.go:16
- [x] **Focus Ring Stroke**
  - implementation:
    - textarea.go:124
  - tests:
    - textarea_test.go:33

---



### Demos
- [x] **1. Multi-Line Text Area**
  - implementation:
    - component: textarea.go:50
    - demo: demo.md:13
  - tests:
    - textarea_test.go:15

## Code Structure & Entry Points
- `textarea.go`: Primary component implementation.
- `textarea_test.go`: Unit test suite.
- `demo.md`: Component interactive demo snippets.
- `AGENTS.md`: Component specification.