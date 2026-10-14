# Vendor Due Diligence

Separate vendor claims from evidence, assumptions, dependencies, lock-in, operational impact and questions to test.

## Authority

The workflow uses a model for analysis/proposal. **Human review remains authoritative.**

## Files

- `manifest.json` — versioned machine-readable workflow contract.
- `SKILL.md` — model instructions.
- `example-input.md` — fictional source material.
- `example-output.md` — one conforming worked output.

## Use

```bash
go run ./cmd/workbench render -workflow vendor-due-diligence -input workflows/vendor-due-diligence/example-input.md
```

Validate an output:

```bash
go run ./cmd/workbench check -workflow vendor-due-diligence -output workflows/vendor-due-diligence/example-output.md
```
