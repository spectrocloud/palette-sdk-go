"""Hybrid router: simple queries → single specialist, journeys → multi-agent crew."""

import json
import logging
import os
import re
import time
from pathlib import Path
from typing import Any

import httpx
from crewai import Agent, Crew, Process, Task

from ..protocol import DebugMessage, TextMessage

from ..config import settings
from ..session import Session
from ..tools import create_tool

logger = logging.getLogger(__name__)

SKILLS_DIR = Path(__file__).resolve().parent.parent / "skills"
REFS_DIR = SKILLS_DIR / "refs"
API_SKILLS_DIR = SKILLS_DIR / "api"
JOURNEY_SKILLS_DIR = SKILLS_DIR / "journeys"


def _load(path: Path) -> str:
    return path.read_text() if path.exists() else ""


# ---------------------------------------------------------------------------
# Agent config
# ---------------------------------------------------------------------------

AGENT_CONFIG = {
    "infra": {
        "role": "Infra Provisioner",
        "goal": "Handle cluster lifecycle and cloud account operations.",
        "api_skills": [
            "cluster_common_api", "cloud_account_api", "aws_infra_api", "edge_api",
        ],
        "journey_skills": ["deploying_eks_cluster", "deploying_edge_cluster_agent", "deploying_edge_cluster_appliance"],
        "refs": ["provisioning.md", "cloud_types.md", "edge.md"],
        "tools": [
            "cluster_common_api", "eks_cluster_api", "aws_cluster_api",
            "azure_cluster_api", "aks_cluster_api", "gcp_cluster_api",
            "gke_cluster_api", "vmware_cluster_api", "maas_cluster_api",
            "edge_cluster_api", "cloud_account_api", "aws_infra_api", "edge_api",
        ],
        "keywords": [
            "deploy", "teardown", "kubeconfig",
            "eks", "aws", "azure", "aks", "gcp", "gke", "vsphere", "vmware",
            "maas", "edge", "cloud account", "provision", "node",
            "vpc", "vpcs", "subnet", "subnets", "key pair", "key pairs",
            "security group", "security groups", "availability zone",
            "ec2", "instance", "instances", "ami",
        ],
    },
    "platform_admin": {
        "role": "Platform Admin",
        "goal": "Manage projects and platform-wide settings.",
        "api_skills": ["project_api", "platform_settings_api"],
        "journey_skills": [],
        "refs": [],
        "tools": ["project_api", "platform_settings_api"],
        "keywords": [
            "project", "sso", "oidc", "saml", "session timeout", "login banner",
            "password policy", "fips", "tenant", "platform setting",
        ],
    },
    "app_workload": {
        "role": "App & Workload Specialist",
        "goal": "Manage cluster profiles, registries, and virtual machines.",
        "api_skills": ["profile_api", "registry_api", "vm_api"],
        "journey_skills": [],
        "refs": ["profiles_ref.md", "vmo.md"],
        "tools": ["profile_api", "registry_api", "vm_api"],
        "keywords": [
            "profile", "profiles", "pack", "packs",
            "registry", "registries", "helm", "oci", "ecr",
            "virtual machine", "vm", "kubevirt", "vmo",
            "version", "monitoring", "logging", "security",
            "prometheus", "grafana", "datadog", "thanos",
            "addon", "add-on", "ingress", "load balancer",
            "upstream", "supported",
        ],
    },
    "operations": {
        "role": "Operations Specialist",
        "goal": "Inspect clusters, configure backups, and run compliance scans.",
        "api_skills": ["cluster_common_api", "diagnostics_api"],
        "journey_skills": [],
        "refs": ["provisioning.md", "profiles_ref.md"],
        "tools": ["cluster_common_api", "diagnostics_api"],
        "keywords": [
            "inspect", "backup", "backups", "compliance", "scan",
            "alert", "alerts", "certificate", "certificates",
            "cert", "certs", "repave", "os patch", "health",
            "status", "diagnostic", "diagnostics", "monitor",
        ],
    },
    "access_admin": {
        "role": "Access Admin",
        "goal": "Manage users, teams, roles, and API keys.",
        "api_skills": ["user_api", "team_api", "role_api", "api_key_api"],
        "journey_skills": [],
        "refs": ["rbac.md"],
        "tools": ["user_api", "team_api", "role_api", "api_key_api"],
        "keywords": [
            "user", "users", "team", "teams", "role", "roles",
            "api key", "api keys", "onboard", "rbac",
            "permission", "permissions", "assign role", "invite",
        ],
    },
}

# ---------------------------------------------------------------------------
# Journey routing — matches journey verbs + domain keywords
# ---------------------------------------------------------------------------

JOURNEY_PATTERNS = [
    {
        "name": "create_cloud_account",
        "patterns": [
            r"create.*cloud\s*account", r"create.*aws\s*account",
            r"create.*azure\s*account", r"create.*gcp\s*account",
            r"create.*vsphere\s*account", r"create.*maas\s*account",
            r"add.*cloud\s*account", r"new.*cloud\s*account",
            r"set up.*cloud\s*account", r"setup.*cloud\s*account",
            r"add.*aws\s*account", r"add.*azure\s*account", r"add.*gcp\s*account",
            r"new.*aws\s*account", r"new.*azure\s*account", r"new.*gcp\s*account",
            r"aws.*account.*create", r"azure.*account.*create", r"gcp.*account.*create",
            r"aws.*cloud.*create", r"azure.*cloud.*create", r"gcp.*cloud.*create",
        ],
        "agents": ["infra"],
    },
    {
        "name": "create_cluster_profile",
        "patterns": [
            r"create.*cluster\s*profile", r"create.*infra\s*profile",
            r"build.*cluster\s*profile", r"build.*infra\s*profile",
            r"new.*cluster\s*profile", r"new.*infra\s*profile",
            r"create.*profile.*eks", r"create.*profile.*aws",
            r"create.*profile.*azure", r"create.*profile.*gcp",
            r"create.*eks.*profile", r"create.*aws.*profile",
        ],
        "agents": ["app_workload"],
    },
    {
        "name": "prepare_and_deploy_edge",
        "patterns": [
            # Combined: create VMs AND launch cluster in one flow
            r"create.*\d+.*vm.*cluster", r"create.*\d+.*aws.*cluster",
            r"launch.*\d+.*vm.*cluster", r"launch.*\d+.*aws.*cluster",
            r"\d+.*vm.*then.*cluster", r"\d+.*aws.*then.*cluster",
            r"create.*vm.*launch.*cluster", r"create.*aws.*launch.*cluster",
            r"\d+.*agent.*mode.*cluster", r"\d+.*aws.*vm.*agent.*mode.*cluster",
        ],
        "agents": ["infra", "app_workload"],
    },
    {
        "name": "deploy_edge_cluster_agent",
        "patterns": [
            # Cluster only (no VM creation)
            r"deploy.*edge.*cluster", r"launch.*edge.*cluster",
            r"create.*edge.*cluster(?!.*profile)",
            r"edge.*cluster.*agent\s*mode",
            r"agent\s*mode.*edge.*cluster",
            r"using.*registered.*edge", r"using.*edge\s*host",
        ],
        "agents": ["infra", "app_workload"],
    },
    {
        "name": "prepare_edge_hosts_agent",
        "patterns": [
            # Host/VM only (no cluster creation)
            r"create.*edge\s*host", r"prepare.*edge\s*host",
            r"provision.*edge\s*host",
            r"create.*\d+.*edge.*host", r"prepare.*\d+.*edge.*host",
            r"launch.*\d+.*agent.*mode", r"create.*\d+.*agent.*mode",
            r"launch.*\d+.*edge.*vm", r"launch.*\d+.*edge.*instance",
            r"agent\s*mode.*vm", r"agent\s*mode.*instance",
            r"launch.*agent.*vm", r"launch.*agent.*instance",
        ],
        "agents": ["infra"],
    },
    {
        "name": "deploy_edge_cluster_appliance",
        "patterns": [
            r"deploy.*edge.*appliance.*cluster", r"create.*edge.*cluster.*appliance",
            r"edge.*cluster.*appliance\s*mode",
            r"appliance\s*mode.*edge.*cluster",
        ],
        "agents": ["infra", "app_workload"],
    },
    {
        "name": "deploy_eks_cluster",
        "patterns": [
            r"deploy.*eks", r"create.*eks.*cluster(?!\s*profile)", r"provision.*eks",
            r"set up.*eks.*cluster(?!\s*profile)", r"launch.*eks", r"build.*eks.*cluster(?!\s*profile)",
            r"deploy.*new.*eks", r"i want.*deploy.*eks", r"i want.*create.*eks.*cluster(?!\s*profile)",
            r"i need.*eks.*cluster(?!\s*profile)",
        ],
        "agents": ["infra", "app_workload"],
    },
    # Add more journeys here as they are created:
    # {"name": "deploy_aws_cluster", "patterns": [...], "agents": ["infra", "app_workload"]},
    # {"name": "migrate_vms_from_vmware", "patterns": [...], "agents": ["infra", "app_workload"]},
    # {"name": "onboard_new_team", "patterns": [...], "agents": ["access_admin", "platform_admin"]},
]

# ---------------------------------------------------------------------------
# Shared context
# ---------------------------------------------------------------------------

CONCEPTS = _load(SKILLS_DIR / "concepts.md")

AGENT_RULES = (
    "RULES:\n"
    "1. Call ONE tool with the EXACT input_json from your TOOL REFERENCE, then give your Final Answer.\n"
    "2. Tool results are shown as rich UI — do NOT repeat data.\n"
    "3. If the request matches a LIMITATION, say so directly.\n"
    "4. Pass input_json as a JSON string.\n"
    "5. Do NOT call multiple tools. One tool call, then Final Answer.\n"
)


def _fetch_tool_descriptions() -> dict[str, str]:
    try:
        resp = httpx.get(f"{settings.tool_server_url}/tools", timeout=5.0)
        resp.raise_for_status()
        return {t["name"]: t["description"] for t in resp.json()}
    except Exception as e:
        logger.warning("Could not fetch tools from server: %s", e)
        return {}


def _build_backstory(cfg: dict) -> str:
    parts = [f"CONCEPTS:\n{CONCEPTS}\n"]
    parts.append(AGENT_RULES)

    # API skills only — compact, no refs (saves tokens)
    api_skills = [_load(API_SKILLS_DIR / f"{s}.md") for s in cfg.get("api_skills", [])]
    api_skills = [s for s in api_skills if s]
    if api_skills:
        parts.append("TOOLS:\n\n" + "\n\n".join(api_skills))

    return "\n".join(parts)


def _build_agent(key: str, cfg: dict, tool_descs: dict, session: Session | None) -> Agent:
    tools = [
        create_tool(t, tool_descs.get(t, t), t, session)
        for t in cfg["tools"]
    ]
    return Agent(
        role=cfg["role"],
        goal=cfg["goal"],
        backstory=_build_backstory(cfg),
        tools=tools,
        llm=settings.crewai_llm,
        verbose=True,
        max_iter=2,
    )


# ---------------------------------------------------------------------------
# Router
# ---------------------------------------------------------------------------

