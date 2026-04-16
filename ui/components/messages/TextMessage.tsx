"use client";

import ReactMarkdown from "react-markdown";
import type { TextMessage as TextMessageType } from "@/lib/types";

export function TextMessage({ message }: { message: TextMessageType }) {
  return (
    <div
      className="text-[13px] leading-[1.7] [&_p]:mb-2 [&_p:last-child]:mb-0 [&_strong]:font-semibold [&_ul]:list-disc [&_ul]:pl-5 [&_ul]:mb-2 [&_ol]:list-decimal [&_ol]:pl-5 [&_ol]:mb-2 [&_li]:mb-1 [&_a]:underline"
      style={{ color: "#111626" }}
    >
      <ReactMarkdown
        components={{
          code: ({ children }) => (
            <code
              className="px-1.5 py-0.5 rounded text-[12px] font-mono"
              style={{ backgroundColor: "rgba(1,33,33,0.06)", color: "#111626" }}
            >
              {children}
            </code>
          ),
          a: ({ children, href }) => (
            <a href={href} style={{ color: "#1F7A78" }}>{children}</a>
          ),
        }}
      >
        {message.content}
      </ReactMarkdown>
    </div>
  );
}
