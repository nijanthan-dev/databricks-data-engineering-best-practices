package skillvalidator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateSkillAcceptsCurrentRepo(t *testing.T) {
	root := repoRoot(t)

	result, err := Validate(root)
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
	if !result.Valid {
		t.Fatalf("expected valid result, got issues: %v", result.Issues)
	}
}

func TestValidateSkillReportsMissingRequiredTerm(t *testing.T) {
	root := copyRepoFixture(t)
	skillPath := filepath.Join(root, "skills", SkillName, "SKILL.md")
	contents := readFile(t, skillPath)
	contents = strings.ReplaceAll(contents, "Unity Catalog", "Catalog")
	writeFile(t, skillPath, contents)

	result, err := Validate(root)
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
	if result.Valid {
		t.Fatal("expected invalid result")
	}
	assertIssue(t, result.Issues, "missing terms: Unity Catalog")
}

func TestValidateSkillReportsMissingCommunityLabel(t *testing.T) {
	root := copyRepoFixture(t)
	skillPath := filepath.Join(root, "skills", SkillName, "SKILL.md")
	contents := readFile(t, skillPath)
	contents = strings.ReplaceAll(contents, "community-informed", "operational")
	writeFile(t, skillPath, contents)

	result, err := Validate(root)
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
	if result.Valid {
		t.Fatal("expected invalid result")
	}
	assertIssue(t, result.Issues, "missing terms: community-informed")
}

func TestValidateSkillReportsFolderNameMismatch(t *testing.T) {
	root := copyRepoFixture(t)
	oldDir := filepath.Join(root, "skills", SkillName)
	newDir := filepath.Join(root, "skills", "wrong-name")
	if err := os.Rename(oldDir, newDir); err != nil {
		t.Fatalf("rename fixture skill dir: %v", err)
	}

	result, err := Validate(root)
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
	if result.Valid {
		t.Fatal("expected invalid result")
	}
	assertIssue(t, result.Issues, "missing skill directory")
}

func TestValidateSkillReportsDescriptionPrefix(t *testing.T) {
	root := copyRepoFixture(t)
	skillPath := filepath.Join(root, "skills", SkillName, "SKILL.md")
	contents := readFile(t, skillPath)
	contents = strings.Replace(contents, "description: Use when ", "description: For ", 1)
	writeFile(t, skillPath, contents)

	result, err := Validate(root)
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
	if result.Valid {
		t.Fatal("expected invalid result")
	}
	assertIssue(t, result.Issues, "description must start with 'Use when '")
}

func TestValidateSkillReportsFrontmatterNameMismatch(t *testing.T) {
	root := copyRepoFixture(t)
	skillPath := filepath.Join(root, "skills", SkillName, "SKILL.md")
	contents := readFile(t, skillPath)
	contents = strings.Replace(contents, "name: "+SkillName, "name: wrong-name", 1)
	writeFile(t, skillPath, contents)

	result, err := Validate(root)
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
	if result.Valid {
		t.Fatal("expected invalid result")
	}
	assertIssue(t, result.Issues, "name must match folder")
}

func repoRoot(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			t.Fatal("repo root not found")
		}
		wd = parent
	}
}

func copyRepoFixture(t *testing.T) string {
	t.Helper()
	source := repoRoot(t)
	target := t.TempDir()
	for _, rel := range []string{
		filepath.Join("skills", SkillName, "SKILL.md"),
	} {
		src := filepath.Join(source, rel)
		dst := filepath.Join(target, rel)
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			t.Fatalf("mkdir fixture: %v", err)
		}
		writeFile(t, dst, readFile(t, src))
	}
	return target
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return string(data)
}

func writeFile(t *testing.T, path string, contents string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func assertIssue(t *testing.T, issues []string, want string) {
	t.Helper()
	for _, issue := range issues {
		if strings.Contains(issue, want) {
			return
		}
	}
	t.Fatalf("expected issue containing %q, got %v", want, issues)
}
