import Post from "./post";
export default function Feed({ posts }) {
  if (!posts) {
    return <div>No Posts</div>;
  }
  return (
    <>
      {posts.map((item) => (
        <Post post={item} key={item.id} />
      ))}
    </>
  );
}
