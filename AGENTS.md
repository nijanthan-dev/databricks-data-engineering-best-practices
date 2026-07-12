# Agent Instructions

## Mandatory Codex Baseline

These rules are self-contained. Do not depend on or assume `~/.codex/AGENTS.md` was loaded.

- Read this file before planning, tool use, edits, or Git/GitHub actions.
- Be extremely concise in interactions and commit messages.
- Diagnose first. Once cause is clear, stop broad exploration and execute.
- Batch independent reads/searches. Prefer `rg`, filenames, counts, selected lines, summaries, or scripts over full dumps.
- Default each tool result to 100 lines or 20 KB maximum. Expand only relevant sections. Store noisy logs in temporary files and return summaries plus exceptions.
- Request only needed API/CLI fields. If output truncates, narrow the query; never rerun unchanged with a larger cap.
- Exceed output limits when required for correctness, security, or debugging.
- Do not reread unchanged files, logs, PR state, or tool output. Recheck only when state may have changed or a required gate demands it.
- Start with zero skills. Load exactly one only when clearly applicable, explicitly requested, or required. Prefer the narrowest skill; never load precautionary or overlapping skills.
- Before loading a second skill, state the distinct missing capability and why current tools/skill cannot cover it. Maximum two unless user or higher-priority instructions require more.
- At first skill use, report `Skills: <name> — <reason>`. Final response must report every skill used; report `Skills: none` if none.
- Use subagents only for independent work. Give narrow briefs and minimal history. Skip delegation when coordination costs exceed the work.
- Prefer deterministic scripts for repeated audits, polling, parsing, and verification. Return aggregates plus exceptions, not raw records.
- Keep the strongest model for ambiguous, high-risk, architecture, debugging, and review work. Use cheaper/faster models only for bounded mechanical work with deterministic validation.
- Never reduce required checks, security review, test coverage, or live-state verification to save tokens.
- Before editing, verify cwd, branch, and worktree match the task. Preserve unrelated user changes.
- Use standard branch prefixes. For parallel/automation work, prefer a fresh worktree from current `origin/main`.
- Use GitHub CLI for GitHub. Never run `gh auth refresh` unless explicitly requested. If a scope is missing, report it and stop.
- Before acting on PR review state, recheck head, latest reviews, unresolved threads, and checks live. Eyes/ack reactions are not approval.
- Fix valid review comments, reply with the fix, resolve the thread, and scan for the same bug class before re-review.
- Before merge, run Gitleaks on local history/current tree and PR diff; inspect remote PR state for secrets. If a secret reached remote Git, stop and remove it from branch history.
- Squash merge only. After merge, verify PR/issue closure and relevant main CI, update main, and clean safe merged branches/worktrees.
- Planning/issue-only requests stop at planning/issues unless implementation is explicit. End plans with extremely concise unresolved questions, if any.
- Fix issues introduced by current work and failures blocking required validation. Report unrelated pre-existing failures clearly; do not expand scope to fix them unless the user authorizes it.
- Never print, log, commit, screenshot, or comment secrets. Do not ask users to paste raw secrets.

## Repository Scope

- Keep guidance explicit about repository boundaries and ownership.
- Default shared, long-lived resources and schemas to Terraform.
- Never let Terraform and Databricks Asset Bundles own the same object.
- Treat bundle-destroy behavior as a required design review question.
- Keep changes verifiable by agents through documented local commands and focused checks.
