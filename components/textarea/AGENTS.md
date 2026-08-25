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
    - textarea_test.go:12
- [x] **Prefilled Text Area**
  - implementation:
    - textarea.go:45
  - tests:
    - textarea_test.go:12
- [x] **Disabled Text Area**
  - implementation:
    - textarea.go:73
  - tests:
    - textarea_test.go:22

### Capabilities & Features
- [x] **Multi-Line Text Editing**
  - implementation:
    - textarea.go:44
  - tests:
    - textarea_test.go:12
- [x] **Placeholder Text**
  - implementation:
    - textarea.go:107
  - tests:
    - textarea_test.go:12
- [x] **Focus Ring Stroke**
  - implementation:
    - textarea.go:124
  - tests:
    - textarea_test.go:12

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/textarea-demo.tsx
  - implementation:
    - component: textarea.go:66
    - demo: demo.go:20
  - tests:
    - textarea_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/textarea-button.tsx
  - implementation:
    - component: textarea.go:66
    - demo: demo.go:22
  - tests:
    - textarea_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/textarea-disabled.tsx
  - implementation:
    - component: textarea.go:73
    - demo: demo.go:21
  - tests:
    - textarea_test.go:22
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/textarea-field.tsx
  - implementation:
    - component: textarea.go:66
    - demo: demo.go:20
  - tests:
    - textarea_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/textarea-invalid.tsx
  - implementation:
    - component: textarea.go:66
    - demo: demo.go:20
  - tests:
    - textarea_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/textarea-rtl.tsx
  - implementation:
    - component: textarea.go:66
    - demo: demo.go:20
  - tests:
    - textarea_test.go:12

---

## Code Structure & Entry Points
- `textarea.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `textarea_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
