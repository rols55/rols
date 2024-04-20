'use client';

class Socket {
  constructor(url) {
    this.socket = new WebSocket( url ? url : `ws://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/ws`);
    this.socket.onopen = () => console.log("Connected to WebSocket server");
    this.socket.onclose = () => console.log("Disconnected from WebSocket server");
    this.socket.onmessage = (e) => {
      var data = JSON.parse(e.data)
      switch (data.type) {
        case "message":
          if (data.group_id > 0) {
            this.handleOnGroupMessage(data)
          } else {
            this.handleOnMessage(data)
          }
          break;
        case "notification": 
          this.handleOnNotification(data);
          break;
        case "connected":
          this.handleOnConnected(data);
          break;
        case "disconnected":
          this.handleOnDisconnected(data);
          break;
        case "get_online":
          this.handleOnGetOnline(data);
          break;
        default:
          this.handleOnDefault(data);
      }
    }

    this.onMessageHandlers = new Map([]);

    this.handleOnMessage = (data) => {
      console.log("handleOnMessage received data:", data);
      for (const [key, value] of this.onMessageHandlers) {
        if (key === data.reciver_uuid || key === data.sender_uuid) {
          value(data);
          break;
        }
      }
    }

    this.handleOnGroupMessage = (data) => {
      console.log("handleOnGroupMessage received data:", data);
    }

    this.handleOnNotification = (data) => {
      console.log("handleOnNotification received data:", data);
    }

    this.handleOnConnected = (data) => {
      console.log("handleOnConnected received data:", data);
    }

    this.handleOnDisconnected = (data) => {
      console.log("handleOnDisconnected received data:", data);
    }

    this.handleOnGetOnline = (data) => {
      console.log("handleOnGetOnline received data:", data);
    }

    this.handleOnDefault = (data) => {
      console.log("handleOnDefault received data:", data);
    }
  }

  addOnMessageHandler(id, handler) {
    this.onMessageHandlers.set(id, handler);
  }

  removeOnMessageHandler(id) {
    this.onMessageHandlers.delete(id);
  }

  send(data) {
    console.log("Sending data: " + JSON.stringify(data));
    const sendData = JSON.stringify(data);
    this.socket.send(sendData);
  }
  
  close() {
    this.socket.close();
  }
}

import React, { createContext, useContext, useEffect, useState } from 'react';

const WebSocketContext = createContext();

export const WebSocketProvider = ({ children }) => {
  const [socket, setSocket] = useState(null);

  useEffect(() => {
    const newSocket = new Socket(`ws://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/ws`);
    setSocket(newSocket);

    return () => {
      newSocket.close();
    };
  }, []);

  return (
    <WebSocketContext.Provider value={socket}>
      {children}
    </WebSocketContext.Provider>
  );
};

export const useWebSocket = () => useContext(WebSocketContext);