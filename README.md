# Enterprise AI Workbench

**Small, reusable AI workflows for turning messy enterprise information into decisions, registers, briefs and technical artefacts.**

Enterprise AI work is usually messier than the software demo. The hard part is often not getting a model to produce prose; it is making evidence, assumptions, contradictions, ownership, uncertainty and acceptance visible enough that somebody else can rely on the work.

The workbench is built around one rule:

> **Models analyse and propose. Humans or named deterministic controls accept.**

The workflows are plain Markdown and can be used with any capable LLM. A small Go CLI adds the engineering layer: versioned workflow manifests, catalogue validation, reproducible run bundles, hashes, output contracts and review records.

## What is included

| Workflow | Purpose |
|---|---|
| `research-to-decision` | Turn research material into evidence, unknowns, options and a defensible decision brief. |
| `meeting-to-decisions` | Separate decisions from actions, assumptions, disagreements, deferrals and unknowns. |
| `notes-to-project-register` | Turn project notes into a deliberately small operational register. |
| `documents-to-risks` | Extract candidate risks as cause, event and consequence for human review. |
| `policy-to-obligations` | Extract normative obligations, triggers, evidence requirements and exceptions. |
| `vendor-due-diligence` | Separate vendor claims from evidence, dependencies, lock-in and questions to test. |
| `unknowns-first` | Distinguish known, inferred, assumed, contradictory and unknown. |
| `messy-input-to-structured-signal` | Create a traceable candidate structure from mixed-quality enterprise material. |

Every workflow contains its own README, `manifest.json`, `SKILL.md`, fictional example input and conforming example output.

## Install

Requires Go 1.23+.

```bash
go install github.com/jonathanselby/enterprise-ai-workbench/cmd/workbench@latest
```

Or run directly from a clone:

```bash
go run ./cmd/workbench version
```

## Start here

List the workflows:

```bash
workbench list
```

Validate the catalogue:

```bash
workbench validate
```

Render a reproducible run bundle:

```bash
workbench render \
  -workflow research-to-decision \
  -input workflows/research-to-decision/example-input.md \
  -out /tmp/research-run.md
```

The bundle contains the workflow version, authority boundary, SHA-256 hashes of the skill and input, the required output contract, instructions and source input.

Send that bundle to the model of your choice. The workbench deliberately does **not** contain provider integrations.

Check the returned output structurally:

```bash
workbench check \
  -workflow research-to-decision \
  -output workflows/research-to-decision/example-output.md
```

Record human acceptance:

```bash
workbench review \
  -workflow research-to-decision \
  -input workflows/research-to-decision/example-input.md \
  -output workflows/research-to-decision/example-output.md \
  -decision accepted \
  -reviewer "Architecture Review" \
  -reason "Evidence and recommendation reviewed"
```

The resulting record hashes the exact workflow instructions, input and output that were reviewed.

## Why the workbench does not call a model

That is deliberate separation of concerns.

The workbench owns **method and provenance**. A model layer owns **inference**. A human or deterministic control owns **acceptance**.

```text
WORKFLOW METHOD
      |
      v
RUN BUNDLE + PROVENANCE
      |
      v
MODEL / MODEL ROUTER
      |
      v
CANDIDATE OUTPUT
      |
      v
CONTRACT CHECK
      |
      v
HUMAN / DETERMINISTIC ACCEPTANCE
```

If model routing is required, the rendered bundle can be handed to [`three-lane-model-harness`](https://github.com/jonathanselby/three-lane-model-harness) or another provider layer. Keeping that concern separate means a workflow does not change merely because the organisation changes model provider.

## What the CLI proves — and what it does not

The CLI can prove that:

- a named workflow/version was used;
- the skill instructions and input hash to known values;
- the output contains the workflow's required structural sections;
- a reviewer recorded an accepted/rejected decision against specific input/output hashes.

It **cannot** prove that:

- the source material was complete;
- the model's claims are true;
- the reviewer is really the person named in a record;
- the human decision was correct.

That restraint is intentional. Provenance is not truth.

## Methods

Three short methods underpin the workflows:

- [`intent-before-prompt`](methods/intent-before-prompt.md) — define the work before polishing the prompt;
- [`human-review-test`](methods/human-review-test.md) — make review substantive rather than ceremonial;
- [`accepted-useful-work`](methods/accepted-useful-work.md) — keep acceptance authority outside the model.

## Repository structure

```text
enterprise-ai-workbench/
├── README.md
├── PRD.md
├── EDD.md
├── architecture.txt
├── VERSION
├── CHANGELOG.md
├── LICENSE
├── cmd/workbench/
├── internal/workbench/
├── methods/
├── references/
└── workflows/
    ├── research-to-decision/
    ├── meeting-to-decisions/
    ├── notes-to-project-register/
    ├── documents-to-risks/
    ├── policy-to-obligations/
    ├── vendor-due-diligence/
    ├── unknowns-first/
    └── messy-input-to-structured-signal/
```

## Design choices

- **Markdown first.** A workflow remains understandable without the CLI.
- **Machine-readable manifest.** Tooling can validate and enumerate workflows without parsing prose.
- **Provider-neutral.** No model SDK is embedded in the method repository.
- **Deterministic provenance.** Unchanged workflow instructions and input produce the same run ID.
- **Acceptance outside the model.** The model cannot declare its own work authoritative.
- **Fictional examples only.** No employer, board or client material is published.
- **Standard library only.** The executable remains small and easy to build.

## Limitations

This is not an agent framework, security control, records-management system or autonomous decision engine. Output checking is structural. Reviewer identity is not authenticated. Review records are not digitally signed. Sensitive inputs remain the operator's responsibility.

## Exit codes

See [`references/exit-codes.md`](references/exit-codes.md).

## Licence

MIT. See [`LICENSE`](LICENSE).
