# Breadcrumb Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/breadcrumb](https://ui.shadcn.com/docs/components/breadcrumb)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `breadcrumb` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Standard Path Trail**
  - implementation:
    - breadcrumb.go:21
  - tests:
    - breadcrumb_test.go:12
- [x] **Active Page Link**
  - implementation:
    - breadcrumb.go:73
  - tests:
    - breadcrumb_test.go:12

### Capabilities & Features
- [x] **Horizontal Flex Layout**
  - implementation:
    - breadcrumb.go:94
  - tests:
    - breadcrumb_test.go:12
- [x] **Slash / Custom Separators**
  - implementation:
    - breadcrumb.go:88
  - tests:
    - breadcrumb_test.go:26
- [x] **Interactive Item Pointer Click**
  - implementation:
    - breadcrumb.go:68
  - tests:
    - breadcrumb_test.go:54

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/breadcrumb-basic.tsx
  - implementation:
    - component: breadcrumb.go:21
    - demo: demo.go:19
  - tests:
    - breadcrumb_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/breadcrumb-demo.tsx
  - implementation:
    - component: breadcrumb.go:21
    - demo: demo.go:19
  - tests:
    - breadcrumb_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/breadcrumb-dropdown.tsx
  - implementation:
    - component: breadcrumb.go:21
    - demo: demo.go:44
  - tests:
    - breadcrumb_test.go:40
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/breadcrumb-ellipsis.tsx
  - implementation:
    - component: breadcrumb.go:21
    - demo: demo.go:36
  - tests:
    - breadcrumb_test.go:40
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/breadcrumb-link.tsx
  - implementation:
    - component: breadcrumb.go:68
    - demo: demo.go:51
  - tests:
    - breadcrumb_test.go:54
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/breadcrumb-rtl.tsx
  - implementation:
    - component: breadcrumb.go:21
    - demo: demo.go:19
  - tests:
    - breadcrumb_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/breadcrumb-separator.tsx
  - implementation:
    - component: breadcrumb.go:88
    - demo: demo.go:27
  - tests:
    - breadcrumb_test.go:26

---

## Code Structure & Entry Points
- `breadcrumb.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `breadcrumb_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
