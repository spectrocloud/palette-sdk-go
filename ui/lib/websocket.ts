import type { ClientMessage, ServerMessage } from "./types";

const WS_URL = process.env.NEXT_PUBLIC_WS_URL || "ws://localhost:8080/ws/chat";
const RECONNECT_DELAY = 2000;
const MAX_RECONNECT_ATTEMPTS = 10;

export type WSStatus = "connecting" | "connected" | "disconnected" | "error";

export interface WSClient {
  send: (msg: ClientMessage) => void;
  close: () => void;
  status: WSStatus;
}

export function createWSClient(
  onMessage: (msg: ServerMessage) => void,
  onStatusChange: (status: WSStatus) => void
): WSClient {
  let ws: WebSocket | null = null;
  let reconnectAttempts = 0;
  let status: WSStatus = "disconnected";
  let closed = false;

  function setStatus(s: WSStatus) {
    status = s;
    onStatusChange(s);
  }

  function connect() {
    if (closed) return;
    setStatus("connecting");

    ws = new WebSocket(WS_URL);

    ws.onopen = () => {
      reconnectAttempts = 0;
      setStatus("connected");
    };

    ws.onmessage = (event) => {
      try {
        const msg = JSON.parse(event.data) as ServerMessage;
        onMessage(msg);
      } catch {
        console.error("Failed to parse WebSocket message:", event.data);
      }
    };

    ws.onclose = () => {
      if (closed) return;
      setStatus("disconnected");
      if (reconnectAttempts < MAX_RECONNECT_ATTEMPTS) {
        reconnectAttempts++;
        setTimeout(connect, RECONNECT_DELAY);
      }
    };

    ws.onerror = () => {
      setStatus("error");
    };
  }

  connect();

  return {
    send(msg: ClientMessage) {
      if (ws?.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify(msg));
      }
    },
    close() {
      closed = true;
      ws?.close();
    },
    get status() {
      return status;
    },
  };
}
