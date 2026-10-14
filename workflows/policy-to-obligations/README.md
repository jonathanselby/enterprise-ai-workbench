# Policy to Obligations

Extract normative statements, owners, triggers, evidence requirements, exceptions and ambiguity from policy text.

## Authority

The workflow uses a model for analysis/proposal. **Human review remains authoritative.**

## Files

- `manifest.json` — versioned machine-readable workflow contract.
- `SKILL.md` — model instructions.
- `example-input.md` — fictional source material.
- `example-output.md` — one conforming worked output.

## Use

```bash
go run ./cmd/workbench render -workflow policy-to-obligations -input workflows/policy-to-obligations/example-input.md
```

Validate an output:

```bash
go run ./cmd/workbench check -workflow policy-to-obligations -output workflows/policy-to-obligations/example-output.md
```