def _match_journey(message: str) -> dict | None:
    """Check if the message matches a known journey pattern.
    First tries regex, then falls back to LLM for typo tolerance."""
    msg_lower = message.lower()

    # Try regex first (fast, exact)
    for journey in JOURNEY_PATTERNS:
        for pattern in journey["patterns"]:
            if re.search(pattern, msg_lower):
                return journey

    # Fallback: use LLM to match against journey names (handles typos)
    try:
        import litellm
        journey_list = "\n".join(f"- {j['name']}: {', '.join(j['patterns'][:2])}" for j in JOURNEY_PATTERNS)
        resp = litellm.completion(
            model=settings.crewai_llm,
            messages=[{"role": "user", "content": (
                f"The user said: \"{message}\"\n\n"
                f"Does this match any of these journeys?\n{journey_list}\n\n"
                "Respond with EXACTLY the journey name, or 'none' if no match.\n"
                "Be tolerant of typos and rephrasings."
            )}],
            temperature=0,
            max_tokens=50,
        )
        matched_name = resp.choices[0].message.content.strip().lower()
        for journey in JOURNEY_PATTERNS:
            if journey["name"] == matched_name:
                logger.info("LLM matched journey: %s (regex failed)", matched_name)
                return journey
    except Exception as e:
        logger.warning("LLM journey matching failed: %s", e)

    return None


# ---------------------------------------------------------------------------
# Intent + specialist classification (single LLM call)
# ---------------------------------------------------------------------------

# Cache for the LLM classification to avoid redundant calls
_classification_cache: dict[str, tuple[str, str]] = {}


def _classify_with_llm(message: str) -> tuple[str, str]:
    """Use one LLM call to classify both intent and specialist agent.
    Returns (intent, agent_key)."""
    import litellm

    # Check journey match first (regex is fast and reliable for this)
    journey = _match_journey(message)
    if journey:
        # Still use LLM to distinguish "tell me about deploying EKS" vs "deploy EKS"
        pass

    agent_list = "\n".join(
        f"- {key}: {cfg['role']} — tools: {', '.join(cfg['tools'])}"
        for key, cfg in AGENT_CONFIG.items()
    )

    prompt = (
        f"Classify this user message for a Kubernetes management platform (Palette).\n\n"
        f"Message: \"{message}\"\n\n"
        f"1. INTENT — pick exactly one:\n"
        f"   - direct: User wants to fetch and display data (list, show, get, search)\n"
        f"   - interpret: User wants analysis, reasoning, comparison, troubleshooting, or asks a question about the data (why, is any, how many, compare, recommend, unhealthy, EOL, best practice)\n"
        f"   - journey: User wants to execute a multi-step workflow (deploy, create, provision, set up, launch, build)\n\n"
        f"2. AGENT — pick the best specialist:\n{agent_list}\n\n"
        f"IMPORTANT routing hints:\n"
        f"- AWS VMs, EC2 instances, VPCs, subnets, security groups, NLB → infra (aws_infra_api)\n"
        f"- KubeVirt/virtual machines INSIDE a cluster → app_workload (vm_api)\n"
        f"- Packs, profiles, registries, pack versions → app_workload (profile_api)\n"
        f"- Cluster health, status, diagnostics → operations\n"
        f"- Users, teams, roles, API keys → access_admin\n\n"
        f"Respond with EXACTLY two lines:\n"
        f"INTENT: <direct|interpret|journey>\n"
        f"AGENT: <agent_key>\n"
    )

    try:
        response = litellm.completion(
            model=settings.crewai_llm,
            messages=[{"role": "user", "content": prompt}],
            temperature=0,
            max_tokens=30,
        )
        text = response.choices[0].message.content.strip()
        logger.info("LLM classify: %s → %s", message[:80], text)

        intent = "direct"
        agent = "infra"
        for line in text.split("\n"):
            line = line.strip()
            if line.upper().startswith("INTENT:"):
                val = line.split(":", 1)[1].strip().lower()
                if val in ("direct", "interpret", "journey"):
                    intent = val
            elif line.upper().startswith("AGENT:"):
                val = line.split(":", 1)[1].strip().lower()
                if val in AGENT_CONFIG:
                    agent = val

        # Override: if journey regex matched and LLM says journey, trust it
        # If LLM says interpret but journey matched, it's an informational query about the journey
        if journey and intent == "journey":
            pass  # LLM confirmed
        elif journey and intent != "journey":
            pass  # LLM says it's informational, not execution — trust LLM

        return intent, agent

    except Exception as e:
        logger.warning("LLM classification failed: %s — falling back to keyword matching", e)
        return _classify_fallback(message)


def _classify_fallback(message: str) -> tuple[str, str]:
    """Keyword-based fallback when LLM classification fails."""
    msg_lower = message.lower()

    # Intent
    journey = _match_journey(msg_lower)
    if journey:
        # Check if informational
        info_signals = {"how to", "how do i", "steps", "explain", "tell me about", "guide", "walk me through"}
        if any(s in msg_lower for s in info_signals):
            intent = "interpret"
        else:
            intent = "journey"
    elif any(v in msg_lower for v in ("compare", "recommend", "why", "analyze", "is any",
                                       "are any", "how many", "what is", "unhealthy", "eol")):
        intent = "interpret"
    else:
        intent = "direct"

    # Agent — simple keyword scoring
    scores: dict[str, int] = {}
    for key, cfg in AGENT_CONFIG.items():
        score = sum(len(kw) for kw in cfg.get("keywords", []) if kw in msg_lower)
        scores[key] = score
    agent = max(scores, key=scores.get)  # type: ignore
    if scores[agent] == 0:
        agent = "infra"

    return intent, agent


def classify_intent(message: str) -> str:
    """Classify user intent using LLM."""
    if message not in _classification_cache:
        _classification_cache[message] = _classify_with_llm(message)
        # Keep cache small
        if len(_classification_cache) > 100:
            _classification_cache.clear()
    return _classification_cache[message][0]


def _match_specialist(message: str) -> str:
    """Match message to the best specialist agent using LLM."""
    if message not in _classification_cache:
        _classification_cache[message] = _classify_with_llm(message)
        if len(_classification_cache) > 100:
            _classification_cache.clear()
    return _classification_cache[message][1]


# ---------------------------------------------------------------------------
# Crew builders
# ---------------------------------------------------------------------------

def _build_single_agent_crew(agent_key: str, message: str, session: Session | None) -> Crew:
    """Fast path: one specialist, one tool call."""
    tool_descs = _fetch_tool_descriptions()
    cfg = AGENT_CONFIG[agent_key]
    agent = _build_agent(agent_key, cfg, tool_descs, session)

    task = Task(
        description=(
            f"User request: {message}\n\n"
            "Call the ONE most appropriate tool, then give your Final Answer."
        ),
        expected_output="A brief response.",
        agent=agent,
    )

    return Crew(
        agents=[agent],
        tasks=[task],
        process=Process.sequential,
        verbose=True,
    )


def _build_journey_crew(journey: dict, message: str, session: Session | None) -> Crew:
    """Complex path: coordinator + relevant specialists."""
    tool_descs = _fetch_tool_descriptions()
    agent_keys = journey["agents"]

    agents = []
    for key in agent_keys:
        cfg = AGENT_CONFIG[key]
        agents.append(_build_agent(key, cfg, tool_descs, session))

    # Load the journey skill doc for the coordinator
    journey_skill_file = journey["name"].replace("deploy_", "deploying_").replace("migrate_", "migrating_").replace("setup_", "setting_up_").replace("onboard_", "onboarding_")
    journey_skill = _load(JOURNEY_SKILLS_DIR / f"{journey_skill_file}.md")
    if not journey_skill:
        # Try exact name
        journey_skill = _load(JOURNEY_SKILLS_DIR / f"{journey['name']}.md")

    coordinator = Agent(
        role="Coordinator",
        goal="Orchestrate the journey by delegating steps to specialists.",
        backstory=(
            f"CONCEPTS:\n{CONCEPTS}\n\n"
            f"You are coordinating the '{journey['name']}' journey.\n"
            f"Available specialists: {', '.join(a.role for a in agents)}\n\n"
            f"JOURNEY STEPS:\n\n{journey_skill}\n\n"
            "Follow the steps above. Delegate each step to the right specialist."
        ),
        llm=settings.crewai_llm,
        verbose=True,
        allow_delegation=True,
    )

    task = Task(
        description=(
            f"User request: {message}\n\n"
            "Follow the JOURNEY STEPS from your instructions. "
            "Delegate each step to the appropriate specialist agent."
        ),
        expected_output="A summary of what was accomplished at each step.",
    )

    return Crew(
        agents=agents,
        tasks=[task],
        process=Process.hierarchical,
        manager_agent=coordinator,
        verbose=True,
    )


# ---------------------------------------------------------------------------
# Direct call (bypass CrewAI for simple queries)
# ---------------------------------------------------------------------------

def _build_direct_call(agent_key: str, session: Session | None) -> dict:
    """Build a direct-call context: tools + a pick_tool function that uses one LLM call."""
    import litellm

    tool_descs = _fetch_tool_descriptions()
    cfg = AGENT_CONFIG[agent_key]

    # Build tool instances
    tools_dict = {}
    for t in cfg["tools"]:
        tools_dict[t] = create_tool(t, tool_descs.get(t, t), t, session)

    # Build the backstory for tool selection
    backstory = _build_backstory(cfg)

    def pick_tool(user_message: str) -> tuple[str | None, str]:
        """Use one LLM call to pick the tool and generate input_json. Includes conversation history."""
        tool_list = "\n".join(f"- {name}: {tool_descs.get(name, name)}" for name in cfg["tools"])

        system = (
            f"{backstory}\n\n"
            f"Available tools:\n{tool_list}\n\n"
            "Respond with EXACTLY two lines:\n"
            "TOOL: <tool_name>\n"
            "INPUT: <json_string>\n\n"
            "If no tool fits, respond with:\n"
            "TOOL: none\n"
            "INPUT: <your text answer to the user>\n\n"
            "IMPORTANT: When the user asks about pack versions, supported software, "
            "or what version of something Palette supports, ALWAYS use profile_api with "
            "search_packs to look it up. For example, to find Prometheus versions: "
            '{"mode": "search_packs", "filter": {"name": {"eq": "prometheus"}}}\n'
            "You can also search by addOnType, layer, environment, etc. "
            "NEVER say you don't have access to version information — use the tools.\n"
        )

        # Build messages with conversation history
        messages = [{"role": "system", "content": system}]
        if session:
            # Include recent conversation history (last 10 turns)
            messages.extend(session.conversation_history[-20:])
        messages.append({"role": "user", "content": user_message})

        response = litellm.completion(
            model=settings.crewai_llm,
            messages=messages,
            temperature=0,
        )

        text = response.choices[0].message.content.strip()
        logger.info("LLM pick_tool response: %s", text[:200])

        # Parse TOOL: and INPUT: lines
        tool_name = None
        tool_input = "{}"
        for line in text.split("\n"):
            line = line.strip()
            if line.upper().startswith("TOOL:"):
                tool_name = line.split(":", 1)[1].strip()
                if tool_name.lower() == "none":
                    tool_name = None
            elif line.upper().startswith("INPUT:"):
                tool_input = line.split(":", 1)[1].strip()

        return tool_name, tool_input

    return {"tools": tools_dict, "pick_tool": pick_tool}


