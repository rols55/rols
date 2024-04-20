import User from "./profile/user";
import { usePathname } from "next/navigation";
import apiFetch from "./apiFetch";
export default function Following({ user }) {
  const handleUnFollow = async ({ target }) => {
    try {
      await apiFetch(`/followers/unfollow?uuid=${user.follower}`, null, {
        method: "POST",
      }).then(() => {
        target.parentElement.remove();
      });
    } catch (error) {
      alert(error);
    }
  };
  const pathname = usePathname().split("/")[2];
  return (
    <div className="element-of-modal">
      <User uuid={user.follower}>{user.username}</User>
      {!pathname &&
        (user.accepted ? (
          <button onClick={handleUnFollow}>Unfollow</button>
        ) : (
          <button onClick={handleUnFollow}>Pending/Cancel</button>
        ))}
    </div>
  );
}
