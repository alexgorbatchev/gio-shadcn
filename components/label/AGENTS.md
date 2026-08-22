# Label & Typography Component Specification (`gio-shadcn`)

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
- [x] **Heading 1 (H1)**
  - implementation:
    - label.go:167
  - tests:
    - label_test.go:13
- [x] **Heading 2 (H2)**
  - implementation:
    - label.go:168
  - tests:
    - label_test.go:20
- [x] **Heading 3 (H3)**
  - implementation:
    - label.go:169
  - tests:
    - label_test.go:27
- [x] **Heading 4 (H4)**
  - implementation:
    - label.go:170
  - tests:
    - label_test.go:34
- [x] **Body Paragraph (P)**
  - implementation:
    - label.go:171
  - tests:
    - label_test.go:41
- [x] **Muted & Small Text**
  - implementation:
    - label.go:172
  - tests:
    - label_test.go:48

### Capabilities & Features
- [x] **Typography Font Scale**
  - implementation:
    - label.go:197
  - tests:
    - label_test.go:56
- [x] **Theme Color Integration**
  - implementation:
    - label.go:214
  - tests:
    - label_test.go:69

---

## Code Structure & Entry Points
- `label.go`: Primary component widget layout and state logic.
- `label_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
