'use client'
import { useState } from 'react';

export default function ChatButton(params) {
  const [isChatboxOpen, setIsChatboxOpen] = useState(false);

  const handleChatboxToggle = () => {
    setIsChatboxOpen(prevState => !prevState); // Toggle the chatbox state
  };

  return null
}