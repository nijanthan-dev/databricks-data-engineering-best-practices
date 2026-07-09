package skillvalidator

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	SkillName = "databricks-data-engineering-best-practices"
	sourceURL = "https://docs.databricks.com/aws/en/developers/best-practices"
)

var requiredTerms = []string{
	"Declarative Automation Bundles",
	"Unity Catalog",
	"service principals",
	"OIDC",
	"Lakeflow Spark Declarative Pipelines",
	"Databricks",
}

var (
	namePattern        = regexp.MustCompile(`(?m)^name:\s*([a-z0-9-]+)\s*$`)
	descriptionPattern = regexp.MustCompile(`(?m)^description:\s*(.+)\s*$`)
)

type Result struct {
	Valid  bool
	Issues []string
}

func Validate(root string) (Result, error) {
	skillDir := filepath.Join(root, "skills", SkillName)
	skillPath := filepath.Join(skillDir, "SKILL.md")

	data, err := os.ReadFile(skillPath)
	if err != nil {
		if os.IsNotExist(err) {
			return Result{Valid: false, Issues: []string{"missing skill directory or SKILL.md"}}, nil
		}
		return Result{}, fmt.Errorf("read skill: %w", err)
	}

	text := string(data)
	issues := validateText(text, filepath.Base(skillDir))
	return Result{Valid: len(issues) == 0, Issues: issues}, nil
}

func validateText(text string, folderName string) []string {
	var issues []string
	if !strings.HasPrefix(text, "---\n") {
		return []string{"missing YAML frontmatter"}
	}

	parts := strings.SplitN(text, "---\n", 3)
	if len(parts) < 3 {
		return []string{"frontmatter not closed"}
	}

	frontmatter := parts[1]
	body := parts[2]

	nameMatch := namePattern.FindStringSubmatch(frontmatter)
	if nameMatch == nil {
		issues = append(issues, "invalid or missing name")
	} else {
		name := nameMatch[1]
		if name != folderName {
			issues = append(issues, "name must match folder")
		}
		if len(name) > 64 || strings.HasPrefix(name, "-") || strings.HasSuffix(name, "-") || strings.Contains(name, "--") {
			issues = append(issues, "name violates spec")
		}
	}

	descriptionMatch := descriptionPattern.FindStringSubmatch(frontmatter)
	if descriptionMatch == nil {
		issues = append(issues, "missing description")
	} else {
		description := strings.TrimSpace(descriptionMatch[1])
		if len(description) > 1024 {
			issues = append(issues, "description over 1024 chars")
		}
		if !strings.HasPrefix(description, "Use when ") {
			issues = append(issues, "description must start with 'Use when '")
		}
	}

	var missing []string
	for _, term := range requiredTerms {
		if !strings.Contains(text, term) {
			missing = append(missing, term)
		}
	}
	if len(missing) > 0 {
		issues = append(issues, "missing terms: "+strings.Join(missing, ", "))
	}

	if !strings.Contains(text, sourceURL) {
		issues = append(issues, "missing Databricks source URL")
	}

	if len(strings.Split(strings.TrimSuffix(body, "\n"), "\n")) > 500 {
		issues = append(issues, "SKILL.md body too long")
	}

	return issues
}
