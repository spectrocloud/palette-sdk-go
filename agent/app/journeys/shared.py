"""Reusable journey building blocks — pick-or-create patterns."""

from __future__ import annotations

import json
import logging
import time
from typing import Any

from ..protocol import DebugMessage, SelectMessage, TextMessage
from ..session import Session
from ..tools.base import PaletteTool

logger = logging.getLogger(__name__)


def ensure_cloud_account(
    provider: str,
    tools: dict[str, PaletteTool],
    session: Session,
) -> dict | None:
    """List cloud accounts, let user pick one, or create new. Returns selected account dict."""
    tool = tools.get("cloud_account_api")
    if not tool:
        return None

    session.send_sync(DebugMessage(
        event="journey_step", data={"action": f"ensure_{provider}_cloud_account"}, timestamp=time.time(),
    ))

    # Fetch accounts
    tool.silent = True
    raw = tool._run(json.dumps({"mode": "list_accounts"}))
    tool.silent = False

    accounts = _extract_list(raw)
    if not accounts:
        session.send_sync(TextMessage(
            content=f"No {provider.upper()} cloud accounts found. Please create one in the Palette UI and try again.",
            agent="Palette Assistant",
        ))
        return None

    # Build options with "Create new" at the end
    options = []
    for acc in accounts:
        label = acc.get("name", acc.get("uid", "unknown"))
        kind = acc.get("kind", "")
        if kind:
            label += f" ({kind})"
        options.append({"value": json.dumps(acc), "label": label})

    select_id = f"pick_account_{int(time.time())}"
    session.send_sync(SelectMessage(
        id=select_id,
        title=f"Select {provider.upper()} Cloud Account",
        description="Choose the cloud account for this deployment",
        options=options,
        agent="Palette Assistant",
    ))

    selected = session.wait_for_input_sync(select_id)
    try:
        return json.loads(selected) if isinstance(selected, str) else selected
    except (json.JSONDecodeError, TypeError):
        return {"name": str(selected)}


def ensure_cluster_profile(
    cloud_type: str,
    tools: dict[str, PaletteTool],
    session: Session,
) -> dict | None:
    """List profiles for a cloud type, let user pick. Returns selected profile dict."""
    tool = tools.get("profile_api")
    if not tool:
        return None

    session.send_sync(DebugMessage(
        event="journey_step", data={"action": f"ensure_{cloud_type}_profile"}, timestamp=time.time(),
    ))

    # Search for matching profiles
    tool.silent = True
    raw = tool._run(json.dumps({
        "mode": "search_profiles",
        "filter": {"filter": {"environment": [cloud_type], "profileType": ["cluster"]}},
    }))
    tool.silent = False

    profiles = _extract_list(raw)
    if not profiles:
        session.send_sync(TextMessage(
            content=(
                f"No {cloud_type.upper()}-compatible infra profiles found.\n\n"
                "To deploy a cluster, you need an infra profile with the right cloud type. "
                "Please create one in the Palette UI:\n"
                "1. Go to **Profiles** → **Create Profile**\n"
                "2. Select type **Infra**\n"
                f"3. Choose cloud type **{cloud_type.upper()}**\n"
                "4. Select packs for OS, Kubernetes, CNI, and CSI layers\n"
                "5. Publish the profile\n\n"
                "Then come back and try again."
            ),
            agent="Palette Assistant",
        ))
        return None

    options = []
    for p in profiles:
        label = p.get("name", p.get("uid", "unknown"))
        version = p.get("version", "")
        if version:
            label += f" v{version}"
        options.append({"value": json.dumps(p), "label": label})

    select_id = f"pick_profile_{int(time.time())}"
    session.send_sync(SelectMessage(
        id=select_id,
        title=f"Select {cloud_type.upper()} Cluster Profile",
        description="Choose the infra profile for this deployment",
        options=options,
        agent="Palette Assistant",
    ))

    selected = session.wait_for_input_sync(select_id)
    try:
        return json.loads(selected) if isinstance(selected, str) else selected
    except (json.JSONDecodeError, TypeError):
        return {"name": str(selected)}


def collect_form(
    title: str,
    fields: list[dict],
    tool: PaletteTool,
    session: Session,
) -> dict | None:
    """Show a form and wait for user input. Returns collected data."""
    from ..protocol import FormMessage, FormField, FieldType

    form_fields = []
    for f in fields:
        ft = FieldType.STRING
        if f.get("type") == "number":
            ft = FieldType.NUMBER
        elif f.get("type") == "boolean":
            ft = FieldType.BOOLEAN
        elif f.get("type") == "password":
            ft = FieldType.PASSWORD
        form_fields.append(FormField(
            name=f["name"],
            label=f.get("label", f["name"].replace("_", " ").title()),
            type=ft,
            required=f.get("required", False),
            default=f.get("default"),
            description=f.get("description", ""),
        ))

    form_id = f"form_{int(time.time())}"
    session.send_sync(FormMessage(
        id=form_id,
        title=title,
        fields=form_fields,
        submit_label="Continue",
        agent="Palette Assistant",
    ))

    return session.wait_for_input_sync(form_id)


def confirm_action(
    title: str,
    description: str,
    session: Session,
    destructive: bool = False,
) -> bool:
    """Show confirmation dialog. Returns True if confirmed."""
    from ..protocol import ConfirmMessage

    confirm_id = f"confirm_{int(time.time())}"
    session.send_sync(ConfirmMessage(
        id=confirm_id,
        title=title,
        description=description,
        destructive=destructive,
        agent="Palette Assistant",
    ))

    return bool(session.wait_for_input_sync(confirm_id))


def _extract_list(raw: str) -> list[dict]:
    """Extract a flat list of dicts from a tool response."""
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
        # Flatten each item
        return [_flatten(r) for r in data]
    if isinstance(data, dict):
        if "items" in data and isinstance(data["items"], list):
            return [_flatten(r) for r in data["items"]]
        for v in data.values():
            found = _find_list(v, depth + 1)
            if found:
                return found
    return []


_SKIP = {"annotations", "labels", "creationTimestamp", "deletionTimestamp",
         "lastModifiedTimestamp", "permissions", "ownerUid", "tenantUid",
         "scopeVisibility", "projectUid", "scope", "packs", "draft", "logoUrl"}


def _flatten(obj: dict, depth: int = 0) -> dict:
    if depth > 5:
        return {}
    flat: dict[str, Any] = {}
    for k, v in obj.items():
        if k in _SKIP:
            continue
        if isinstance(v, (str, int, float, bool)):
            if v != "" and v is not None:
                flat[k] = v
        elif isinstance(v, dict):
            flat.update(_flatten(v, depth + 1))
    return flat
