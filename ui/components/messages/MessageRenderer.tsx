"use client";

import type { ServerMessage } from "@/lib/types";
import { TextMessage } from "./TextMessage";
import { FormMessage } from "./FormMessage";
import { SelectMessage } from "./SelectMessage";
import { TableMessage } from "./TableMessage";
import { ConfirmMessage } from "./ConfirmMessage";
import { StatusMessage } from "./StatusMessage";
import { ToolCallMessage } from "./ToolCallMessage";
import { PlanMessage } from "./PlanMessage";

interface MessageRendererProps {
  message: ServerMessage;
  onFormSubmit: (id: string, data: Record<string, unknown>) => void;
  onSelectResponse: (id: string, value: string | string[]) => void;
  onConfirmResponse: (id: string, confirmed: boolean) => void;
}

export function MessageRenderer({
  message,
  onFormSubmit,
  onSelectResponse,
  onConfirmResponse,
}: MessageRendererProps) {
  switch (message.type) {
    case "text":
      return <TextMessage message={message} />;
    case "form":
      return <FormMessage message={message} onSubmit={onFormSubmit} />;
    case "select":
      return <SelectMessage message={message} onSelect={onSelectResponse} />;
    case "table":
      return <TableMessage message={message} />;
    case "confirm":
      return <ConfirmMessage message={message} onRespond={onConfirmResponse} />;
    case "status":
      return <StatusMessage message={message} />;
    case "tool_call":
      return <ToolCallMessage message={message} />;
    case "plan":
      return <PlanMessage message={message} />;
    case "error":
      return (
        <div className="flex items-start gap-2 text-[13px]">
          <div
            className="w-[7px] h-[7px] rounded-full mt-1.5 shrink-0"
            style={{ backgroundColor: "#EB5757" }}
          />
          <div style={{ color: "#EB5757" }}>
            <span className="font-semibold">Error:</span> {message.message}
          </div>
        </div>
      );
    default:
      return null;
  }
}
