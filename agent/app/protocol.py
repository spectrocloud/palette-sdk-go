"""Structured message protocol — source of truth for Python side.

Defines all server→client and client→server message types that flow
over the WebSocket connection between the frontend and the agent backend.
"""

from __future__ import annotations

from enum import Enum
from typing import Any, Literal

from pydantic import BaseModel, Field


# ---------------------------------------------------------------------------
# Common
# ---------------------------------------------------------------------------


class FieldType(str, Enum):
    STRING = "string"
    NUMBER = "number"
    BOOLEAN = "boolean"
    TEXTAREA = "textarea"
    PASSWORD = "password"
    SELECT = "select"


class FormField(BaseModel):
    name: str
    label: str
    type: FieldType = FieldType.STRING
    required: bool = False
    description: str = ""
    default: Any = None
    enum: list[str] | None = None
    placeholder: str = ""


class TableColumn(BaseModel):
    key: str
    label: str
    sortable: bool = False


# ---------------------------------------------------------------------------
# Server → Client messages
# ---------------------------------------------------------------------------


class TextMessage(BaseModel):
    type: Literal["text"] = "text"
    content: str
    agent: str = ""


class FormMessage(BaseModel):
    type: Literal["form"] = "form"
    id: str
    title: str
    description: str = ""
    fields: list[FormField]
    submit_label: str = "Submit"
    agent: str = ""


class SelectMessage(BaseModel):
    type: Literal["select"] = "select"
    id: str
    title: str
    description: str = ""
    options: list[dict[str, str]]  # [{value, label}]
    multiple: bool = False
    agent: str = ""


class TableMessage(BaseModel):
    type: Literal["table"] = "table"
    title: str
    columns: list[TableColumn]
    rows: list[dict[str, Any]]
    agent: str = ""


class ConfirmMessage(BaseModel):
    type: Literal["confirm"] = "confirm"
    id: str
    title: str
    description: str
    confirm_label: str = "Confirm"
    cancel_label: str = "Cancel"
    destructive: bool = False
    agent: str = ""


class StatusMessage(BaseModel):
    type: Literal["status"] = "status"
    message: str
    progress: float | None = None  # 0.0–1.0
    done: bool = False
    agent: str = ""


class ToolCallMessage(BaseModel):
    type: Literal["tool_call"] = "tool_call"
    tool: str
    input_summary: str = ""
    result_summary: str = ""
    status: Literal["running", "complete", "error"] = "running"
    agent: str = ""


class ErrorMessage(BaseModel):
    type: Literal["error"] = "error"
    message: str
    agent: str = ""


class DebugMessage(BaseModel):
    type: Literal["debug"] = "debug"
    event: str
    data: dict[str, Any] = {}
    timestamp: float = 0.0


class PlanStep(BaseModel):
    id: str
    label: str
    description: str = ""
    status: Literal["pending", "active", "done", "skipped", "error"] = "pending"
    icon: str = ""  # emoji or icon name
    children: list["PlanStep"] = []


class PlanMessage(BaseModel):
    """Rich visual plan — rendered as a flowchart/checklist with status."""
    type: Literal["plan"] = "plan"
    id: str
    title: str
    description: str = ""
    steps: list[PlanStep]
    agent: str = ""


class PlanUpdateMessage(BaseModel):
    """Update the status of a step in an existing plan."""
    type: Literal["plan_update"] = "plan_update"
    plan_id: str
    step_id: str
    status: Literal["pending", "active", "done", "skipped", "error"] = "done"
    message: str = ""


# Union of all server messages for deserialization
ServerMessage = (
    TextMessage
    | FormMessage
    | SelectMessage
    | TableMessage
    | ConfirmMessage
    | StatusMessage
    | ToolCallMessage
    | ErrorMessage
    | DebugMessage
    | PlanMessage
    | PlanUpdateMessage
)


# ---------------------------------------------------------------------------
# Client → Server messages
# ---------------------------------------------------------------------------


class ChatInput(BaseModel):
    type: Literal["chat"] = "chat"
    content: str


class FormSubmitInput(BaseModel):
    type: Literal["form_submit"] = "form_submit"
    id: str
    data: dict[str, Any]


class SelectResponseInput(BaseModel):
    type: Literal["select_response"] = "select_response"
    id: str
    value: str | list[str]


class ConfirmResponseInput(BaseModel):
    type: Literal["confirm_response"] = "confirm_response"
    id: str
    confirmed: bool


ClientMessage = ChatInput | FormSubmitInput | SelectResponseInput | ConfirmResponseInput


def parse_client_message(raw: dict[str, Any]) -> ClientMessage:
    """Parse a raw dict into the correct ClientMessage subtype."""
    msg_type = raw.get("type")
    match msg_type:
        case "chat":
            return ChatInput(**raw)
        case "form_submit":
            return FormSubmitInput(**raw)
        case "select_response":
            return SelectResponseInput(**raw)
        case "confirm_response":
            return ConfirmResponseInput(**raw)
        case _:
            raise ValueError(f"Unknown client message type: {msg_type}")
