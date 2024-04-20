"use client";
import { useState, useEffect } from "react";
import Link from "next/link";

export default function People() {
  const [data, setData] = useState([]);
  useEffect(() => {
    const fetchData = async () => {
    const cookies = document.cookie;
    if (cookies) {
      try {
        const response = await fetch(`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/users`, { credentials: 'include' });
        if (response.ok) {
            const data = await response.json();
            setData(data);
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }}
    };
    fetchData();
  }, []); 
  const allUsers = data.users;

  return (
    <>
    <div className="all-people">
        <div className="people-title">
        <h2>All people</h2>
        </div>
      {allUsers ? (
        allUsers.map((user) => (
            <div className="people" key={user.uuid}>
               <div> 
                    <img className="people-photo"
                    src={`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/getimageuser/${user.image}`}
                    />
                </div> 
                <div>
                    <div className="people-username">
                    <Link href={`/profile/${user.uuid}`}>
                        <h3>{user.username}</h3>
                    </Link>    
                    </div>
                    <p className="people-name">{user.firstname} {user.lastname}</p>
                </div>
    </div>
  ))
) : (
  <h4>No people found</h4>
)}
      </div>
      </>
  );
}
