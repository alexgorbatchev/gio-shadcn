# Carousel Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/carousel](https://ui.shadcn.com/docs/components/carousel)  
**Official shadcn Source Spec (.mdx):** [https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/carousel.mdx](https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/carousel.mdx)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `carousel` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Horizontal Slide Carousel**
  - implementation:
    - carousel.go:34
  - tests:
    - carousel_test.go:13

### Capabilities & Features
- [x] **Active Slide Index Tracking**
  - implementation:
    - carousel.go:61
  - tests:
    - carousel_test.go:25
- [x] **Next / Previous Controls**
  - implementation:
    - carousel.go:78
  - tests:
    - carousel_test.go:37
- [x] **Slide Item Container**
  - implementation:
    - carousel.go:70
  - tests:
    - carousel_test.go:56

---



### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/carousel-api.tsx
  - implementation:
    - component: carousel.go:34
    - demo: demo.go:13
  - tests:
    - carousel_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/carousel-demo.tsx
  - implementation:
    - component: carousel.go:34
    - demo: demo.go:13
  - tests:
    - carousel_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/carousel-multiple.tsx
  - implementation:
    - component: carousel.go:34
    - demo: demo.go:13
  - tests:
    - carousel_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/carousel-orientation.tsx
  - implementation:
    - component: carousel.go:34
    - demo: demo.go:13
  - tests:
    - carousel_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/carousel-plugin.tsx
  - implementation:
    - component: carousel.go:34
    - demo: demo.go:13
  - tests:
    - carousel_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/carousel-size.tsx
  - implementation:
    - component: carousel.go:34
    - demo: demo.go:13
  - tests:
    - carousel_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/carousel-spacing.tsx
  - implementation:
    - component: carousel.go:34
    - demo: demo.go:13
  - tests:
    - carousel_test.go:13

## Code Structure & Entry Points
- `carousel.go`: Primary component implementation.
- `carousel_test.go`: Unit test suite.
- `demo.md`: Component interactive demo snippets.
- `AGENTS.md`: Component specification.