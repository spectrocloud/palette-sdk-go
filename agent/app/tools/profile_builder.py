"""Profile builder tool — fetches available packs and presents a single form
with pre-populated dropdowns and smart defaults for creating a cluster profile."""

from __future__ import annotations

import json
import logging
import time
from typing import Any

from ..config import settings
from ..defaults import get_pack_defaults, get_preferred_registry, get_naming_template, load_defaults
from ..protocol import (
    DebugMessage,
    FieldType,
    FormField,
    FormMessage,
    TextMessage,
)
from ..session import Session
from .base import PaletteTool

logger = logging.getLogger(__name__)


def build_profile_form(
    cloud_type: str,
    tools: dict[str, PaletteTool],
    session: Session,
    edge_mode: str | None = None,
) -> dict | None:
    """Fetch packs for each layer, build a single form with defaults, return collected data.
    For edge-native, edge_mode can be 'agent' or 'appliance' to select the right OS pack."""

    profile_tool = tools.get("profile_api")
    if not profile_tool:
        session.send_sync(TextMessage(content="profile_api tool not available.", agent="Palette Assistant"))
        return None

    session.send_sync(DebugMessage(
        event="profile_builder",
        data={"action": "fetching_packs", "cloud_type": cloud_type},
        timestamp=time.time(),
    ))

    session.send_sync(TextMessage(
        content=f"Fetching available packs for **{cloud_type.upper()}** profiles...",
        agent="Palette Assistant",
    ))

    # Determine which layers to include from defaults (skip layers not in defaults)
    all_layers = ["os", "k8s", "cni", "csi"]
    pack_defaults = dict(get_pack_defaults(cloud_type))
    layers = [l for l in all_layers if l in pack_defaults]
    if not layers:
        layers = all_layers  # fallback if no defaults at all

    # Fetch packs for each layer
    layer_packs: dict[str, list[dict]] = {}

    for layer in layers:
        profile_tool.silent = True
        search_input = json.dumps({
            "mode": "search_packs",
            "filter": {"type": ["spectro"], "layer": [layer], "environment": [cloud_type]},
            "sort": [],
        })
        logger.info("Profile builder: searching %s packs with %s", layer, search_input[:150])
        raw = profile_tool._run(search_input)
        profile_tool.silent = False

        all_packs = _extract_packs(raw)

        # Only show packs that explicitly support this cloud type
        # Exclude packs that only have "all" — they're too generic
        packs = [p for p in all_packs if cloud_type in p.get("cloudTypes", [])]

        # Sort: packs that ONLY support this cloud type first (most specific),
        # then packs that support multiple types
        packs.sort(key=lambda p: len(p.get("cloudTypes", [])))

        layer_packs[layer] = packs
        logger.info("Profile builder: %s layer → %d packs for %s", layer, len(packs), cloud_type)

    # Get defaults from manifests/defaults/
    defaults = dict(get_pack_defaults(cloud_type))  # copy so we can override

    # For edge, override OS default based on agent vs appliance mode
    if cloud_type == "edge-native" and edge_mode:
        if edge_mode == "agent":
            defaults["os"] = "byoi-agent-mode"
        else:  # appliance (default)
            defaults["os"] = "edge-native-byoi"
    elif cloud_type == "edge-native" and not edge_mode:
        # Default to appliance mode OS
        defaults["os"] = "edge-native-byoi"

    logger.info("Profile builder: defaults for %s (edge_mode=%s) = %s", cloud_type, edge_mode, defaults)

    # Build form fields
    fields = []

    # Profile name from naming template
    all_defaults = load_defaults(cloud_type)
    suggested_name = get_naming_template("profile", cloud_type)
    profile_version = all_defaults.get("profile", {}).get("version", "1.0.0")
    fields.append(FormField(
        name="profile_name",
        label="Profile Name",
        type=FieldType.STRING,
        required=True,
        default=suggested_name,
        description=f"Name for your new {cloud_type.upper()} infra profile",
    ))

    # Version
    fields.append(FormField(
        name="version",
        label="Profile Version",
        type=FieldType.STRING,
        required=True,
        default=profile_version,
        description="Semantic version for this profile",
    ))

    # Pack selection for each layer
    layer_labels = {
        "os": "Operating System",
        "k8s": "Kubernetes Distribution",
        "cni": "CNI (Networking)",
        "csi": "CSI (Storage)",
    }

    for layer in layers:
        packs = layer_packs.get(layer, [])
        if not packs:
            fields.append(FormField(
                name=f"{layer}_pack",
                label=f"{layer_labels.get(layer, layer)} Pack",
                type=FieldType.STRING,
                required=True,
                description=f"No {layer} packs found for {cloud_type}. Enter pack name manually.",
            ))
            continue

        # Build enum options: "name|version|uid" (value), display as "DisplayName vVersion"
        options = []
        for p in packs:
            name = p.get("name", "unknown")
            display = p.get("displayName", name)
            version = p.get("version", "latest")
            uid = p.get("uid", "")
            value = f"{name}|{version}|{uid}"
            options.append({"value": value, "label": f"{display} v{version}", "name": name})

        # Pick default from manifests/defaults/providers/{cloud_type}.yaml
        preferred = defaults.get(layer, "")
        preferred_version = defaults.get(f"{layer}_version", "")
        default = ""
        if preferred:
            # If a specific version is requested, match name+version first
            if preferred_version:
                for o in options:
                    if o["name"] == preferred and preferred_version in o["value"]:
                        default = o["value"]
                        break
            # Fall back to matching by name only
            if not default:
                for o in options:
                    if o["name"] == preferred:
                        default = o["value"]
                        break
        if not default and options:
            default = options[0]["value"]

        enum_values = [o["value"] for o in options]

        fields.append(FormField(
            name=f"{layer}_pack",
            label=f"{layer_labels.get(layer, layer)}",
            type=FieldType.SELECT,
            required=True,
            default=default,
            enum=enum_values,
            description=f"Select the {layer_labels.get(layer, layer).lower()} pack",
        ))

    # Send the form
    form_id = f"profile_builder_{int(time.time())}"
    session.register_pending(form_id)

    session.send_sync(FormMessage(
        id=form_id,
        title=f"Create {cloud_type.upper()} Infra Profile",
        description="Review the defaults and adjust as needed. All layers are pre-filtered for your cloud type.",
        fields=fields,
        submit_label="Create Profile",
        agent="Palette Assistant",
    ))

    logger.info("Profile builder: waiting for form submission id=%s", form_id)
    form_data = session.wait_for_input_sync(form_id, timeout=600.0)
    logger.info("Profile builder: form submitted: %s", str(form_data)[:200])

    return form_data


