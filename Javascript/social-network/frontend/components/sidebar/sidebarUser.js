"use client";
import { useState } from "react";
import ReactDOM from "react-dom";
import UserChat from "../chat/UserChat";

export default function SidebarUser({ user }) {
  const [isOpen, setIsOpen] = useState(false);

  const toggleChatbox = () => {
    setIsOpen((prevState) => !prevState);
  };

  return (
    <div className="user">
      <div className="status"></div>
      <h3 onClick={toggleChatbox} className="tooltip">
        {user.username}
        <span className="tooltiptext">
          {user.firstname} {user.lastname}
        </span>
      </h3>
      {isOpen &&
        ReactDOM.createPortal(
          <UserChat user={user} onClose={toggleChatbox}/>,
          document.getElementById("chatBox")
        )}
    </div>
  );
}
