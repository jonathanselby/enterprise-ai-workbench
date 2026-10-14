# Documents to Risks

Extract candidate risks from documents and express them as cause, event and consequence for human review.

## Authority

The workflow uses a model for analysis/proposal. **Human review remains authoritative.**

## Files

- `manifest.json` — versioned machine-readable workflow contract.
- `SKILL.md` — model instructions.
- `example-input.md` — fictional source material.
- `example-output.md` — one conforming worked output.

## Use

```bash
go run ./cmd/workbench render -workflow documents-to-risks -input workflows/documents-to-risks/example-input.md
```

Validate an output:

```bash
go run ./cmd/workbench check -workflow documents-to-risks -output workflows/documents-to-risks/example-output.md
```
