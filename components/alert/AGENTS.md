# Alert Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/alert](https://ui.shadcn.com/docs/components/alert)  
**Official shadcn Source Spec (.mdx):** [https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/alert.mdx](https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/alert.mdx)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `alert` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Default / Info Alert**
  - implementation:
    - alert.go:42
  - tests:
    - alert_test.go:13
- [x] **Destructive / Warning Alert**
  - implementation:
    - alert.go:64
  - tests:
    - alert_test.go:23

### Capabilities & Features
- [x] **Title Header**
  - implementation:
    - alert.go:88
  - tests:
    - alert_test.go:33
- [x] **Description Body**
  - implementation:
    - alert.go:102
  - tests:
    - alert_test.go:42
- [x] **Variant Background Styling**
  - implementation:
    - alert.go:121
  - tests:
    - alert_test.go:51
- [x] **Border Stroke**
  - implementation:
    - alert.go:124
  - tests:
    - alert_test.go:66

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-action.tsx
  - implementation:
    - component: alert.go:42
    - demo: demo.go:9
  - tests:
    - alert_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-basic.tsx
  - implementation:
    - component: alert.go:42
    - demo: demo.go:9
  - tests:
    - alert_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-colors.tsx
  - implementation:
    - component: alert.go:121
    - demo: demo.go:15
  - tests:
    - alert_test.go:51
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-demo.tsx
  - implementation:
    - component: alert.go:42
    - demo: demo.go:9
  - tests:
    - alert_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-destructive.tsx
  - implementation:
    - component: alert.go:64
    - demo: demo.go:15
  - tests:
    - alert_test.go:23
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-rtl.tsx
  - implementation:
    - component: alert.go:42
    - demo: demo.go:9
  - tests:
    - alert_test.go:13

---

## Code Structure & Entry Points
- `alert.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `alert_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
