# Radio Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/radio-group](https://ui.shadcn.com/docs/components/radio-group)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `radio` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Selected Radio Option**
  - implementation:
    - radio.go:37
  - tests:
    - radio_test.go:12
- [x] **Unselected Radio Option**
  - implementation:
    - radio.go:37
  - tests:
    - radio_test.go:22
- [x] **Disabled Radio Option**
  - implementation:
    - radio.go:73
  - tests:
    - radio_test.go:32

### Capabilities & Features
- [x] **Pointer Click Selection**
  - implementation:
    - radio.go:56
  - tests:
    - radio_test.go:12
- [x] **Outer Circle & Inner Selected Dot Ellipse**
  - implementation:
    - radio.go:98
  - tests:
    - radio_test.go:12

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/radio-group-demo.tsx
  - implementation:
    - component: radio.go:50
    - demo: demo.go:27
  - tests:
    - radio_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/radio-group-disabled.tsx
  - implementation:
    - component: radio.go:73
    - demo: demo.go:45
  - tests:
    - radio_test.go:32
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/radio-fields.tsx
  - implementation:
    - component: radio.go:50
    - demo: demo.go:27
  - tests:
    - radio_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/radio-group-choice-card.tsx
  - implementation:
    - component: radio.go:50
    - demo: demo.go:27
  - tests:
    - radio_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/radio-group-description.tsx
  - implementation:
    - component: radio.go:50
    - demo: demo.go:27
  - tests:
    - radio_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/radio-group-invalid.tsx
  - implementation:
    - component: radio.go:50
    - demo: demo.go:27
  - tests:
    - radio_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/radio-group-rtl.tsx
  - implementation:
    - component: radio.go:50
    - demo: demo.go:27
  - tests:
    - radio_test.go:12

---

## Code Structure & Entry Points
- `radio.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `radio_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
