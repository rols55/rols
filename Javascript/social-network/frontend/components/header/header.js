import Link from "next/link";
import Logout from "@/components/logout";
import 'material-symbols';
import Notifications from './notifications2';
import { CheckSession } from "@/middleware";



export default async function Header() {
  const authed = await CheckSession();
  return (
    <>
      <header>
        <nav className="nav">
          <div className="header-left">
            <Link href="/">Social Network</Link>
          </div>
          <div className="header-right">
            {authed ? (
              <div className="header-links">
                <Link href="/people">People</Link>
                <Link href="/groups">Groups</Link>
                <Link href="/profile">Profile</Link>
                <Notifications/>
                <Logout />
              </div>
            ) : (
              <div className="header-links">
                <Link href="/login">Login</Link>
                <Link href="/register">Register</Link>
              </div>
            )}
          </div>
        </nav>
      </header>
      <div style={{ height: 50 }} />
    </>
  );
}
