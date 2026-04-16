"use client";

import type { ReactNode } from "react";

const AGENT_STYLES: Record<string, { bg: string; border: string; dot: string }> = {
  user: { bg: "rgba(31,122,120,0.06)", border: "rgba(31,122,120,0.15)", dot: "#1F7A78" },
  system: { bg: "#FAF9F8", border: "#E8EBEE", dot: "#777" },
  Coordinator: { bg: "rgba(31,122,120,0.06)", border: "rgba(31,122,120,0.15)", dot: "#1F7A78" },
  "Infra Provisioner": { bg: "rgba(64,155,246,0.06)", border: "rgba(64,155,246,0.15)", dot: "#409BF6" },
  "Platform Admin": { bg: "rgba(33,150,83,0.06)", border: "rgba(33,150,83,0.15)", dot: "#219653" },
  "App & Workload": { bg: "rgba(242,153,74,0.06)", border: "rgba(242,153,74,0.15)", dot: "#F2994A" },
  Operations: { bg: "rgba(95,213,249,0.06)", border: "rgba(95,213,249,0.15)", dot: "#5FD5F9" },
  "Access Admin": { bg: "rgba(235,87,87,0.06)", border: "rgba(235,87,87,0.15)", dot: "#EB5757" },
};

interface MessageBubbleProps {
  agent: string;
  children: ReactNode;
}

export function MessageBubble({ agent, children }: MessageBubbleProps) {
  const isUser = agent === "user";
  const style = AGENT_STYLES[agent] || { bg: "#ffffff", border: "#DEE1EA", dot: "#777" };

  return (
    <div className={`flex ${isUser ? "justify-end" : "justify-start"}`}>
      <div
        className="rounded px-4 py-3 max-w-[85%]"
        style={{
          backgroundColor: style.bg,
          border: `1px solid ${style.border}`,
        }}
      >
        {!isUser && agent && (
          <div className="flex items-center gap-2 mb-2">
            <div
              className="w-[7px] h-[7px] rounded-full shrink-0"
              style={{ backgroundColor: style.dot }}
            />
            <span
              className="text-[11px] font-semibold uppercase tracking-wide"
              style={{ color: "#545f7e" }}
            >
              {agent}
            </span>
          </div>
        )}
        {children}
      </div>
    </div>
  );
}
