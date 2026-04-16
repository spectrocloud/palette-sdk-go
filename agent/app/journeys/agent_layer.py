"""Agentic layer for journeys — wraps deterministic steps with LLM intelligence.

The deterministic Python code handles step ordering, form rendering, and API calls.
This layer adds:
- Error interpretation and recovery suggestions
- Mid-journey question handling
- Smart recommendations based on context
"""

from __future__ import annotations

import json
import logging
from typing import Any

import litellm

from ..config import settings
from ..defaults import load_defaults
from ..protocol import DebugMessage, TextMessage
from ..session import Session

logger = logging.getLogger(__name__)

# Shared concepts for LLM context
_CONCEPTS = ""


def _get_concepts() -> str:
    global _CONCEPTS
    if not _CONCEPTS:
        from pathlib import Path
        concepts_path = Path(__file__).resolve().parent.parent / "skills" / "concepts.md"
        if concepts_path.exists():
            _CONCEPTS = concepts_path.read_text()
    return _CONCEPTS


def handle_error(
    error_message: str,
    step_description: str,
    journey_context: dict,
    session: Session,
) -> str | None:
    """LLM interprets an error and suggests recovery.
    Returns a recovery action string or None if no recovery possible."""

    session.send_sync(DebugMessage(
        event="agent_thinking",
        data={"purpose": "error_recovery", "error": error_message[:100]},
        timestamp=__import__("time").time(),
    ))

    prompt = (
        f"CONTEXT:\n{_get_concepts()}\n\n"
        f"JOURNEY: {journey_context.get('journey_name', 'unknown')}\n"
        f"CURRENT STEP: {step_description}\n"
        f"COLLECTED SO FAR: {json.dumps(journey_context.get('collected', {}), default=str)[:500]}\n\n"
        f"ERROR: {error_message}\n\n"
        "You are helping a user with a Palette platform operation that hit an error.\n"
        "Analyze the error and respond with EXACTLY one of:\n\n"
        "RECOVERY: <action description>\n"
        "MESSAGE: <helpful message to show the user>\n\n"
        "Or if no recovery is possible:\n"
        "MESSAGE: <explanation of what went wrong and what the user should do>\n\n"
        "Common recoveries:\n"
        "- Quota exceeded → suggest deleting unused resources\n"
        "- Permission denied → explain what permissions are needed\n"
        "- Resource not found → suggest creating it\n"
        "- Invalid input → explain what's wrong with the input\n"
        "- Server error → suggest retrying or contacting support\n"
    )

    try:
        response = litellm.completion(
            model=settings.crewai_llm,
            messages=[{"role": "user", "content": prompt}],
            temperature=0,
        )
        text = response.choices[0].message.content.strip()
        logger.info("Agent error recovery: %s", text[:200])

        recovery = None
        message = None
        for line in text.split("\n"):
            line = line.strip()
            if line.upper().startswith("RECOVERY:"):
                recovery = line.split(":", 1)[1].strip()
            elif line.upper().startswith("MESSAGE:"):
                message = line.split(":", 1)[1].strip()

        if message:
            session.send_sync(TextMessage(content=message, agent="Palette Assistant"))

        return recovery

    except Exception as e:
        logger.warning("Agent error recovery failed: %s", e)
        session.send_sync(TextMessage(
            content=f"An error occurred: {error_message}",
            agent="Palette Assistant",
        ))
        return None


def recommend_selection(
    options: list[dict],
    context: str,
    session: Session,
) -> str | None:
    """LLM recommends which option to pick from a list.
    Returns the recommended option name or None."""

    if not options:
        return None

    options_text = "\n".join(
        f"- {o.get('name', o.get('label', str(o)[:50]))}: {json.dumps({k: v for k, v in o.items() if k not in ('name', 'label') and isinstance(v, (str, int, float, bool))})}"
        for o in options[:10]
    )

    prompt = (
        f"CONTEXT: {context}\n\n"
        f"OPTIONS:\n{options_text}\n\n"
        "Which option would you recommend and why? Respond in one sentence.\n"
        "Format: RECOMMEND: <option name> — <reason>"
    )

    try:
        response = litellm.completion(
            model=settings.crewai_llm,
            messages=[{"role": "user", "content": prompt}],
            temperature=0,
            max_tokens=100,
        )
        text = response.choices[0].message.content.strip()

        for line in text.split("\n"):
            if line.upper().startswith("RECOMMEND:"):
                recommendation = line.split(":", 1)[1].strip()
                session.send_sync(TextMessage(
                    content=f"*Recommendation: {recommendation}*",
                    agent="Palette Assistant",
                ))
                return recommendation.split("—")[0].strip() if "—" in recommendation else None

    except Exception as e:
        logger.warning("Agent recommendation failed: %s", e)

    return None


def summarize_journey(
    journey_name: str,
    collected: dict,
    session: Session,
) -> None:
    """LLM generates a natural summary of what was accomplished."""

    prompt = (
        f"A user just completed the '{journey_name}' journey on the Palette platform.\n"
        f"Here's what was done: {json.dumps(collected, default=str)[:1000]}\n\n"
        "Write a brief, friendly summary (2-3 sentences) of what was accomplished "
        "and what the user should do next. Be specific about resource names."
    )

    try:
        response = litellm.completion(
            model=settings.crewai_llm,
            messages=[{"role": "user", "content": prompt}],
            temperature=0.3,
            max_tokens=200,
        )
        summary = response.choices[0].message.content.strip()
        session.send_sync(TextMessage(content=summary, agent="Palette Assistant"))

    except Exception as e:
        logger.warning("Agent summary failed: %s", e)
