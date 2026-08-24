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
    - dropdownmenu_test.go:13

### Capabilities & Features
- [x] **Menu Item List**
  - implementation:
    - dropdownmenu.go:21
  - tests:
    - dropdownmenu_test.go:34
- [x] **Keyboard Shortcut Badges**
  - implementation:
    - dropdownmenu.go:136
  - tests:
    - dropdownmenu_test.go:41
- [x] **Open / Close State Toggle**
  - implementation:
    - dropdownmenu.go:70
  - tests:
    - dropdownmenu_test.go:48
- [x] **Hover Highlight State**
  - implementation:
    - dropdownmenu.go:147
  - tests:
    - dropdownmenu_test.go:48

---


### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-avatar.tsx
  - implementation:
    - component: dropdownmenu.go:61
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-basic.tsx
  - implementation:
    - component: dropdownmenu.go:61
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-checkboxes-icons.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-checkboxes.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-complex.tsx
  - implementation:
    - component: dropdownmenu.go:61
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-demo.tsx
  - implementation:
    - component: dropdownmenu.go:61
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-destructive.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-icons.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-radio-group.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-shortcuts.tsx
  - implementation:
    - component: dropdownmenu.go:147
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:34
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/dropdown-menu-submenu.tsx
  - implementation:
    - component: dropdownmenu.go:61
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:41
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-basic.tsx
  - implementation:
    - component: dropdownmenu.go:61
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-checkboxes.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-demo.tsx
  - implementation:
    - component: dropdownmenu.go:61
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-destructive.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-groups.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-icons.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-radio.tsx
  - implementation:
    - component: dropdownmenu.go:82
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-shortcuts.tsx
  - implementation:
    - component: dropdownmenu.go:147
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:34
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-sides.tsx
  - implementation:
    - component: dropdownmenu.go:61
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:41
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/context-menu-submenu.tsx
  - implementation:
    - component: dropdownmenu.go:61
    - demo: demo.go:14
  - tests:
    - dropdownmenu_test.go:41

## Code Structure & Entry Points
- `dropdownmenu.go`: Primary component widget layout and state logic.
- `dropdownmenu_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.md`: Component interactive demo snippets and layout specs (`components/dropdownmenu/demo.md`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
