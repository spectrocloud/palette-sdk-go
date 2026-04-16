"""Defaults loader — reads from manifests/defaults/ and merges hierarchically.

Usage:
    from app.defaults import load_defaults

    d = load_defaults("eks")
    d["packs"]["os"]          # "amazon-linux-eks"
    d["cluster"]["region"]    # "us-east-1"
    d["preferred_registry"]   # "Palette Repo" (from global)
    d["backup"]["schedule"]   # "0 2 * * *" (from global)
"""

from __future__ import annotations

import logging
from pathlib import Path
from typing import Any

import yaml

logger = logging.getLogger(__name__)

_DEFAULTS_DIR = Path(__file__).resolve().parent.parent.parent / "manifests" / "defaults"

# Cache
_cache: dict[str, dict] = {}


def _deep_merge(base: dict, override: dict) -> dict:
    """Merge override into base recursively. Override wins on conflicts."""
    result = base.copy()
    for key, value in override.items():
        if key in result and isinstance(result[key], dict) and isinstance(value, dict):
            result[key] = _deep_merge(result[key], value)
        else:
            result[key] = value
    return result


def _load_yaml(path: Path) -> dict:
    if not path.exists():
        return {}
    try:
        return yaml.safe_load(path.read_text()) or {}
    except Exception as e:
        logger.warning("Failed to load %s: %s", path, e)
        return {}


def load_defaults(cloud_type: str) -> dict[str, Any]:
    """Load merged defaults for a cloud type.

    Merge order (later wins):
    1. global.yaml
    2. shared/*.yaml (all shared files)
    3. providers/{cloud_type}.yaml
    """
    if cloud_type in _cache:
        return _cache[cloud_type]

    # 1. Global
    result = _load_yaml(_DEFAULTS_DIR / "global.yaml")

    # 2. Shared (merge all)
    shared_dir = _DEFAULTS_DIR / "shared"
    if shared_dir.exists():
        for f in sorted(shared_dir.glob("*.yaml")):
            shared = _load_yaml(f)
            result = _deep_merge(result, shared)

    # 3. Provider-specific
    provider = _load_yaml(_DEFAULTS_DIR / "providers" / f"{cloud_type}.yaml")
    result = _deep_merge(result, provider)

    _cache[cloud_type] = result
    return result


def get_pack_defaults(cloud_type: str) -> dict[str, str]:
    """Get preferred pack names per layer for a cloud type."""
    d = load_defaults(cloud_type)
    return d.get("packs", {})


def get_cluster_defaults(cloud_type: str) -> dict[str, Any]:
    """Get cluster configuration defaults for a cloud type."""
    d = load_defaults(cloud_type)
    return d.get("cluster", {})


def get_preferred_registry() -> str:
    """Get the preferred registry name."""
    d = load_defaults("_global_only_")
    if not d.get("preferred_registry"):
        d = _load_yaml(_DEFAULTS_DIR / "global.yaml")
    return d.get("preferred_registry", "Palette Repo")


def get_naming_template(resource: str, cloud_type: str) -> str:
    """Get naming template for a resource type."""
    d = load_defaults(cloud_type)
    naming = d.get("naming", {})
    template = naming.get(resource, f"{cloud_type}-{resource}")
    return template.replace("{cloud_type}", cloud_type)
