# Enterprise AI Workbench — Engineering Design Document

## 1. Purpose

The Enterprise AI Workbench is a small Go CLI plus a catalogue of reusable Markdown workflows. It makes AI-assisted enterprise work repeatable without making the workbench itself a model provider, router or agent framework.

> **The workflow structures model work. Authority remains outside the model.**

## 2. Provider-neutral boundary

`render` produces a run bundle that can be pasted into a capable LLM, consumed by automation, or handed to a separate routing layer such as `three-lane-model-harness`. Workflow method is decoupled from provider transport.

## 3. Component model

```text
workflow files -> workbench validate/render -> external model -> model output -> workbench check -> human/named deterministic review -> review record
```

## 4. Workflow contract

Every workflow contains `README.md`, `manifest.json`, `SKILL.md`, `example-input.md`, and `example-output.md`.

Manifest fields: `id`, `version`, `title`, `purpose`, `authority`, `skill`, `required_output_sections`.

## 5. Catalogue validation

Validation checks manifest readability; directory/manifest ID agreement; required metadata; supported authority; non-empty skill; unique output sections; and required workflow files.

## 6. Run bundle

A rendered bundle contains deterministic run ID, workflow ID/version, authority, SHA-256 skill/input hashes, required sections, instructions and original input. The run ID is derived from workflow identity/version plus skill and input hashes, not wall-clock time.

## 7. Output checking

`check` confirms all required level-two Markdown headings exist. This is structural validation only and is not presented as proof of factual correctness.

## 8. Review and acceptance

`review` records run ID, workflow/version, decision, reviewer, reason, input/output/skill hashes and review time. An accepted result is refused if the deterministic output contract fails.

## 9. CLI and exit codes

Commands: `version`, `list`, `validate`, `render`, `check`, `review`.

Exit codes: 0 success; 1 deterministic contract failure; 2 usage error; 3 file/manifest/execution error.

## 10. State

No database or service state. Files are ordinary inputs; review records are emitted to stdout for caller-controlled persistence.

## 11. Security boundary

No network calls or credentials. Hashes are provenance/change references, not confidentiality or identity. `reviewer` is not authenticated. No DLP, prompt-injection defence or policy enforcement is implemented.

## 12. Tests

Unit tests cover deterministic run IDs, provenance, output contracts and review behaviour. CI builds the CLI, validates the full catalogue, renders a bundle, checks every worked output and records an example acceptance.

## 13. Limitations

V0.1 does not execute a model, automatically judge semantic correctness, authenticate reviewers, sign records, persist runs, provide a web interface or infer which workflow should be used.
