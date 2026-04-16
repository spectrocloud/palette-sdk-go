"""Deploy EKS cluster journey — Python-driven, uses shared building blocks."""

from __future__ import annotations

import json
import logging
import time

from ..protocol import DebugMessage, TextMessage
from ..session import Session
from ..tools.base import PaletteTool
from .shared import ensure_cloud_account, ensure_cluster_profile, collect_form, confirm_action

logger = logging.getLogger(__name__)


def run(tools: dict[str, PaletteTool], session: Session) -> str:
    """Execute the EKS deployment journey step by step."""

    session.send_sync(TextMessage(
        content="Starting EKS cluster deployment. I'll guide you through each step.",
        agent="Palette Assistant",
    ))

    # Step 1: Ensure AWS cloud account
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 1, "description": "Select cloud account"}, timestamp=time.time(),
    ))
    account = ensure_cloud_account("aws", tools, session)
    if not account:
        return "Journey stopped: no cloud account available."

    session.send_sync(TextMessage(
        content=f"Using cloud account: **{account.get('name', 'unknown')}**",
        agent="Palette Assistant",
    ))

    # Step 2: Ensure EKS-compatible profile
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 2, "description": "Select cluster profile"}, timestamp=time.time(),
    ))
    profile = ensure_cluster_profile("eks", tools, session)
    if not profile:
        return "Journey stopped: no EKS-compatible profile available."

    session.send_sync(TextMessage(
        content=f"Using cluster profile: **{profile.get('name', 'unknown')}**",
        agent="Palette Assistant",
    ))

    # Step 3: Collect cluster configuration
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 3, "description": "Collect cluster config"}, timestamp=time.time(),
    ))
    config = collect_form(
        title="EKS Cluster Configuration",
        fields=[
            {"name": "cluster_name", "label": "Cluster Name", "required": True},
            {"name": "region", "label": "AWS Region", "required": True, "default": "us-east-1"},
            {"name": "instance_type", "label": "Node Instance Type", "default": "t3.large"},
            {"name": "node_count", "label": "Node Count", "type": "number", "default": "3"},
        ],
        tool=tools.get("eks_cluster_api"),
        session=session,
    )
    if not config:
        return "Journey stopped: no configuration provided."

    # Step 4: Confirm
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 4, "description": "Confirm deployment"}, timestamp=time.time(),
    ))
    confirmed = confirm_action(
        title="Deploy EKS Cluster",
        description=(
            f"**Cluster:** {config.get('cluster_name', 'unnamed')}\n"
            f"**Region:** {config.get('region', 'us-east-1')}\n"
            f"**Instance Type:** {config.get('instance_type', 't3.large')}\n"
            f"**Node Count:** {config.get('node_count', 3)}\n"
            f"**Cloud Account:** {account.get('name', 'unknown')}\n"
            f"**Profile:** {profile.get('name', 'unknown')}\n\n"
            "Proceed with deployment?"
        ),
        session=session,
    )
    if not confirmed:
        return "Deployment cancelled by user."

    # Step 5: Create the cluster
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 5, "description": "Creating cluster"}, timestamp=time.time(),
    ))

    tool = tools.get("eks_cluster_api")
    if tool:
        create_input = {
            "mode": "create_eks",
            "cluster": {
                "metadata": {"name": config.get("cluster_name", "unnamed")},
                "spec": {
                    "cloudAccountUid": account.get("uid", ""),
                    "profiles": [{"uid": profile.get("uid", "")}],
                    "cloudConfig": {
                        "region": config.get("region", "us-east-1"),
                    },
                },
            },
        }
        tool._run(json.dumps(create_input))

    # Post-deployment
    session.send_sync(TextMessage(
        content=(
            "**EKS cluster deployment initiated!**\n\n"
            f"Cluster **{config.get('cluster_name')}** is now provisioning. This typically takes 10-15 minutes.\n\n"
            "**Next steps:**\n"
            "- Monitor status: ask me \"show cluster status\" or \"is my cluster healthy?\"\n"
            "- Get kubeconfig: ask me \"get kubeconfig for [cluster name]\"\n"
            "- Add add-on profiles: ask me \"list profiles\" to see available add-ons"
        ),
        agent="Palette Assistant",
    ))

    return ""
