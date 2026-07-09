# Skill Evaluation

This follows the writing-skills RED/GREEN/REFACTOR pattern for a reference skill.

## RED scenarios

Baseline behaviors to guard against:

1. Agent reviews a Databricks repo and only checks notebooks, missing bundles, workspace isolation, identities, and CI/CD.
2. Agent suggests Terraform for every Databricks resource, ignoring bundle-first deployment.
3. Agent accepts notebook-heavy business logic because "it works in the workspace."
4. Agent says "no issue" when prod jobs run as a personal user and secrets are in examples.

## Expected baseline failures

- Narrow notebook-only review.
- Generic cloud IaC advice not specific to Databricks bundles.
- No environment isolation checks.
- No source credit.
- No production testing or observability gate.

## GREEN checks

With the skill loaded, the agent should:

- Name Databricks-specific surfaces: Declarative Automation Bundles, Unity Catalog, personal schemas, service principals, OIDC, dynamic task values, Lakeflow Spark Declarative Pipelines.
- Recommend bundles for Databricks workspace resources and Terraform for cloud/admin resources.
- Flag notebook-heavy business logic and recommend `src/` modules or `.sql` files.
- Require unit, bundle validation, and staging integration gates.
- Credit Databricks when publishing guidance.

## Prompt

```text
Review this Databricks repo design: all logic is in notebooks, prod jobs run as Alice, CI only runs yamllint, Terraform creates jobs, dev and prod share a catalog, and secrets appear in README examples. Give high-priority findings and fixes.
```

Pass condition: output flags all major failures above and proposes Databricks-specific fixes.
