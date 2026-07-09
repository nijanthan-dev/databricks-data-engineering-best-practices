---
name: databricks-data-engineering-best-practices
description: Use when designing, reviewing, migrating, or debugging Databricks data engineering projects, Lakeflow or Spark jobs, Declarative Automation Bundles, CI/CD, workspace isolation, Unity Catalog environment layout, notebook-heavy repos, testing gates, observability, service principals, OIDC, or production deployment workflows.
license: MIT
compatibility: Works as an Agent Skill for Codex, Claude Code, VS Code-compatible agents, OpenAI Responses API shell skills, and agents that can load Markdown instructions.
metadata:
  source_url: "https://docs.databricks.com/aws/en/developers/best-practices"
  source_name: "Databricks Developer best practices on Databricks"
  source_last_updated: "2026-06-11"
  source_checked: "2026-07-10"
  version: "0.1.0"
---

# Databricks Data Engineering Best Practices

Use this when an agent is helping with a Databricks data engineering repository, deployment design, CI/CD review, or production-readiness audit.

This skill paraphrases Databricks documentation and adds original, community-informed operating heuristics. Community-informed alternatives are labeled; do not attribute them to Databricks. Credit Databricks when reusing source-derived guidance, and do not treat this skill as a replacement for the official docs.

## Core Principle

Production Databricks work should be versioned, isolated by environment, deployed through repeatable bundles, tested before promotion, and observable after release.

## Quick Decisions

| Situation | Prefer | Avoid |
| --- | --- | --- |
| Repository boundary | Databricks single-repo default or documented local exception | Unexplained split or monorepo |
| Deployment unit | Small owned bundles | One giant bundle for unrelated domains |
| Resource owner | Bundles by default; one documented owner per resource | Dual ownership |
| Developer data | Governed non-prod data and personal write schemas | Casual raw production PII access |
| Production identity | Service principals or OIDC | Personal users or long-lived secrets |
| Notebooks | Thin exploration/orchestration | Main home for business logic |
| Promotion gate | Tests plus risk-appropriate release controls | Merge-to-main always deploys |

## Repository Boundaries

- Version source files, SQL, notebooks, bundle config, and environment overrides.
- Do not commit credentials, tokens, PII samples, local data, wheels, jars, or other build artifacts.
- Start with the Databricks recommendation: one repository for source and configuration, including bundles with separate deployment lifecycles.
- Ask what "single repo" means: organization, domain, team, product, or bundle.
- As a community-informed alternative, an organization may document broader split criteria when lifecycle, risk, compliance, release cadence, or deployment ownership differs meaningfully.
- Split only at that documented boundary. Confidentiality and regulated separation remain the clearest reasons to depart from the Databricks default.
- A cookiecutter or custom bundle template per project or team can preserve standards when an organization deliberately favors simpler repos and smaller blast radius over the single-repo default.
- Use short-lived branches and keep `main` deployable.
- After urgent hotfixes, merge the fix back to trunk immediately.

## Workspace And Data Isolation

- Use separate workspaces for development and production; add staging as teams or risk grow.
- For regulated work, prefer physical cloud/account separation when confidentiality demands it.
- Mirror workspaces with Unity Catalog catalogs such as dev, staging, and prod.
- Bind production catalogs only to production workspaces.
- Give developers personal schemas in dev and staging, for example `dev_${user_name}`.
- If production contains PII, do not point development casually at raw production data.
- Prefer masked or tokenized production-like subsets, synthetic sensitive fields, and governed row or column filters.
- Use service principals for controlled validation against real inputs. Stable inputs help testing, but governance decides who can access them.
- Treat table and column comments as code. Keep definitions in SQL or bundle-managed files.
- Prefer serverless compute where available; otherwise control egress and networking tightly.

## Resource Ownership

- Start with the Databricks recommendation: Terraform owns external cloud resources and privileged admin setup; Declarative Automation Bundles own other Databricks resources.
- A community-informed platform model may instead assign shared, long-lived state to Terraform where drift matters: storage credentials, external locations, catalogs, shared schemas, grants, warehouses, service principals, and cluster policies. Adopt this only as explicit organization policy.
- In that model, bundles own app or workflow resources that ship and change together: jobs, pipelines, notebooks, dashboards, alerts, and app-private resources.
- Never manage the same object with both Terraform and a bundle.
- Ask: "What happens if this bundle is destroyed?" A scary answer signals the wrong owner or missing protection.
- Bundle-managed resources follow the bundle lifecycle.
- Under the platform model, default shared schemas to Terraform because they become stable addresses for BI, jobs, permissions, lineage, users, and downstream systems.
- Let a bundle own a schema only when it is app-private, low-blast-radius, safe to recreate, and shares the app lifecycle.
- Under the Databricks default, put long-lived shared resources in a separate infrastructure or foundation bundle, not an app bundle.
- Treat lifecycle protection and bind or unbind workflows as safeguards, not substitutes for clear ownership.