def _compact_tool_result(tool_result: str, max_chars: int = 8000) -> str:
    """Compact a JSON tool result by extracting key scalar fields from each item.
    This avoids sending huge nested JSON to the LLM."""
    try:
        data = json.loads(tool_result)
    except (json.JSONDecodeError, TypeError):
        return tool_result[:max_chars]

    result_data = data.get("data", data)
    if not isinstance(result_data, dict):
        return tool_result[:max_chars]

    # Find the list in the data
    items = None
    for v in result_data.values():
        if isinstance(v, list) and v and isinstance(v[0], dict):
            items = v
            break

    if not items:
        return tool_result[:max_chars]

    # Extract key scalars from each item — skip noisy/misleading fields
    _skip = {"annotations", "labels", "creationTimestamp", "deletionTimestamp",
             "lastModifiedTimestamp", "ownerReferences", "managedFields",
             "finalizers", "permissions", "ownerUid", "tenantUid", "logoUrl",
             "packs", "draft", "projectMeta", "cloudAccountMeta",
             "scopeVisibility", "scope", "lastSyncedTimestamp",
             # Skip misleading boolean fields that confuse the LLM
             "isAvailable", "isHost", "isBrownfield",
             # Skip noisy metric/cost fields
             "cost", "hourlyRate", "location", "coordinates",
             "repave", "notifications", "virtual", "fips",
             }

    def _scalars(obj, depth=0):
        flat = {}
        if depth > 3:
            return flat
        for k, v in obj.items():
            if k in _skip:
                continue
            if isinstance(v, (str, int, float, bool)) and v != "":
                # Skip false booleans — they add noise
                if v is False:
                    continue
                flat[k] = v
            elif isinstance(v, dict):
                flat.update(_scalars(v, depth + 1))
        return flat

    rows = [_scalars(item) for item in items]
    compact = json.dumps(rows, indent=1)
    if len(compact) > max_chars:
        compact = compact[:max_chars] + f"\n... ({len(rows)} total items, truncated)"
    return f"{len(rows)} items:\n{compact}"


def _web_search(query: str, max_results: int = 5) -> str:
    """Search the web using DuckDuckGo. Returns formatted results.
    Appends 'software release' to disambiguate from pop culture results."""
    try:
        from ddgs import DDGS
        # Add context to avoid pop-culture hits for names like "Thanos"
        search_query = f"{query} software release github"
        results = DDGS().text(search_query, max_results=max_results, timeout=10)
        if not results:
            return ""
        lines = []
        for r in results:
            lines.append(f"**{r.get('title', '')}**")
            lines.append(r.get("body", ""))
            lines.append(f"Source: {r.get('href', '')}\n")
        return "\n".join(lines)
    except Exception as e:
        logger.warning("Web search failed: %s", e)
        return ""


def answer_question(user_message: str, agent_key: str) -> str:
    """Answer a question using domain knowledge, general knowledge, or web search."""
    import litellm

    cfg = AGENT_CONFIG.get(agent_key, {})
    backstory = _build_backstory(cfg)

    # Also load any relevant journey skills
    journey_context = ""
    for j in JOURNEY_PATTERNS:
        for p in j["patterns"]:
            if re.search(p, user_message.lower()):
                skill_file = j["name"].replace("deploy_", "deploying_").replace(
                    "migrate_", "migrating_").replace("setup_", "setting_up_").replace(
                    "onboard_", "onboarding_")
                content = _load(JOURNEY_SKILLS_DIR / f"{skill_file}.md")
                if content:
                    journey_context = f"\n\nJOURNEY REFERENCE:\n{content}"
                break

    # Also load refs for deeper context
    refs_context = ""
    for r in cfg.get("refs", []):
        content = _load(REFS_DIR / r)
        if content:
            refs_context += f"\n\n{content}"

    # Check if this is a question that might need web search
    # (versions, latest releases, comparisons with upstream, current info)
    needs_web = any(kw in user_message.lower() for kw in [
        "latest version", "upstream", "current version", "newest",
        "released", "release", "compare version", "up to date",
        "what version", "which version",
    ])

    web_context = ""
    if needs_web:
        logger.info("Searching web for: %s", user_message[:100])
        web_results = _web_search(user_message)
        if web_results:
            web_context = f"\n\nWEB SEARCH RESULTS:\n{web_results}"

    prompt = (
        f"{backstory}"
        f"{refs_context}"
        f"{journey_context}"
        f"{web_context}\n\n"
        f"User question: {user_message}\n\n"
        "Provide a clear, detailed answer. "
        "If the question is about a process or workflow, explain the steps. "
        "Use Palette domain knowledge when relevant, but also use your general knowledge "
        "for questions about upstream software versions, open-source projects, comparisons, "
        "or industry best practices. Do NOT refuse to answer just because it's not strictly "
        "about Palette — if you know the answer, share it. "
        "If web search results are provided above, use them to give accurate, current information."
    )

    response = litellm.completion(
        model=settings.crewai_llm,
        messages=[{"role": "user", "content": prompt}],
        temperature=0,
    )

    return response.choices[0].message.content.strip()


def interpret_result(user_message: str, tool_name: str, tool_result: str) -> str:
    """Second LLM call: interpret/analyze tool results for the user."""
    import litellm

    # Web search if user asks about versions/upstream
    web_context = ""
    needs_web = any(kw in user_message.lower() for kw in [
        "latest version", "upstream", "current version", "newest",
        "released", "release", "compare version", "up to date",
    ])
    if needs_web:
        web_results = _web_search(user_message)
        if web_results:
            web_context = f"\n\nWEB SEARCH RESULTS:\n{web_results}"

    # Compact the tool result for the LLM — extract key fields only
    compact_result = _compact_tool_result(tool_result)

    prompt = (
        f"DOMAIN CONCEPTS:\n{CONCEPTS}\n\n"
        f"The user asked: {user_message}\n\n"
        f"Tool '{tool_name}' returned this data:\n{compact_result}\n\n"
        f"{web_context}\n\n"
        "IMPORTANT: Analyze ALL items in the data above, not just the first few. "
        "Based on this data, provide a concise, helpful analysis answering the user's question. "
        "Focus on insights, not raw data. Be specific and actionable. "
        "If the user asks about upstream/latest versions of open-source software, "
        "use web search results and your general knowledge to compare with what Palette supports."
    )

    response = litellm.completion(
        model=settings.crewai_llm,
        messages=[{"role": "user", "content": prompt}],
        temperature=0,
    )

    return response.choices[0].message.content.strip()


# ---------------------------------------------------------------------------
# Journey runner (LLM-orchestrated, tools handle interactivity)
# ---------------------------------------------------------------------------

def run_journey(journey: dict, user_message: str, session: Session | None) -> str:
    """Run a journey. Python-driven for interactive flows, LLM for others."""

    if not session:
        return "Session required for journeys."

    # Build tools for all involved agents
    tool_descs = _fetch_tool_descriptions()
    all_tools: dict[str, Any] = {}
    for agent_key in journey["agents"]:
        cfg = AGENT_CONFIG[agent_key]
        for t in cfg["tools"]:
            if t not in all_tools:
                all_tools[t] = create_tool(t, tool_descs.get(t, t), t, session)

    # Python-driven journeys
    if journey["name"] == "create_cloud_account":
        return _run_create_account_journey(journey, user_message, session, all_tools)
    if journey["name"] == "create_cluster_profile":
        return _run_create_profile_journey(journey, user_message, session, all_tools)
    if journey["name"] == "deploy_eks_cluster":
        return _run_deploy_eks_journey(journey, user_message, session, all_tools)
    if journey["name"] == "prepare_and_deploy_edge":
        result = _run_prepare_edge_hosts_agent(journey, user_message, session, all_tools)
        if result:  # Non-empty means error
            return result
        return _run_deploy_edge_cluster(journey, user_message, session, all_tools)
    if journey["name"] == "prepare_edge_hosts_agent":
        return _run_prepare_edge_hosts_agent(journey, user_message, session, all_tools)
    if journey["name"] == "deploy_edge_cluster_agent":
        return _run_deploy_edge_cluster(journey, user_message, session, all_tools)

    # Fallback: LLM-orchestrated
    return _run_llm_journey(journey, user_message, session)


def _run_create_account_journey(journey: dict, user_message: str, session: Session, all_tools: dict) -> str:
    """Python-driven cloud account creation."""
    from ..tools.account_builder import build_account_form

    # Detect provider from user message
    msg_lower = user_message.lower()
    provider = "aws"  # default
    for p in ["aws", "azure", "gcp", "vsphere", "maas"]:
        if p in msg_lower:
            provider = p
            break

    session.send_sync(TextMessage(
        content=f"Let's create a new **{provider.upper()}** cloud account.",
        agent="Palette Assistant",
    ))

    result = build_account_form(provider, all_tools, session)
    if result:
        return ""
    return "Account creation cancelled or failed."


def _run_create_profile_journey(journey: dict, user_message: str, session: Session, all_tools: dict) -> str:
    """Python-driven profile creation journey."""
    from ..tools.profile_builder import build_profile_form, create_profile_from_form

    # Use LLM to detect cloud type from natural language
    import litellm
    cloud_type_resp = litellm.completion(
        model=settings.crewai_llm,
        messages=[{"role": "user", "content": (
            f"The user said: \"{user_message}\"\n\n"
            "What Palette cloud type are they referring to? "
            "Respond with EXACTLY one of: eks, aws, azure, aks, gcp, gke, vmware, edge-native, maas\n\n"
            "Hints:\n"
            "- edge, edge agent, edge appliance, edge native, edge cluster → edge-native\n"
            "- amazon eks, eks → eks\n"
            "- aws, amazon → aws\n"
            "- vsphere, vcenter → vmware\n"
            "- bare metal → maas\n"
            "- If unclear, respond: unknown"
        )}],
        temperature=0,
        max_tokens=20,
    )
    cloud_type = cloud_type_resp.choices[0].message.content.strip().lower()
    logger.info("LLM detected cloud type: %s from: %s", cloud_type, user_message[:50])

    valid_types = {"eks", "aws", "azure", "aks", "gcp", "gke", "vmware", "edge-native", "maas"}
    if cloud_type not in valid_types:
        # LLM couldn't determine — ask the user
        from ..protocol import SelectMessage
        ct_id = f"pick_cloud_type_{int(time.time())}"
        session.register_pending(ct_id)
        session.send_sync(SelectMessage(
            id=ct_id, title="Select Cloud Type",
            description="What cloud type is this profile for?",
            options=[
                {"value": "eks", "label": "Amazon EKS"},
                {"value": "aws", "label": "AWS IaaS"},
                {"value": "azure", "label": "Azure IaaS"},
                {"value": "aks", "label": "Azure AKS"},
                {"value": "gcp", "label": "Google Cloud"},
                {"value": "gke", "label": "Google GKE"},
                {"value": "vmware", "label": "VMware vSphere"},
                {"value": "edge-native", "label": "Edge Native"},
                {"value": "maas", "label": "MAAS Bare Metal"},
            ],
            agent="Palette Assistant",
        ))
        cloud_type = session.wait_for_input_sync(ct_id, timeout=600.0)
        if not cloud_type:
            return "No cloud type selected."

    # For edge, detect agent vs appliance mode to set correct OS default
    edge_mode = None
    if cloud_type == "edge-native":
        msg_lower = user_message.lower()
        if "agent" in msg_lower:
            edge_mode = "agent"
        elif "appliance" in msg_lower:
            edge_mode = "appliance"

    mode_label = f" ({edge_mode} mode)" if edge_mode else ""
    session.send_sync(TextMessage(
        content=f"I'll help you create a **{cloud_type.upper()}{mode_label}** infra profile. Fetching available packs...",
        agent="Palette Assistant",
    ))

    form_data = build_profile_form(cloud_type, all_tools, session, edge_mode=edge_mode)
    if not form_data:
        return "Profile creation cancelled."

    result = create_profile_from_form(cloud_type, form_data, all_tools, session)
    if result:
        return ""
    return "Profile creation failed."


