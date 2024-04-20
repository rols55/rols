import Profile from "@/components/profile/profile";
export default function OtherProfile({ params }) {
  return <Profile uuid={params.uuid} />;
}
