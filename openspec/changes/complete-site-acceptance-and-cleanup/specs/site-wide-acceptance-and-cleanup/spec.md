## ADDED Requirements

### Requirement: Module-by-Module Full-Site Acceptance Matrix

The project MUST maintain and execute a site-wide acceptance matrix for the active local development baseline, covering each active business module individually.

#### Scenario: Acceptance matrix covers active modules

- **GIVEN** the recovered Vue 3 local development baseline is active
- **WHEN** the full-site acceptance change is executed
- **THEN** the change MUST define and run acceptance checks for each active module, including spot trading, contract, OTC, crowdfunding, content pages, and user-center-adjacent active flows
- **AND** each module MUST be classified in the final report as `passed`, `passed_with_local_limitations`, `blocked`, or `deferred`

### Requirement: Announcement Detail Must Settle Without Prolonged Broken Loading State

The active announcement detail page MUST resolve to stable content presentation in local development without getting stuck in a misleading loading experience.

#### Scenario: Announcement detail hydrates into content

- **GIVEN** announcement content APIs are available in the local backend
- **WHEN** a user opens an active announcement detail route
- **THEN** the page MUST transition from loading to content presentation without remaining stuck in an indefinite or visibly broken loading state
- **AND** any short-lived placeholder behavior MUST not obscure whether the page succeeded or failed

### Requirement: Local Acceptance Must Not Be Polluted by Avoidable Third-Party Chat Noise

The local development baseline MUST isolate or gracefully handle the third-party customer-support script so acceptance results are not obscured by avoidable console noise or page disruption.

#### Scenario: Local acceptance runs without third-party chat pollution

- **GIVEN** the site is running in the local development environment
- **WHEN** a user or tester performs the acceptance walkthrough
- **THEN** the third-party chat integration MUST either load cleanly or be gracefully suppressed for local development
- **AND** its failure MUST NOT create repeated blocking console noise that interferes with diagnosing business behavior

### Requirement: Active Vue 3 Pages Must Be Free Of Material Migration Residue

The active Vue 3 pages MUST not retain visible style, encoding, or legacy artifacts that materially degrade acceptance of the recovered product.

#### Scenario: Active page residue is cleaned

- **GIVEN** an active Vue 3 page is part of the acceptance matrix
- **WHEN** that page is reviewed during the full-site acceptance pass
- **THEN** visible encoding corruption, broken styling, or obsolete legacy remnants that affect the user experience MUST be repaired or explicitly deferred in the final report

### Requirement: Remaining Business Domains Must Reach Explicit Local Closure Status

The remaining business domains beyond the core recovered spot flow MUST each reach an explicit local closure status backed by browser-led verification.

#### Scenario: Contract, OTC, and crowdfunding are classified by evidence

- **GIVEN** contract, OTC, crowdfunding, and other remaining active business routes are present in the local baseline
- **WHEN** the full-site acceptance pass is completed
- **THEN** each domain MUST have browser-led validation evidence recorded in the change report
- **AND** any domain that cannot be fully accepted MUST identify its concrete blocker rather than remaining unclassified
