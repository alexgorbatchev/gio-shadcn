# Dialog Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/dialog](https://ui.shadcn.com/docs/components/dialog)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `dialog` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Modal Dialog Window**
  - implementation:
    - dialog.go:50
  - tests:
    - dialog_test.go:13

### Capabilities & Features
- [x] **Dark Backdrop Overlay & Outside Click-To-Close**
  - implementation:
    - dialog.go:108
  - tests:
    - dialog_test.go:13
- [x] **Header Title & Description**
  - implementation:
    - dialog.go:142
  - tests:
    - dialog_test.go:24
- [x] **Confirm & Cancel Actions**
  - implementation:
    - dialog.go:167
  - tests:
    - dialog_test.go:37
- [x] **Custom Content Body Widget**
  - implementation:
    - dialog.go:158
  - tests:
    - dialog_test.go:49

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dialog-demo.tsx
  - implementation:
    - component: dialog.go:50
    - demo: demo.go:28
  - tests:
    - dialog_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-dialog-demo.tsx
  - implementation:
    - component: dialog.go:50
    - demo: demo.go:54
  - tests:
    - dialog_test.go:24
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/alert-dialog-destructive.tsx
  - implementation:
    - component: dialog.go:50
    - demo: demo.go:73
  - tests:
    - dialog_test.go:37
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dialog-close-button.tsx
  - implementation:
    - component: dialog.go:50
    - demo: demo.go:92
  - tests:
    - dialog_test.go:49
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dialog-scrollable-content.tsx
  - implementation:
    - component: dialog.go:50
    - demo: demo.go:114
  - tests:
    - dialog_test.go:49

---

## Code Structure & Entry Points
- `dialog.go`: Primary component widget layout and state logic.
- `dialog_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Exported interactive demo widget (`Demo`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
