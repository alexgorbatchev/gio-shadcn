# Card Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/card](https://ui.shadcn.com/docs/components/card)  
**Official shadcn Source Spec (.mdx):** [https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/card.mdx](https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/card.mdx)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `card` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Default Card**
  - implementation:
    - card.go:68
  - tests:
    - card_test.go:13
- [x] **Header + Content + Footer Layout**
  - implementation:
    - card.go:74
  - tests:
    - card_test.go:20

### Capabilities & Features
- [x] **Card Title**
  - implementation:
    - card.go:217
  - tests:
    - card_test.go:42
- [x] **Card Description**
  - implementation:
    - card.go:238
  - tests:
    - card_test.go:49
- [x] **Card Content Area**
  - implementation:
    - card.go:256
  - tests:
    - card_test.go:56
- [x] **Card Footer**
  - implementation:
    - card.go:272
  - tests:
    - card_test.go:63
- [x] **Border Stroke & Rounded Radius**
  - implementation:
    - card.go:136
  - tests:
    - card_test.go:70

---


### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/card-demo.tsx
  - implementation:
    - component: card.go:68
    - demo: demo.go:13
  - tests:
    - card_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/card-edge-to-edge.tsx
  - implementation:
    - component: card.go:68
    - demo: demo.go:13
  - tests:
    - card_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/card-image.tsx
  - implementation:
    - component: card.go:68
    - demo: demo.go:13
  - tests:
    - card_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/card-rtl.tsx
  - implementation:
    - component: card.go:68
    - demo: demo.go:13
  - tests:
    - card_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/card-small.tsx
  - implementation:
    - component: card.go:68
    - demo: demo.go:13
  - tests:
    - card_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/card-spacing.tsx
  - implementation:
    - component: card.go:68
    - demo: demo.go:13
  - tests:
    - card_test.go:13

## Code Structure & Entry Points
- `card.go`: Primary component widget layout and state logic.
- `card_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.md`: Component interactive demo snippets and layout specs (`components/card/demo.md`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
