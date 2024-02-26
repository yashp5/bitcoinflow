class WebSocketService {
  ws = null;
  listeners = [];
  url = ""; // Store the URL to reconnect
  reconnectInterval = 5000; // Time between reconnection attempts in milliseconds
  reconnectAttempts = 0;
  maxReconnectAttempts = 10; // Maximum number of reconnection attempts

  constructor() {}

  connect(url) {
    this.url = url; // Store the URL for reconnection
    if (this.ws) {
      console.log("WebSocket is already connected or connecting.");
      return;
    }

    this.ws = new WebSocket(url);

    this.ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      this.listeners.forEach((listener) => listener(message));
    };

    this.ws.onopen = () => {
      console.log("WebSocket connection established");
      this.reconnectAttempts = 0; // Reset reconnect attempts on successful connection
    };

    this.ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    this.ws.onclose = () => {
      console.log("WebSocket connection closed");
      this.ws = null;
      this.tryReconnect();
    };
  }

  tryReconnect() {
    if (this.reconnectAttempts < this.maxReconnectAttempts) {
      console.log(
        `Attempting to reconnect... (${this.reconnectAttempts + 1}/${
          this.maxReconnectAttempts
        })`
      );
      setTimeout(() => {
        this.reconnectAttempts++;
        this.connect(this.url);
      }, this.reconnectInterval);
    } else {
      console.log("Max reconnection attempts reached. Giving up.");
    }
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
