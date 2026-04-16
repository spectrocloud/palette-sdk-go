"use client";

import Image from "next/image";

export function Sidebar() {
  return (
    <aside
      className="w-[240px] flex flex-col min-h-screen shrink-0"
      style={{ backgroundColor: "#012121", color: "#ffffff" }}
    >
      {/* Logo */}
      <div className="px-5 py-5" style={{ borderBottom: "1px solid rgba(255,255,255,0.1)" }}>
        <div className="flex items-center gap-3">
          <Image src="/logo.svg" alt="Palette" width={28} height={34} />
          <div>
            <h1 className="text-[15px] font-semibold leading-tight" style={{ color: "#ffffff" }}>
              Palette Agentic
            </h1>
            <p className="text-[10px] font-light" style={{ color: "rgba(255,255,255,0.5)" }}>
              Agentic Management
            </p>
          </div>
        </div>
      </div>

      {/* Navigation */}
      <nav className="flex-1 px-3 py-4">
        <div className="space-y-0.5">
          <button
            className="w-full text-left px-3 py-2 rounded text-[13px] font-medium"
            style={{ backgroundColor: "rgba(31,122,120,0.2)", color: "#ffffff" }}
          >
            Chat
          </button>
        </div>

        <div className="mt-8">
          <p
            className="text-[10px] font-medium uppercase tracking-wider mb-3 px-3"
            style={{ color: "rgba(255,255,255,0.3)" }}
          >
            Agents
          </p>
          <div className="space-y-0.5">
            <AgentItem label="Coordinator" color="#1F7A78" />
            <AgentItem label="Infra Provisioner" color="#409BF6" />
            <AgentItem label="Platform Admin" color="#219653" />
            <AgentItem label="App & Workload" color="#F2994A" />
            <AgentItem label="Operations" color="#5FD5F9" />
            <AgentItem label="Access Admin" color="#EB5757" />
          </div>
        </div>
      </nav>

      {/* Footer */}
      <div className="px-5 py-4" style={{ borderTop: "1px solid rgba(255,255,255,0.1)" }}>
        <p className="text-[10px]" style={{ color: "rgba(255,255,255,0.3)" }}>
          Palette SDK v0.1.0
        </p>
      </div>
    </aside>
  );
}

function AgentItem({ label, color }: { label: string; color: string }) {
  return (
    <div
      className="flex items-center gap-2.5 px-3 py-1.5 text-[12px]"
      style={{ color: "rgba(255,255,255,0.5)" }}
    >
      <div
        className="w-[7px] h-[7px] rounded-full shrink-0"
        style={{ backgroundColor: color }}
      />
      {label}
    </div>
  );
}
