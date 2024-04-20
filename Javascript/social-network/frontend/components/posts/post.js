import User from "@/components/profile/user";
import Link from "next/link";
export default function Post({ post }) {
  return (
    <div id={post.id} className="post">
      <User uuid={post.uuid}>{post.author}</User>
      <Link href={`/post/${post.id}`}>
        <h2>{post.title}</h2>
        <p>{post.text}</p>
        <p>{new Date(post.creation_date).toLocaleString()}</p>
        </Link>
        {post.image && (
          <img
            style={{
                width: "100%",
                maxWidth: "700px",
                margin: "auto",
            }}
            src={`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/getimage/${post.image}`}
          />
        )}
      <br />
      <Link href={`/post/${post.id}`}>Comments: {post.comments_count ? post.comments_count : 0}</Link>
    </div>
  );
}
