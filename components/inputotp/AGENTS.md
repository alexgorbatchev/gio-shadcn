# Input OTP Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/input-otp](https://ui.shadcn.com/docs/components/input-otp)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `inputotp` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **4-Digit PIN Code**
  - implementation:
    - inputotp.go:42
  - tests:
    - inputotp_test.go:13
- [x] **6-Digit PIN Code**
  - implementation:
    - inputotp.go:44
  - tests:
    - inputotp_test.go:22

### Capabilities & Features
- [x] **Digit Box Layout**
  - implementation:
    - inputotp.go:109
  - tests:
    - inputotp_test.go:31
- [x] **Active Cell Highlight**
  - implementation:
    - inputotp.go:124
  - tests:
    - inputotp_test.go:46

---



### Demos
- [x] **1. 6-Digit OTP PIN Input**
  - implementation:
    - component: inputotp.go:50
    - demo: demo.md:13
  - tests:
    - inputotp_test.go:15

## Code Structure & Entry Points
- `inputotp.go`: Primary component implementation.
- `inputotp_test.go`: Unit test suite.
- `demo.md`: Component interactive demo snippets.
- `AGENTS.md`: Component specification.