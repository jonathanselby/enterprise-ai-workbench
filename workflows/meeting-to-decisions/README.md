# Meeting to Decisions

Separate what a meeting actually decided from actions, assumptions, disagreements, deferrals and unknowns.

## Authority

The workflow uses a model for analysis/proposal. **Human review remains authoritative.**

## Files

- `manifest.json` — versioned machine-readable workflow contract.
- `SKILL.md` — model instructions.
- `example-input.md` — fictional source material.
- `example-output.md` — one conforming worked output.

## Use

From the repository root:

```bash
go run ./cmd/workbench render -workflow meeting-to-decisions -input workflows/meeting-to-decisions/example-input.md
```

Validate an output:

```bash
go run ./cmd/workbench check -workflow meeting-to-decisions -output workflows/meeting-to-decisions/example-output.md
```
