package workbench

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func LoadWorkflow(root, id string) (Workflow, error) {
	dir := filepath.Join(root, id)
	manifestPath := filepath.Join(dir, "manifest.json")
	b, err := os.ReadFile(manifestPath)
	if err != nil {
		return Workflow{}, fmt.Errorf("read manifest: %w", err)
	}
	var m Manifest
	if err := json.Unmarshal(b, &m); err != nil {
		return Workflow{}, fmt.Errorf("decode manifest: %w", err)
	}
	skillName := m.Skill
	if strings.TrimSpace(skillName) == "" {
		skillName = "SKILL.md"
	}
	skill, err := os.ReadFile(filepath.Join(dir, skillName))
	if err != nil {
		return Workflow{}, fmt.Errorf("read skill: %w", err)
	}
	return Workflow{Dir: dir, Manifest: m, Skill: string(skill)}, nil
}

func List(root string) ([]CatalogEntry, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	var out []CatalogEntry
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		wf, err := LoadWorkflow(root, entry.Name())
		if err != nil {
			return nil, fmt.Errorf("%s: %w", entry.Name(), err)
		}
		m := wf.Manifest
		out = append(out, CatalogEntry{ID: m.ID, Version: m.Version, Title: m.Title, Authority: m.Authority, Purpose: m.Purpose})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

func ValidateCatalog(root string) ValidationResult {
	entries, err := os.ReadDir(root)
	if err != nil {
		return ValidationResult{Valid: false, Issues: []ValidationIssue{{Path: root, Code: "read", Message: err.Error()}}}
	}
	var issues []ValidationIssue
	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		count++
		id := entry.Name()
		wf, err := LoadWorkflow(root, id)
		if err != nil {
			issues = append(issues, ValidationIssue{Workflow: id, Path: "manifest/skill", Code: "read", Message: err.Error()})
			continue
		}
		issues = append(issues, validateWorkflow(wf, id)...)
	}
	if count == 0 {
		issues = append(issues, ValidationIssue{Path: root, Code: "empty", Message: "no workflow directories found"})
	}
	return ValidationResult{Valid: len(issues) == 0, Issues: issues}
}

func validateWorkflow(wf Workflow, dirID string) []ValidationIssue {
	m := wf.Manifest
	var issues []ValidationIssue
	add := func(path, code, msg string) {
		issues = append(issues, ValidationIssue{Workflow: dirID, Path: path, Code: code, Message: msg})
	}
	if strings.TrimSpace(m.ID) == "" {
		add("id", "required", "workflow id is required")
	}
	if m.ID != dirID {
		add("id", "mismatch", "workflow id must match directory name")
	}
	if strings.TrimSpace(m.Version) == "" {
		add("version", "required", "workflow version is required")
	}
	if strings.TrimSpace(m.Title) == "" {
		add("title", "required", "workflow title is required")
	}
	if strings.TrimSpace(m.Purpose) == "" {
		add("purpose", "required", "workflow purpose is required")
	}
	if m.Authority != "human" && m.Authority != "deterministic" {
		add("authority", "invalid", "authority must be human or deterministic")
	}
	if strings.TrimSpace(wf.Skill) == "" {
		add("skill", "empty", "skill instructions cannot be empty")
	}
	if len(m.RequiredOutputSections) == 0 {
		add("required_output_sections", "required", "at least one output section is required")
	}
	seen := map[string]bool{}
	for i, s := range m.RequiredOutputSections {
		s = strings.TrimSpace(s)
		if s == "" {
			add(fmt.Sprintf("required_output_sections[%d]", i), "required", "section name cannot be empty")
			continue
		}
		if seen[s] {
			add(fmt.Sprintf("required_output_sections[%d]", i), "duplicate", "section names must be unique")
		}
		seen[s] = true
	}
	for _, name := range []string{"README.md", "example-input.md", "example-output.md"} {
		if _, err := os.Stat(filepath.Join(wf.Dir, name)); err != nil {
			add(name, "missing", "required workflow file is missing")
		}
	}
	return issues
}
