# Collapsible Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/collapsible](https://ui.shadcn.com/docs/components/collapsible)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `collapsible` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Expanded State**
  - implementation:
    - collapsible.go:40
  - tests:
    - collapsible_test.go:13
- [x] **Collapsed State**
  - implementation:
    - collapsible.go:94
  - tests:
    - collapsible_test.go:49

### Capabilities & Features
- [x] **Trigger Button Header**
  - implementation:
    - collapsible.go:80
  - tests:
    - collapsible_test.go:23
- [x] **Content Body Visibility Toggle**
  - implementation:
    - collapsible.go:94
  - tests:
    - collapsible_test.go:35

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/collapsible-basic.tsx
  - implementation:
    - component: collapsible.go:40
    - demo: demo.go:24
  - tests:
    - collapsible_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/collapsible-demo.tsx
  - implementation:
    - component: collapsible.go:40
    - demo: demo.go:32
  - tests:
    - collapsible_test.go:23
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/collapsible-file-tree.tsx
  - implementation:
    - component: collapsible.go:94
    - demo: demo.go:40
  - tests:
    - collapsible_test.go:35
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/collapsible-settings.tsx
  - implementation:
    - component: collapsible.go:40
    - demo: demo.go:68
  - tests:
    - collapsible_test.go:49

---

## Code Structure & Entry Points
- `collapsible.go`: Primary component widget layout and state logic.
- `collapsible_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Exported interactive demo widget (`Demo`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
