"""Per-WebSocket-connection session state."""

from __future__ import annotations

import asyncio
import logging
import threading
import uuid
from typing import Any

from starlette.websockets import WebSocket

from .protocol import ServerMessage

logger = logging.getLogger(__name__)


class _PendingInput:
    """Thread-safe container for waiting on user input."""

    def __init__(self):
        self.event = threading.Event()
        self.data: Any = None

    def set(self, data: Any) -> None:
        self.data = data
        self.event.set()

    def wait(self, timeout: float = 300.0) -> Any:
        if not self.event.wait(timeout=timeout):
            raise TimeoutError(f"No response within {timeout}s")
        return self.data


class Session:
    """Manages state for a single WebSocket connection."""

    def __init__(self, ws: WebSocket) -> None:
        self.id = str(uuid.uuid4())
        self.ws = ws
        self._pending: dict[str, _PendingInput] = {}
        self._lock = threading.Lock()
        self._loop = asyncio.get_event_loop()
        self.tool_sent_messages = False
        self.chat_queue: asyncio.Queue = asyncio.Queue()
        self.conversation_history: list[dict] = []  # [{role, content}]
        self.artifacts: dict[str, Any] = {}  # Data passed between journeys

    async def send(self, message: ServerMessage) -> None:
        """Send a structured message to the client (async)."""
        await self.ws.send_json(message.model_dump())

    def send_sync(self, message: ServerMessage) -> None:
        """Send a message from a sync thread context."""
        self.tool_sent_messages = True
        asyncio.run_coroutine_threadsafe(self.send(message), self._loop)
        # Brief yield to let event loop process
        import time
        time.sleep(0.05)

    def register_pending(self, request_id: str) -> _PendingInput:
        """Register a pending input BEFORE sending the UI message."""
        with self._lock:
            pending = _PendingInput()
            self._pending[request_id] = pending
            return pending

    def wait_for_input_sync(self, request_id: str, timeout: float = 300.0) -> Any:
        """Block the current thread until the client responds."""
        with self._lock:
            pending = self._pending.get(request_id)
        if pending is None:
            # Not pre-registered, create now
            pending = self.register_pending(request_id)

        logger.info("Session %s: waiting for input id=%s", self.id, request_id)
        result = pending.wait(timeout=timeout)
        logger.info("Session %s: got input id=%s", self.id, request_id)

        with self._lock:
            self._pending.pop(request_id, None)
        return result

    def resolve_input(self, request_id: str, data: Any) -> bool:
        """Resolve a pending input with the client's response."""
        with self._lock:
            pending = self._pending.get(request_id)
        if pending is None:
            logger.warning("No pending input for id=%s", request_id)
            return False
        pending.set(data)
        return True

    # Keep old async methods for compatibility but they're unused now
    def request_input(self, request_id: str):
        return self.register_pending(request_id)
