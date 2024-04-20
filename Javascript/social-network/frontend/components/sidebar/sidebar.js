"use client";
import SidebarUser from "./sidebarUser";
import SidebarGroup from "./sidebarGroup";
import apiFetch from "../apiFetch";
import { useState, useEffect } from "react";
import Link from "next/link";

export default function Sidebar() {
  const [data, setData] = useState([]);
  const [groupData, setGroupData] = useState([]);
  useEffect(() => {
    const fetchData = async () => {
      try {
        const jsonData = await apiFetch("/chatusers");
        setData(jsonData);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };
    const fetchGroupData = async () => {
      try {
        const jsonGroupData = await apiFetch("/groups/requests");
        setGroupData(jsonGroupData);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchGroupData();
    fetchData();
  }, []); // Empty dependency array means this effect runs only once after the initial render
  const users = data.users;
  const groups = groupData.group_requests; // TODO: Add groups fetch
  //console.log(groups);
  return (
    <>
      <Link
        className="sidebar-titles"
        href={{
          pathname: "/profile",
          query: { selector: "following" },
        }}
      >
        Private chat
      </Link>
      <div className="sidebar-items">
        {users ? (
          users.map((user) => <SidebarUser key={user.uuid} user={user} />)
        ) : (
          <h4>No followers</h4>
        )}
      </div>
      <Link
        className="sidebar-titles"
        href={{
          pathname: "/profile",
          query: { selector: "groups" },
        }}
      >
        Group chats
      </Link>
      <div className="sidebar-items">
        {groups ? (
          groups.map((group) => (
            <SidebarGroup key={group.group_id} group={group} />
          ))
        ) : (
          <h4>Not in any groups</h4>
        )}
      </div>
    </>
  );
}