def create_profile_from_form(
    cloud_type: str,
    form_data: dict,
    tools: dict[str, PaletteTool],
    session: Session,
) -> dict | None:
    """Create and publish a profile from the builder form data."""

    profile_tool = tools.get("profile_api")
    if not profile_tool:
        return None

    profile_name = form_data.get("profile_name", f"{cloud_type}-infra-profile")
    version = form_data.get("version", "1.0.0")

    # Parse pack selections: "name|version|uid"
    packs = []
    for layer in ["os", "k8s", "cni", "csi"]:
        pack_value = form_data.get(f"{layer}_pack", "")
        if "|" in str(pack_value):
            parts = str(pack_value).split("|")
            packs.append({
                "name": parts[0],
                "tag": parts[1] if len(parts) > 1 else "latest",
                "uid": parts[2] if len(parts) > 2 else "",
                "type": "spectro",
            })
        elif pack_value:
            packs.append({
                "name": str(pack_value),
                "tag": "latest",
                "type": "spectro",
            })

    if len(packs) < 3:
        session.send_sync(TextMessage(
            content="Error: At least 3 layers (OS, Kubernetes, CNI) are required.",
            agent="Palette Assistant",
        ))
        return None

    # Fetch pack values for each selected pack
    session.send_sync(TextMessage(
        content="Fetching pack configurations...",
        agent="Palette Assistant",
    ))
    for pack in packs:
        pack_uid = pack.get("uid", "")
        if pack_uid:
            profile_tool.silent = True
            raw = profile_tool._run(json.dumps({"mode": "get_pack", "uid": pack_uid}))
            profile_tool.silent = False
            try:
                result = json.loads(raw)
                pack_data = result.get("data", {}).get("pack", {})
                pack_values = pack_data.get("packValues", [])
                # Find the values entry matching this pack UID
                for pv in pack_values:
                    if pv.get("packUid") == pack_uid or pv.get("values"):
                        values = pv.get("values", "")
                        if values:
                            pack["values"] = values
                            logger.info("Pack %s: loaded %d chars of values", pack["name"], len(values))
                            break
                # Also carry forward annotations from pack values
                for pv in pack_values:
                    if pv.get("annotations"):
                        pack.setdefault("annotations", {}).update(pv["annotations"])
            except Exception as e:
                logger.warning("Failed to fetch values for pack %s: %s", pack["name"], e)

    # Patch OS pack values for agent mode: set system.uri to "NA"
    # Without this, the profile save will fail for BYOOS/agent-mode OS packs.
    for pack in packs:
        if pack.get("name", "") in ("edge-native-byoi", "byoi-agent-mode") and pack.get("values"):
            values = pack["values"]
            if "system.uri" in values:
                # Replace empty system.uri with "NA"
                import re as _re
                values = _re.sub(
                    r'(system\.uri:\s*)""\s*$',
                    r'\1"NA"',
                    values,
                    flags=_re.MULTILINE,
                )
                # Also handle system.uri: (no value)
                values = _re.sub(
                    r'(system\.uri:\s*)$',
                    r'\1"NA"',
                    values,
                    flags=_re.MULTILINE,
                )
                pack["values"] = values
                logger.info("Patched system.uri to NA for pack %s", pack["name"])

    # Create profile
    session.send_sync(TextMessage(
        content=f"Creating profile **{profile_name}**...",
        agent="Palette Assistant",
    ))

    create_input = {
        "mode": "create_profile",
        "profile": {
            "metadata": {"name": profile_name},
            "spec": {
                "version": version,
                "template": {
                    "cloudType": cloud_type,
                    "type": "cluster",
                    "packs": packs,
                },
            },
        },
    }

    raw = profile_tool._run(json.dumps(create_input))
    try:
        result = json.loads(raw)
    except Exception:
        result = {"status": "error", "message": raw[:200]}

    if result.get("status") == "error":
        session.send_sync(TextMessage(
            content=f"Failed to create profile: {result.get('message', 'unknown error')}",
            agent="Palette Assistant",
        ))
        return None

    profile_uid = result.get("data", {}).get("profile_uid", "")

    # Publish
    if profile_uid:
        session.send_sync(TextMessage(
            content="Publishing profile...",
            agent="Palette Assistant",
        ))
        profile_tool._run(json.dumps({"mode": "publish_profile", "uid": profile_uid}))

    session.send_sync(TextMessage(
        content=(
            f"Profile **{profile_name}** created and published successfully!\n"
            f"UID: `{profile_uid}`\n"
            f"Cloud type: {cloud_type.upper()}\n"
            f"Packs: {', '.join(p['name'] for p in packs)}"
        ),
        agent="Palette Assistant",
    ))

    return {"name": profile_name, "uid": profile_uid, "version": version}




