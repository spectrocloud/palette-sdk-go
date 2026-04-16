"""Journey preview tool — shows a step-by-step checklist with prereq status."""

from __future__ import annotations

import json
import logging
import time
from typing import Any

import httpx

from ..config import settings
from ..protocol import DebugMessage, TextMessage
from ..session import Session
from ..tools.base import PaletteTool

logger = logging.getLogger(__name__)


def preview_journey(
    journey_name: str,
    journey_skill: str,
    tools: dict[str, PaletteTool],
    session: Session,
) -> str:
    """Check prerequisites and show a step-by-step plan with status indicators."""

    session.send_sync(DebugMessage(
        event="journey_preview",
        data={"journey": journey_name},
        timestamp=time.time(),
    ))

    # Run prereq checks
    checks = [
        {
            "name": "AWS Cloud Account",
            "tool": "cloud_account_api",
            "input": {"mode": "list_accounts"},
            "check": lambda rows: len(rows) > 0,
            "found_msg": lambda rows: f"{len(rows)} account(s) found",
            "missing_msg": "No AWS accounts — you'll need to create one",
            "create_action": "Create via cloud_account_api (mode: create_aws)",
        },
        {
            "name": "EKS Infra Profile",
            "tool": "profile_api",
            "input": {"mode": "search_profiles", "filter": {"filter": {"environment": ["eks"], "profileType": ["cluster"]}}},
            "check": lambda rows: len(rows) > 0,
            "found_msg": lambda rows: f"{len(rows)} profile(s) found",
            "missing_msg": "No EKS infra profiles — you'll need to create one",
            "create_action": "Create in Palette UI: Profiles → Create → Infra → EKS",
        },
    ]

    # Build the checklist
    lines = [f"## Journey: Deploy EKS Cluster\n"]
    all_ready = True

    for i, check in enumerate(checks):
        tool = tools.get(check["tool"])
        if not tool:
            lines.append(f"  {i+1}. [ ] **{check['name']}** — tool `{check['tool']}` not available")
            all_ready = False
            continue

        # Call tool silently
        tool.silent = True
        try:
            raw = tool._run(json.dumps(check["input"]))
            rows = _extract_list(raw)
            passed = check["check"](rows)
        except Exception as e:
            logger.warning("Prereq check %s failed: %s", check["name"], e)
            rows = []
            passed = False
        finally:
            tool.silent = False

        if passed:
            msg = check["found_msg"](rows)
            lines.append(f"  {i+1}. [x] **{check['name']}** — {msg}")
        else:
            msg = check["missing_msg"]
            action = check["create_action"]
            lines.append(f"  {i+1}. [ ] **{check['name']}** — {msg}")
            lines.append(f"     → *Action needed:* {action}")
            all_ready = False

    # Journey steps
    lines.append("")
    lines.append("### Deployment Steps")
    lines.append("")

    steps = [
        ("Select cloud account", "Pick from existing accounts", True),
        ("Select cluster profile", "Pick an EKS-compatible infra profile", True),
        ("Configure cluster", "Provide: cluster name, AWS region, instance type, node count", False),
        ("Review & confirm", "Review all selections before deploying", False),
        ("Create cluster", "Call Palette API to create the EKS cluster", False),
    ]

    for i, (step, detail, is_prereq) in enumerate(steps):
        prefix = "[x]" if is_prereq and all_ready else "[ ]"
        lines.append(f"  {i+1}. {prefix} **{step}** — {detail}")

    lines.append("")
    if all_ready:
        lines.append("**All prerequisites met.** Ready to proceed with deployment.")
        lines.append('Say "proceed" or "go ahead" to start.')
    else:
        lines.append("**Some prerequisites are missing.** Resolve the items above before deploying.")
        lines.append("I can help you create the missing resources — just ask.")

    content = "\n".join(lines)
    session.send_sync(TextMessage(content=content, agent="Palette Assistant"))

    return json.dumps({"status": "ok", "message": "preview", "data": {"ready": all_ready}})


def _extract_list(raw: str) -> list[dict]:
    """Extract flat list from tool response."""
    try:
        result = json.loads(raw)
    except (json.JSONDecodeError, TypeError):
        return []

    data = result.get("data", {})
    return _find_list(data)


def _find_list(data: Any, depth: int = 0) -> list[dict]:
    if depth > 3:
        return []
    if isinstance(data, list) and data and isinstance(data[0], dict):
        return data
    if isinstance(data, dict):
        if "items" in data and isinstance(data["items"], list):
            return data["items"]
        for v in data.values():
            found = _find_list(v, depth + 1)
            if found:
                return found
    return []
