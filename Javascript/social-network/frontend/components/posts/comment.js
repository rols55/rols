import User from "@/components/profile/user";
export default function Comment({comment}) {
  return (
    <div id={comment.id} className="comment">
      <User params={comment} />
      <p> {comment.text}</p>
      <p> {new Date(comment.creation_date).toLocaleString()}</p>
      {comment.image && (
                             <img
                             style={{
                                 width: "100%",
                                 maxWidth: "700px",
                                 margin: "auto",
                             }}
                             src={`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/getimage/${comment.image}`}
                         />
      )}
    </div>
  );
}
