package workbench

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func SHA256Text(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func RunID(m Manifest, skill, input string) string {
	seed := m.ID + "\n" + m.Version + "\n" + SHA256Text(skill) + "\n" + SHA256Text(input)
	return "run-" + SHA256Text(seed)[:16]
}

func Render(wf Workflow, input string) (string, error) {
	if issues := validateWorkflow(wf, wf.Manifest.ID); len(issues) > 0 {
		return "", fmt.Errorf("workflow is invalid: %s", issues[0].Message)
	}
	m := wf.Manifest
	meta := BundleMetadata{
		RunID: RunID(m, wf.Skill, input), WorkflowID: m.ID, WorkflowVersion: m.Version,
		Authority: m.Authority, SkillSHA256: SHA256Text(wf.Skill), InputSHA256: SHA256Text(input),
		RequiredOutputSections: append([]string(nil), m.RequiredOutputSections...),
	}
	metaJSON, _ := json.MarshalIndent(meta, "", "  ")
	var b strings.Builder
	b.WriteString("# Enterprise AI Workbench Run Bundle\n\n")
	b.WriteString("## Provenance\n\n```json\n")
	b.Write(metaJSON)
	b.WriteString("\n```\n\n")
	b.WriteString("## Authority Boundary\n\n")
	if m.Authority == "human" {
		b.WriteString("The model may analyse and propose. A human reviewer remains authoritative for acceptance.\n\n")
	} else {
		b.WriteString("The model may analyse and propose. Acceptance must be performed by the named deterministic check outside the model.\n\n")
	}
	b.WriteString("## Workflow Instructions\n\n")
	b.WriteString(strings.TrimSpace(wf.Skill))
	b.WriteString("\n\n## Required Output Contract\n\n")
	for _, s := range m.RequiredOutputSections {
		b.WriteString("- `## " + s + "`\n")
	}
	b.WriteString("\n## Input\n\n")
	b.WriteString(strings.TrimSpace(input))
	b.WriteString("\n")
	return b.String(), nil
}

func CheckOutput(m Manifest, output string) ValidationResult {
	var issues []ValidationIssue
	for _, section := range m.RequiredOutputSections {
		heading := "## " + strings.TrimSpace(section)
		if !containsHeading(output, heading) {
			issues = append(issues, ValidationIssue{Workflow: m.ID, Path: section, Code: "missing_section", Message: "required output section is missing"})
		}
	}
	return ValidationResult{Valid: len(issues) == 0, Issues: issues}
}

func containsHeading(output, heading string) bool {
	for _, line := range strings.Split(output, "\n") {
		if strings.TrimSpace(line) == heading {
			return true
		}
	}
	return false
}

func Review(wf Workflow, input, output, decision, reviewer, reason string, now time.Time) (ReviewRecord, error) {
	decision = strings.ToLower(strings.TrimSpace(decision))
	if decision != "accepted" && decision != "rejected" {
		return ReviewRecord{}, fmt.Errorf("decision must be accepted or rejected")
	}
	if strings.TrimSpace(reviewer) == "" {
		return ReviewRecord{}, fmt.Errorf("reviewer is required")
	}
	if strings.TrimSpace(reason) == "" {
		return ReviewRecord{}, fmt.Errorf("reason is required")
	}
	if decision == "accepted" {
		if result := CheckOutput(wf.Manifest, output); !result.Valid {
			return ReviewRecord{}, fmt.Errorf("cannot accept output that fails output contract: %s", result.Issues[0].Path)
		}
	}
	return ReviewRecord{
		RunID: RunID(wf.Manifest, wf.Skill, input), WorkflowID: wf.Manifest.ID, WorkflowVersion: wf.Manifest.Version,
		Decision: decision, Reviewer: reviewer, Reason: reason, InputSHA256: SHA256Text(input),
		OutputSHA256: SHA256Text(output), SkillSHA256: SHA256Text(wf.Skill), ReviewedAt: now.UTC(),
	}, nil
}

func ReadText(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
