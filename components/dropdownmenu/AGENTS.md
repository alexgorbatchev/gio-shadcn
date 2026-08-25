# Dropdown Menu Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/dropdown-menu](https://ui.shadcn.com/docs/components/dropdown-menu)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `dropdownmenu` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Action Dropdown Menu**
  - implementation:
    - dropdownmenu.go:35
  - tests:
    - dropdownmenu_test.go:12

### Capabilities & Features
- [x] **Menu Item List**
  - implementation:
    - dropdownmenu.go:21
  - tests:
    - dropdownmenu_test.go:12
- [x] **Trigger Button Toggle**
  - implementation:
    - dropdownmenu.go:88
  - tests:
    - dropdownmenu_test.go:36
- [x] **Keyboard Shortcut Badges**
  - implementation:
    - dropdownmenu.go:193
  - tests:
    - dropdownmenu_test.go:12
- [x] **Open / Close State Toggle**
  - implementation:
    - dropdownmenu.go:82
  - tests:
    - dropdownmenu_test.go:36
- [x] **Hover Highlight State**
  - implementation:
    - dropdownmenu.go:175
  - tests:
    - dropdownmenu_test.go:12

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-demo.tsx
  - implementation:
    - component: dropdownmenu.go:88
    - demo: demo.go:19
  - tests:
    - dropdownmenu_test.go:36
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-basic.tsx
  - implementation:
    - component: dropdownmenu.go:35
    - demo: demo.go:32
  - tests:
    - dropdownmenu_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-checkboxes.tsx
  - implementation:
    - component: dropdownmenu.go:35
    - demo: demo.go:42
  - tests:
    - dropdownmenu_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-shortcuts.tsx
  - implementation:
    - component: dropdownmenu.go:193
    - demo: demo.go:50
  - tests:
    - dropdownmenu_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-avatar.tsx
  - implementation:
    - component: dropdownmenu.go:35
    - demo: demo.go:19
  - tests:
    - dropdownmenu_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-destructive.tsx
  - implementation:
    - component: dropdownmenu.go:35
    - demo: demo.go:19
  - tests:
    - dropdownmenu_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-rtl.tsx
  - implementation:
    - component: dropdownmenu.go:35
    - demo: demo.go:19
  - tests:
    - dropdownmenu_test.go:12

---

## Code Structure & Entry Points
- `dropdownmenu.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `dropdownmenu_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
