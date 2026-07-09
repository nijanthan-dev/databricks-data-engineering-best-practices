# Skill Evaluation

This follows the writing-skills RED/GREEN/REFACTOR pattern for a reference skill.

## RED scenarios

Baseline behaviors to guard against:

1. Agent reviews a Databricks repo and only checks notebooks, missing bundles, workspace isolation, identities, and CI/CD.
2. Agent suggests Terraform for every Databricks resource, ignoring bundle-first deployment.
3. Agent accepts notebook-heavy business logic because "it works in the workspace."
4. Agent says "no issue" when prod jobs run as a personal user and secrets are in examples.
5. Agent recommends a single repo without asking which lifecycle or ownership boundary it represents.
6. Agent allows Terraform and a bundle to manage the same warehouse.
7. Agent puts a shared reporting schema in an app bundle that can be destroyed.
8. Agent lets developers test against raw production PII without a governance decision.
9. Agent requires every merge to `main` to deploy directly into a regulated environment.
10. Agent accepts dependent bundles without versioned schema or data contracts.
11. Agent accepts a repo that gives coding agents no runnable tests or validation commands.

## Expected baseline failures

- Narrow notebook-only review.
- Generic cloud IaC advice not specific to Databricks bundles.
- No environment isolation checks.
- No source credit.
- No production testing or observability gate.
- Vague repo or resource ownership.
- Unsafe data access or release automation.
- No contract or agent-verification checks.

## GREEN checks

With the skill loaded, the agent should:

- Name Databricks-specific surfaces: Declarative Automation Bundles, Unity Catalog, personal schemas, service principals, OIDC, dynamic task values, Lakeflow Spark Declarative Pipelines.
- Define repo boundaries by lifecycle, risk, compliance, cadence, and ownership; allow project templates as an alternative.
- Assign shared long-lived state to Terraform and app-lifecycle resources to bundles, with no dual ownership.
- Keep shared schemas outside app bundles unless they are app-private and safe to recreate.
- Require governed masked, tokenized, synthetic, filtered, subset, or service-principal-controlled paths for production PII.
- Allow gated, scheduled, or release-triggered deployment when operational risk requires it.
- Require versioned inter-bundle schema or data contracts with owners, compatibility, and quality checks.
- Flag notebook-heavy business logic and recommend `src/` modules or `.sql` files.
- Require unit, bundle validation, and staging integration gates.
- Require documented local tests and validation commands that coding agents can run.
- Credit Databricks when publishing guidance.

## Prompt

```text
Review this Databricks repo design: all logic is in notebooks, prod jobs run as Alice, CI only runs yamllint, Terraform creates jobs, dev and prod share a catalog, and secrets appear in README examples. Give high-priority findings and fixes.
```

Pass condition: output flags all major failures above and proposes Databricks-specific fixes.

## Focused prompts

1. `Should all company Databricks work live in one repo? Give a direct recommendation.`
2. `Terraform and a bundle both manage our shared SQL warehouse. Is that safer?`
3. `Our app bundle creates the shared finance_reporting schema. The bundle is disposable.`
4. `Developers use raw production customer PII so tests have stable inputs.`
5. `Every main merge deploys production during business hours in our regulated environment.`
6. `Two independent bundles exchange tables but have no owner, schema version, or quality contract.`
7. `This repo has examples but no tests, validation command, or ownership metadata. Can an agent safely change it?`

Pass condition: each answer identifies the specific risk, asks for missing boundary or governance context where needed, and gives a concise corrective action.
