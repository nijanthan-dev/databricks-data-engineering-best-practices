#!/usr/bin/env python3
from pathlib import Path
import re
import sys

ROOT = Path(__file__).resolve().parents[1]
SKILL_DIR = ROOT / "skills" / "databricks-data-engineering-best-practices"
SKILL = SKILL_DIR / "SKILL.md"


def fail(message: str) -> None:
    print(f"FAIL: {message}")
    sys.exit(1)


text = SKILL.read_text(encoding="utf-8")
if not text.startswith("---\n"):
    fail("missing YAML frontmatter")

parts = text.split("---\n", 2)
if len(parts) < 3:
    fail("frontmatter not closed")

frontmatter = parts[1]
body = parts[2]

name_match = re.search(r"^name:\s*([a-z0-9-]+)\s*$", frontmatter, re.M)
if not name_match:
    fail("invalid or missing name")

name = name_match.group(1)
if name != SKILL_DIR.name:
    fail("name must match folder")

if len(name) > 64 or name.startswith("-") or name.endswith("-") or "--" in name:
    fail("name violates spec")

description_match = re.search(r"^description:\s*(.+)\s*$", frontmatter, re.M)
if not description_match:
    fail("missing description")

description = description_match.group(1).strip()
if len(description) > 1024:
    fail("description over 1024 chars")
if not description.startswith("Use when "):
    fail("description must start with 'Use when '")

required_terms = [
    "Declarative Automation Bundles",
    "Unity Catalog",
    "service principals",
    "OIDC",
    "Lakeflow Spark Declarative Pipelines",
    "Databricks",
]
missing = [term for term in required_terms if term not in text]
if missing:
    fail("missing terms: " + ", ".join(missing))

if "https://docs.databricks.com/aws/en/developers/best-practices" not in text:
    fail("missing Databricks source URL")

if len(body.splitlines()) > 500:
    fail("SKILL.md body too long")

print("PASS: skill valid")
