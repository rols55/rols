import User from "./profile/user";
import { usePathname } from "next/navigation";
import apiFetch from "./apiFetch";
import { useState } from "react";
export default function Followers({ user }) {
  const [accepted, setAccepted] = useState(user.accepted);
  const handleFollow = async () => {
    try {
      await apiFetch(`/followers/accept?uuid=${user.follower}`, null, {
        method: "POST",
      }).then(() => {
        setAccepted(true);
      });
    } catch (error) {
      alert(error);
    }
  };
  const handleUnFollow = async ({ target }) => {
    try {
      await apiFetch(`/followers/cancel?uuid=${user.follower}`, null, {
        method: "POST",
      }).then(() => {
        target.parentElement.parentElement.remove();
      });
    } catch (error) {
      alert(error);
    }
  };
  const pathname = usePathname().split("/")[2];
  return (
    <div className="element-of-modal">
      <User uuid={user.follower}>{user.username}</User>
      {!pathname && !accepted && (
        <div className="followers-actions">
          <button onClick={handleFollow}>Accept</button>
          <button onClick={handleUnFollow}>Decline</button>
        </div>
      )}
    </div>
  );
}
