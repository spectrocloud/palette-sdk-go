"""FastAPI + WebSocket backend for Palette Agentic."""

from __future__ import annotations

import asyncio
import logging
import time

from fastapi import FastAPI, WebSocket, WebSocketDisconnect
from fastapi.middleware.cors import CORSMiddleware

from .config import settings
from .crews.palette_crew import (
    classify_intent,
    interpret_result,
    answer_question,
    _match_journey,
    _match_specialist,
    _build_direct_call,
    run_journey,
)
from .protocol import (
    DebugMessage,
    ErrorMessage,
    StatusMessage,
    TextMessage,
    parse_client_message,
)
from .session import Session

logging.basicConfig(level=settings.log_level.upper())
logger = logging.getLogger(__name__)

app = FastAPI(title="Palette Agentic", version="0.1.0")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/health")
async def health() -> dict[str, str]:
    return {"status": "ok"}


async def _input_listener(ws: WebSocket, session: Session):
    """Background task: reads WebSocket messages and resolves pending inputs."""
    try:
        while True:
            raw = await ws.receive_json()
            try:
                msg = parse_client_message(raw)
            except ValueError:
                continue

            if msg.type == "form_submit":
                session.resolve_input(msg.id, msg.data)
            elif msg.type == "select_response":
                session.resolve_input(msg.id, msg.value)
            elif msg.type == "confirm_response":
                session.resolve_input(msg.id, msg.confirmed)
            elif msg.type == "chat":
                # Queue chat messages for the main handler
                session.chat_queue.put_nowait(msg)
    except Exception:
        pass  # Connection closed


@app.websocket("/ws/chat")
async def websocket_chat(ws: WebSocket) -> None:
    await ws.accept()
    session = Session(ws)
    logger.info("New session: %s", session.id)

    # Start background input listener
    listener_task = asyncio.create_task(_input_listener(ws, session))

    await session.send(TextMessage(content="Connected. Ready to help.", agent="system"))

    try:
        while True:
            # Wait for chat messages from the listener
            msg = await session.chat_queue.get()

            session.tool_sent_messages = False
            t0 = time.time()

            # Store user message in conversation history
            session.conversation_history.append({"role": "user", "content": msg.content})

            intent = classify_intent(msg.content)
            agent_key = _match_specialist(msg.content)

            await session.send(DebugMessage(
                event="route",
                data={"intent": intent, "agent": agent_key},
                timestamp=time.time(),
            ))

            try:
                if intent == "journey":
                    journey = _match_journey(msg.content)
                    await session.send(DebugMessage(
                        event="journey_start",
                        data={"journey": journey["name"], "agents": journey["agents"]},
                        timestamp=time.time(),
                    ))
                    summary = await asyncio.to_thread(
                        run_journey, journey, msg.content, session
                    )
                    if summary:
                        await session.send(TextMessage(content=summary, agent="Palette Assistant"))

                else:
                    direct = _build_direct_call(agent_key, session)

                    await session.send(DebugMessage(
                        event="llm_start",
                        data={"purpose": "pick tool", "llm": settings.crewai_llm},
                        timestamp=time.time(),
                    ))

                    tool_name, tool_input = await asyncio.to_thread(
                        direct["pick_tool"], msg.content
                    )

                    await session.send(DebugMessage(
                        event="llm_done",
                        data={"tool": tool_name, "input": str(tool_input)[:150]},
                        timestamp=time.time(),
                    ))

                    if tool_name and tool_name in direct["tools"]:
                        tool = direct["tools"][tool_name]

                        tool_result = await asyncio.to_thread(tool._run, tool_input)

                        # Store in conversation history
                        session.conversation_history.append({
                            "role": "assistant",
                            "content": f"Called {tool_name}. Result: {str(tool_result)[:500]}",
                        })

                        # Run analysis if intent is interpret, or if the user asked a question
                        _msg_lower = msg.content.lower().strip()
                        is_question = (
                            "?" in msg.content
                            or _msg_lower.startswith(("what", "which", "how", "why", "is ", "are ", "do ", "does ", "can ", "any"))
                        )
                        if intent == "interpret" or is_question:
                            await session.send(StatusMessage(
                                message="Analyzing results...", done=False, agent=agent_key,
                            ))
                            try:
                                await session.send(DebugMessage(
                                    event="llm_start",
                                    data={"purpose": "interpret results"},
                                    timestamp=time.time(),
                                ))
                                analysis = await asyncio.to_thread(
                                    interpret_result, msg.content, tool_name, tool_result
                                )
                                await session.send(TextMessage(content=analysis, agent=agent_key))
                            except Exception as e:
                                logger.error("interpret_result failed: %s", e)
                                await session.send(TextMessage(
                                    content=f"Analysis failed: {e}",
                                    agent=agent_key,
                                ))
                            finally:
                                await session.send(StatusMessage(
                                    message="", done=True, agent=agent_key,
                                ))
                                await session.send(DebugMessage(
                                    event="llm_done",
                                    data={"purpose": "interpret"},
                                    timestamp=time.time(),
                                ))
                    else:
                        # No tool matched — use LLM knowledge to answer
                        await session.send(DebugMessage(
                            event="llm_start",
                            data={"purpose": "answer from knowledge"},
                            timestamp=time.time(),
                        ))
                        answer = await asyncio.to_thread(
                            answer_question, msg.content, agent_key
                        )
                        await session.send(DebugMessage(
                            event="llm_done",
                            data={"purpose": "answer", "length": len(answer)},
                            timestamp=time.time(),
                        ))
                        await session.send(TextMessage(content=answer, agent=agent_key))

                await session.send(DebugMessage(
                    event="done",
                    data={"elapsed_sec": round(time.time() - t0, 2)},
                    timestamp=time.time(),
                ))

            except Exception as e:
                logger.exception("Error handling request")
                await session.send(ErrorMessage(message=f"Error: {e}", agent="system"))

    except (WebSocketDisconnect, asyncio.CancelledError):
        logger.info("Session %s disconnected", session.id)
    except Exception:
        logger.exception("WebSocket error for session %s", session.id)
    finally:
        listener_task.cancel()


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host=settings.host, port=settings.port)
