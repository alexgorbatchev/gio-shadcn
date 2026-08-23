# Titlebar Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/titlebar](https://ui.shadcn.com/docs/components/titlebar)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `titlebar` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Custom Window Titlebar**
  - implementation:
    - titlebar.go:87
  - tests:
    - titlebar_test.go:13

### Capabilities & Features
- [x] **Title Text**
  - implementation:
    - titlebar.go:121
  - tests:
    - titlebar_test.go:15
- [x] **Window Control Buttons (Close, Minimize, Maximize)**
  - implementation:
    - titlebar.go:142
  - tests:
    - titlebar_test.go:27
- [x] **Window Drag Region**
  - implementation:
    - titlebar.go:175
  - tests:
    - titlebar_test.go:27

---



### Demos
- [x] **1. Window Title Bar**
  - implementation:
    - component: titlebar.go:50
    - demo: demo.md:13
  - tests:
    - titlebar_test.go:15

## Code Structure & Entry Points
- `titlebar.go`: Primary component implementation.
- `titlebar_test.go`: Unit test suite.
- `demo.md`: Component interactive demo snippets.
- `AGENTS.md`: Component specification.