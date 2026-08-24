# Scroll Area Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/scroll-area](https://ui.shadcn.com/docs/components/scroll-area)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `scrollarea` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Custom Scrollable Container**
  - implementation:
    - scrollarea.go:31
  - tests:
    - scrollarea_test.go:13

### Capabilities & Features
- [x] **Vertical / Horizontal Scroll Tracks**
  - implementation:
    - scrollarea.go:34
  - tests:
    - scrollarea_test.go:19
- [x] **Custom Scroll Thumb Bar**
  - implementation:
    - scrollarea.go:68
  - tests:
    - scrollarea_test.go:33

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/scroll-area-demo.tsx
  - implementation:
    - component: scrollarea.go:48
    - demo: demo.go:13
  - tests:
    - scrollarea_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/scroll-area-horizontal-demo.tsx
  - implementation:
    - component: scrollarea.go:48
    - demo: demo.go:13
  - tests:
    - scrollarea_test.go:24

---

## Code Structure & Entry Points
- `scrollarea.go`: Primary component implementation.
- `scrollarea_test.go`: Unit test suite.
- `demo.go`: Modular component interactive demo layout (`components/scrollarea/demo.go`).
- `AGENTS.md`: Component specification.