# Unknowns First

Force analysis to distinguish known facts from inference, assumptions, contradictions and missing evidence.

## Authority

The workflow uses a model for analysis/proposal. **Human review remains authoritative.**

## Files

- `manifest.json` — versioned machine-readable workflow contract.
- `SKILL.md` — model instructions.
- `example-input.md` — fictional source material.
- `example-output.md` — one conforming worked output.

## Use

```bash
go run ./cmd/workbench render -workflow unknowns-first -input workflows/unknowns-first/example-input.md
```

Validate an output:

```bash
go run ./cmd/workbench check -workflow unknowns-first -output workflows/unknowns-first/example-output.md
```
