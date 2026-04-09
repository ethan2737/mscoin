## ADDED Requirements

### Requirement: Exchange and swap pages mount safely
The exchange and swap route entries SHALL mount without fatal runtime exceptions caused by null DOM assumptions, missing required props, or legacy runtime contract reads.

#### Scenario: Exchange route opens
- **WHEN** the user navigates to `/exchange/:pair`
- **THEN** the page renders its primary trading surface without fatal mount or render exceptions

#### Scenario: Swap route opens
- **WHEN** the user navigates to `/swapindex/:pair`
- **THEN** the page renders its primary trading surface without fatal mount or render exceptions

### Requirement: Realtime market wiring uses the supported runtime path
Exchange and swap realtime feeds SHALL use the supported socket/runtime path for the active Vue 3 baseline rather than hardcoded legacy localhost paths.

#### Scenario: Exchange realtime connection initializes
- **WHEN** the exchange page starts its market feed
- **THEN** it connects using the supported runtime path and does not hardcode `io('http://localhost')`
