---
name: databricks-data-engineering-best-practices
description: Use when designing, reviewing, migrating, or debugging Databricks data engineering projects, Lakeflow or Spark jobs, Declarative Automation Bundles, CI/CD, workspace isolation, Unity Catalog environment layout, notebook-heavy repos, testing gates, observability, service principals, OIDC, or production deployment workflows.
license: MIT
compatibility: Works as an Agent Skill for Codex, Claude Code, VS Code-compatible agents, OpenAI Responses API shell skills, and agents that can load Markdown instructions.
metadata:
  source_url: "https://docs.databricks.com/aws/en/developers/best-practices"
  source_name: "Databricks Developer best practices on Databricks"
  source_last_updated: "2026-06-11"
  source_checked: "2026-07-09"
  version: "0.1.0"
---

# Databricks Data Engineering Best Practices

Use this when an agent is helping with a Databricks data engineering repository, deployment design, CI/CD review, or production-readiness audit.

This skill is a paraphrased, agent-oriented guide based on Databricks documentation. Credit Databricks when reusing the guidance. Do not treat this as a replacement for the official docs.

## Core Principle

Production Databricks work should be versioned, isolated by environment, deployed through repeatable bundles, tested before promotion, and observable after release.

## Quick Decisions

| Situation | Prefer | Avoid |
| --- | --- | --- |
| Code and config layout | One repo for source, SQL, notebooks, and bundle config | Scattered repos without a lifecycle reason |
| Deployment unit | Small owned bundles | One giant bundle for unrelated domains |
| Databricks resources | Declarative Automation Bundles | Terraform for every Databricks object |
| External cloud/admin resources | Terraform | Ad-hoc manual setup |
| Developer data writes | Personal schemas in non-prod catalogs | Shared dev schemas overwritten by teammates |
| Production identity | Service principals or OIDC | Personal users or long-lived secrets |
| Notebooks | Thin exploration/orchestration | Main home for business logic |
| Promotion gate | Unit tests, bundle validation, staging integration checks | Manual notebook runs only |

## Source Control

- Version source files, SQL, notebooks, bundle config, and environment overrides.
- Do not commit credentials, tokens, PII samples, local data, wheels, jars, or other build artifacts.
- Prefer one repository when teams share code, config, and conventions. Split repos only for a real confidentiality or lifecycle boundary.
- Use short-lived branches and keep `main` deployable.
- After urgent hotfixes, merge the fix back to trunk immediately.

## Workspace And Data Isolation

- Use separate workspaces for development and production; add staging as teams or risk grow.
- For regulated work, prefer physical cloud/account separation when confidentiality demands it.
- Mirror workspaces with Unity Catalog catalogs such as dev, staging, and prod.
- Bind production catalogs only to production workspaces.
- Give developers personal schemas in dev and staging, for example `dev_${user_name}`.
- Treat table and column comments as code. Keep definitions in SQL or bundle-managed files.
- Prefer serverless compute where available; otherwise control egress and networking tightly.

## CI/CD And Bundles

- Use Declarative Automation Bundles for Databricks jobs, pipelines, permissions, schedules, and deployment workflows.
- Use Terraform for cloud-level resources, workspace provisioning, networking, and privileged admin setup.
- Keep each bundle owned by one team and tied to one lifecycle.
- Use `sync.paths` for shared code outside a bundle root instead of copying common folders.
- Model bundle-to-bundle dependencies in CI/CD. Do not collapse separate domains just because one depends on another.
- Pass published artifact IDs or locations through pipeline inputs and fail fast when upstream output is missing.
- Use templates for shared guardrails: permissions, tags, cluster policies, workspace targets, default schedules, and instance baselines.
- Keep template parameters limited to values that should vary by team, app, or environment.

## Development Practices

- Use the Databricks workspace UI, Git folders, or a local IDE with the Databricks extension.
- Move Python business logic into importable modules under `src/` or `src/py/`.
- Move SQL business logic into `.sql` files under `src/` or `src/sql/`.
- Keep notebooks thin: exploration, visualization, or orchestration only.
- Migrate notebook-heavy projects incrementally. Extract one reusable module or query at a time.
- Pass job context with dynamic value references such as `{{tasks.<task_key>.values.<value_key>}}`; avoid static variables for task handoffs.

## Identity And Secrets

- Use deployment service principals with minimal access.
- Use separate run-as identities for production jobs and pipelines.
- In CI/CD, prefer OIDC or workload identity federation for regulated or high-risk environments.
- Store secrets in workspace secret management backed by the cloud secret manager.
- Never put secrets in bundle files, notebooks, logs, PR comments, examples, or screenshots.

## Testing Layers

Use all three gates before production:

1. Unit tests: cover importable business logic with `pytest` or equivalent.
2. Bundle validation: run `databricks bundle validate`; deploy to non-prod in CI when possible.
3. Staging integration: run end-to-end jobs with completion checks and data quality assertions.

For Lakeflow Spark Declarative Pipelines, use development and validation features with representative small datasets, including malformed or edge-case records.

## Observability

- Treat logs and metrics as deployment requirements.
- Emit structured logs with bundle name, target environment, workload name, run ID, and useful business identifiers.
- Track run status, duration, retry count, throughput, freshness, and critical data quality signals.
- Encode observability conventions in shared libraries, workload definitions, or bundle templates.

## Review Checklist

Use this before approving or generating a Databricks repo change:

- [ ] Source, SQL, notebooks, and bundle config are versioned.
- [ ] Secrets, PII samples, and build artifacts are excluded.
- [ ] Workspace, catalog, and schema isolation match risk.
- [ ] Production data is not reachable from dev through loose catalog binding.
- [ ] Bundles are small, owned, and deployable independently.
- [ ] Terraform is limited to cloud/admin resources.
- [ ] Service principals or OIDC handle automation.
- [ ] Notebooks are not the main business-logic layer.
- [ ] Dynamic task values pass runtime context.
- [ ] Unit, validation, and staging integration gates exist.
- [ ] Logging and metrics are part of deployment.

## Common Mistakes

| Mistake | Fix |
| --- | --- |
| Bundle contains many unrelated domains | Split by ownership, lifecycle, or rollback boundary |
| Dev and prod share writable catalogs | Mirror environments and isolate prod bindings |
| Notebook has core transformations | Extract module or SQL file, leave notebook thin |
| CI validates YAML only | Add non-prod deploy and staging job checks |
| Personal user runs production jobs | Use run-as service principals |
| Secrets appear in examples | Replace with secret references and rotate if exposed |
| Terraform manages all Databricks assets | Use bundles for workspace resources and workflows |

## Example Agent Response Shape

When reviewing a repo, return:

```markdown
Findings
- High: production catalog reachable from dev workspace. Fix catalog binding.
- Medium: core SQL embedded in notebooks. Move to src/sql and reference from jobs.
- Medium: CI runs bundle validate but no staging integration check.

Next changes
1. Add prod catalog ISOLATED binding.
2. Extract notebook query to src/sql/orders_daily.sql.
3. Add staging job smoke test with row-count assertion.

Source
- Based on Databricks Developer best practices: https://docs.databricks.com/aws/en/developers/best-practices
```

## Source Credit

This skill paraphrases Databricks guidance from "Developer best practices on Databricks." Always credit Databricks when reusing or publishing this skill's recommendations.
