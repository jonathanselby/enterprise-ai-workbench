# Research to Decision

Turn a broad research question into evidence, contradictions, unknowns, options and a defensible decision brief.

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
go run ./cmd/workbench render -workflow research-to-decision -input workflows/research-to-decision/example-input.md
```

Validate an output:

```bash
go run ./cmd/workbench check -workflow research-to-decision -output workflows/research-to-decision/example-output.md
```
