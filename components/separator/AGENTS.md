# Separator Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/separator](https://ui.shadcn.com/docs/components/separator)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `separator` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Horizontal Separator**
  - implementation:
    - separator.go:58
  - tests:
    - separator_test.go:13
- [x] **Vertical Separator**
  - implementation:
    - separator.go:64
  - tests:
    - separator_test.go:25

### Capabilities & Features
- [x] **1px Divider Stroke Line**
  - implementation:
    - separator.go:76
  - tests:
    - separator_test.go:33
- [x] **Muted Border Theme Color**
  - implementation:
    - separator.go:71
  - tests:
    - separator_test.go:25

---



### Demos
- [x] **1. Horizontal Separator Divider**
  - implementation:
    - component: separator.go:50
    - demo: demo.md:13
  - tests:
    - separator_test.go:15

## Code Structure & Entry Points
- `separator.go`: Primary component implementation.
- `separator_test.go`: Unit test suite.
- `demo.md`: Component interactive demo snippets.
- `AGENTS.md`: Component specification.