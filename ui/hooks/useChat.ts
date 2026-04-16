"use client";

import { useCallback, useEffect, useRef, useState } from "react";
import type {
  ClientMessage,
  DebugMessage,
  ServerMessage,
} from "@/lib/types";
import { createWSClient, type WSClient, type WSStatus } from "@/lib/websocket";

export function useChat() {
  const [messages, setMessages] = useState<ServerMessage[]>([]);
  const [debugMessages, setDebugMessages] = useState<DebugMessage[]>([]);
  const [status, setStatus] = useState<WSStatus>("disconnected");
  const [isThinking, setIsThinking] = useState(false);
  const [statusText, setStatusText] = useState<string>("");
  const wsRef = useRef<WSClient | null>(null);

  useEffect(() => {
    const client = createWSClient(
      (msg) => {
        // Debug messages go to separate list
        if (msg.type === "debug") {
          setDebugMessages((prev) => [...prev, msg as DebugMessage]);
          return;
        }

        // Status messages control the status indicator (independent of isThinking)
        if (msg.type === "status") {
          const status = msg as { done?: boolean; message?: string };
          if (status.done) {
            setStatusText("");
          } else {
            setStatusText(status.message || "Processing...");
          }
          return;
        }

        // Any visible message from the server clears the initial thinking spinner
        // but does NOT clear statusText (which has its own done signal)
        if (msg.type !== "tool_call" && msg.type !== "plan_update") {
          setIsThinking(false);
        }

        setMessages((prev) => {
          // Skip tool_call messages
          if (msg.type === "tool_call") {
            return prev;
          }

          // Plan updates: update the step status in the existing plan message
          if (msg.type === "plan_update") {
            const update = msg as import("@/lib/types").PlanUpdateMessage;
            return prev.map((m) => {
              if (m.type === "plan" && m.id === update.plan_id) {
                const plan = m as import("@/lib/types").PlanMessage;
                return {
                  ...plan,
                  steps: plan.steps.map((s) =>
                    s.id === update.step_id ? { ...s, status: update.status } : s
                  ),
                };
              }
              return m;
            });
          }

          // Skip duplicate error messages
          if (msg.type === "error") {
            const lastErr = [...prev].reverse().find((m) => m.type === "error");
            if (lastErr && lastErr.type === "error" && lastErr.message === (msg as { message: string }).message) {
              return prev;
            }
          }

          return [...prev, msg];
        });
      },
      (s) => setStatus(s)
    );
    wsRef.current = client;
    return () => client.close();
  }, []);

  const send = useCallback((msg: ClientMessage) => {
    wsRef.current?.send(msg);
    if (msg.type === "chat") {
      setIsThinking(true);
      setMessages((prev) => [
        ...prev,
        { type: "text", content: msg.content, agent: "user" } as ServerMessage,
      ]);
    }
  }, []);

  const sendChat = useCallback(
    (content: string) => send({ type: "chat", content }),
    [send]
  );

  const sendFormSubmit = useCallback(
    (id: string, data: Record<string, unknown>) =>
      send({ type: "form_submit", id, data }),
    [send]
  );

  const sendSelectResponse = useCallback(
    (id: string, value: string | string[]) =>
      send({ type: "select_response", id, value }),
    [send]
  );

  const sendConfirmResponse = useCallback(
    (id: string, confirmed: boolean) =>
      send({ type: "confirm_response", id, confirmed }),
    [send]
  );

  const clearDebug = useCallback(() => setDebugMessages([]), []);

  return {
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
  };
}
