package workbench

import "time"

type Manifest struct {
	ID                     string   `json:"id"`
	Version                string   `json:"version"`
	Title                  string   `json:"title"`
	Purpose                string   `json:"purpose"`
	Authority              string   `json:"authority"`
	Skill                  string   `json:"skill"`
	RequiredOutputSections []string `json:"required_output_sections"`
}

type Workflow struct {
	Dir      string
	Manifest Manifest
	Skill    string
}

type CatalogEntry struct {
	ID        string `json:"id"`
	Version   string `json:"version"`
	Title     string `json:"title"`
	Authority string `json:"authority"`
	Purpose   string `json:"purpose"`
}

type ValidationIssue struct {
	Workflow string `json:"workflow,omitempty"`
	Path     string `json:"path"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

type ValidationResult struct {
	Valid  bool              `json:"valid"`
	Issues []ValidationIssue `json:"issues,omitempty"`
}

type BundleMetadata struct {
	RunID                  string   `json:"run_id"`
	WorkflowID             string   `json:"workflow_id"`
	WorkflowVersion        string   `json:"workflow_version"`
	Authority              string   `json:"authority"`
	SkillSHA256            string   `json:"skill_sha256"`
	InputSHA256            string   `json:"input_sha256"`
	RequiredOutputSections []string `json:"required_output_sections"`
}

type ReviewRecord struct {
	RunID           string    `json:"run_id"`
	WorkflowID      string    `json:"workflow_id"`
	WorkflowVersion string    `json:"workflow_version"`
	Decision        string    `json:"decision"`
	Reviewer        string    `json:"reviewer"`
	Reason          string    `json:"reason"`
	InputSHA256     string    `json:"input_sha256"`
	OutputSHA256    string    `json:"output_sha256"`
	SkillSHA256     string    `json:"skill_sha256"`
	ReviewedAt      time.Time `json:"reviewed_at"`
}
