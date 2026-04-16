"""Palette tool — bridges to Go tool server with optional interactive UI.

When "interactive": true is passed in the input, the tool presents results
as a select/pick list and waits for the user's choice before returning.
Otherwise it returns data normally (and optionally renders a table in the UI).
"""

from __future__ import annotations

import json
import logging
import time
from typing import Any

import httpx
from crewai.tools import BaseTool

from ..config import settings
from ..protocol import (
    DebugMessage,
    ErrorMessage,
    FormField,
    FormMessage,
    FieldType,
    SelectMessage,
    ServerMessage,
    TableColumn,
    TableMessage,
    TextMessage,
)
from ..session import Session

logger = logging.getLogger(__name__)


class PaletteTool(BaseTool):
    """Generic tool that bridges to a Go tool server endpoint."""

    go_tool_name: str = ""
    session: Session | None = None
    silent: bool = False

    model_config = {"arbitrary_types_allowed": True}

    def _run(self, input_json: str = "{}") -> str:
        """Called by LLM orchestrator. POSTs to Go tool server.

        Special input fields (stripped before sending to Go):
        - "interactive": true — present results as pick list, wait for user selection
        - "client_filter": {...} — filter results client-side
        - "collect": {...fields} — show a form and wait for user input (no tool call)
        """
        # Parse input robustly
        if isinstance(input_json, dict):
            payload = input_json
        elif isinstance(input_json, str):
            s = input_json.strip()
            if s.startswith("```"):
                s = "\n".join(l for l in s.split("\n") if not l.startswith("```")).strip()
            try:
                payload = json.loads(s)
                if not isinstance(payload, dict):
                    payload = {}
            except (json.JSONDecodeError, TypeError):
                payload = {}
        else:
            payload = {}

        logger.info("Tool %s: payload=%s", self.go_tool_name, json.dumps(payload)[:200])

        # Extract special fields before sending to Go
        interactive = payload.pop("interactive", False)
        client_filter = payload.pop("client_filter", None)
        collect_fields = payload.pop("collect", None)
        build_profile = payload.pop("build_profile", None)
        build_account = payload.pop("build_account", None)

        # --- Build account mode: show credential form ---
        if build_account and self.session:
            from .account_builder import build_account_form
            provider = build_account if isinstance(build_account, str) else payload.get("provider", "aws")
            tools_dict = {"cloud_account_api": self}
            result = build_account_form(provider, tools_dict, self.session)
            if result:
                return json.dumps({"status": "ok", "message": f"Account created: {result.get('name')}", "data": {"account": result}})
            return json.dumps({"status": "error", "message": "Account creation cancelled or failed."})

        # --- Build profile mode: fetch packs + show combined form ---
        if build_profile and self.session:
            from .profile_builder import build_profile_form, create_profile_from_form
            cloud_type = build_profile if isinstance(build_profile, str) else payload.get("cloud_type", "eks")
            # Need access to tools dict — get it from session context or build locally
            tools_dict = {"profile_api": self}
            form_data = build_profile_form(cloud_type, tools_dict, self.session)
            if form_data:
                result = create_profile_from_form(cloud_type, form_data, tools_dict, self.session)
                if result:
                    return json.dumps({"status": "ok", "message": f"Profile created: {result.get('name')}", "data": {"profile": result}})
            return json.dumps({"status": "error", "message": "Profile creation cancelled or failed."})

        # --- Collect mode: show form, no tool call ---
        if collect_fields and self.session:
            return self._handle_collect(collect_fields)

        # Debug
        t0 = time.time()
        self._send(DebugMessage(
            event="tool_call",
            data={"tool": self.go_tool_name, "input": payload, "interactive": interactive},
            timestamp=t0,
        ))

        # Call Go tool server
        try:
            with httpx.Client(timeout=120.0) as client:
                resp = client.post(
                    f"{settings.tool_server_url}/tools/{self.go_tool_name}",
                    json=payload,
                )
                resp.raise_for_status()
                raw = resp.text
        except httpx.HTTPStatusError as e:
            error_msg = f"Tool server error ({e.response.status_code}): {e.response.text}"
            logger.error(error_msg)
            self._send(ErrorMessage(message=error_msg, agent=self.name))
            return json.dumps({"status": "error", "message": error_msg})
        except httpx.ConnectError:
            error_msg = f"Cannot connect to tool server at {settings.tool_server_url}"
            logger.error(error_msg)
            self._send(ErrorMessage(message=error_msg, agent=self.name))
            return json.dumps({"status": "error", "message": error_msg})

        # Parse response
        try:
            result = json.loads(raw)
        except json.JSONDecodeError:
            result = {"status": "ok", "message": raw[:500]}

        if result.get("status") == "error":
            self._send(ErrorMessage(message=result.get("message", "Unknown error"), agent=self.name))
            return raw

        # Find list data in result
        data = result.get("data", {})
        rows = self._find_list(data)
        filtered_rows: list[dict] | None = None

        if rows:
            flat_rows = [self._extract_scalars(r) for r in rows]
            flat_rows = [r for r in flat_rows if r]

            if client_filter and flat_rows:
                flat_rows = self._apply_filter(flat_rows, client_filter)
                filtered_rows = flat_rows

            # --- Interactive mode: present as pick list, wait for selection ---
            if interactive and flat_rows and self.session:
                return self._handle_interactive(flat_rows, result)

            # --- Normal mode: show table ---
            if flat_rows:
                self._send_table(flat_rows, result)
            else:
                self._send(TextMessage(
                    content="No results found.",
                    agent=self.name,
                ))

        else:
            # Check if the data contains an empty list or null values (no results)
            has_empty = any(
                v is None or (isinstance(v, list) and len(v) == 0)
                for v in data.values()
            ) if isinstance(data, dict) else False

            if has_empty:
                self._send(TextMessage(
                    content="No results found.",
                    agent=self.name,
                ))
            else:
                # Scalar/non-list data — send as text
                self._send(TextMessage(
                    content=result.get("message", json.dumps(result, indent=2)),
                    agent=self.name,
                ))

        # Debug
        elapsed = time.time() - t0
        self._send(DebugMessage(
            event="tool_result",
            data={"tool": self.go_tool_name, "status": "ok", "elapsed_sec": round(elapsed, 3)},
            timestamp=time.time(),
        ))

        # If client_filter was applied, return only filtered data for LLM interpretation
        if filtered_rows is not None:
            return json.dumps({
                "status": "ok",
                "message": result.get("message", ""),
                "data": {"items": filtered_rows},
            })

        return raw

    def _handle_interactive(self, flat_rows: list[dict], result: dict) -> str:
        """Present results as a select list and wait for user's choice."""
        # Build options — use 'name' as label, full row as value
        options = []
        for row in flat_rows:
            label = str(row.get("name", row.get("uid", str(row)[:50])))
            # Add extra info to label
            extras = [f"{k}={v}" for k, v in row.items() if k not in ("name", "uid") and v][:2]
            if extras:
                label += f" ({', '.join(extras)})"
            options.append({"value": json.dumps(row), "label": label})

        select_id = f"select_{self.go_tool_name}_{int(time.time())}"

        # Register BEFORE sending to avoid race condition
        self.session.register_pending(select_id)

        self._send(SelectMessage(
            id=select_id,
            title=result.get("message", "Select an option").replace(" completed.", "").replace("_", " ").title(),
            description=f"Choose from {len(options)} options",
            options=options,
            agent=self.name,
        ))

        # Wait for user selection (future already registered)
        logger.info("Tool %s: waiting for user selection on id=%s", self.go_tool_name, select_id)
        selected = self.session.wait_for_input_sync(select_id, timeout=600.0)
        logger.info("Tool %s: user selected=%s", self.go_tool_name, str(selected)[:100])

        # Parse selection back to dict
        try:
            selected_data = json.loads(selected) if isinstance(selected, str) else selected
        except (json.JSONDecodeError, TypeError):
            selected_data = selected

        # Return the selection as the tool result
        return json.dumps({
            "status": "ok",
            "message": f"Selected: {selected_data.get('name', selected) if isinstance(selected_data, dict) else selected}",
            "data": {"selected": selected_data},
        })

    def _handle_collect(self, fields_spec: dict) -> str:
        """Show a form and wait for user input. Returns collected data."""
        form_id = f"form_{self.go_tool_name}_{int(time.time())}"

        fields = []
        for name, spec in fields_spec.items():
            if isinstance(spec, str):
                spec = {"type": "string", "label": spec}
            ft = FieldType.STRING
            spec_type = spec.get("type", "string")
            if spec_type == "number":
                ft = FieldType.NUMBER
            elif spec_type == "boolean":
                ft = FieldType.BOOLEAN
            elif spec_type == "select":
                ft = FieldType.SELECT

            fields.append(FormField(
                name=name,
                label=spec.get("label", name.replace("_", " ").title()),
                type=ft,
                required=spec.get("required", False),
                default=spec.get("default"),
                description=spec.get("description", ""),
                enum=spec.get("enum"),
            ))

        self._send(FormMessage(
            id=form_id,
            title=fields_spec.get("_title", "Provide Details"),
            fields=fields,
            submit_label="Continue",
            agent=self.name,
        ))

        form_data = self.session.wait_for_input_sync(form_id)
        logger.info("Tool %s: form data=%s", self.go_tool_name, str(form_data)[:200])

        return json.dumps({
            "status": "ok",
            "message": "Form submitted.",
            "data": {"collected": form_data},
        })

    def _send_table(self, flat_rows: list[dict], result: dict) -> None:
        """Send a table message to the UI."""
        columns = list(flat_rows[0].keys())
        title = result.get("message", "Results").replace(" completed.", "")
        self._send(TableMessage(
            title=title.replace("_", " ").title(),
            columns=[
                TableColumn(key=c, label=c.replace("_", " ").title(), sortable=True)
                for c in columns
            ],
            rows=flat_rows,
            agent=self.name,
        ))

    def _send(self, msg: ServerMessage) -> None:
        if self.session:
            if self.silent and msg.type != "debug":
                return
            self.session.send_sync(msg)

    def _find_list(self, data: Any, depth: int = 0) -> list[dict] | None:
        if depth > 3:
            return None
        if isinstance(data, list) and len(data) > 0 and isinstance(data[0], dict):
            return data
        if isinstance(data, dict):
            if "items" in data and isinstance(data["items"], list):
                return data["items"]
            for value in data.values():
                found = self._find_list(value, depth + 1)
                if found:
                    return found
        return None

    def _extract_scalars(self, obj: dict) -> dict:
        flat: dict[str, Any] = {}
        self._collect_scalars(obj, flat)
        return flat

    def _apply_filter(self, rows: list[dict], filters: dict) -> list[dict]:
        return [r for r in rows if all(
            str(filters[k]).lower() in str(r.get(k, "")).lower()
            for k in filters
        )]

    _SKIP_FIELDS = {
        "annotations", "labels", "creationTimestamp", "deletionTimestamp",
        "lastModifiedTimestamp", "ownerReferences", "managedFields",
        "selfLink", "resourceVersion", "generation", "finalizers",
        "permissions", "ownerUid", "tenantUid", "scopeVisibility",
        "projectUid", "scope", "lastSyncedTimestamp",
        "packs", "draft", "logoUrl",
        "projectMeta", "cloudAccountMeta",
    }

    def _collect_scalars(self, obj: dict, flat: dict, depth: int = 0) -> None:
        if depth > 5:
            return
        for key, value in obj.items():
            if key in self._SKIP_FIELDS:
                continue
            if isinstance(value, (str, int, float, bool)):
                if value != "" and value is not None:
                    # Don't overwrite top-level fields (e.g. metadata.name) with
                    # nested duplicates (e.g. projectMeta.name).
                    if key not in flat or depth == 0:
                        flat[key] = value
            elif isinstance(value, dict):
                self._collect_scalars(value, flat, depth + 1)
            elif key == "registries" and isinstance(value, list) and value:
                # Extract latest version from registries array (pack search results)
                latest = ""
                for reg in value:
                    if isinstance(reg, dict):
                        v = reg.get("latestVersion", "")
                        if v and (not latest or v > latest):
                            latest = v
                if latest:
                    flat["latestVersion"] = latest


def create_tool(name: str, description: str, go_tool_name: str, session: Session | None = None) -> PaletteTool:
    return PaletteTool(
        name=name,
        description=f"{description}. See skill docs for exact input_json.",
        go_tool_name=go_tool_name,
        session=session,
    )