def _run_deploy_eks_journey(journey: dict, user_message: str, session: Session, all_tools: dict) -> str:
    """Python-driven EKS deploy journey with agentic error handling and recommendations."""
    from ..tools.profile_builder import build_profile_form, create_profile_from_form
    from ..tools.account_builder import build_account_form
    from ..defaults import load_defaults, get_pack_defaults
    from ..journeys.agent_layer import handle_error, recommend_selection, summarize_journey

    defaults = load_defaults("eks")
    journey_context = {"journey_name": "deploy_eks_cluster", "collected": {}}

    # Show the plan first
    from ..protocol import PlanMessage, PlanStep, PlanUpdateMessage, ConfirmMessage

    plan_id = f"deploy_eks_plan_{int(time.time())}"
    session.send_sync(PlanMessage(
        id=plan_id,
        title="Deploy EKS Cluster",
        description="Here's what we'll do step by step:",
        steps=[
            PlanStep(id="account", label="Select AWS Cloud Account", description="Pick existing or create new", icon="\u2601"),
            PlanStep(id="profile", label="Select Cluster Profile", description="Pick existing EKS profile or create new", icon="\U0001f4cb"),
            PlanStep(id="config", label="Configure Cluster", description="Name, region, instance type, node count", icon="\u2699"),
            PlanStep(id="confirm", label="Review & Confirm", description="Review all selections", icon="\u2714"),
            PlanStep(id="create", label="Create Cluster", description="Call Palette API to provision", icon="\U0001f680"),
        ],
        agent="Palette Assistant",
    ))

    # Wait for user to agree with the plan
    confirm_plan_id = f"confirm_plan_{int(time.time())}"
    session.register_pending(confirm_plan_id)
    session.send_sync(ConfirmMessage(
        id=confirm_plan_id,
        title="Proceed with this plan?",
        description="I'll guide you through each step interactively.",
        confirm_label="Let's go",
        cancel_label="Cancel",
        agent="Palette Assistant",
    ))
    if not session.wait_for_input_sync(confirm_plan_id, timeout=600.0):
        return "Journey cancelled."

    def update_step(step_id: str, status: str):
        session.send_sync(PlanUpdateMessage(plan_id=plan_id, step_id=step_id, status=status, agent="Palette Assistant"))

    # Step 1: Ensure cloud account
    update_step("account", "active")
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 1, "description": "Select or create AWS cloud account"},
        timestamp=time.time(),
    ))

    account_tool = all_tools.get("cloud_account_api")
    if account_tool:
        account_tool.silent = True
        raw = account_tool._run(json.dumps({"mode": "list_accounts"}))
        account_tool.silent = False

        accounts = _extract_flat_list(raw)
        if accounts:
            # Agent recommends best account
            recommend_selection(accounts, "EKS cluster deployment — need an AWS cloud account", session)

            from ..protocol import SelectMessage
            options = [
                {"value": json.dumps(a), "label": f"{a.get('name', 'unknown')} ({a.get('kind', '')})"}
                for a in accounts
            ]
            options.append({"value": "__create_new__", "label": "+ Create new AWS account"})

            select_id = f"pick_account_{int(time.time())}"
            session.register_pending(select_id)
            session.send_sync(SelectMessage(
                id=select_id, title="Select AWS Cloud Account",
                description="Choose an existing account or create a new one",
                options=options, agent="Palette Assistant",
            ))
            selected = session.wait_for_input_sync(select_id, timeout=600.0)

            if selected == "__create_new__":
                account = build_account_form("eks", all_tools, session)
            else:
                try:
                    account = json.loads(selected) if isinstance(selected, str) else selected
                except Exception:
                    account = {"name": str(selected)}
        else:
            session.send_sync(TextMessage(
                content="No cloud accounts found. Let's create one.",
                agent="Palette Assistant",
            ))
            account = build_account_form("eks", all_tools, session)
    else:
        account = None

    if not account:
        return "Journey stopped: no cloud account available."

    update_step("account", "done")
    session.send_sync(TextMessage(
        content=f"Using cloud account: **{account.get('name', 'unknown')}**",
        agent="Palette Assistant",
    ))

    # Step 2: Ensure cluster profile
    update_step("profile", "active")
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 2, "description": "Select or create EKS cluster profile"},
        timestamp=time.time(),
    ))

    profile_tool = all_tools.get("profile_api")
    if profile_tool:
        profile_tool.silent = True
        raw = profile_tool._run(json.dumps({
            "mode": "search_profiles",
            "filter": {"filter": {"environment": ["eks"], "profileType": ["cluster"]}},
        }))
        profile_tool.silent = False

        profiles = _extract_flat_list(raw)
        if profiles:
            recommend_selection(profiles, "EKS cluster deployment — need an EKS-compatible infra profile", session)

            from ..protocol import SelectMessage
            options = [
                {"value": json.dumps(p), "label": f"{p.get('name', 'unknown')} v{p.get('version', '')}"}
                for p in profiles
            ]
            options.append({"value": "__create_new__", "label": "+ Create new EKS profile"})

            select_id = f"pick_profile_{int(time.time())}"
            session.register_pending(select_id)
            session.send_sync(SelectMessage(
                id=select_id, title="Select EKS Cluster Profile",
                description="Choose an existing profile or create a new one",
                options=options, agent="Palette Assistant",
            ))
            selected = session.wait_for_input_sync(select_id, timeout=600.0)

            if selected == "__create_new__":
                form_data = build_profile_form("eks", all_tools, session)
                profile = create_profile_from_form("eks", form_data, all_tools, session) if form_data else None
            else:
                try:
                    profile = json.loads(selected) if isinstance(selected, str) else selected
                except Exception:
                    profile = {"name": str(selected)}
        else:
            session.send_sync(TextMessage(
                content="No EKS profiles found. Let's create one.",
                agent="Palette Assistant",
            ))
            form_data = build_profile_form("eks", all_tools, session)
            profile = create_profile_from_form("eks", form_data, all_tools, session) if form_data else None
    else:
        profile = None

    if not profile:
        return "Journey stopped: no cluster profile available."

    update_step("profile", "done")
    session.send_sync(TextMessage(
        content=f"Using cluster profile: **{profile.get('name', 'unknown')}**",
        agent="Palette Assistant",
    ))

    # Step 3: Collect cluster config
    update_step("config", "active")
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 3, "description": "Configure cluster"},
        timestamp=time.time(),
    ))

    cluster_defaults = defaults.get("cluster", {})
    from ..protocol import FormMessage, FormField, FieldType

    config_form_id = f"cluster_config_{int(time.time())}"
    session.register_pending(config_form_id)
    session.send_sync(FormMessage(
        id=config_form_id,
        title="EKS Cluster Configuration",
        description="Review defaults and adjust as needed.",
        fields=[
            FormField(name="cluster_name", label="Cluster Name", type=FieldType.STRING, required=True,
                      default="eks-cluster"),
            FormField(name="region", label="AWS Region", type=FieldType.STRING, required=True,
                      default=cluster_defaults.get("region", "us-east-1")),
            FormField(name="instance_type", label="Node Instance Type", type=FieldType.STRING,
                      default=cluster_defaults.get("instance_type", "t3.large")),
            FormField(name="node_count", label="Node Count", type=FieldType.NUMBER, required=True,
                      default=str(cluster_defaults.get("node_count", 3))),
            FormField(name="ssh_key", label="SSH Key Pair Name", type=FieldType.STRING,
                      default="", description="AWS SSH key pair name (optional)"),
        ],
        submit_label="Review & Deploy",
        agent="Palette Assistant",
    ))
    config = session.wait_for_input_sync(config_form_id, timeout=600.0)
    if not config:
        return "Journey stopped: no configuration provided."

    # Step 4: Confirm
    update_step("config", "done")
    update_step("confirm", "active")
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 4, "description": "Confirm deployment"},
        timestamp=time.time(),
    ))

    confirm_id = f"deploy_confirm_{int(time.time())}"
    session.register_pending(confirm_id)
    session.send_sync(ConfirmMessage(
        id=confirm_id,
        title="Deploy EKS Cluster",
        description=(
            f"**Cluster:** {config.get('cluster_name', 'unnamed')}\n"
            f"**Region:** {config.get('region', 'us-east-1')}\n"
            f"**Instance Type:** {config.get('instance_type', 't3.large')}\n"
            f"**Nodes:** {config.get('node_count', 3)}\n"
            f"**Account:** {account.get('name', '?')}\n"
            f"**Profile:** {profile.get('name', '?')}\n\n"
            "Proceed with deployment?"
        ),
        agent="Palette Assistant",
    ))
    confirmed = session.wait_for_input_sync(confirm_id, timeout=600.0)
    if not confirmed:
        return "Deployment cancelled."

    # Step 5: Create cluster
    update_step("confirm", "done")
    update_step("create", "active")
    session.send_sync(DebugMessage(
        event="journey_step", data={"step": 5, "description": "Creating cluster"},
        timestamp=time.time(),
    ))

    eks_tool = all_tools.get("eks_cluster_api")
    if eks_tool:
        region = config.get("region", "us-east-1")
        instance_type = config.get("instance_type", "t3.large")
        node_count = int(config.get("node_count", 3))

        create_input = json.dumps({
            "mode": "create_eks",
            "cluster": {
                "metadata": {
                    "name": config.get("cluster_name", "unnamed"),
                    "annotations": {},
                    "labels": {},
                },
                "spec": {
                    "cloudAccountUid": account.get("uid", ""),
                    "edgeHostUid": "",
                    "profiles": [{"uid": profile.get("uid", "")}],
                    "fargateProfiles": [],
                    "cloudConfig": {
                        "region": region,
                        "sshKeyName": config.get("ssh_key", ""),
                        "vpcId": "",
                        "endpointAccess": {
                            "private": True,
                            "public": True,
                            "publicCIDRs": ["0.0.0.0/0"],
                            "privateCIDRs": ["0.0.0.0/0"],
                        },
                    },
                    "machinePoolConfig": [
                        {
                            "cloudConfig": {
                                "instanceType": instance_type,
                                "azs": [],
                                "rootDeviceSize": 60,
                                "capacityType": "on-demand",
                                "amiType": "AL2023_x86_64_STANDARD",
                                "subnets": [],
                            },
                            "poolConfig": {
                                "name": "worker-pool",
                                "labels": ["worker"],
                                "taints": [],
                                "additionalLabels": {},
                                "additionalAnnotations": {},
                                "nodeRepaveInterval": 0,
                                "updateStrategy": {
                                    "type": "RollingUpdateScaleOut",
                                },
                                "machinePoolProperties": {
                                    "archType": "amd64",
                                },
                                "size": node_count,
                            },
                        },
                    ],
                },
            },
        })
        raw_result = eks_tool._run(create_input)

        # Check for errors — let the agent interpret and suggest recovery
        try:
            result = json.loads(raw_result)
            if result.get("status") == "error":
                update_step("create", "error")
                journey_context["collected"] = {
                    "account": account.get("name"),
                    "profile": profile.get("name"),
                    "config": config,
                }
                recovery = handle_error(
                    result.get("message", "Unknown error"),
                    "Create EKS cluster",
                    journey_context,
                    session,
                )
                return recovery or "Cluster creation failed."
        except Exception:
            pass

    update_step("create", "done")

    # Agent generates a natural summary
    journey_context["collected"] = {
        "cluster_name": config.get("cluster_name"),
        "region": config.get("region"),
        "account": account.get("name"),
        "profile": profile.get("name"),
        "instance_type": config.get("instance_type"),
        "node_count": config.get("node_count"),
    }
    summarize_journey("deploy_eks_cluster", journey_context["collected"], session)

    return ""


