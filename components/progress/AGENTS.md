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
- [x] **0% Fill Level**
  - implementation:
    - progress.go:34
  - tests:
    - progress_test.go:13
- [x] **65% Fill Level**
  - implementation:
    - progress.go:88
  - tests:
    - progress_test.go:20
- [x] **100% Fill Level**
  - implementation:
    - progress.go:88
  - tests:
    - progress_test.go:27

### Capabilities & Features
- [x] **Clamped Corner Radius (`heightPx / 2`)**
  - implementation:
    - progress.go:78
  - tests:
    - progress_test.go:34
- [x] **Background Track & Filled Bar Layout**
  - implementation:
    - progress.go:85
  - tests:
    - progress_test.go:48
- [x] **GPU Color State Reset**
  - implementation:
    - progress.go:96
  - tests:
    - progress_test.go:34

### Demos
- [x] **Progress Bar (0%)**
  - implementation:
    - progress.go:34
  - tests:
    - progress_test.go:13
- [x] **Progress Bar (35%)**
  - implementation:
    - progress.go:88
  - tests:
    - progress_test.go:20
- [x] **Progress Bar (65%)**
  - implementation:
    - progress.go:88
  - tests:
    - progress_test.go:20
- [x] **Progress Bar (100%)**
  - implementation:
    - progress.go:88
  - tests:
    - progress_test.go:27

---

## Code Structure & Entry Points
- `progress.go`: Primary component widget layout and state logic.
- `progress_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
