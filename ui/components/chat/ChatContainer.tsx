"use client";

import { useState } from "react";
import { useChat } from "@/hooks/useChat";
import { MessageList } from "./MessageList";
import { ChatInput } from "./ChatInput";
import { DebugPanel } from "../debug/DebugPanel";

export function ChatContainer() {
  const {
    messages,
    debugMessages,
    status,
    isThinking,
    statusText,
    sendChat,
    sendFormSubmit,
    sendSelectResponse,
    sendConfirmResponse,
    clearDebug,
  } = useChat();
  const [showDebug, setShowDebug] = useState(false);

  const statusColor =
    status === "connected"
      ? "#2D864E"
      : status === "connecting"
      ? "#F2994A"
      : "#EB5757";

  return (
    <div className="flex h-screen" style={{ backgroundColor: "#F7F9FF" }}>
      {/* Main chat area */}
      <div className="flex-1 flex flex-col min-w-0">
        {/* Header */}
        <header
          className="px-6 py-3 flex items-center justify-between shrink-0"
          style={{
            backgroundColor: "#ffffff",
            borderBottom: "1px solid #DEE1EA",
          }}
        >
          <h2 className="text-[14px] font-medium" style={{ color: "#111626" }}>
            Chat
          </h2>
          <div className="flex items-center gap-4">
            {/* Debug toggle */}
            <button
              onClick={() => setShowDebug(!showDebug)}
              className="flex items-center gap-1.5 text-[11px] font-medium px-2.5 py-1 rounded transition-colors"
              style={{
                backgroundColor: showDebug ? "rgba(31,122,120,0.1)" : "transparent",
                color: showDebug ? "#1F7A78" : "#777",
                border: `1px solid ${showDebug ? "#1F7A78" : "#DEE1EA"}`,
              }}
            >
              <svg width="12" height="12" viewBox="0 0 12 12" fill="none">
                <path d="M6 1v10M1 6h10M3 3l6 6M9 3L3 9" stroke="currentColor" strokeWidth="1" strokeLinecap="round"/>
              </svg>
              Debug
              {debugMessages.length > 0 && (
                <span
                  className="text-[9px] px-1.5 py-0.5 rounded-full font-semibold"
                  style={{ backgroundColor: "#1F7A78", color: "#fff" }}
                >
                  {debugMessages.length}
                </span>
              )}
            </button>

            {/* Connection status */}
            <div className="flex items-center gap-2">
              <div
                className="w-[7px] h-[7px] rounded-full"
                style={{ backgroundColor: statusColor }}
              />
              <span
                className="text-[11px] font-medium capitalize"
                style={{ color: "#545f7e" }}
              >
                {status}
              </span>
            </div>
          </div>
        </header>

        {/* Messages */}
        <MessageList
          messages={messages}
          isThinking={isThinking}
          statusText={statusText}
          onFormSubmit={sendFormSubmit}
          onSelectResponse={sendSelectResponse}
          onConfirmResponse={sendConfirmResponse}
        />

        {/* Input */}
        <ChatInput onSend={sendChat} disabled={status !== "connected"} />
      </div>

      {/* Debug panel (right side, collapsible) */}
      {showDebug && (
        <div
          className="shrink-0"
          style={{ width: "340px", borderLeft: "1px solid #DEE1EA" }}
        >
          <DebugPanel messages={debugMessages} onClear={clearDebug} />
        </div>
      )}
    </div>
  );
}
