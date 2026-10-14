# Intent Before Prompt

A prompt is an implementation detail. Start by defining the work.

Before asking a model to do anything, write down:

1. **Objective** — what decision, artefact or reduction in uncertainty is needed?
2. **Source boundary** — what material may the model use, and what is outside scope?
3. **Authority** — who or what is allowed to accept the result as true or actionable?
4. **Evidence standard** — what must be directly supported, and what may be inferred?
5. **Output contract** — what structure will make the result usable by the next person or system?
6. **Failure behaviour** — what should happen when evidence is missing or contradictory?

This avoids a common failure mode: polishing the wording of a prompt before deciding what good work would look like.

The workbench manifests encode these ideas in small form: purpose, authority and required output sections. The skill then tells the model how to perform the analysis.
