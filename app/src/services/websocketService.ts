class WebSocketService {
  ws: WebSocket;
  listeners: Array<(message: any) => void>;
  constructor() {
    this.ws = null;
    this.listeners = [];
  }

  connect(url) {
    if (this.ws) {
      console.log("WebSocket is already connected.");
      return;
    }

    this.ws = new WebSocket(url);

    this.ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      this.listeners.forEach((listener) => listener(message));
    };

    this.ws.onopen = () => {
      console.log("WebSocket connection established");
    };

    this.ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    this.ws.onclose = () => {
      console.log("WebSocket connection closed");
      this.ws = null;
    };
  }

  addListener(listener) {
    this.listeners.push(listener);
  }

  removeListener(listener) {
    const index = this.listeners.indexOf(listener);
    if (index > -1) {
      this.listeners.splice(index, 1);
    }
  }

  sendMessage(message) {
    if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
      console.error("WebSocket is not connected.");
      return;
    }
    this.ws.send(JSON.stringify(message));
  }
}

// Export a singleton instance of the WebSocketService
const webSocketService = new WebSocketService();
export default webSocketService;
