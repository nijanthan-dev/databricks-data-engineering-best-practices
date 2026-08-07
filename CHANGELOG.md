# Changelog

## 1.2.0 - 2026-08-08

- Renamed the user-facing skill title to Databricks Developer Best Practices.
- Rechecked the official Databricks source; it remains last updated 2026-07-10.
- Kept the technical skill ID stable for clean Agent Skills installation.

## 1.1.0 - 2026-08-01

- Refreshed source metadata and expanded workspace isolation, bundle lifecycle, developer tooling, and testing guidance from the Databricks page updated 2026-07-10.

## 1.0.0 - 2026-07-10

- First stable release.
- Clarified repository boundaries and project-template alternatives.
- Defined Terraform, bundle, schema, and lifecycle ownership rules.
- Added governed PII access, inter-bundle contracts, environment overrides, and gated deployment guidance.
- Added agent-friendly repository guidance and focused evaluation scenarios.
- Incorporated community feedback through original, source-credit-safe wording.
- Distinguished official Databricks defaults from community-informed operating alternatives.

## 0.1.0 - 2026-07-09

- Initial installable Agent Skill.
- Added Databricks source credit.
- Added validation and skill evaluation notes.
- Replaced Python validator with typed Go validator and tests.
- Added local pre-commit validation.
- Strengthened local Go validation with vet and race/shuffle tests.
- Replaced manual install docs with tested `npx skills add` commands.
- Added Databricks Genie Agent setup guidance based on official docs.
