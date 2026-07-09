# Contributing

## Scope

Keep this skill focused on Databricks data engineering delivery practices: repo layout, bundles, CI/CD, workspaces, Unity Catalog, identity, testing, and observability.

## Rules

- Credit Databricks for source-derived guidance.
- Label community-informed alternatives; do not attribute them to Databricks.
- Paraphrase. Do not copy Databricks docs wholesale.
- Check the official source before changing recommendations.
- Keep `SKILL.md` concise and agent-oriented.
- Run validation before opening a PR.

```bash
go vet ./...
go test -race -shuffle=on ./...
go run ./cmd/validate-skill
```

## PR checklist

- [ ] Source page checked.
- [ ] Databricks credit preserved.
- [ ] Skill name and folder still match.
- [ ] Description remains under 1024 chars.
- [ ] No secrets, tokens, PII, or build artifacts.
- [ ] Validation passes.
