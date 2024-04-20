"use client";
import { useState, useEffect } from "react";
import Link from "next/link";
import Feed from "@/components/posts/feed";
import Followers from "../followers";
import Following from "../following";
import Groups from "../groups/Groups";
import GroupInvites from "../groups/GroupInvites";
import { getUUID } from "@/app/util";
import { useSearchParams } from 'next/navigation'

export default function SelectorButtons({
  follows,
  posts,
  groups,
  selector,
  uuid,
}) {

  const [selectedButton, setSelectedButton] = useState("feed"); // Set default option
  const searchParams = useSearchParams()

  useEffect(() => {
    const referrer = searchParams.get('referrer')
    if (referrer) {
      // Set the selectedButton state based on the current query parameter
      setSelectedButton(referrer);
    } else {
      setSelectedButton("feed"); // Default to "feed"
    }
  }, [searchParams]);
  

  const handleButtonClick = (secion) => {
    setSelectedButton(secion);
  };
  return (
    <div className="under-profile">
      <div className="button-row">
        <div
          href="profile/feed"
          className="button-27"
          role="button"
          onClick={() => handleButtonClick("feed")}
        >
          Feed
        </div>
        <button
          className="button-27"
          onClick={() => handleButtonClick("groups")}
        >
          Groups
        </button>
        <button
          className="button-27"
          onClick={() => handleButtonClick("following")}
        >
          Following
        </button>
        <button
          className="button-27"
          onClick={() => handleButtonClick("followers")}
        >
          Followers
        </button>
      </div>
      {selectedButton === "feed" && <Feed posts={posts} />}
      {selectedButton === "groups" && (<>
        <GroupInvites />
        <h3>My Groups</h3>
        <Groups uuid={uuid ? uuid : getUUID()} />
      </>)}
      {selectedButton === "following" && (
        <div className="following-list">
          {follows && follows.following ? (
            follows.following.map((following) => (
              <Following key={following.follower} user={following} uuid />
            ))
          ) : (
            <div>Not following anyone</div>
          )}
        </div>
      )}
{selectedButton === "followers" && follows && (
    <div className="following-list">
    {follows.followers ? (
      follows.followers.map((follower) => (
        <Followers key={follower.follower} user={follower} uuid={uuid} />
      ))
    ) : (
      <p>No followers found</p>
    )}
  </div>
)}
    </div>
  );
}
