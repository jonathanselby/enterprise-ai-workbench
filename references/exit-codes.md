# CLI Exit Codes

| Code | Meaning |
|---:|---|
| 0 | command succeeded; validation/check passed where relevant |
| 1 | workflow catalogue or model output failed a deterministic contract check |
| 2 | CLI usage error |
| 3 | file, manifest, parsing or execution error |

These codes are intended for shell scripts and CI pipelines.
