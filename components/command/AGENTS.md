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
    - command.go:35
  - tests:
    - command_test.go:13

### Capabilities & Features
- [x] **Search Input Filter (`searchEditor`)**
  - implementation:
    - command.go:88
  - tests:
    - command_test.go:24
- [x] **Filtered Results List**
  - implementation:
    - command.go:102
  - tests:
    - command_test.go:42
- [x] **Keyboard Shortcut Badges**
  - implementation:
    - command.go:151
  - tests:
    - command_test.go:52
- [x] **Hover State Highlight**
  - implementation:
    - command.go:161
  - tests:
    - command_test.go:24

### Demos (Official shadcn Demos)
- [x] **1. Command Dialog**
  - implementation:
    - command.go:76
  - tests:
    - command_test.go:13
- [x] **2. Search Filter List**
  - implementation:
    - command.go:102
  - tests:
    - command_test.go:24
- [x] **3. Keyboard Shortcuts**
  - implementation:
    - command.go:151
  - tests:
    - command_test.go:52
- [x] **4. Action Items Selection**
  - implementation:
    - command.go:106
  - tests:
    - command_test.go:42

---

## Code Structure & Entry Points
- `command.go`: Primary component widget layout and state logic.
- `command_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
