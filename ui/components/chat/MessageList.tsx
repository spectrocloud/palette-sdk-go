"use client";

import { useEffect, useRef } from "react";
import type { ServerMessage } from "@/lib/types";
import { MessageBubble } from "./MessageBubble";
import { MessageRenderer } from "../messages/MessageRenderer";

interface MessageListProps {
  messages: ServerMessage[];
  isThinking?: boolean;
  statusText?: string;
  onFormSubmit: (id: string, data: Record<string, unknown>) => void;
  onSelectResponse: (id: string, value: string | string[]) => void;
  onConfirmResponse: (id: string, confirmed: boolean) => void;
}

export function MessageList({
  messages,
  isThinking,
  statusText,
  onFormSubmit,
  onSelectResponse,
  onConfirmResponse,
}: MessageListProps) {
  const bottomRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages, isThinking, statusText]);

  return (
    <div className="flex-1 overflow-y-auto p-6">
      <div className="max-w-4xl mx-auto space-y-4">
        {messages.map((msg, i) => (
          <MessageBubble key={i} agent={msg.agent}>
            <MessageRenderer
              message={msg}
              onFormSubmit={onFormSubmit}
              onSelectResponse={onSelectResponse}
              onConfirmResponse={onConfirmResponse}
            />
          </MessageBubble>
        ))}
        {(isThinking || statusText) && <ThinkingIndicator text={statusText} />}
        <div ref={bottomRef} />
      </div>
    </div>
  );
}

function ThinkingIndicator({ text }: { text?: string }) {
  return (
    <div className="flex items-center gap-3 px-4 py-3">
      <div
        className="w-[28px] h-[28px] rounded-full flex items-center justify-center shrink-0"
        style={{ backgroundColor: "#1F7A78" }}
      >
        <svg
          className="animate-spin"
          width="14"
          height="14"
          viewBox="0 0 14 14"
          fill="none"
        >
          <circle
            cx="7"
            cy="7"
            r="5.5"
            stroke="rgba(255,255,255,0.3)"
            strokeWidth="2"
          />
          <path
            d="M12.5 7a5.5 5.5 0 00-5.5-5.5"
            stroke="#ffffff"
            strokeWidth="2"
            strokeLinecap="round"
          />
        </svg>
      </div>
      <span
        className="text-[13px] font-medium"
        style={{ color: "#545f7e" }}
      >
        {text || "Processing..."}
      </span>
    </div>
  );
}