def _run_prepare_edge_hosts_agent(journey: dict, user_message: str, session: Session, all_tools: dict) -> str:
    """Python-driven edge host preparation in agent mode on AWS."""
    from ..tools.edge_userdata import generate_agent_userdata
    from ..defaults import load_defaults
    from ..protocol import PlanMessage, PlanStep, PlanUpdateMessage, ConfirmMessage, FormMessage, FormField, FieldType, SelectMessage
    import os
    import re as _re

    defaults = load_defaults("edge-native")
    agent_defaults = defaults.get("agent", {})

    # Extract host count from message if mentioned
    host_count = 3
    match = _re.search(r"(\d+)\s*(host|node|vm|instance|machine)", user_message.lower())
    if match:
        host_count = int(match.group(1))

    # Show plan
    plan_id = f"edge_agent_plan_{int(time.time())}"
    session.send_sync(PlanMessage(
        id=plan_id,
        title="Prepare Edge Hosts (Agent Mode)",
        description=f"Launching {host_count} edge hosts on AWS with stylus agent",
        steps=[
            PlanStep(id="token", label="Create Registration Token", icon="\U0001f511"),
            PlanStep(id="infra", label="Select AWS Infrastructure", description="VPC, subnet, key pair, security group", icon="\u2601"),
            PlanStep(id="config", label="Configure Hosts", description="Count, instance type, disk, site name", icon="\u2699"),
            PlanStep(id="launch", label=f"Launch {host_count} EC2 Instances", description="With stylus agent user-data", icon="\U0001f680"),
            PlanStep(id="verify", label="Verify Host Registration", description="Check hosts registered in Palette", icon="\u2714"),
        ],
        agent="Palette Assistant",
    ))

    confirm_id = f"edge_plan_confirm_{int(time.time())}"
    session.register_pending(confirm_id)
    session.send_sync(ConfirmMessage(
        id=confirm_id, title="Proceed with edge host preparation?",
        description=f"This will create a registration token and launch {host_count} EC2 instances with the stylus agent.",
        confirm_label="Let's go", cancel_label="Cancel", agent="Palette Assistant",
    ))
    if not session.wait_for_input_sync(confirm_id, timeout=600.0):
        return "Journey cancelled."

    def update_step(step_id: str, status: str):
        session.send_sync(PlanUpdateMessage(plan_id=plan_id, step_id=step_id, status=status, agent="Palette Assistant"))

    # Step 1: Create registration token
    update_step("token", "active")
    edge_tool = all_tools.get("edge_api")
    token_value = ""
    if edge_tool:
        edge_tool.silent = True
        raw = edge_tool._run(json.dumps({"mode": "create_token", "token_name": "palette-agentic"}))
        edge_tool.silent = False
        try:
            result = json.loads(raw)
            if result.get("status") == "ok":
                token_value = result.get("data", {}).get("token_value", "")
                session.send_sync(TextMessage(content=f"Registration token created.", agent="Palette Assistant"))
            else:
                from ..journeys.agent_layer import handle_error
                handle_error(result.get("message", ""), "Create registration token", {"journey_name": "edge_agent"}, session)
                return "Failed to create registration token."
        except Exception:
            pass
    update_step("token", "done")

    # Step 2: Select AWS infrastructure
    update_step("infra", "active")
    aws_tool = all_tools.get("aws_infra_api")
    vpc = None
    subnet = None
    key_pair = None

    if aws_tool:
        # Select VPC
        aws_tool.silent = True
        raw = aws_tool._run(json.dumps({"mode": "list_vpcs"}))
        aws_tool.silent = False
        vpcs = _extract_flat_list(raw)
        if vpcs:
            options = [{"value": json.dumps(v), "label": f"{v.get('name', v.get('id', '?'))} ({v.get('cidr', '')})"} for v in vpcs]
            sid = f"pick_vpc_{int(time.time())}"
            session.register_pending(sid)
            session.send_sync(SelectMessage(id=sid, title="Select VPC", description="Where to launch edge hosts", options=options, agent="Palette Assistant"))
            selected = session.wait_for_input_sync(sid, timeout=600.0)
            try:
                vpc = json.loads(selected) if isinstance(selected, str) else selected
            except Exception:
                vpc = {"id": str(selected)}

        # Select subnet
        if vpc:
            aws_tool.silent = True
            raw = aws_tool._run(json.dumps({"mode": "list_subnets", "vpc_id": vpc.get("id", "")}))
            aws_tool.silent = False
            subnets = _extract_flat_list(raw)
            if subnets:
                options = [{"value": json.dumps(s), "label": f"{s.get('name', s.get('id', '?'))} ({s.get('cidr', '')}) {s.get('availabilityZone', '')}"} for s in subnets]
                sid = f"pick_subnet_{int(time.time())}"
                session.register_pending(sid)
                session.send_sync(SelectMessage(id=sid, title="Select Subnet", options=options, agent="Palette Assistant"))
                selected = session.wait_for_input_sync(sid, timeout=600.0)
                try:
                    subnet = json.loads(selected) if isinstance(selected, str) else selected
                except Exception:
                    subnet = {"id": str(selected)}

        # Select key pair
        aws_tool.silent = True
        raw = aws_tool._run(json.dumps({"mode": "list_key_pairs"}))
        aws_tool.silent = False
        keys = _extract_flat_list(raw)
        if keys:
            options = [{"value": k.get("name", ""), "label": k.get("name", "?")} for k in keys]
            sid = f"pick_key_{int(time.time())}"
            session.register_pending(sid)
            session.send_sync(SelectMessage(id=sid, title="Select SSH Key Pair", options=options, agent="Palette Assistant"))
            key_pair = session.wait_for_input_sync(sid, timeout=600.0)

        # Security group: offer to create one with edge ports or pick existing
        vpc_id = vpc.get("id", "") if vpc else ""
        security_group = None

        # Show required ports info
        session.send_sync(TextMessage(
            content=(
                "**Edge cluster hosts require these ports:**\n\n"
                "| Port | Protocol | Direction | Purpose |\n"
                "|------|----------|-----------|--------|\n"
                "| 22 | TCP | Inbound | SSH access |\n"
                "| 443 | TCP | Outbound | Palette management |\n"
                "| 5080 | TCP | Inbound | Agent / Local UI |\n"
                "| 5082 | TCP | Inbound | Content bundle upload |\n"
                "| 6443 | TCP | Inbound + Inter-node | Kubernetes API |\n"
                "| 2379-2380 | TCP | Inter-node | etcd |\n"
                "| 10250 | TCP | Inter-node | Kubelet API |\n"
                "| 10259 | TCP | Inter-node | kube-scheduler |\n"
                "| 10257 | TCP | Inter-node | kube-controller-manager |\n"
                "| 30000-32767 | TCP | Inter-node | NodePort Services |\n"
                "| 5473, 9098 | TCP | Inter-node | Calico |\n"
            ),
            agent="Palette Assistant",
        ))

        from ..protocol import ConfirmMessage
        sg_choice_id = f"sg_choice_{int(time.time())}"
        session.register_pending(sg_choice_id)
        session.send_sync(ConfirmMessage(
            id=sg_choice_id,
            title="Create a new security group with these ports?",
            description="I'll create a security group with all required edge cluster ports pre-configured.",
            confirm_label="Create new SG",
            cancel_label="Pick existing SG",
            agent="Palette Assistant",
        ))
        create_new_sg = session.wait_for_input_sync(sg_choice_id, timeout=600.0)

        if create_new_sg:
            # Create SG with all edge ports
            sg_name = f"palette-edge-{config.get('site_name', 'cluster') if 'config' in dir() else 'cluster'}-{int(time.time()) % 10000}"
            sg_spec = {
                "mode": "create_security_group",
                "sg_spec": {
                    "name": sg_name,
                    "description": "Palette edge cluster ports - managed by Palette Agentic",
                    "vpcId": vpc_id,
                    "rules": [
                        # Inbound
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 22, "toPort": 22, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 5080, "toPort": 5080, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 5082, "toPort": 5082, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 6443, "toPort": 6443, "cidr": "0.0.0.0/0"},
                        # Inter-node (using VPC CIDR or 0.0.0.0/0 for simplicity)
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 2379, "toPort": 2380, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 10250, "toPort": 10250, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 10257, "toPort": 10257, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 10259, "toPort": 10259, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 30000, "toPort": 32767, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 5473, "toPort": 5473, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 9098, "toPort": 9098, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "udp", "fromPort": 8472, "toPort": 8472, "cidr": "0.0.0.0/0"},
                        {"direction": "ingress", "protocol": "tcp", "fromPort": 4240, "toPort": 4240, "cidr": "0.0.0.0/0"},
                        # Outbound (443 for Palette management — default SG allows all outbound)
                        {"direction": "egress", "protocol": "tcp", "fromPort": 443, "toPort": 443, "cidr": "0.0.0.0/0"},
                    ],
                },
            }
            session.send_sync(TextMessage(content=f"Creating security group **{sg_name}**...", agent="Palette Assistant"))
            aws_tool.silent = True
            raw = aws_tool._run(json.dumps(sg_spec))
            aws_tool.silent = False
            try:
                result = json.loads(raw)
                if result.get("status") == "ok":
                    sg_data = result.get("data", {}).get("security_group", {})
                    security_group = sg_data.get("id", "")
                    session.send_sync(TextMessage(content=f"Security group created: **{sg_name}** (`{security_group}`)", agent="Palette Assistant"))
                else:
                    session.send_sync(TextMessage(content=f"Failed to create SG: {result.get('message', '')}", agent="Palette Assistant"))
            except Exception as e:
                session.send_sync(TextMessage(content=f"Failed to create SG: {e}", agent="Palette Assistant"))
        else:
            # Pick existing SG
            aws_tool.silent = True
            raw = aws_tool._run(json.dumps({"mode": "list_security_groups", "vpc_id": vpc_id}))
            aws_tool.silent = False
            sgs = _extract_flat_list(raw)
            if sgs:
                options = [{"value": sg.get("id", ""), "label": f"{sg.get('name', '?')} — {sg.get('description', '')[:40]}"} for sg in sgs]
                sid = f"pick_sg_{int(time.time())}"
                session.register_pending(sid)
                session.send_sync(SelectMessage(id=sid, title="Select Security Group", description="Ensure it has the required edge ports open", options=options, agent="Palette Assistant"))
                security_group = session.wait_for_input_sync(sid, timeout=600.0)

    update_step("infra", "done")

    # Step 3: Collect config
    update_step("config", "active")
    config_id = f"edge_config_{int(time.time())}"
    session.register_pending(config_id)
    session.send_sync(FormMessage(
        id=config_id,
        title="Edge Host Configuration",
        fields=[
            FormField(name="host_count", label="Number of Hosts", type=FieldType.NUMBER, required=True, default=str(host_count)),
            FormField(name="ami_id", label="AMI ID", type=FieldType.STRING, required=True, default=agent_defaults.get("base_ami", "ami-00de3875b03809ec5"), description="Ubuntu 22.04 LTS AMI"),
            FormField(name="instance_type", label="Instance Type", type=FieldType.STRING, default=agent_defaults.get("instance_type", "t3.xlarge")),
            FormField(name="disk_size", label="Root Disk Size (GB)", type=FieldType.NUMBER, default=str(agent_defaults.get("disk_size_gb", 100))),
            FormField(name="site_name", label="Site Name (tag)", type=FieldType.STRING, default="edge-site-01"),
        ],
        submit_label="Launch Hosts",
        agent="Palette Assistant",
    ))
    config = session.wait_for_input_sync(config_id, timeout=600.0)
    if not config:
        return "Journey cancelled."
    host_count = int(config.get("host_count", host_count))
    update_step("config", "done")

    # Step 4: Pre-create edge hosts in Palette + Launch EC2 instances
    update_step("launch", "active")
    palette_endpoint = os.getenv("PALETTE_HOST", "api.spectrocloud.com")
    project_uid = os.getenv("PALETTE_PROJECT_UID", "")

    # Generate host names
    site_name = config.get("site_name", "edge-site")
    host_names = [f"{site_name}-{i+1}" for i in range(host_count)]

    # Pre-create edge hosts in Palette so cluster can reference them before VMs register
    edge_tool = all_tools.get("edge_api")
    if edge_tool:
        session.send_sync(TextMessage(
            content=f"Pre-creating **{host_count}** edge hosts in Palette...",
            agent="Palette Assistant",
        ))
        edge_tool.silent = True
        raw = edge_tool._run(json.dumps({
            "mode": "create_edge_host",
            "host_names": host_names,
            "tags": {"site": site_name, "managed-by": "palette-agentic", "infra": "aws"},
        }))
        edge_tool.silent = False
        try:
            result = json.loads(raw)
            if result.get("status") == "ok":
                session.send_sync(TextMessage(
                    content=f"Pre-created {len(host_names)} edge hosts: {', '.join(host_names)}",
                    agent="Palette Assistant",
                ))
            else:
                session.send_sync(TextMessage(
                    content=f"Warning: pre-creation failed: {result.get('message', '')}. Hosts will register on boot.",
                    agent="Palette Assistant",
                ))
        except Exception as e:
            session.send_sync(TextMessage(
                content=f"Warning: pre-creation failed: {e}. Hosts will register on boot.",
                agent="Palette Assistant",
            ))

    launched = []
    edge_host_ids = host_names  # Use the site names as edge host IDs
    if aws_tool:
        for i in range(host_count):
            edge_host_site_name = host_names[i]
            host_name = f"edge-agent-{site_name}-{i+1}"

            # Generate per-host user-data
            userdata = generate_agent_userdata(
                token=token_value,
                palette_endpoint=f"https://{palette_endpoint}",
                project_uid=project_uid,
                edge_id="",  # Let Palette assign based on site name
                site_name=edge_host_site_name,
                tags={
                    "site": site_name,
                    "managed-by": "palette-agentic",
                    "infra": "aws",
                    "name": host_name,
                },
            )

            session.send_sync(TextMessage(
                content=f"Launching host {i+1}/{host_count}: **{host_name}** (edge host: `{edge_host_site_name}`)",
                agent="Palette Assistant",
            ))

            spec = {
                "mode": "launch_instance",
                "spec": {
                    "name": host_name,
                    "imageId": config.get("ami_id", "ami-00de3875b03809ec5"),
                    "instanceType": config.get("instance_type", "t3.xlarge"),
                    "diskSizeGb": int(config.get("disk_size", 100)),
                    "subnetId": subnet.get("id", "") if subnet else "",
                    "keyName": key_pair if isinstance(key_pair, str) else "",
                    "securityGroups": [security_group] if security_group else [],
                    "userData": userdata,
                },
            }
            aws_tool.silent = True
            raw = aws_tool._run(json.dumps(spec))
            aws_tool.silent = False
            try:
                result = json.loads(raw)
                if result.get("status") == "ok":
                    inst = result.get("data", {}).get("instance", {})
                    inst["edge_id"] = edge_host_site_name
                    launched.append(inst)
                else:
                    session.send_sync(TextMessage(content=f"Failed to launch {host_name}: {result.get('message', '')}", agent="Palette Assistant"))
            except Exception as e:
                session.send_sync(TextMessage(content=f"Failed to launch {host_name}: {e}", agent="Palette Assistant"))

    session.send_sync(TextMessage(
        content=f"Launched **{len(launched)}/{host_count}** EC2 instances.",
        agent="Palette Assistant",
    ))

    session.send_sync(TextMessage(
        content=(
            "The stylus agent installation script was injected via user-data. "
            "Each host will automatically install the agent and register with Palette on first boot."
        ),
        agent="Palette Assistant",
    ))
    update_step("launch", "done")
    update_step("verify", "done")  # No need to wait — edge IDs are pre-assigned

    # Show edge host IDs for cluster creation
    host_table = "\n".join(
        f"| {l.get('name', f'host-{i+1}')} | `{l.get('edge_id', '?')[:12]}...` | `{l.get('id', '?')}` |"
        for i, l in enumerate(launched)
    )

    # Store artifacts for the cluster creation journey
    session.artifacts["edge_host_ids"] = edge_host_ids
    session.artifacts["edge_hosts"] = launched
    session.artifacts["edge_token"] = token_value
    session.artifacts["edge_site_name"] = config.get("site_name", "edge-site-01")

    session.send_sync(TextMessage(
        content=(
            f"**Edge host preparation complete!**\n\n"
            f"**{len(launched)}** hosts launched with pre-assigned edge IDs:\n\n"
            f"| Name | Edge ID | EC2 Instance |\n"
            f"|------|---------|-------------|\n"
            f"{host_table}\n\n"
            f"- Site: {config.get('site_name', 'edge-site-01')}\n"
            f"- Agent auto-installs on first boot and registers with the pre-assigned edge ID\n\n"
            "**You can create the cluster immediately** using these edge IDs — no need to wait for registration.\n"
            "The cluster will start provisioning as soon as the hosts register.\n\n"
            "**Next steps:**\n"
            "- Deploy a cluster: \"deploy edge cluster using these hosts\"\n"
            "- Check registration: \"list edge hosts\""
        ),
        agent="Palette Assistant",
    ))

    return ""


