# Progress Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/progress](https://ui.shadcn.com/docs/components/progress)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `progress` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Linear Progress Bar Variant**
  - implementation:
    - progress.go:37
  - tests:
    - progress_test.go:12

### Capabilities & Features
- [x] **Background Track Drawing**
  - implementation:
    - progress.go:82
  - tests:
    - progress_test.go:12
- [x] **Filled Progress Portion (0% to 100%)**
  - implementation:
    - progress.go:88
  - tests:
    - progress_test.go:22
- [x] **Clamped Radius Half-Height Safety**
  - implementation:
    - progress.go:76
  - tests:
    - progress_test.go:12

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/progress-demo.tsx
  - implementation:
    - component: progress.go:54
    - demo: demo.go:23
  - tests:
    - progress_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/progress-controlled.tsx
  - implementation:
    - component: progress.go:54
    - demo: demo.go:24
  - tests:
    - progress_test.go:22
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/progress-label.tsx
  - implementation:
    - component: progress.go:54
    - demo: demo.go:23
  - tests:
    - progress_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/progress-rtl.tsx
  - implementation:
    - component: progress.go:54
    - demo: demo.go:23
  - tests:
    - progress_test.go:12

---

## Code Structure & Entry Points
- `progress.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `progress_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
