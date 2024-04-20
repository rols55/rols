import UserChat from "./UserChat";

export default function Chat({ user, onClose }) {
  return (
    <div className="chat">
      <h3 style={{ marginLeft: "10px" }}>{user.username}</h3>
      <button onClick={onClose} className="close-button">
        ×
      </button>
    </div>
  );
}
