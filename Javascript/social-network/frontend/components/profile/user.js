import Link from "next/link";
export default function User({ uuid, children }) {
  return (
    <Link href={`/profile/${uuid}`}>
      <h3>{children}</h3>
    </Link>
  );
}
