# Label Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/label](https://ui.shadcn.com/docs/components/label)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `label` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **H1**
  - implementation:
    - label.go:69
  - tests:
    - label_test.go:14
- [x] **H2**
  - implementation:
    - label.go:70
  - tests:
    - label_test.go:15
- [x] **H3**
  - implementation:
    - label.go:68
  - tests:
    - label_test.go:16
- [x] **H4**
  - implementation:
    - label.go:69
  - tests:
    - label_test.go:17
- [x] **Body Paragraph**
  - implementation:
    - label.go:134
  - tests:
    - label_test.go:17
- [x] **Lead**
  - implementation:
    - label.go:184
  - tests:
    - label_test.go:14
- [x] **Large**
  - implementation:
    - label.go:21
  - tests:
    - label_test.go:15
- [x] **Muted**
  - implementation:
    - label.go:188
  - tests:
    - label_test.go:16
- [x] **Small**
  - implementation:
    - label.go:20
  - tests:
    - label_test.go:17

### Capabilities & Features
- [x] **Typography Font Scale**
  - implementation:
    - label.go:3
  - tests:
    - label_test.go:13
- [x] **Font Weight Configuration**
  - implementation:
    - label.go:105
  - tests:
    - label_test.go:14
- [x] **Theme Color Integration**
  - implementation:
    - label.go:11
  - tests:
    - label_test.go:12

---

## Code Structure & Entry Points
- `label.go`: Primary component widget layout and state logic.
- `label_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
