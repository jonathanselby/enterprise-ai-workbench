# Documents to Risks

## Candidate Risks
| Cause | Event | Consequence |
|---|---|---|
| The first release depends on a single identity connector and customer failover has not been exercised in a production-like environment. | The identity connector fails or cannot be recovered within the operational window. | Staff may be unable to use the rostering service and sites may depend on manual entry; a prolonged incident near payroll cut-off could affect payroll processing. |

## Evidence
- Single identity connector in the first release.
- Vendor staging recovery test exists.
- Customer has not run a production-like failover exercise.
- Manual roster entry is documented for up to four hours.
- Payroll cut-off is 3pm Friday.

## Existing controls
- Vendor-tested recovery in staging.
- Manual roster entry at each site for up to four hours.

## Unknowns
- Recovery-time performance in the customer environment.
- What happens if manual operation exceeds four hours.
- Whether payroll has an independent recovery or late-submission process.

## Review notes
A human risk owner should determine whether this becomes a formal risk, its rating, owner and required treatment.
