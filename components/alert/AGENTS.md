# Alert Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/alert](https://ui.shadcn.com/docs/components/alert)

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
    - alert_test.go:25
- [x] **Destructive / Warning Alert**
  - implementation:
    - alert.go:64
  - tests:
    - alert_test.go:60

### Capabilities & Features
- [x] **Title Header**
  - implementation:
    - alert.go:88
  - tests:
    - alert_test.go:35
- [x] **Description Body**
  - implementation:
    - alert.go:102
  - tests:
    - alert_test.go:46
- [x] **Action Widget Support**
  - implementation:
    - alert.go:128
  - tests:
    - alert_test.go:13
- [x] **Icon Support**
  - implementation:
    - alert.go:121
  - tests:
    - alert_test.go:25
- [x] **Border & Background Drawing**
  - implementation:
    - alert.go:136
  - tests:
    - alert_test.go:82

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-action.tsx
  - implementation:
    - component: alert.go:128
    - demo: demo.go:32
  - tests:
    - alert_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-basic.tsx
  - implementation:
    - component: alert.go:42
    - demo: demo.go:42
  - tests:
    - alert_test.go:25
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-colors.tsx
  - implementation:
    - component: alert.go:136
    - demo: demo.go:49
  - tests:
    - alert_test.go:35
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-demo.tsx
  - implementation:
    - component: alert.go:42
    - demo: demo.go:56
  - tests:
    - alert_test.go:46
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-destructive.tsx
  - implementation:
    - component: alert.go:64
    - demo: demo.go:68
  - tests:
    - alert_test.go:60
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-rtl.tsx
  - implementation:
    - component: alert.go:42
    - demo: demo.go:76
  - tests:
    - alert_test.go:72

---

## Code Structure & Entry Points
- `alert.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `alert_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
