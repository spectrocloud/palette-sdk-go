"use client";

import type { StatusMessage as StatusMessageType } from "@/lib/types";

export function StatusMessage({ message }: { message: StatusMessageType }) {
  const percent =
    message.progress != null ? Math.round(message.progress * 100) : null;

  return (
    <div className="space-y-2">
      <div className="flex items-center gap-2.5 text-[13px]">
        {!message.done && (
          <div
            className="w-4 h-4 rounded-full animate-spin"
            style={{ border: "2px solid #1F7A78", borderTopColor: "transparent" }}
          />
        )}
        {message.done && (
          <div
            className="w-4 h-4 rounded-full flex items-center justify-center"
            style={{ backgroundColor: "#2D864E" }}
          >
            <svg width="10" height="10" viewBox="0 0 10 10" fill="none">
              <path d="M2 5L4 7L8 3" stroke="white" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
            </svg>
          </div>
        )}
        <span
          className="font-medium"
          style={{ color: message.done ? "#2D864E" : "#545f7e" }}
        >
          {message.message}
        </span>
      </div>

      {percent != null && (
        <div className="w-full h-1.5 rounded-full overflow-hidden" style={{ backgroundColor: "#E8EBEE" }}>
          <div
            className="h-full rounded-full transition-all duration-500 ease-out"
            style={{ width: `${percent}%`, backgroundColor: "#1F7A78" }}
          />
        </div>
      )}
    </div>
  );
}
