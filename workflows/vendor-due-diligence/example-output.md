# Vendor Due Diligence

## Claims
- Platform is "enterprise grade".
- Workflow definitions can be exported as JSON.
- Higher API limits are available by arrangement.

## Evidence
- Proposal states retention options of 30, 90 or 365 days.
- Private networking is limited to the Enterprise tier.
- Stated API limit is 600 calls per minute.
- A SOC 2 Type II report was supplied and is 14 months old.

## Assumptions
- JSON export may support migration, but the proposal does not establish semantic portability to another workflow engine.

## Dependencies
- Enterprise tier is required if private networking is a requirement.
- Workloads above 600 API calls per minute depend on a separately agreed limit.

## Lock-in and exit
- Workflow definitions are exportable, but execution-history portability and migration assistance are not established.
- Exit assistance is not priced.

## Operational impact
Retention tier and API limits could affect auditability and throughput. Private networking changes the required commercial tier.

## Questions to test
- What exactly is contained in the JSON export, including secrets, connections and custom logic?
- Can execution history be exported in bulk?
- What tested rate limit is available and under what commercial terms?
- Is a current assurance report available?
- What exit assistance can be contracted and priced?

## Decision notes
No purchase decision is supported by this extract alone. The open portability, assurance, throughput and exit questions should be resolved against explicit decision criteria.
