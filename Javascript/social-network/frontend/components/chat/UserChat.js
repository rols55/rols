"use client";
import React, { useState, useEffect } from 'react';
import getUserUUID from "@/app/util";
import styles from "./page.module.css";
import EmojiPickerButton from "../groups/emojiPickerButton";
import apiFetch from '../apiFetch';
import { useWebSocket } from '@/app/WebSocketProvider';

const UserChat = ({ user, onClose}) => {
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');
  const [uuid, setUUID] = useState("");
  const [offset, setOffset] = useState(0);
  const socket = useWebSocket();

  useEffect(() => {
    if (!socket) return; 
    const handleOnMessage = (data) => {
      setMessages((prevMessages) => [data,...prevMessages]);
    }
    socket.addOnMessageHandler(user.uuid, handleOnMessage)
    return () => {
      socket.removeOnMessageHandler(user.uuid);
    };
  }, [socket]);

  useEffect(() => {
    getUserUUID().then((res) => { setUUID(res); });
    if (user) {
        const fetchMessages = () => {
        apiFetch("/history?target="+user.uuid+"&offset="+offset).then((resp) => {
            if (resp.history) {
            setMessages(resp.history);
            }
        });
        };
        fetchMessages();
    }
  }, []);

  const sendMessage = (e) => {
    e.preventDefault();
    if (newMessage.trim() !== '') {
      socket.send({
        type: "message",
        reciver_uuid: user.uuid,
        message_text: newMessage,
      });
      setNewMessage('');
    }
  };

  const getMessageTitle = (message) => {
    if (message.sender_uuid !== uuid) {
      return user.username +"@"+ new Date(message.timestamp).toLocaleString();
    }
    return "Me@"+ new Date(message.timestamp).toLocaleString();
  }

  const handleEmojiPicker = (emoji) => {
    setNewMessage((prevMsg) => prevMsg + emoji);
  };

  return (
    <div className={styles.chatContainer}>
      <div className={styles.chatTitleContainer}>
        <h3 className={styles.chatTitle}>{user.username}</h3>
        <button onClick={onClose} className={styles.chatClose}>
          ×
        </button>
      </div>
      <div className={styles.messageContainer}>
        {messages.map((message, index) => (
          <div key={index} className={message.sender_uuid === uuid ? styles.messageSender : styles.message}>
            <span>{getMessageTitle(message)}</span>
            <div>{message.message_text}</div>
          </div>
        ))}
      </div>
      <form onSubmit={sendMessage} className={styles.inputForm}>
        <input
          type="text"
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
          placeholder="Type a message..."
          className={styles.input}
        />
        <EmojiPickerButton
          handleEmojiPicker={handleEmojiPicker}
          className={styles.emojiButton}
        />
        <button type="submit" className={styles.button}>Send</button>
      </form>
    </div>
  );
};

export default UserChat;