def _run_deploy_edge_cluster(journey: dict, user_message: str, session: Session, all_tools: dict) -> str:
    """Python-driven edge cluster creation using registered/pending edge hosts."""
    from ..tools.profile_builder import build_profile_form, create_profile_from_form
    from ..defaults import load_defaults
    from ..protocol import PlanMessage, PlanStep, PlanUpdateMessage, ConfirmMessage, FormMessage, FormField, FieldType, SelectMessage

    defaults = load_defaults("edge-native")

    # Check for artifacts from prepare step
    edge_host_ids = session.artifacts.get("edge_host_ids", [])
    edge_hosts = session.artifacts.get("edge_hosts", [])
    site_name = session.artifacts.get("edge_site_name", "edge-cluster")

    # Show plan
    plan_id = f"edge_cluster_plan_{int(time.time())}"
    has_hosts = len(edge_host_ids) > 0
    session.send_sync(PlanMessage(
        id=plan_id,
        title="Deploy Edge Cluster",
        description=f"{'Using ' + str(len(edge_host_ids)) + ' pre-assigned edge hosts' if has_hosts else 'Select registered edge hosts'}",
        steps=[
            PlanStep(id="hosts", label="Select Edge Hosts",
                     description=f"{len(edge_host_ids)} hosts from previous step" if has_hosts else "Pick from registered hosts",
                     status="done" if has_hosts else "pending", icon="\U0001f5a5"),
            PlanStep(id="profile", label="Select Cluster Profile", description="Edge-native infra profile", icon="\U0001f4cb"),
            PlanStep(id="config", label="Configure Cluster", description="VIP, node pools", icon="\u2699"),
            PlanStep(id="confirm", label="Review & Confirm", icon="\u2714"),
            PlanStep(id="create", label="Create Cluster", icon="\U0001f680"),
        ],
        agent="Palette Assistant",
    ))

    confirm_plan_id = f"edge_plan_confirm_{int(time.time())}"
    session.register_pending(confirm_plan_id)
    session.send_sync(ConfirmMessage(
        id=confirm_plan_id, title="Proceed with edge cluster deployment?",
        description=f"{'Using ' + str(len(edge_host_ids)) + ' edge hosts from the previous step.' if has_hosts else 'You will select edge hosts from those registered in Palette.'}",
        confirm_label="Let's go", cancel_label="Cancel", agent="Palette Assistant",
    ))
    if not session.wait_for_input_sync(confirm_plan_id, timeout=600.0):
        return "Journey cancelled."

    def update_step(step_id: str, status: str):
        session.send_sync(PlanUpdateMessage(plan_id=plan_id, step_id=step_id, status=status, agent="Palette Assistant"))

    # Step 1: Select edge hosts (skip if artifacts exist)
    if not has_hosts:
        update_step("hosts", "active")
        edge_tool = all_tools.get("edge_api")
        if edge_tool:
            edge_tool.silent = True
            raw = edge_tool._run(json.dumps({"mode": "list_edge_hosts"}))
            edge_tool.silent = False
            # Parse full host objects (not flattened) for rich info
            try:
                result = json.loads(raw)
                all_hosts_raw = result.get("data", {}).get("edge_hosts", [])
            except Exception:
                all_hosts_raw = []

            # Filter: current project + ready/unused
            current_project = os.getenv("PALETTE_PROJECT_UID", "")
            hosts = []
            for h in all_hosts_raw:
                meta = h.get("metadata", {})
                spec = h.get("spec", {})
                status = h.get("status", {})
                ann = meta.get("annotations", {})

                # Project filter
                host_project = ann.get("projectUid", "")
                if current_project and host_project != current_project:
                    continue

                # State filter — only ready/unused
                state = status.get("state", "").lower()
                if state not in ("ready", "registered", "unused", "unpaired", ""):
                    continue

                device = spec.get("device", {})
                hosts.append({
                    "uid": meta.get("uid", ""),
                    "name": meta.get("name", ""),
                    "state": status.get("state", "unknown"),
                    "health": status.get("health", {}).get("state", "unknown") if isinstance(status.get("health"), dict) else str(status.get("health", "unknown")),
                    "labels": meta.get("labels", {}),
                    "cpu_cores": device.get("cpu", {}).get("cores", 0),
                    "memory_mb": device.get("memory", {}).get("sizeInMB", 0),
                    "arch": device.get("archType", ""),
                })

            if not hosts:
                total = len(all_hosts_raw)
                session.send_sync(TextMessage(
                    content=(
                        f"No available edge hosts found in this project.\n"
                        f"({total} hosts total across all projects, but none are ready/unassigned in project `{current_project}`).\n\n"
                        "Prepare edge hosts first with: \"create edge hosts in agent mode\""
                    ),
                    agent="Palette Assistant",
                ))
                return "No edge hosts available."

            # Multi-select with rich labels
            options = []
            for h in hosts:
                uid = h["uid"]
                name = h["name"]
                tags = ", ".join(f"{k}={v}" for k, v in h.get("labels", {}).items()) if h.get("labels") else "no tags"
                cpu = h.get("cpu_cores", "?")
                mem_gb = round(h.get("memory_mb", 0) / 1024, 1)
                state = h.get("state", "?")
                health = h.get("health", "?")

                label = f"{name} [{state}/{health}] {cpu} CPU, {mem_gb}GB RAM — {tags}"
                options.append({
                    "value": json.dumps({"uid": uid, "name": name}),
                    "label": label,
                })

            sid = f"pick_hosts_{int(time.time())}"
            session.register_pending(sid)
            session.send_sync(SelectMessage(
                id=sid,
                title="Select Edge Hosts",
                description=f"{len(hosts)} available hosts. Select the ones to use for this cluster.",
                options=options,
                multiple=True,
                agent="Palette Assistant",
            ))
            selected_raw = session.wait_for_input_sync(sid, timeout=600.0)

            # Parse multi-select response (list of JSON strings)
            selected_hosts = []
            if isinstance(selected_raw, list):
                for s in selected_raw:
                    try:
                        selected_hosts.append(json.loads(s) if isinstance(s, str) else s)
                    except Exception:
                        selected_hosts.append({"uid": str(s)})
            elif isinstance(selected_raw, str):
                try:
                    selected_hosts.append(json.loads(selected_raw))
                except Exception:
                    selected_hosts.append({"uid": selected_raw})

            if not selected_hosts:
                session.send_sync(TextMessage(content="No hosts selected.", agent="Palette Assistant"))
                return "No hosts selected."

            edge_host_ids = [h.get("uid", "") for h in selected_hosts]
            edge_hosts = selected_hosts
            session.send_sync(TextMessage(
                content=f"Selected **{len(edge_host_ids)}** edge hosts.",
                agent="Palette Assistant",
            ))
        update_step("hosts", "done")
    else:
        session.send_sync(TextMessage(
            content=f"Using **{len(edge_host_ids)}** edge hosts from previous preparation step.",
            agent="Palette Assistant",
        ))

    # Step 2: Select or create edge cluster profile
    update_step("profile", "active")
    profile_tool = all_tools.get("profile_api")
    profile = None
    if profile_tool:
        profile_tool.silent = True
        raw = profile_tool._run(json.dumps({
            "mode": "search_profiles",
            "filter": {"filter": {"environment": ["edge-native"], "profileType": ["cluster"]}},
        }))
        profile_tool.silent = False
        profiles = _extract_flat_list(raw)
        if profiles:
            from ..journeys.agent_layer import recommend_selection
            recommend_selection(profiles, "Edge native cluster — need an edge-compatible infra profile", session)

            options = [
                {"value": json.dumps(p), "label": f"{p.get('name', '?')} v{p.get('version', '')}"}
                for p in profiles
            ]
            options.append({"value": "__create_new__", "label": "+ Create new edge profile"})

            sid = f"pick_edge_profile_{int(time.time())}"
            session.register_pending(sid)
            session.send_sync(SelectMessage(
                id=sid, title="Select Edge Cluster Profile",
                description="Choose an edge-native infra profile",
                options=options, agent="Palette Assistant",
            ))
            selected = session.wait_for_input_sync(sid, timeout=600.0)

            if selected == "__create_new__":
                form_data = build_profile_form("edge-native", all_tools, session)
                profile = create_profile_from_form("edge-native", form_data, all_tools, session) if form_data else None
            else:
                try:
                    profile = json.loads(selected) if isinstance(selected, str) else selected
                except Exception:
                    profile = {"name": str(selected)}
        else:
            session.send_sync(TextMessage(content="No edge profiles found. Let's create one.", agent="Palette Assistant"))
            form_data = build_profile_form("edge-native", all_tools, session)
            profile = create_profile_from_form("edge-native", form_data, all_tools, session) if form_data else None

    if not profile:
        return "No cluster profile available."

    update_step("profile", "done")
    session.send_sync(TextMessage(
        content=f"Using profile: **{profile.get('name', 'unknown')}**",
        agent="Palette Assistant",
    ))

    # Step 3: Create NLB + configure cluster
    update_step("config", "active")
    total_hosts = len(edge_host_ids)

    # Collect cluster name first
    config_id = f"edge_cluster_config_{int(time.time())}"
    session.register_pending(config_id)
    session.send_sync(FormMessage(
        id=config_id,
        title="Edge Cluster Configuration",
        description=f"All {total_hosts} hosts will be in a single pool (control plane + worker).",
        fields=[
            FormField(name="cluster_name", label="Cluster Name", type=FieldType.STRING, required=True,
                      default=f"edge-{site_name}"),
        ],
        submit_label="Continue",
        agent="Palette Assistant",
    ))
    config = session.wait_for_input_sync(config_id, timeout=600.0)
    if not config:
        return "Journey cancelled."

    # Create NLB for K8s API (port 6443) — only for AWS-hosted edge hosts
    # Detect infra type from edge host labels
    is_aws = False
    for h in edge_hosts:
        labels = h.get("labels", {})
        if labels.get("infra") == "aws":
            is_aws = True
            break
    # Also check artifacts
    if not is_aws:
        for h in session.artifacts.get("edge_hosts", []):
            if isinstance(h, dict) and h.get("edge_id", ""):
                # Hosts from prepare step are always AWS in this flow
                is_aws = bool(os.getenv("AWS_ACCESS_KEY_ID"))
                break

    aws_tool = all_tools.get("aws_infra_api")
    vip = ""
    nlb_arn = ""
    if aws_tool and is_aws:
        # Get instance IDs — from artifacts (same session) or by querying AWS tags
        instance_ids = [h.get("id", "") for h in session.artifacts.get("edge_hosts", []) if h.get("id")]
        subnet_id = ""
        vpc_id = ""

        if instance_ids:
            # From artifacts
            for h in session.artifacts.get("edge_hosts", []):
                if h.get("subnetId") and not subnet_id:
                    subnet_id = h["subnetId"]
                if h.get("vpcId") and not vpc_id:
                    vpc_id = h["vpcId"]
        else:
            # No artifacts — find our EC2 instances by tag
            session.send_sync(TextMessage(
                content="Looking up EC2 instances tagged by Palette Agentic...",
                agent="Palette Assistant",
            ))
            aws_tool.silent = True
            raw = aws_tool._run(json.dumps({
                "mode": "list_instances",
                "filters": {"tag:managed-by": "palette-agentic", "instance-state-name": "running"},
            }))
            aws_tool.silent = False
            tagged_instances = _extract_flat_list(raw)

            if tagged_instances:
                # Match edge hosts by name tag
                for inst in tagged_instances:
                    inst_name = inst.get("name", "")
                    # Check if this instance's name matches any of our selected edge hosts
                    for eid in edge_host_ids:
                        if eid in inst_name or inst_name.endswith(eid):
                            instance_ids.append(inst.get("id", ""))
                            if not subnet_id:
                                subnet_id = inst.get("subnetId", "")
                            if not vpc_id:
                                vpc_id = inst.get("vpcId", "")
                            break

                if instance_ids:
                    session.send_sync(TextMessage(
                        content=f"Found **{len(instance_ids)}** matching EC2 instances.",
                        agent="Palette Assistant",
                    ))
                else:
                    # Fallback: use all tagged instances
                    instance_ids = [i.get("id", "") for i in tagged_instances if i.get("id")]
                    if tagged_instances:
                        subnet_id = tagged_instances[0].get("subnetId", "")
                        vpc_id = tagged_instances[0].get("vpcId", "")
                    session.send_sync(TextMessage(
                        content=f"Using all **{len(instance_ids)}** Palette-managed instances.",
                        agent="Palette Assistant",
                    ))

        if instance_ids and subnet_id:
            nlb_name = f"edge-{config.get('cluster_name', 'cluster')}-nlb"[:32]

            session.send_sync(TextMessage(
                content=f"Creating NLB **{nlb_name}** for K8s API (port 6443)...",
                agent="Palette Assistant",
            ))

            aws_tool.silent = True
            raw = aws_tool._run(json.dumps({
                "mode": "create_nlb",
                "nlb_spec": {
                    "name": nlb_name,
                    "vpcId": vpc_id,
                    "subnetIds": [subnet_id],
                    "port": 6443,
                    "targetIds": instance_ids,
                },
            }))
            aws_tool.silent = False

            try:
                result = json.loads(raw)
                if result.get("status") == "ok":
                    nlb_data = result.get("data", {}).get("nlb", {})
                    vip = nlb_data.get("lbDns", "")
                    nlb_arn = nlb_data.get("lbArn", "")
                    session.send_sync(TextMessage(
                        content=(
                            f"NLB created: **{nlb_name}**\n"
                            f"- DNS: `{vip}`\n"
                            f"- Targets: {len(instance_ids)} instances on port 6443\n"
                            f"- This will be used as the cluster VIP"
                        ),
                        agent="Palette Assistant",
                    ))
                else:
                    session.send_sync(TextMessage(
                        content=f"Failed to create NLB: {result.get('message', '')}. You'll need to provide a VIP manually.",
                        agent="Palette Assistant",
                    ))
            except Exception as e:
                session.send_sync(TextMessage(
                    content=f"Failed to create NLB: {e}. You'll need to provide a VIP manually.",
                    agent="Palette Assistant",
                ))

    # If NLB failed, ask for VIP manually
    if not vip:
        vip_id = f"edge_vip_{int(time.time())}"
        session.register_pending(vip_id)
        session.send_sync(FormMessage(
            id=vip_id,
            title="Cluster VIP",
            fields=[
                FormField(name="vip", label="Virtual IP (VIP)", type=FieldType.STRING, required=True,
                          description="IP or DNS for the cluster API endpoint"),
            ],
            submit_label="Continue",
            agent="Palette Assistant",
        ))
        vip_data = session.wait_for_input_sync(vip_id, timeout=600.0)
        vip = vip_data.get("vip", "") if vip_data else ""

    config["vip"] = vip
    config["nlb_arn"] = nlb_arn

    # Store NLB ARN for cleanup if needed
    session.artifacts["nlb_arn"] = nlb_arn

    update_step("config", "done")

    # Step 4: Confirm
    update_step("confirm", "active")
    confirm_id = f"edge_deploy_confirm_{int(time.time())}"
    session.register_pending(confirm_id)
    session.send_sync(ConfirmMessage(
        id=confirm_id,
        title="Deploy Edge Cluster",
        description=(
            f"**Cluster:** {config.get('cluster_name', 'unnamed')}\n"
            f"**VIP:** {config.get('vip', 'not set')}\n"
            f"**Profile:** {profile.get('name', '?')}\n"
            f"**Hosts:** {total_hosts} (single pool — all CP + worker)\n"
            f"**Host IDs:** {', '.join(h[:12] + '...' for h in edge_host_ids)}\n\n"
            "Proceed?"
        ),
        agent="Palette Assistant",
    ))
    if not session.wait_for_input_sync(confirm_id, timeout=600.0):
        return "Deployment cancelled."

    update_step("confirm", "done")

    # Step 5: Create cluster
    update_step("create", "active")
    edge_cluster_tool = all_tools.get("edge_cluster_api")
    if edge_cluster_tool:
        # Build the edge native cluster entity
        # Single pool: all hosts as CP + worker (default edge pattern)
        all_host_uids = edge_host_ids
        create_input = json.dumps({
            "mode": "create_edge_native",
            "cluster": {
                "metadata": {"name": config.get("cluster_name", "edge-cluster")},
                "spec": {
                    "cloudConfig": {
                        "controlPlaneEndpoint": {
                            "host": config.get("vip", ""),
                            "type": "IP",
                        },
                    },
                    "machinepoolconfig": [
                        {
                            "poolConfig": {
                                "name": "master-pool",
                                "size": len(all_host_uids),
                                "isControlPlane": True,
                                "labels": ["master"],
                            },
                            "cloudConfig": {
                                "edgeHosts": [{"hostUid": uid} for uid in all_host_uids],
                            },
                        },
                    ],
                    "profiles": [{"uid": profile.get("uid", "")}],
                },
            },
        })

        raw_result = edge_cluster_tool._run(create_input)

        # Check for errors
        try:
            result = json.loads(raw_result)
            if result.get("status") == "error":
                update_step("create", "error")
                from ..journeys.agent_layer import handle_error
                handle_error(
                    result.get("message", "Unknown error"),
                    "Create edge cluster",
                    {"journey_name": "deploy_edge_cluster", "collected": config},
                    session,
                )
                return "Cluster creation failed."
        except Exception:
            pass

    update_step("create", "done")

    session.send_sync(TextMessage(
        content=(
            f"**Edge cluster deployment initiated!**\n\n"
            f"Cluster **{config.get('cluster_name')}** is being provisioned.\n"
            f"- VIP: {config.get('vip')}\n"
            f"- {total_hosts} nodes (single pool — CP + worker)\n"
            f"- Profile: {profile.get('name')}\n\n"
            "The cluster will start provisioning as edge hosts register and become available.\n\n"
            "**Next steps:**\n"
            "- Check status: \"show cluster status\"\n"
            "- List edge hosts: \"list edge hosts\""
        ),
        agent="Palette Assistant",
    ))

    # Clear artifacts
    session.artifacts.pop("edge_host_ids", None)
    session.artifacts.pop("edge_hosts", None)

    return ""


