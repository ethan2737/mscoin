## 1. Route And Flow Inventory

- [x] 1.1 Inventory the active Vue 3 route tree and identify the primary business flow covered by each top-level route group.
- [x] 1.2 Map shared shell behaviors in `App.vue`, including header navigation, language switching, dropdowns, auth checks, and mobile menu behavior.
- [x] 1.3 Document which homepage, auth, notice/help, exchange, OTC, swap, and user-center regions are expected to be static versus dynamic.

## 2. Data And Dependency Tracing

- [x] 2.1 Trace the HTTP dependencies for each audited flow, including request path, auth requirement, dev proxy mapping, and expected backend service/port.
- [x] 2.2 Trace websocket or socket.io dependencies for realtime regions and record the expected runtime contract.
- [x] 2.3 Flag flows that currently rely on fallback data, placeholder content, or migration stubs instead of real backend data.

## 3. Verification And Evidence Gathering

- [x] 3.1 Verify the highest-risk flows in-browser, including homepage interactions, language/menu controls, login, register, and exchange entry.
- [x] 3.2 Reconcile observed failures with the corresponding code paths to classify them as frontend regressions, dependency gaps, or unresolved contract issues.
- [x] 3.3 Capture reproducible evidence for each confirmed finding, including affected route/component, trigger steps, and blocking dependency if applicable.

## 4. Reporting And Remediation Planning

- [x] 4.1 Produce a frontend business-flow and data-flow audit report under the OpenSpec change directory so it can be reused during later fix changes.
- [x] 4.2 Summarize findings by severity and by defect class, including broken rendering, broken interaction, broken data contract, and blocked-by-backend cases.
- [x] 4.3 Derive the recommended remediation sequence and propose follow-up OpenSpec changes or implementation steps from the report.
