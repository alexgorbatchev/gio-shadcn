# Resizable Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/resizable](https://ui.shadcn.com/docs/components/resizable)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `resizable` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Horizontal Split Panels**
  - implementation:
    - resizable.go:34
  - tests:
    - resizable_test.go:13

### Capabilities & Features
- [x] **Left Widget Container**
  - implementation:
    - resizable.go:58
  - tests:
    - resizable_test.go:25
- [x] **Right Widget Container**
  - implementation:
    - resizable.go:73
  - tests:
    - resizable_test.go:25
- [x] **Divider Handle Bar**
  - implementation:
    - resizable.go:66
  - tests:
    - resizable_test.go:32

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/resizable-demo.tsx
  - implementation:
    - component: resizable.go:47
    - demo: demo.go:13
  - tests:
    - resizable_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/resizable-handle.tsx
  - implementation:
    - component: resizable.go:66
    - demo: demo.go:13
  - tests:
    - resizable_test.go:21
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/resizable-vertical.tsx
  - implementation:
    - component: resizable.go:47
    - demo: demo.go:13
  - tests:
    - resizable_test.go:21

---

## Code Structure & Entry Points
- `resizable.go`: Primary component implementation.
- `resizable_test.go`: Unit test suite.
- `demo.go`: Modular component interactive demo layout (`components/resizable/demo.go`).
- `AGENTS.md`: Component specification.