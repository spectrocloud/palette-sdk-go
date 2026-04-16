// Protocol types — mirrors agent/app/protocol.py

// ---------------------------------------------------------------------------
// Common
// ---------------------------------------------------------------------------

export type FieldType =
  | "string"
  | "number"
  | "boolean"
  | "textarea"
  | "password"
  | "select";

export interface FormField {
  name: string;
  label: string;
  type: FieldType;
  required: boolean;
  description: string;
  default?: unknown;
  enum?: string[];
  placeholder?: string;
}

export interface TableColumn {
  key: string;
  label: string;
  sortable?: boolean;
}

// ---------------------------------------------------------------------------
// Server → Client messages
// ---------------------------------------------------------------------------

export interface TextMessage {
  type: "text";
  content: string;
  agent: string;
}

export interface FormMessage {
  type: "form";
  id: string;
  title: string;
  description?: string;
  fields: FormField[];
  submit_label: string;
  agent: string;
}

export interface SelectMessage {
  type: "select";
  id: string;
  title: string;
  description?: string;
  options: { value: string; label: string }[];
  multiple?: boolean;
  agent: string;
}

export interface TableMessage {
  type: "table";
  title: string;
  columns: TableColumn[];
  rows: Record<string, unknown>[];
  agent: string;
}

export interface ConfirmMessage {
  type: "confirm";
  id: string;
  title: string;
  description: string;
  confirm_label: string;
  cancel_label: string;
  destructive?: boolean;
  agent: string;
}

export interface StatusMessage {
  type: "status";
  message: string;
  progress?: number;
  done?: boolean;
  agent: string;
}

export interface ToolCallMessage {
  type: "tool_call";
  tool: string;
  input_summary?: string;
  result_summary?: string;
  status: "running" | "complete" | "error";
  agent: string;
}

export interface ErrorMessage {
  type: "error";
  message: string;
  agent: string;
}

export interface DebugMessage {
  type: "debug";
  event: string;
  data: Record<string, unknown>;
  timestamp: number;
  agent: string;
}

export interface PlanStep {
  id: string;
  label: string;
  description?: string;
  status: "pending" | "active" | "done" | "skipped" | "error";
  icon?: string;
  children?: PlanStep[];
}

export interface PlanMessage {
  type: "plan";
  id: string;
  title: string;
  description?: string;
  steps: PlanStep[];
  agent: string;
}

export interface PlanUpdateMessage {
  type: "plan_update";
  plan_id: string;
  step_id: string;
  status: "pending" | "active" | "done" | "skipped" | "error";
  message?: string;
  agent: string;
}

export type ServerMessage =
  | TextMessage
  | FormMessage
  | SelectMessage
  | TableMessage
  | ConfirmMessage
  | StatusMessage
  | ToolCallMessage
  | ErrorMessage
  | DebugMessage
  | PlanMessage
  | PlanUpdateMessage;

// ---------------------------------------------------------------------------
// Client → Server messages
// ---------------------------------------------------------------------------

export interface ChatInput {
  type: "chat";
  content: string;
}

export interface FormSubmitInput {
  type: "form_submit";
  id: string;
  data: Record<string, unknown>;
}

export interface SelectResponseInput {
  type: "select_response";
  id: string;
  value: string | string[];
}

export interface ConfirmResponseInput {
  type: "confirm_response";
  id: string;
  confirmed: boolean;
}

export type ClientMessage =
  | ChatInput
  | FormSubmitInput
  | SelectResponseInput
  | ConfirmResponseInput;
