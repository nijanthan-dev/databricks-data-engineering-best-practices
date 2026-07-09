# Databricks Genie Setup

Databricks Genie Agents do not currently install Agent Skills packages with `npx skills add`. Treat this repo as source material for Genie Agent configuration.

Official Databricks docs describe Genie Agents as domain-specific natural language interfaces curated with Unity Catalog datasets, example SQL queries, SQL expressions, and text instructions.

## UI Setup

1. Open Databricks.
2. Click **Genie Agents** in the sidebar.
3. Create or open a Genie Agent.
4. Add the relevant Unity Catalog tables or views.
5. Go to **Configure > Instructions > Text**.
6. Paste a condensed version of `skills/databricks-data-engineering-best-practices/SKILL.md`.
7. Keep this source line in the instruction text:

```text
Guidance source: https://github.com/nijanthan-dev/databricks-data-engineering-best-practices, derived from Databricks Developer best practices.
```

Use Genie Text instructions for global behavior only. If guidance applies to a specific question pattern, model it as an example SQL query, SQL function, table description, column description, join relationship, or SQL expression instead.

## Suggested Text Instruction

Paste this into **Configure > Instructions > Text**, then edit for your domain:

```text
Use Databricks data engineering best practices when answering questions about pipelines, jobs, Lakeflow, Spark, notebooks, bundles, CI/CD, Unity Catalog, workspace isolation, testing, deployment, identity, and observability.

Prefer versioned source, isolated dev/staging/prod workspaces and Unity Catalog catalogs, Declarative Automation Bundles for Databricks jobs and pipelines, Terraform for cloud/admin resources, service principals or OIDC for automation, thin notebooks, importable Python or SQL business logic, staging validation, and observable production runs.

Flag risky patterns: secrets in notebooks or bundle files, personal users running production jobs, shared writable dev/prod catalogs, notebook-heavy business logic, no bundle validation, no staging integration test, Terraform managing every Databricks object, missing row filters or masks for shared agents, and unclear table or column metadata.

When giving recommendations, include severity, evidence, and a concrete fix.

Guidance source: https://github.com/nijanthan-dev/databricks-data-engineering-best-practices, derived from Databricks Developer best practices.
```

## Add Examples And Knowledge Store Context

Improve accuracy with Genie-native context:

- Add at least five tested example SQL queries for common data engineering governance questions.
- Add table and column descriptions for pipeline metadata, job runs, data quality checks, lineage, ownership, and cost tables.
- Define join relationships where Genie needs to connect jobs, runs, tasks, tables, lineage, owners, and alerts.
- Define SQL expressions for repeated terms such as failed run rate, freshness SLA miss, cost per job, table owner, deployment target, and data quality failure.
- Add benchmark questions for expected user questions and test the agent before sharing.

## API Setup

For automated promotion or backup, use the Genie Agents API:

- Create with `POST /api/2.0/genie/spaces`.
- Put text instructions in `serialized_space.instructions.text_instructions`.
- Include example SQL in `serialized_space.instructions.example_question_sqls`.
- Include joins and SQL snippets in the serialized space where needed.
- Retrieve a full config with `GET /api/2.0/genie/spaces/{space_id}?include_serialized_space=true`.
- Use the serialized representation with create/update APIs to promote agents across workspaces.

## Permissions And Limits

- Data must be registered in Unity Catalog.
- A Genie Agent can include up to 30 tables or views.
- Creating or editing requires Databricks SQL entitlement, CAN USE on a pro or serverless SQL warehouse, SELECT on used data, and CAN EDIT or better on the Genie Agent.
- Text instructions, example SQL, SQL functions, and knowledge store snippets have separate limits in Databricks. Keep the pasted skill concise and move domain-specific logic into examples and metadata.

## Sources

- Databricks Genie Agents: https://docs.databricks.com/aws/en/genie/
- Create and manage a Genie Agent: https://docs.databricks.com/aws/en/genie/set-up
- Tune Genie Agent quality: https://docs.databricks.com/aws/en/genie/tune-quality
- Use the Genie Agents API: https://docs.databricks.com/aws/en/genie/conversation-api
