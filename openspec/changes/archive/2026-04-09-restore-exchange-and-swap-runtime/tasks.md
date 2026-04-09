## 1. Runtime Stabilization

- [x] 1.1 Audit exchange and swap boot paths for null access, missing prop assumptions, and broken pair initialization.
- [x] 1.2 Refactor exchange and swap runtime wiring to use the supported Vue 3 runtime/socket contract.

## 2. Surface Recovery

- [x] 2.1 Restore safe rendering of the core exchange surface, including market list and key summary/depth regions.
- [x] 2.2 Restore safe rendering of the core swap surface and verify shared market/contract components still mount.

## 3. Verification

- [x] 3.1 Verify `/exchange/:pair` renders without fatal runtime exceptions.
- [x] 3.2 Verify `/swapindex/:pair` renders without fatal runtime exceptions.
