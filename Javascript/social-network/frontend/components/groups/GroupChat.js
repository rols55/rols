"use client";
import React, { useState, useEffect } from 'react';
import getUserUUID from "@/app/util";
import styles from "./page.module.css";
import EmojiPickerButton from "./emojiPickerButton";
import apiFetch from '../apiFetch';
import { useWebSocket } from '@/app/WebSocketProvider';

const GroupChat = ({groupId , members}) => {
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');
  const [uuid, setUUID] = useState("");
  const [offset, setOffset] = useState(0);
  const socket = useWebSocket();

  useEffect(() => {
    if (!socket) return;
    socket.handleOnGroupMessage = (data) => {
      setMessages((prevMessages) => [data,...prevMessages]);
    }
    return () => {
      socket.handleOnGroupMessage = null;
    };
  }, [socket]);

  useEffect(() => {
    getUserUUID().then((res) => { setUUID(res); });
 
    const fetchMessages = () => {
      apiFetch("/history/group?target="+groupId+"&offset="+offset).then((resp) => {
        if (resp.history) {
          setMessages(resp.history);
        }
      });
    };
    fetchMessages();
  }, []);

  const sendMessage = (e) => {
    e.preventDefault();
    if (newMessage.trim() !== '') {
      socket.send({
        type: "message",
        group_id: groupId,
        message_text: newMessage,
        reciver_uuid: "",
      });
      setNewMessage('');
    }
  };

  const getMessageTitle = (message) => {
    const member = members.find(m => m.uuid === message.sender_uuid);
    if (member) {
      return member.username +"@"+ new Date(message.timestamp).toLocaleString();
    }
  }
  
  const handleEmojiPicker = (emoji) => {
    setNewMessage((prevMsg) => prevMsg + emoji);
  };

  return (
    <div className={styles.chatContainer}>
      <h3>Group Chat</h3>
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
        <input
          type="text"
          value={uuid}
          onChange={(e) => setUUID(e.target.value)}
          className={styles.input}
          hidden
        />
        <button type="submit" className={styles.button}>Send</button>
      </form>
    </div>
  );
};

export default GroupChat;