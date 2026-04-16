"use client";

import { useEffect, useRef } from "react";
import type { DebugMessage } from "@/lib/types";

const EVENT_STYLES: Record<string, { color: string; icon: string; label: string }> = {
  route: { color: "#1F7A78", icon: "\u2192", label: "ROUTE" },
  crew_start: { color: "#409BF6", icon: "\u25B6", label: "CREW" },
  tool_call: { color: "#F2994A", icon: "\u2699", label: "TOOL" },
  tool_result: { color: "#219653", icon: "\u2713", label: "RESULT" },
  crew_done: { color: "#219653", icon: "\u2714", label: "DONE" },
  error: { color: "#EB5757", icon: "\u2717", label: "ERROR" },
};

interface DebugPanelProps {
  messages: DebugMessage[];
  onClear: () => void;
}

export function DebugPanel({ messages, onClear }: DebugPanelProps) {
  const bottomRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  return (
    <div
      className="flex flex-col h-full"
      style={{ backgroundColor: "#0B3433", color: "#e0e0e0" }}
    >
      {/* Header */}
      <div
        className="flex items-center justify-between px-4 py-2 shrink-0"
        style={{ borderBottom: "1px solid rgba(255,255,255,0.1)" }}
      >
        <span className="text-[11px] font-semibold uppercase tracking-wider" style={{ color: "#1F7A78" }}>
          Debug
        </span>
        <button
          onClick={onClear}
          className="text-[10px] px-2 py-0.5 rounded"
          style={{ color: "rgba(255,255,255,0.4)", backgroundColor: "rgba(255,255,255,0.05)" }}
        >
          Clear
        </button>
      </div>

      {/* Events */}
      <div className="flex-1 overflow-y-auto px-3 py-2">
        {messages.length === 0 && (
          <p className="text-[11px]" style={{ color: "rgba(255,255,255,0.3)" }}>
            No events yet. Send a message to see the execution trace.
          </p>
        )}
        {messages.map((msg, i) => (
          <DebugEvent key={i} msg={msg} prevTimestamp={i > 0 ? messages[i - 1].timestamp : null} />
        ))}
        <div ref={bottomRef} />
      </div>
    </div>
  );
}

function DebugEvent({ msg, prevTimestamp }: { msg: DebugMessage; prevTimestamp: number | null }) {
  const style = EVENT_STYLES[msg.event] || { color: "#777", icon: "\u2022", label: msg.event.toUpperCase() };
  const delta = prevTimestamp ? ((msg.timestamp - prevTimestamp) * 1000).toFixed(0) : null;
  const time = new Date(msg.timestamp * 1000).toLocaleTimeString("en-US", {
    hour12: false,
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });

  return (
    <div className="mb-2" style={{ fontSize: "11px", lineHeight: "1.5" }}>
      {/* Header line */}
      <div className="flex items-center gap-1.5">
        <span style={{ color: style.color }}>{style.icon}</span>
        <span className="font-semibold" style={{ color: style.color }}>{style.label}</span>
        <span style={{ color: "rgba(255,255,255,0.3)" }}>{time}</span>
        {delta && (
          <span style={{ color: "rgba(255,255,255,0.2)" }}>+{delta}ms</span>
        )}
      </div>

      {/* Data */}
      <div className="ml-5 mt-0.5" style={{ color: "rgba(255,255,255,0.6)" }}>
        {Object.entries(msg.data).map(([key, value]) => (
          <div key={key} className="flex gap-1.5">
            <span style={{ color: "rgba(255,255,255,0.35)" }}>{key}:</span>
            <span className="font-mono" style={{ wordBreak: "break-all" }}>
              {typeof value === "object" ? JSON.stringify(value) : String(value)}
            </span>
          </div>
        ))}
      </div>
    </div>
  );
}
