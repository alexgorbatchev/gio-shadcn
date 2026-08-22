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
- [x] **Linear Progress Bar**
  - implementation:
    - progress.go:34
  - tests:
    - progress_test.go:13

### Capabilities & Features
- [x] **Background Track**
  - implementation:
    - progress.go:83
  - tests:
    - progress_test.go:33
- [x] **Filled Progress Portion**
  - implementation:
    - progress.go:89
  - tests:
    - progress_test.go:25
- [x] **Clamped Radius Half-Height**
  - implementation:
    - progress.go:76
  - tests:
    - progress_test.go:33

---

### Demos
- [x] **Progress Bar (0%)**
  - implementation:
    - progress.go:3
  - tests:
    - progress_test.go:2
- [x] **Progress Bar (35%)**
  - implementation:
    - progress.go:4
  - tests:
    - progress_test.go:1
- [x] **Progress Bar (65%)**
  - implementation:
    - progress.go:2
  - tests:
    - progress_test.go:2
- [x] **Progress Bar (100%)**
  - implementation:
    - progress.go:3
  - tests:
    - progress_test.go:1

## Code Structure & Entry Points
- `progress.go`: Primary component widget layout and state logic.
- `progress_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