def _extract_flat_list(raw: str) -> list[dict]:
    """Extract and flatten a list from tool response."""
    try:
        result = json.loads(raw)
    except Exception:
        return []
    data = result.get("data", {})
    rows = _find_list_deep(data)
    return [_flatten_item(r) for r in rows if _flatten_item(r)]


def _find_list_deep(data: Any, depth: int = 0) -> list[dict]:
    if depth > 3:
        return []
    if isinstance(data, list) and data and isinstance(data[0], dict):
        return data
    if isinstance(data, dict):
        if "items" in data and isinstance(data["items"], list):
            return data["items"]
        for v in data.values():
            found = _find_list_deep(v, depth + 1)
            if found:
                return found
    return []


_SKIP_FIELDS = {"annotations", "labels", "creationTimestamp", "deletionTimestamp",
                "lastModifiedTimestamp", "permissions", "ownerUid", "tenantUid",
                "scopeVisibility", "projectUid", "scope", "packs", "draft", "logoUrl"}


def _flatten_item(obj: dict, depth: int = 0) -> dict:
    if depth > 5:
        return {}
    flat: dict[str, Any] = {}
    for k, v in obj.items():
        if k in _SKIP_FIELDS:
            continue
        if isinstance(v, (str, int, float, bool)):
            if v != "" and v is not None:
                flat[k] = v
        elif isinstance(v, dict):
            flat.update(_flatten_item(v, depth + 1))
    return flat