def _extract_packs(raw: str) -> list[dict]:
    """Extract pack list with name, version, uid from tool response.
    Prefers Palette Registry when multiple registries exist."""
    try:
        result = json.loads(raw)
    except (json.JSONDecodeError, TypeError):
        return []

    data = result.get("data", {})
    items = _find_list(data)

    flat = []
    for item in items:
        spec = item.get("spec", item)
        name = spec.get("name", "")
        display_name = spec.get("displayName", name)
        layer = spec.get("layer", "")

        # Get UID and version — pick the registry with the latest version
        registries = spec.get("registries", [])
        uid = ""
        version = ""
        if registries:
            # Sort by version descending, pick the one with the latest
            best_reg = registries[0]
            best_ver = best_reg.get("latestVersion", "0")
            for r in registries:
                r_ver = r.get("latestVersion", "0")
                if _compare_versions(r_ver, best_ver) > 0:
                    best_reg = r
                    best_ver = r_ver
            uid = best_reg.get("latestPackUid", "")
            version = best_reg.get("latestVersion", "")

        cloud_types = spec.get("cloudTypes", [])

        if name:
            flat.append({
                "name": name,
                "displayName": display_name,
                "version": version,
                "uid": uid,
                "layer": layer,
                "cloudTypes": cloud_types,
            })
    return flat


def _compare_versions(a: str, b: str) -> int:
    """Compare two version strings. Returns >0 if a>b, <0 if a<b, 0 if equal."""
    def parts(v: str) -> list[int]:
        try:
            return [int(x) for x in v.split(".")]
        except (ValueError, AttributeError):
            return [0]
    pa, pb = parts(a), parts(b)
    for i in range(max(len(pa), len(pb))):
        va = pa[i] if i < len(pa) else 0
        vb = pb[i] if i < len(pb) else 0
        if va != vb:
            return va - vb
    return 0


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


def _flatten_pack(obj: dict, depth: int = 0) -> dict:
    if depth > 4:
        return {}
    flat: dict[str, Any] = {}
    for k, v in obj.items():
        if k in ("annotations", "labels", "logoUrl", "readme", "schema", "values",
                  "presets", "macros", "template"):
            continue
        if isinstance(v, (str, int, float, bool)):
            if v != "" and v is not None:
                flat[k] = v
        elif isinstance(v, dict):
            flat.update(_flatten_pack(v, depth + 1))
    return flat
