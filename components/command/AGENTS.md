# Command Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/command](https://ui.shadcn.com/docs/components/command)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `command` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Command Palette Search Box**
  - implementation:
    - command.go:61
  - tests:
    - command_test.go:13

### Capabilities & Features
- [x] **Search Input Filter**
  - implementation:
    - command.go:94
  - tests:
    - command_test.go:13
- [x] **Command Item List with Groups**
  - implementation:
    - command.go:121
  - tests:
    - command_test.go:37
- [x] **Shortcut Badges & Icons**
  - implementation:
    - command.go:189
  - tests:
    - command_test.go:25

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/command-basic.tsx
  - implementation:
    - component: command.go:61
    - demo: demo.go:22
  - tests:
    - command_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/command-demo.tsx
  - implementation:
    - component: command.go:61
    - demo: demo.go:32
  - tests:
    - command_test.go:25
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/command-dialog.tsx
  - implementation:
    - component: command.go:61
    - demo: demo.go:45
  - tests:
    - command_test.go:37
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/command-groups.tsx
  - implementation:
    - component: command.go:121
    - demo: demo.go:56
  - tests:
    - command_test.go:37
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/command-shortcuts.tsx
  - implementation:
    - component: command.go:189
    - demo: demo.go:67
  - tests:
    - command_test.go:25

---

## Code Structure & Entry Points
- `command.go`: Primary component widget layout and state logic.
- `command_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Exported interactive demo widget (`Demo`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
