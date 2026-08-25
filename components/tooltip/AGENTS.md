# Tooltip Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/tooltip](https://ui.shadcn.com/docs/components/tooltip)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `tooltip` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Tooltip Callout**
  - implementation:
    - tooltip.go:43
  - tests:
    - tooltip_test.go:12

### Capabilities & Features
- [x] **Callout Text Box**
  - implementation:
    - tooltip.go:68
  - tests:
    - tooltip_test.go:12
- [x] **Dark Overlay Background**
  - implementation:
    - tooltip.go:81
  - tests:
    - tooltip_test.go:12

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/tooltip-demo.tsx
  - implementation:
    - component: tooltip.go:43
    - demo: demo.go:16
  - tests:
    - tooltip_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/tooltip-sides.tsx
  - implementation:
    - component: tooltip.go:43
    - demo: demo.go:17
  - tests:
    - tooltip_test.go:22
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/tooltip-keyboard.tsx
  - implementation:
    - component: tooltip.go:43
    - demo: demo.go:17
  - tests:
    - tooltip_test.go:22
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/tooltip-disabled.tsx
  - implementation:
    - component: tooltip.go:43
    - demo: demo.go:16
  - tests:
    - tooltip_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/tooltip-rtl.tsx
  - implementation:
    - component: tooltip.go:43
    - demo: demo.go:16
  - tests:
    - tooltip_test.go:12

---

## Code Structure & Entry Points
- `tooltip.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `tooltip_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
