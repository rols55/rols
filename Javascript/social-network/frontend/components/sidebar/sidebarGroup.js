
"use client";
import { useState } from "react";
import ReactDOM from "react-dom";
import Chat from "../chat/chat";

export default function SidebarGroup({ group }) {
  const [isOpen, setIsOpen] = useState(false);

  const toggleChatbox = () => {
    setIsOpen((prevState) => !prevState);
  };

  return (
    <div className="group">
      <div className="status"></div>
      <h3 onClick={toggleChatbox} className="tooltip">
        {group.group_title}
        <span className="tooltiptext">
          {group.group_title}
        </span>
      </h3>
      {isOpen &&
        ReactDOM.createPortal(
          <Chat key={group.group_id} name={group.group_title}  onClose={toggleChatbox}/>,
          document.getElementById("chatBox")
        )}
    </div>
  );
}
