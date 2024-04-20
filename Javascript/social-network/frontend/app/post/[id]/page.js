"use client"
import { useEffect, useState } from 'react';
import Post from "@/components/posts/post";
import Reply from "@/components/posts/reply";
import Comment from "@/components/posts/comment";

export default function PostPage(params) {
  const [postData, setPostData] = useState(null);
  
  useEffect(() => {
    async function fetchData() {
      try {
        const cookies = document.cookie;
        if (cookies) {
          const res = await fetch(`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/post/${params.params.id}`, {
            credentials: 'include',
          });
          if (res.ok) {
            const data = await res.json();
            setPostData(data);
          }
        }
      } catch (err) {
        console.error(err);
      }
    }
    fetchData();
  }, [params.params.id]); // Make sure to include params.params.id as a dependency

  if (!postData || !postData.post) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <Post post={postData.post} />
      <Reply postId={params.params.id} /> 
      {postData.comments && postData.comments.map((comment) => (
        <Comment key={comment.id} comment={comment} />
      ))}
    </>
  );
}

