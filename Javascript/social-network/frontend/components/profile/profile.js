"use client";
import { useState, useEffect } from "react";
import About from "@/components/profile/about";
import _ from "@/styles/profile.css";
import apiFetch from "../apiFetch";
import getUserUUID from "@/app/util";
import Link from "next/link";
import SelectorButtons from "@/components/profile/selectorButtons";
import MyModal from "../modal";

export default function Profile({ searchParams, uuid }) {
  const [profile, setProfile] = useState(null);
  const [posts, setPosts] = useState(null);
  const [follows, setFollows] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [displayFollow, setDisplayFollow] = useState(null);
  const [myUuid, setMyUuid] = useState("");
  const [followerOnclick, setfollowerOnclick] = useState(null);
  const [show, setShow] = useState(false);
  const [isOpen, setIsOpen] = useState(false);

  useEffect(() => {
    if (!uuid) {
      const fetchProfile = async () => {
        try {
          apiFetch("/profile").then((resp) => {
            setProfile(resp.user);
            setPosts(resp.posts);
            setShow(resp.show);
          });
          apiFetch("/followers", setFollows);
        } catch (error) {
          setError(error);
        } finally {
          setIsLoading(false);
        }
      };
      fetchProfile();
    } else {
      const fetchProfile = async () => {
        try {
          apiFetch(`/otherprofile?uuid=${uuid}`).then((resp) => {
            setProfile(resp.user);
            setPosts(resp.posts);
            setShow(resp.show);
          });
          await apiFetch(`/otherfollowers?uuid=${uuid}`, setFollows);
        } catch (error) {
          setError(error);
        } finally {
          setIsLoading(false);
        }
      };
      fetchProfile();
    }
  }, []);

  useEffect(() => {
    if (follows) {
      const getUuid = async () => {
        const myUuid = await getUserUUID();
        setMyUuid(myUuid);
      };
      getUuid();
      let follower;
      if (follows.followers) {
        const foundFollower = follows.followers.find(
          (follower) => follower.follower === myUuid
        );
        if (foundFollower) {
          follower = foundFollower.accepted;
          if (follower) {
            setDisplayFollow("Unfollow");
          } else {
            setDisplayFollow("Pending/Cancel");
          }
          setfollowerOnclick(() => handleUnFollow);
        } else {
          setDisplayFollow("Follow");
          setfollowerOnclick(() => handleFollow);
        }
      } else {
        setDisplayFollow("Follow");
        setfollowerOnclick(() => handleFollow);
      }
    }
  }, [follows, myUuid]);

  const handleFollow = async () => {
    try {
      await apiFetch(`/followers/request?uuid=${uuid}`, null, {
        method: "POST",
      }).then(() => {
        window.location.reload();
      });
    } catch (error) {
      alert(error);
    }
  };

  const handleUnFollow = async () => {
    try {
      await apiFetch(`/followers/unfollow?uuid=${uuid}`, null, {
        method: "POST",
      }).then(() => {
        window.location.reload();
      });
    } catch (error) {
      alert(error);
    }
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }
  if (!profile) {
    return <div>User doesn&apos;t exists</div>;
  }
  return (
    <>
      <div className="profile">
      <img
            className="photo"
            src={`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/getimageuser/${profile.image}`}
          />
        <div className="name-line">
          <h2 className="title">
            {profile.firstname} {profile.lastname}{" "}
            {profile.nickname ? `(${profile.nickname})` : ""}
          </h2>
          {uuid ? (
            <button
              className="button-27 follow-button"
              onClick={followerOnclick}
            >
              {displayFollow}
            </button>
          ) : (
            <>
              <button
                onClick={() => setIsOpen(true)}
                className="button-27 follow-button"
              >
                Edit profile
              </button>
              <MyModal
                isOpen={isOpen}
                setIsOpen={setIsOpen}
                profile={profile}
              />
            </>
          )}
        </div>
        <nav
          className="Profile-stats"
          style={{ display: "flex", justifyContent: "space-around" }}
        ></nav>
        <About profile={profile} />
        <div className="about-me">
          <h3>About me:</h3>
          {profile.aboutme && <p>{profile.aboutme}</p>}
        </div>
      </div>

      {profile.public || show ? (
        <SelectorButtons
          follows={follows}
          posts={posts}
          selector={searchParams && searchParams.selector}
          uuid={uuid}
        />
      ) : (
        <div style={{ textAlign: "center" }}>This user is private!</div>
      )}
    </>
  );
}
