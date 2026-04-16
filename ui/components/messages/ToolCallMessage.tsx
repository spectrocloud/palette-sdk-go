"use client";

import { useState } from "react";
import type { ToolCallMessage as ToolCallMessageType } from "@/lib/types";

export function ToolCallMessage({
  message,
}: {
  message: ToolCallMessageType;
}) {
  const [expanded, setExpanded] = useState(false);

  const dotColor =
    message.status === "running"
      ? "#1F7A78"
      : message.status === "complete"
      ? "#2D864E"
      : "#EB5757";

  return (
    <div className="text-[12px]">
      <button
        onClick={() => setExpanded(!expanded)}
        className="flex items-center gap-2 w-full transition-colors"
        style={{ color: "#545f7e" }}
        onMouseOver={(e) => (e.currentTarget.style.color = "#111626")}
        onMouseOut={(e) => (e.currentTarget.style.color = "#545f7e")}
      >
        {message.status === "running" ? (
          <div
            className="w-3 h-3 rounded-full animate-spin shrink-0"
            style={{ border: "1.5px solid #1F7A78", borderTopColor: "transparent" }}
          />
        ) : (
          <div
            className="w-3 h-3 rounded-full shrink-0"
            style={{ backgroundColor: dotColor }}
          />
        )}
        <span className="font-mono font-medium">{message.tool}</span>
        <svg
          className={`w-3 h-3 ml-auto transition-transform ${expanded ? "rotate-180" : ""}`}
          fill="none"
          viewBox="0 0 12 12"
        >
          <path d="M3 5L6 8L9 5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
      </button>

      {expanded && (
        <div
          className="mt-2 ml-5 pl-3 space-y-1"
          style={{ borderLeft: "2px solid #E8EBEE", color: "#545f7e" }}
        >
          {message.input_summary && (
            <p><span className="font-medium" style={{ color: "#111626" }}>Input:</span> {message.input_summary}</p>
          )}
          {message.result_summary && (
            <p><span className="font-medium" style={{ color: "#111626" }}>Result:</span> {message.result_summary}</p>
          )}
        </div>
      )}
    </div>
  );
}
