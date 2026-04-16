"use client";

import type { PlanMessage as PlanMessageType, PlanStep } from "@/lib/types";

const STATUS_STYLES: Record<string, { dot: string; line: string; text: string }> = {
  pending: { dot: "#DEE1EA", line: "#DEE1EA", text: "#777" },
  active: { dot: "#1F7A78", line: "#1F7A78", text: "#111626" },
  done: { dot: "#2D864E", line: "#2D864E", text: "#111626" },
  skipped: { dot: "#F2994A", line: "#F2994A", text: "#777" },
  error: { dot: "#EB5757", line: "#EB5757", text: "#EB5757" },
};

const STATUS_ICONS: Record<string, string> = {
  pending: "\u25CB",   // ○
  active: "\u25CF",    // ●
  done: "\u2713",      // ✓
  skipped: "\u2014",   // —
  error: "\u2717",     // ✗
};

export function PlanMessage({ message }: { message: PlanMessageType }) {
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

      <div className="space-y-0">
        {message.steps.map((step, i) => (
          <StepRow
            key={step.id}
            step={step}
            isLast={i === message.steps.length - 1}
          />
        ))}
      </div>
    </div>
  );
}

function StepRow({ step, isLast }: { step: PlanStep; isLast: boolean }) {
  const style = STATUS_STYLES[step.status] || STATUS_STYLES.pending;
  const icon = step.icon || STATUS_ICONS[step.status] || STATUS_ICONS.pending;
  const isActive = step.status === "active";

  return (
    <div className="flex">
      {/* Vertical line + dot */}
      <div className="flex flex-col items-center mr-3" style={{ width: "24px" }}>
        <div
          className="w-6 h-6 rounded-full flex items-center justify-center text-[12px] font-bold shrink-0"
          style={{
            backgroundColor: step.status === "done" ? style.dot : "transparent",
            border: `2px solid ${style.dot}`,
            color: step.status === "done" ? "#fff" : style.dot,
            boxShadow: isActive ? `0 0 0 2px #fff, 0 0 0 4px ${style.dot}` : undefined,
          }}
        >
          {icon}
        </div>
        {!isLast && (
          <div
            className="w-0.5 flex-1 min-h-[20px]"
            style={{ backgroundColor: style.line }}
          />
        )}
      </div>

      {/* Content */}
      <div className="pb-4 pt-0.5 flex-1">
        <p
          className={`text-[13px] ${isActive ? "font-semibold" : "font-medium"}`}
          style={{ color: style.text }}
        >
          {step.label}
        </p>
        {step.description && (
          <p className="text-[11px] mt-0.5" style={{ color: "#777" }}>
            {step.description}
          </p>
        )}

        {/* Children (sub-steps) */}
        {step.children && step.children.length > 0 && (
          <div className="mt-2 ml-1 space-y-0">
            {step.children.map((child, j) => (
              <StepRow
                key={child.id}
                step={child}
                isLast={j === step.children!.length - 1}
              />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
