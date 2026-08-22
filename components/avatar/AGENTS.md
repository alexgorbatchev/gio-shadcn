# Avatar Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/avatar](https://ui.shadcn.com/docs/components/avatar)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `avatar` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Text Initials Avatar**
  - implementation:
    - avatar.go:13
  - tests:
    - avatar_test.go:31
- [x] **Image Avatar Placeholder**
  - implementation:
    - avatar.go:7
  - tests:
    - avatar_test.go:4
- [x] **Online Status Badge Indicator**
  - implementation:
    - avatar.go:50
  - tests:
    - avatar_test.go:13

### Capabilities & Features
- [x] **Circular Clip Ellipse**
  - implementation:
    - avatar.go:50
  - tests:
    - avatar_test.go:13
- [x] **Custom Sizes (32px, 40px, 56px)**
  - implementation:
    - avatar.go:50
  - tests:
    - avatar_test.go:13
- [x] **Status Dot Indicator**
  - implementation:
    - avatar.go:50
  - tests:
    - avatar_test.go:13

---

## Code Structure & Entry Points
- `avatar.go`: Primary component widget layout and state logic.
- `avatar_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