def _run_llm_journey(journey: dict, user_message: str, session: Session | None) -> str:
    """LLM-orchestrated journey for deploy and other complex flows."""
    import litellm

    # Load journey skill doc
    skill_file = journey["name"].replace("deploy_", "deploying_").replace(
        "migrate_", "migrating_").replace("setup_", "setting_up_").replace(
        "onboard_", "onboarding_")
    journey_skill = _load(JOURNEY_SKILLS_DIR / f"{skill_file}.md")
    if not journey_skill:
        journey_skill = _load(JOURNEY_SKILLS_DIR / f"{journey['name']}.md")

    # Build all tools from involved agents
    tool_descs = _fetch_tool_descriptions()
    all_tools: dict[str, Any] = {}
    for agent_key in journey["agents"]:
        cfg = AGENT_CONFIG[agent_key]
        for t in cfg["tools"]:
            if t not in all_tools:
                all_tools[t] = create_tool(t, tool_descs.get(t, t), t, session)

    # Build compact tool + API skill reference
    all_api_skills = []
    for agent_key in journey["agents"]:
        cfg = AGENT_CONFIG[agent_key]
        for s in cfg.get("api_skills", []):
            content = _load(API_SKILLS_DIR / f"{s}.md")
            if content and content not in all_api_skills:
                all_api_skills.append(content)
    api_ref = "\n\n".join(all_api_skills)

    tool_list = "\n".join(f"- {n}" for n in all_tools)

    system_prompt = (
        f"You are orchestrating a multi-step journey on the Palette platform.\n\n"
        f"DOMAIN:\n{CONCEPTS}\n\n"
        f"JOURNEY GUIDE:\n{journey_skill}\n\n"
        f"TOOL REFERENCE:\n{api_ref}\n\n"
        f"AVAILABLE TOOLS:\n{tool_list}\n\n"
        "YOU MUST FOLLOW THESE RULES EXACTLY:\n\n"
        "RULE 1: Follow the JOURNEY GUIDE steps for the matching variant in EXACT ORDER.\n"
        "Do NOT follow steps from other variants or shared sub-journeys unless the guide tells you to.\n"
        "Do NOT add your own steps. Do NOT skip steps. Do NOT combine steps.\n\n"
        "RULE 2: ALWAYS include '\"interactive\": true' when listing resources for the user to SELECT.\n"
        "CORRECT:   {\"mode\": \"list_accounts\", \"interactive\": true}\n"
        "INCORRECT: {\"mode\": \"list_accounts\"}\n\n"
        "RULE 3: When the guide says to use 'build_profile', call the tool with EXACTLY:\n"
        "INPUT: {\"build_profile\": \"<cloud_type>\"}\n"
        "Do NOT call search_packs separately. build_profile handles everything.\n\n"
        "RULE 4: If a tool returns an empty list, guide the user to create the resource.\n\n"
        "RULE 5: Carry forward data from previous steps (UIDs, names, etc.).\n\n"
        "RESPONSE FORMAT — EXACTLY these three lines:\n"
        "THOUGHT: <one line of reasoning>\n"
        "TOOL: <tool_name>\n"
        "INPUT: <single line valid JSON string>\n\n"
        "When journey is complete:\n"
        "THOUGHT: done\n"
        "TOOL: done\n"
        "INPUT: <final message>"
    )

    # Note: journey preview is handled inside Python-driven journeys (e.g. _run_deploy_eks_journey)
    # LLM-orchestrated journeys rely on the skill doc for step guidance

    messages = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": user_message},
    ]

    max_steps = 15
    for step_num in range(max_steps):
        logger.info("Journey %s step %d: calling LLM", journey["name"], step_num + 1)

        if session:
            session.send_sync(DebugMessage(
                event="journey_step",
                data={"step": step_num + 1, "status": "thinking"},
                timestamp=time.time(),
            ))

        response = litellm.completion(
            model=settings.crewai_llm,
            messages=messages,
            temperature=0,
        )

        llm_text = response.choices[0].message.content.strip()
        logger.info("Journey step %d response:\n%s", step_num + 1, llm_text[:300])

        # Parse response
        thought, tool_name, tool_input = _parse_llm_step(llm_text)

        if session:
            session.send_sync(DebugMessage(
                event="journey_step",
                data={"step": step_num + 1, "thought": thought[:100], "tool": tool_name},
                timestamp=time.time(),
            ))

        # Done
        if tool_name == "__done__":
            if session and tool_input:
                session.send_sync(TextMessage(content=tool_input, agent="Palette Assistant"))
            return ""

        # Execute tool
        if tool_name and tool_name in all_tools:
            tool = all_tools[tool_name]
            tool_result = tool._run(tool_input)

            messages.append({"role": "assistant", "content": llm_text})
            messages.append({"role": "user", "content": (
                f"Tool '{tool_name}' returned:\n{tool_result[:2000]}\n\n"
                "What is the next step? Remember to use \"interactive\": true for selections."
            )})

        elif tool_name:
            # Tool not found
            messages.append({"role": "assistant", "content": llm_text})
            messages.append({"role": "user", "content": (
                f"Tool '{tool_name}' not found. Available tools: {', '.join(all_tools.keys())}. "
                "Please use one of these tools."
            )})
        else:
            # No tool — show thought as message and end
            if session and thought:
                session.send_sync(TextMessage(content=thought, agent="Palette Assistant"))
            break

    return ""


def _parse_llm_step(text: str) -> tuple[str, str | None, str]:
    """Parse THOUGHT/TOOL/INPUT from LLM response."""
    thought = ""
    tool_name = None
    tool_input = "{}"

    for line in text.split("\n"):
        line = line.strip()
        if line.upper().startswith("THOUGHT:"):
            thought = line.split(":", 1)[1].strip()
        elif line.upper().startswith("TOOL:"):
            val = line.split(":", 1)[1].strip()
            if val.lower() == "done":
                tool_name = "__done__"
            elif val.lower() == "none":
                tool_name = None
            else:
                tool_name = val
        elif line.upper().startswith("INPUT:"):
            tool_input = line.split(":", 1)[1].strip()

    return thought, tool_name, tool_input
