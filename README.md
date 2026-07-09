# Databricks Data Engineering Best Practices Skill

Installable Agent Skill for Databricks data engineering work: repo layout, bundles, CI/CD, workspace isolation, testing, observability, and production hygiene.

Source credit: this skill is derived from and credits Databricks documentation, especially [Developer best practices on Databricks](https://docs.databricks.com/aws/en/developers/best-practices). The skill paraphrases that guidance and adds agent-facing workflow structure.

## Install

### Codex / VS Code agent skills

```bash
mkdir -p .agents/skills
cp -R skills/databricks-data-engineering-best-practices .agents/skills/
```

### Claude Code

```bash
mkdir -p ~/.claude/skills
cp -R skills/databricks-data-engineering-best-practices ~/.claude/skills/
```

### OpenAI Responses API

Use the skill folder path in the `skills` array for local shell environments, or zip the folder for hosted/inline skill use.

### Databricks Genie and other agents

If the agent supports Agent Skills, install the folder as-is. If it only supports custom instructions or saved agents, attach or paste `SKILL.md` as the instruction source and keep the `SOURCE_CREDIT.md` link.

## Validate

```bash
go vet ./...
go test -race -shuffle=on ./...
go run ./cmd/validate-skill
```

## Pre-commit

```bash
git config core.hooksPath .githooks
```

The hook formats Go, runs vet, runs race/shuffle tests, and validates the skill.

## Files

- `skills/databricks-data-engineering-best-practices/SKILL.md`: installable skill
- `SOURCE_CREDIT.md`: attribution and source notes
- `DISCOVERY.md`: naming, keywords, and discoverability
- `tests/skill-evaluation.md`: skill TDD scenarios and checks
- `cmd/validate-skill`: local validation CLI
- `internal/skillvalidator`: typed Go validation package and tests

## License

MIT for repo content. Databricks documentation remains owned by Databricks and is credited in `SOURCE_CREDIT.md`.
