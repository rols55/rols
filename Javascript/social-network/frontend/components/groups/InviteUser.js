import React, { useState } from 'react'
import styles from "./page.module.css";
import UserInviteInput from './UserInviteInput';
import apiFetch from "../apiFetch";
import { useRouter } from 'next/navigation'

export default function InviteUser({members, groupId}) {
    const router = useRouter();

    const handleSubmit = async (event) => {
        event.preventDefault();
        const formData = new FormData(event.target);
        try {
          apiFetch("/groups/invite", (resp) => { console.log(resp); router.push('/group/'+groupId)}, { method: 'POST', body: formData });
        } catch (error) {
          console.error('Error creating group:', error);
        }
      };

    return (
        <form className={styles.groupForm} onSubmit={handleSubmit}>
            <input
            type="text"
            name="group_id"
            value={groupId}
            readOnly
            hidden
            />
            <UserInviteInput members={members}/>
            <button className="button-27" type="submit">Send Invites</button>
        </form>
    )
}