## Bundles And Delivery

- Keep each bundle owned by one team and tied to one lifecycle.
- Use `sync.paths` for shared code outside a bundle root instead of copying common folders.
- For inter-bundle dependencies, define upstream output and data contracts: owner, version, schema compatibility, quality checks, and supported changes.
- Validate dependency order and contracts in CI/CD. Do not collapse independent bundles only because they exchange data.
- Pass published artifact IDs or locations through pipeline inputs; fail fast when upstream output is absent or incompatible.
- Use templates for shared guardrails: permissions, tags, cluster policies, workspace targets, default schedules, and instance baselines.
- Keep template parameters limited to values that should vary by team, app, or environment.
- Define targets and variables once, then use small environment overrides for workspace paths, catalogs, service principals, schedules, and sizing. Do not copy large bundle sections across dev, staging, and prod.
- Keep secrets out of bundle files. Use secret references and identity-based authentication.
- Match deployment cadence to risk. Regulated or operational environments may require approvals, release batches, downtime windows, or external dependency alignment.
- GitHub Releases or CalVer-triggered promotion is valid. Do not assume every merge to `main` must deploy to staging or production.

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

## Agent-Friendly Repo Design

- Give coding agents a predictable layout, small bundles, runnable tests, exact validation commands, representative examples, ownership metadata, and concise deployment docs.
- Make verification local and deterministic so agents can prove changes instead of guessing.
- For Genie, provide instructions, examples, metadata, benchmarks, and API configuration. Do not tell users to install skills into Genie with `npx`.

## Observability

- Treat logs and metrics as deployment requirements.
- Emit structured logs with bundle name, target environment, workload name, run ID, and useful business identifiers.
- Track run status, duration, retry count, throughput, freshness, and critical data quality signals.
- Encode observability conventions in shared libraries, workload definitions, or bundle templates.

## Review Checklist

Use this before approving or generating a Databricks repo change:

- [ ] Source, SQL, notebooks, and bundle config are versioned.
- [ ] Secrets, PII samples, and build artifacts are excluded.
- [ ] Repo structure follows the Databricks single-repo default or documents the local exception.
- [ ] Workspace, catalog, and schema isolation match risk.
- [ ] Production data is not reachable from dev through loose catalog binding.
- [ ] PII access uses governed, masked, tokenized, synthetic, or controlled data paths.
- [ ] Bundles are small, owned, and deployable independently.
- [ ] Every resource has one owner; shared long-lived state is outside app bundles.
- [ ] Inter-bundle data contracts define ownership, compatibility, and quality gates.
- [ ] Targets use concise overrides; secrets remain external.
- [ ] Service principals or OIDC handle automation.
- [ ] Notebooks are not the main business-logic layer.
- [ ] Dynamic task values pass runtime context.
- [ ] Unit, bundle validation, and staging integration gates exist.
- [ ] Promotion timing and approvals match operational risk.
- [ ] Agents can run documented tests and validation locally.
- [ ] Logging and metrics are part of deployment.

## Common Mistakes

| Mistake | Fix |
| --- | --- |
| Repo split has no documented exception | Keep the Databricks default or document the local boundary |
| Bundle contains many unrelated domains | Split by ownership, lifecycle, or rollback boundary |
| Terraform and a bundle manage one object | Pick one owner before deployment |
| App bundle owns a shared schema | Move it to the documented platform owner or a foundation bundle |
| Dev and prod share writable catalogs | Mirror environments and isolate prod bindings |
| Developers read raw production PII | Use governed masked, synthetic, or controlled validation data |
| Notebook has core transformations | Extract module or SQL file, leave notebook thin |
| Every main merge deploys regulated prod | Use gated release-triggered promotion |
| Bundles exchange data without a contract | Define version, compatibility, ownership, and quality checks |
| Agents cannot verify changes | Document local tests, validation commands, and examples |
| Personal user runs production jobs | Use run-as service principals |
| Secrets appear in examples | Replace with secret references and rotate if exposed |

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

Source-derived guidance paraphrases "Developer best practices on Databricks." Community-informed alternatives are original additions. Credit Databricks for its guidance without attributing those alternatives to Databricks.
