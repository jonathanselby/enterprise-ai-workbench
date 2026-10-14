package workbench

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func sampleWorkflow(t *testing.T) Workflow {
	t.Helper()
	dir := t.TempDir()
	for _, name := range []string{"README.md", "example-input.md", "example-output.md"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	return Workflow{Dir: dir, Skill: "Do the work.", Manifest: Manifest{
		ID: "sample", Version: "1.0.0", Title: "Sample", Purpose: "Test workflow", Authority: "human", Skill: "SKILL.md",
		RequiredOutputSections: []string{"Known", "Unknown"},
	}}
}

func TestRunIDDeterministic(t *testing.T) {
	wf := sampleWorkflow(t)
	a := RunID(wf.Manifest, wf.Skill, "same input")
	b := RunID(wf.Manifest, wf.Skill, "same input")
	if a != b {
		t.Fatalf("run id changed: %s != %s", a, b)
	}
}

func TestRenderContainsProvenanceAndContract(t *testing.T) {
	wf := sampleWorkflow(t)
	bundle, err := Render(wf, "evidence")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"run-", "\"workflow_id\": \"sample\"", "## Known", "## Unknown", "evidence"} {
		if !strings.Contains(bundle, want) {
			t.Fatalf("bundle missing %q", want)
		}
	}
}

func TestCheckOutput(t *testing.T) {
	wf := sampleWorkflow(t)
	ok := CheckOutput(wf.Manifest, "## Known\nA\n## Unknown\nB")
	if !ok.Valid {
		t.Fatalf("expected valid: %+v", ok.Issues)
	}
	bad := CheckOutput(wf.Manifest, "## Known\nA")
	if bad.Valid || bad.Issues[0].Path != "Unknown" {
		t.Fatalf("expected missing Unknown: %+v", bad)
	}
}

func TestReviewAcceptedRequiresContract(t *testing.T) {
	wf := sampleWorkflow(t)
	_, err := Review(wf, "in", "## Known\nA", "accepted", "Reviewer", "looks good", time.Unix(1, 0))
	if err == nil {
		t.Fatal("expected output contract failure")
	}
}

func TestReviewRecordsHashes(t *testing.T) {
	wf := sampleWorkflow(t)
	r, err := Review(wf, "in", "## Known\nA\n## Unknown\nB", "accepted", "Reviewer", "reviewed", time.Unix(1, 0))
	if err != nil {
		t.Fatal(err)
	}
	if r.RunID == "" || r.InputSHA256 == "" || r.OutputSHA256 == "" || r.SkillSHA256 == "" {
		t.Fatalf("missing provenance: %+v", r)
	}
}
