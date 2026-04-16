"use client";

import { useState } from "react";
import type { SelectMessage as SelectMessageType } from "@/lib/types";

interface SelectMessageProps {
  message: SelectMessageType;
  onSelect: (id: string, value: string | string[]) => void;
}

export function SelectMessage({ message, onSelect }: SelectMessageProps) {
  const [selected, setSelected] = useState<Set<string>>(new Set());
  const [submitted, setSubmitted] = useState(false);
  const isMulti = message.multiple;

  function toggle(value: string) {
    if (isMulti) {
      setSelected((prev) => {
        const next = new Set(prev);
        if (next.has(value)) next.delete(value);
        else next.add(value);
        return next;
      });
    } else {
      setSelected(new Set([value]));
    }
  }

  function handleSubmit() {
    if (selected.size === 0) return;
    setSubmitted(true);
    if (isMulti) {
      onSelect(message.id, Array.from(selected));
    } else {
      onSelect(message.id, Array.from(selected)[0]);
    }
  }

  if (submitted) {
    const labels = Array.from(selected).map(
      (v) => message.options.find((o) => o.value === v)?.label ?? v
    );
    return (
      <div className="text-[13px] italic flex items-center gap-2" style={{ color: "#545f7e" }}>
        <div className="w-[7px] h-[7px] rounded-full" style={{ backgroundColor: "#2D864E" }} />
        Selected: <span className="font-medium not-italic" style={{ color: "#111626" }}>
          {labels.join(", ")}
        </span>
      </div>
    );
  }

  return (
    <div className="space-y-3">
      <div>
        <h3 className="text-[14px] font-semibold" style={{ color: "#111626" }}>
          {message.title}
        </h3>
        {message.description && (
          <p className="text-[12px] mt-1" style={{ color: "#545f7e" }}>
            {message.description}
          </p>
        )}
      </div>

      <div className="grid gap-1.5" style={{ maxWidth: "428px" }}>
        {message.options.map((opt) => {
          const isSelected = selected.has(opt.value);
          return (
            <button
              key={opt.value}
              onClick={() => toggle(opt.value)}
              className="text-left px-3 py-2.5 rounded text-[13px] transition-all"
              style={{
                border: `1px solid ${isSelected ? "#1F7A78" : "#E1DDDA"}`,
                backgroundColor: isSelected ? "rgba(31,122,120,0.05)" : "#ffffff",
                color: isSelected ? "#1F7A78" : "#111626",
                fontWeight: isSelected ? 500 : 400,
              }}
            >
              <span className="flex items-center gap-2.5">
                {isMulti ? (
                  <span
                    className="w-4 h-4 rounded flex items-center justify-center shrink-0 text-[10px]"
                    style={{
                      border: `2px solid ${isSelected ? "#1F7A78" : "#E1DDDA"}`,
                      backgroundColor: isSelected ? "#1F7A78" : "transparent",
                      color: "#fff",
                    }}
                  >
                    {isSelected ? "\u2713" : ""}
                  </span>
                ) : (
                  <span
                    className="w-4 h-4 rounded-full flex items-center justify-center shrink-0"
                    style={{ border: `2px solid ${isSelected ? "#1F7A78" : "#E1DDDA"}` }}
                  >
                    {isSelected && (
                      <span className="w-2 h-2 rounded-full" style={{ backgroundColor: "#1F7A78" }} />
                    )}
                  </span>
                )}
                {opt.label}
              </span>
            </button>
          );
        })}
      </div>

      <button
        onClick={handleSubmit}
        disabled={selected.size === 0}
        className="px-5 py-2 rounded text-[13px] font-medium text-white transition-colors disabled:opacity-40"
        style={{ backgroundColor: "#1F7A78" }}
        onMouseOver={(e) => (e.currentTarget.style.backgroundColor = "#135757")}
        onMouseOut={(e) => (e.currentTarget.style.backgroundColor = "#1F7A78")}
      >
        {isMulti ? `Continue (${selected.size} selected)` : "Continue"}
      </button>
    </div>
  );
}
