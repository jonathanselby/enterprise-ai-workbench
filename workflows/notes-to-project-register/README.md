# Notes to Project Register

Turn unstructured project material into a deliberately small operational register without inventing status or ownership.

## Authority

The workflow uses a model for analysis/proposal. **Human review remains authoritative.**

## Files

- `manifest.json` — versioned machine-readable workflow contract.
- `SKILL.md` — model instructions.
- `example-input.md` — fictional source material.
- `example-output.md` — one conforming worked output.

## Use

```bash
go run ./cmd/workbench render -workflow notes-to-project-register -input workflows/notes-to-project-register/example-input.md
```

Validate an output:

```bash
go run ./cmd/workbench check -workflow notes-to-project-register -output workflows/notes-to-project-register/example-output.md
```
