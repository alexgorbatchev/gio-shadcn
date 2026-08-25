# Spinner Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/spinner](https://ui.shadcn.com/docs/components/spinner)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `spinner` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Activity Loading Spinner**
  - implementation:
    - spinner.go:33
  - tests:
    - spinner_test.go:12

### Capabilities & Features
- [x] **Circular Arc Stroke**
  - implementation:
    - spinner.go:63
  - tests:
    - spinner_test.go:22
- [x] **Spin Animation Indicator**
  - implementation:
    - spinner.go:45
  - tests:
    - spinner_test.go:69

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/spinner-badge.tsx
  - implementation:
    - component: spinner.go:45
    - demo: demo.go:46
  - tests:
    - spinner_test.go:29
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/spinner-button.tsx
  - implementation:
    - component: spinner.go:45
    - demo: demo.go:53
  - tests:
    - spinner_test.go:39
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/spinner-custom.tsx
  - implementation:
    - component: spinner.go:45
    - demo: demo.go:60
  - tests:
    - spinner_test.go:49
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/spinner-demo.tsx
  - implementation:
    - component: spinner.go:45
    - demo: demo.go:61
  - tests:
    - spinner_test.go:69
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/spinner-empty.tsx
  - implementation:
    - component: spinner.go:45
    - demo: demo.go:63
  - tests:
    - spinner_test.go:59
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/spinner-size.tsx
  - implementation:
    - component: spinner.go:33
    - demo: demo.go:67
  - tests:
    - spinner_test.go:12

---

## Code Structure & Entry Points
- `spinner.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `spinner_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
