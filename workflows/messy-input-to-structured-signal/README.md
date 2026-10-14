# Messy Input to Structured Signal

Turn mixed, uneven enterprise material into a traceable structure without pretending the source is cleaner than it is.

## Authority

The workflow uses a model for analysis/proposal. **Human review remains authoritative.**

## Files

- `manifest.json` — versioned machine-readable workflow contract.
- `SKILL.md` — model instructions.
- `example-input.md` — fictional source material.
- `example-output.md` — one conforming worked output.

## Use

```bash
go run ./cmd/workbench render -workflow messy-input-to-structured-signal -input workflows/messy-input-to-structured-signal/example-input.md
```

Validate an output:

```bash
go run ./cmd/workbench check -workflow messy-input-to-structured-signal -output workflows/messy-input-to-structured-signal/example-output.md
```
