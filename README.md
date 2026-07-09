# Databricks Data Engineering Best Practices Skill

Installable Agent Skill for Databricks data engineering work: repo layout, bundles, CI/CD, workspace isolation, testing, observability, and production hygiene.

Source credit: this skill is derived from and credits Databricks documentation, especially [Developer best practices on Databricks](https://docs.databricks.com/aws/en/developers/best-practices). The skill paraphrases that guidance and adds agent-facing workflow structure.

## Install

### Codex + Claude Code

```bash
npx skills add nijanthan-dev/databricks-data-engineering-best-practices --global --agent codex claude-code --skill databricks-data-engineering-best-practices --copy -y
```

This was tested locally. It installs:

- Codex: `~/.agents/skills/databricks-data-engineering-best-practices`
- Claude Code: `~/.claude/skills/databricks-data-engineering-best-practices`

### Other Agents

Use the same command with the agent name your tool supports:

```bash
npx skills add nijanthan-dev/databricks-data-engineering-best-practices --global --agent <agent-name> --skill databricks-data-engineering-best-practices --copy -y
```

To target every supported agent, use `--agent '*'`. Prefer scoped agent names for a quieter install.

### OpenAI Responses API

Use the skill folder path in the `skills` array for local shell environments, or zip `skills/databricks-data-engineering-best-practices`.

### Databricks Genie

Genie Agents do not install Agent Skills packages with `npx skills add`. Configure the skill as Genie Agent instructions instead. See [Databricks Genie setup](docs/databricks-genie.md).

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
- `docs/databricks-genie.md`: Genie Agent setup guidance
- `tests/skill-evaluation.md`: skill TDD scenarios and checks
- `cmd/validate-skill`: local validation CLI
- `internal/skillvalidator`: typed Go validation package and tests

## License

MIT for repo content. Databricks documentation remains owned by Databricks and is credited in `SOURCE_CREDIT.md`.
