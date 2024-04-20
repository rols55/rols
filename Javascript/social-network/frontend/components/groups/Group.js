"use client";
import styles from "./page.module.css";
import { useState, useEffect } from 'react';
import ApproveRequest from "./ApproveGroupRequest";
import RequestInvite from "./RequestInvite";
import GroupChat from "./GroupChat";
import PostForm from "@/components/posts/form";
import Feed from "../posts/feed";
import Events from "../events/Events";
import EventForm from "../events/EventForm";
import getUserUUID from "@/app/util";
import SectionSelector from "./SectionSelector";
import Members from "./Members";
import InviteUser from "./InviteUser";

const Group = ({ res }) => {
    const [uuid, setUUID] = useState('');
    const [isMemeber, setMember] = useState(false);

    useEffect(() => {
      getUserUUID().then((uuid) => setUUID(uuid));
    }, []);

    const onRequestSubmit = () => {
        if (res.group_state === "pending") {
            setMember(true)
        }
    };
    if (!res) {
        return (
            <div className={styles.group}>
                <h3>Select a group</h3>
            </div>
        )
    }
    return (
        <>
            <div className={styles.group}>
                <div className={styles.groupTitleContainer}>
                    <div className={styles.groupTitle}>
                        <h2>{res.group.title}</h2>
                    </div>
                    <div>
                        <RequestInvite group={res.group} state={res.group_state} onRequestSubmit={onRequestSubmit}/>
                    </div>
                </div>
                <p>Description: {res.group.description}</p>
            </div>
            { (res.group_state === "accepted" || isMemeber) && (
                <>
                    <SectionSelector tabs={[
                        {name: "Posts", ref: "posts", content: () => (<>
                            <PostForm groupId={res.group.id} />
                            <Feed posts={res.posts} />
                        </>)},
                        {name: "Events", ref: "events", content: () => (<>
                            <EventForm groupId={res.group.id} />
                            <Events groupId={res.group.id} />
                        </>)},
                        {name: "Chat", ref: "chat", content: () => (<>
                            <GroupChat groupId={res.group.id} members={res.members} />
                        </>)},
                        {name: "Members", ref: "members", content: () => (<>
                            <Members members={res.members} uuid={uuid} />
                        </>)},
                        {name: "Invite", ref: "invite", content: () => (<>
                            <InviteUser members={res.members} groupId={res.group.id} />
                        </>)},
                        ( res.group.user_id == res.user.id ?
                            {name: "Requests", ref: "requests", content: () =>(<>
                                <ApproveRequest groupId={res.group.id}/>
                            </>)}
                        : null),
                        ( process.env.NEXT_PUBLIC_DEV ?
                            {name: "JSON Data", ref: "json", content: () => (<>
                                <h3>JSON Data</h3>
                                <pre>{JSON.stringify(res, null, 2)}</pre>
                            </>)}
                        : null),
                    ]}/>
                </>
            )}
        </>
    );
  };
  
  export default Group;