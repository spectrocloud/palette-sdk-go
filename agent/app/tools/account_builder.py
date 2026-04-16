"""Cloud account builder — presents a form for creating cloud accounts."""

from __future__ import annotations

import json
import logging
import time
from typing import Any

from ..defaults import load_defaults
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

# Provider-specific form fields
_PROVIDER_FIELDS: dict[str, list[dict]] = {
    "aws": {
        "modes": {
            "credentials": {
                "label": "Access Key & Secret",
                "description": "Static IAM access key credentials",
                "fields": [
                    {"name": "account_name", "label": "Account Name", "required": True},
                    {"name": "access_key_id", "label": "Access Key ID", "required": True},
                    {"name": "secret_access_key", "label": "Secret Access Key", "type": "password", "required": True},
                ],
            },
            "sts": {
                "label": "STS Assume Role",
                "description": "Cross-account access via STS assume role — Palette assumes the IAM role directly",
                "fields": [
                    {"name": "account_name", "label": "Account Name", "required": True},
                    {"name": "aws_account_id", "label": "AWS Account ID", "required": True, "description": "12-digit AWS account number"},
                    {"name": "arn", "label": "Role ARN", "required": True, "description": "ARN of the IAM role for Palette to assume"},
                    {"name": "external_id", "label": "External ID", "required": False, "description": "Optional external ID for assume role condition"},
                ],
            },
            "pod_identity": {
                "label": "Pod Identity (No Auth)",
                "description": "Uses EKS Pod Identity or IRSA — no credentials needed",
                "fields": [
                    {"name": "account_name", "label": "Account Name", "required": True},
                ],
            },
        },
    },
    "azure": [
        {"name": "account_name", "label": "Account Name", "required": True},
        {"name": "tenant_id", "label": "Tenant ID", "required": True},
        {"name": "client_id", "label": "Client ID", "required": True},
        {"name": "client_secret", "label": "Client Secret", "type": "password", "required": True},
        {"name": "subscription_id", "label": "Subscription ID", "required": True},
    ],
    "gcp": [
        {"name": "account_name", "label": "Account Name", "required": True},
        {"name": "json_credentials", "label": "Service Account JSON Key", "type": "textarea", "required": True,
         "description": "Paste the full JSON key file contents"},
    ],
    "vsphere": [
        {"name": "account_name", "label": "Account Name", "required": True},
        {"name": "vcenter_server", "label": "vCenter Server", "required": True},
        {"name": "username", "label": "Username", "required": True},
        {"name": "password", "label": "Password", "type": "password", "required": True},
    ],
    "maas": [
        {"name": "account_name", "label": "Account Name", "required": True},
        {"name": "api_endpoint", "label": "API Endpoint", "required": True, "description": "e.g., http://maas:5240/MAAS"},
        {"name": "api_key", "label": "API Key", "type": "password", "required": True},
    ],
}

# Maps provider to create mode and spec builder
_PROVIDER_CREATE: dict[str, dict] = {
    "aws": {
        "mode": "create_aws",
        "build_spec": lambda d: {
            "account": {
                "metadata": {"name": d["account_name"]},
                "spec": {
                    "accessKey": d.get("access_key_id", ""),
                    "secretKey": d.get("secret_access_key", ""),
                    "credentialType": (
                        "sts" if d.get("arn") else
                        "pod-identity" if not d.get("access_key_id") else
                        "secret"
                    ),
                    **({"sts": {"arn": d["arn"], **({"externalId": d["external_id"]} if d.get("external_id") else {})}} if d.get("arn") else {}),
                },
            }
        },
    },
    "azure": {
        "mode": "create_azure",
        "build_spec": lambda d: {
            "account": {
                "metadata": {"name": d["account_name"]},
                "spec": {
                    "tenantId": d["tenant_id"],
                    "clientId": d["client_id"],
                    "clientSecret": d["client_secret"],
                    "subscriptionId": d["subscription_id"],
                },
            }
        },
    },
    "gcp": {
        "mode": "create_gcp",
        "build_spec": lambda d: {
            "account": {
                "metadata": {"name": d["account_name"]},
                "spec": {
                    "jsonCredentials": d["json_credentials"],
                },
            }
        },
    },
    "vsphere": {
        "mode": "create_vsphere",
        "build_spec": lambda d: {
            "account": {
                "metadata": {"name": d["account_name"]},
                "spec": {
                    "vcenterServer": d["vcenter_server"],
                    "username": d["username"],
                    "password": d["password"],
                },
            }
        },
    },
    "maas": {
        "mode": "create_maas",
        "build_spec": lambda d: {
            "account": {
                "metadata": {"name": d["account_name"]},
                "spec": {
                    "apiEndpoint": d["api_endpoint"],
                    "apiKey": d["api_key"],
                },
            }
        },
    },
}


