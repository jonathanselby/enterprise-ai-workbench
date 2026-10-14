# Enterprise AI Workbench — Product Requirements Document

## Problem

Enterprise AI work is often performed as one-off prompting. The result may be useful, but the method, evidence boundary, output structure and acceptance authority are rarely explicit enough for another engineer or reviewer to repeat the work.

## Product objective

Provide a small, reusable workbench for versioned AI-assisted enterprise workflows that turns messy inputs into structured candidate outputs while keeping authority outside the model.

> **Models analyse and propose. Humans or named deterministic controls accept.**

## Users

Engineers, architects, technology leaders, analysts, researchers and governance practitioners building repeatable AI-assisted work without adopting an agent framework.

## In scope

Readable Markdown workflow instructions; machine-readable manifests; workflow versioning; provider-neutral run bundles; deterministic run IDs and hashes; output-section contracts; explicit authority metadata; review records; worked fictional examples; catalogue validation; CLI use; CI tests.

## Out of scope

Autonomous decision making, model routing, agent orchestration, security-policy enforcement, proprietary classification, persistent workflow services, records management.

## V0.1 workflows

Research to decision; meeting to decisions; notes to project register; documents to risks; policy to obligations; vendor due diligence; unknowns first; messy input to structured signal.

## Acceptance criteria

A new engineer can inspect and run workflows without reading Go code; every workflow is complete; catalogue validation catches broken packages; render emits provenance and authority; run IDs are deterministic; output checks detect missing contract sections; accepted reviews require a conforming output; all examples pass CI.
