"use client";

import { useState } from "react";
import type { TableMessage as TableMessageType } from "@/lib/types";

export function TableMessage({ message }: { message: TableMessageType }) {
  const [sortKey, setSortKey] = useState<string>("");
  const [sortAsc, setSortAsc] = useState(true);

  const sortedRows = [...message.rows].sort((a, b) => {
    if (!sortKey) return 0;
    const av = String(a[sortKey] ?? "");
    const bv = String(b[sortKey] ?? "");
    return sortAsc ? av.localeCompare(bv) : bv.localeCompare(av);
  });

  function handleSort(key: string) {
    if (sortKey === key) {
      setSortAsc(!sortAsc);
    } else {
      setSortKey(key);
      setSortAsc(true);
    }
  }

  return (
    <div className="space-y-2">
      <h3 className="text-[13px] font-semibold" style={{ color: "#111626" }}>
        {message.title}
      </h3>
      <div
        className="overflow-x-auto rounded"
        style={{
          border: "1px solid #DEE1EA",
          backgroundColor: "#ffffff",
          boxShadow: "0px 2px 0px rgba(0,0,0,0.03), 0px 0px 2px rgba(0,0,0,0.2)",
        }}
      >
        <table className="w-full text-[13px]">
          <thead>
            <tr style={{ backgroundColor: "#FAF9F8", borderBottom: "1px solid #E8EBEE" }}>
              {message.columns.map((col) => (
                <th
                  key={col.key}
                  className={`text-left px-3 py-2.5 text-[11px] font-semibold uppercase tracking-wider ${
                    col.sortable ? "cursor-pointer select-none" : ""
                  }`}
                  style={{ color: "#545f7e" }}
                  onClick={col.sortable ? () => handleSort(col.key) : undefined}
                  onMouseOver={col.sortable ? (e) => (e.currentTarget.style.color = "#111626") : undefined}
                  onMouseOut={col.sortable ? (e) => (e.currentTarget.style.color = "#545f7e") : undefined}
                >
                  <span className="flex items-center gap-1">
                    {col.label}
                    {col.sortable && sortKey === col.key && (
                      <span style={{ color: "#1F7A78" }}>
                        {sortAsc ? "\u25B2" : "\u25BC"}
                      </span>
                    )}
                  </span>
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {sortedRows.map((row, i) => (
              <tr
                key={i}
                className="transition-colors"
                style={{ borderBottom: i < sortedRows.length - 1 ? "1px solid #E8EBEE" : "none" }}
                onMouseOver={(e) => (e.currentTarget.style.backgroundColor = "#FAF9F8")}
                onMouseOut={(e) => (e.currentTarget.style.backgroundColor = "transparent")}
              >
                {message.columns.map((col) => (
                  <td
                    key={col.key}
                    className="px-3 py-2.5 text-[13px]"
                    style={{ color: "#111626" }}
                  >
                    {String(row[col.key] ?? "")}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
        {sortedRows.length === 0 && (
          <p className="text-center text-[12px] py-6" style={{ color: "#545f7e" }}>
            No data available
          </p>
        )}
      </div>
      <p className="text-[11px]" style={{ color: "#777" }}>
        {sortedRows.length} {sortedRows.length === 1 ? "row" : "rows"}
      </p>
    </div>
  );
}