def build_account_form(
    provider: str,
    tools: dict[str, PaletteTool],
    session: Session,
) -> dict | None:
    """Show a form for creating a cloud account. Returns the created account data."""

    account_tool = tools.get("cloud_account_api")
    if not account_tool:
        session.send_sync(TextMessage(content="cloud_account_api tool not available.", agent="Palette Assistant"))
        return None

    # Map EKS/AKS/GKE to their parent providers for account creation
    provider_map = {"eks": "aws", "aks": "azure", "gke": "gcp"}
    account_provider = provider_map.get(provider, provider)

    provider_config = _PROVIDER_FIELDS.get(account_provider)
    if not provider_config:
        session.send_sync(TextMessage(
            content=f"Account creation for provider '{account_provider}' is not yet supported.",
            agent="Palette Assistant",
        ))
        return None

    defaults = load_defaults(provider)
    suggested_name = f"{account_provider}-account"

    # Check if provider has multiple auth modes
    if isinstance(provider_config, dict) and "modes" in provider_config:
        # Step 1: let user pick auth mode
        from ..protocol import SelectMessage

        modes = provider_config["modes"]
        options = [
            {"value": mode_key, "label": f"{mode_def['label']} — {mode_def['description']}"}
            for mode_key, mode_def in modes.items()
        ]

        select_id = f"account_mode_{int(time.time())}"
        session.register_pending(select_id)

        session.send_sync(SelectMessage(
            id=select_id,
            title=f"{account_provider.upper()} Authentication Mode",
            description="Choose how to authenticate with the cloud provider",
            options=options,
            agent="Palette Assistant",
        ))

        selected_mode = session.wait_for_input_sync(select_id, timeout=600.0)
        if not selected_mode:
            return None

        logger.info("Account builder: auth mode=%s", selected_mode)
        field_defs = modes.get(selected_mode, {}).get("fields", [])
    elif isinstance(provider_config, list):
        # Simple provider — single set of fields
        field_defs = provider_config
    else:
        return None

    # Step 2: show credential form
    fields = []
    for f in field_defs:
        ft = FieldType.STRING
        if f.get("type") == "password":
            ft = FieldType.PASSWORD
        elif f.get("type") == "textarea":
            ft = FieldType.TEXTAREA
        default = suggested_name if f["name"] == "account_name" else f.get("default", "")
        fields.append(FormField(
            name=f["name"],
            label=f["label"],
            type=ft,
            required=f.get("required", False),
            default=default,
            description=f.get("description", ""),
        ))

    form_id = f"account_builder_{int(time.time())}"
    session.register_pending(form_id)

    session.send_sync(FormMessage(
        id=form_id,
        title=f"Create {account_provider.upper()} Cloud Account",
        description="Enter your cloud provider credentials.",
        fields=fields,
        submit_label="Create Account",
        agent="Palette Assistant",
    ))

    logger.info("Account builder: waiting for form submission id=%s", form_id)
    form_data = session.wait_for_input_sync(form_id, timeout=600.0)
    if not form_data:
        return None

    logger.info("Account builder: form submitted for %s", account_provider)

    # Build the create spec
    create_cfg = _PROVIDER_CREATE.get(account_provider)
    if not create_cfg:
        return None

    spec = create_cfg["build_spec"](form_data)
    create_input = {"mode": create_cfg["mode"], **spec}

    session.send_sync(TextMessage(
        content=f"Creating {account_provider.upper()} cloud account **{form_data.get('account_name', '')}**...",
        agent="Palette Assistant",
    ))

    raw = account_tool._run(json.dumps(create_input))
    try:
        result = json.loads(raw)
    except Exception:
        result = {"status": "error", "message": raw[:200]}

    if result.get("status") == "error":
        session.send_sync(TextMessage(
            content=f"Failed to create account: {result.get('message', 'unknown error')}",
            agent="Palette Assistant",
        ))
        return None

    account_uid = result.get("data", {}).get("account_uid", "")
    session.send_sync(TextMessage(
        content=f"Cloud account **{form_data.get('account_name', '')}** created successfully! UID: `{account_uid}`",
        agent="Palette Assistant",
    ))

    return {
        "name": form_data.get("account_name", ""),
        "uid": account_uid,
        "kind": account_provider,
    }
