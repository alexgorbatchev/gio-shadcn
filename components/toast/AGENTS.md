# Toast Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/toast](https://ui.shadcn.com/docs/components/toast)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `toast` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Info Toast Banner**
  - implementation:
    - toast.go:56
  - tests:
    - toast_test.go:12
- [x] **Destructive Toast Banner**
  - implementation:
    - toast.go:69
  - tests:
    - toast_test.go:27

### Capabilities & Features
- [x] **Title Header**
  - implementation:
    - toast.go:103
  - tests:
    - toast_test.go:12
- [x] **Description Body**
  - implementation:
    - toast.go:118
  - tests:
    - toast_test.go:12
- [x] **Lucide Icon Integration**
  - implementation:
    - toast.go:96
  - tests:
    - toast_test.go:12

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/sonner-demo.tsx
  - implementation:
    - component: toast.go:56
    - demo: demo.go:20
  - tests:
    - toast_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/sonner-description.tsx
  - implementation:
    - component: toast.go:56
    - demo: demo.go:20
  - tests:
    - toast_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/sonner-types.tsx
  - implementation:
    - component: toast.go:69
    - demo: demo.go:27
  - tests:
    - toast_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/sonner-position.tsx
  - implementation:
    - component: toast.go:56
    - demo: demo.go:20
  - tests:
    - toast_test.go:12

---

## Code Structure & Entry Points
- `toast.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `toast_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
