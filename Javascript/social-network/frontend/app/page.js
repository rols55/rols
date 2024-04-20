"use client"
import { useEffect, useState } from 'react';
import Feed from "@/components/posts/feed";
import UserForm from "@/components/posts/form";

export default function Home() {
  const [posts, setPosts] = useState(null);

  useEffect(() => {
    const getPosts = async () => {
      try {
        const cookies = document.cookie;
        if (cookies) {
          const res = await fetch(`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/feed`, {
            credentials: 'include'
          });
          if (res.ok) {
            const data = await res.json();
            setPosts(data.posts);
          }
        }
      } catch (err) {
        console.error(err);
      }
    };

    getPosts();
  }, []); // Empty dependency array ensures that this effect runs only once when the component mounts

  return (
    <>
      <UserForm />
      {posts ? <Feed posts={posts} /> : <p>No posts available.</p>}
    </>
  );
}
