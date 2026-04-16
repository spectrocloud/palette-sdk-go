"use client";

import { useState, type KeyboardEvent } from "react";

interface ChatInputProps {
  onSend: (content: string) => void;
  disabled?: boolean;
}

export function ChatInput({ onSend, disabled }: ChatInputProps) {
  const [value, setValue] = useState("");

  function handleSend() {
    const trimmed = value.trim();
    if (!trimmed) return;
    onSend(trimmed);
    setValue("");
  }

  function handleKeyDown(e: KeyboardEvent<HTMLTextAreaElement>) {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      handleSend();
    }
  }

  return (
    <div
      className="p-4"
      style={{ backgroundColor: "#ffffff", borderTop: "1px solid #E8EBEE" }}
    >
      <div className="flex gap-3 max-w-4xl mx-auto">
        <textarea
          value={value}
          onChange={(e) => setValue(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder="Ask Palette Agentic anything..."
          disabled={disabled}
          rows={1}
          style={{
            border: "1px solid #E1DDDA",
            color: "#111626",
            backgroundColor: "#ffffff",
          }}
          className="flex-1 resize-none rounded px-4 py-2.5 text-[13px] placeholder:text-[#777] focus:outline-none focus:border-[#1F7A78] focus:ring-1 focus:ring-[#1F7A78]/30 disabled:opacity-50 disabled:bg-[#F3F0EE]"
        />
        <button
          onClick={handleSend}
          disabled={disabled || !value.trim()}
          className="px-5 py-2.5 rounded text-[13px] font-medium text-white transition-colors disabled:opacity-40 disabled:cursor-not-allowed"
          style={{ backgroundColor: "#1F7A78" }}
          onMouseOver={(e) => (e.currentTarget.style.backgroundColor = "#135757")}
          onMouseOut={(e) => (e.currentTarget.style.backgroundColor = "#1F7A78")}
        >
          Send
        </button>
      </div>
    </div>
  );
}
