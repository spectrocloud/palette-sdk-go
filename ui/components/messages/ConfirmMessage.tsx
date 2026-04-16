"use client";

import { useState } from "react";
import type { ConfirmMessage as ConfirmMessageType } from "@/lib/types";

interface ConfirmMessageProps {
  message: ConfirmMessageType;
  onRespond: (id: string, confirmed: boolean) => void;
}

export function ConfirmMessage({ message, onRespond }: ConfirmMessageProps) {
  const [responded, setResponded] = useState(false);
  const [choice, setChoice] = useState(false);

  function handleRespond(confirmed: boolean) {
    setChoice(confirmed);
    setResponded(true);
    onRespond(message.id, confirmed);
  }

  if (responded) {
    return (
      <div className="text-[13px] italic flex items-center gap-2" style={{ color: "#545f7e" }}>
        <div
          className="w-[7px] h-[7px] rounded-full"
          style={{ backgroundColor: choice ? "#2D864E" : "#777" }}
        />
        {choice ? "Confirmed" : "Cancelled"}
      </div>
    );
  }

  const confirmBg = message.destructive ? "#DD314F" : "#1F7A78";
  const confirmHover = message.destructive ? "#BD344C" : "#135757";

  return (
    <div className="space-y-3">
      <div>
        <h3 className="text-[14px] font-semibold" style={{ color: "#111626" }}>
          {message.title}
        </h3>
        <p className="text-[13px] mt-1 leading-relaxed" style={{ color: "#545f7e" }}>
          {message.description}
        </p>
      </div>

      <div className="flex gap-2">
        <button
          onClick={() => handleRespond(true)}
          className="px-5 py-2 rounded text-[13px] font-medium text-white transition-colors"
          style={{ backgroundColor: confirmBg }}
          onMouseOver={(e) => (e.currentTarget.style.backgroundColor = confirmHover)}
          onMouseOut={(e) => (e.currentTarget.style.backgroundColor = confirmBg)}
        >
          {message.confirm_label}
        </button>
        <button
          onClick={() => handleRespond(false)}
          className="px-5 py-2 rounded text-[13px] font-medium transition-colors"
          style={{
            backgroundColor: "#EDEEF4",
            color: "#111626",
            border: "1px solid #E1DDDA",
          }}
          onMouseOver={(e) => (e.currentTarget.style.backgroundColor = "#F3F0EE")}
          onMouseOut={(e) => (e.currentTarget.style.backgroundColor = "#EDEEF4")}
        >
          {message.cancel_label}
        </button>
      </div>
    </div>
  );
}
