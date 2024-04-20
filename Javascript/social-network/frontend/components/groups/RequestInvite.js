"use client";
import styles from "./page.module.css";
import apiFetch from "../apiFetch";
import { useState, useEffect, useContext } from 'react';
import getUserUUID from "@/app/util";

const RequestInvite = ({group, state, onRequestSubmit}) => {
    const [sent, setSent] = useState(false);

    const handleRequestSubmit = async (e) => {
      e.preventDefault();
      const formData = new FormData(e.target);
      try {
        apiFetch("/groups/request", () => setSent(true), {
          method: 'POST',
          body: formData
        }).then(resp => {onRequestSubmit ? onRequestSubmit(resp): null});
      } catch (error) {
        console.error(error);
      }
    };

    if (state === "rejected") {
      return (
        <p>Cannot join this Group</p>
      )
    }

    if (sent || state === "requested") {
      return (
        <div>
            <p>{state !== "pending" ? "Invite Request Sent" : "Accepted"}</p>
        </div>
      );
    }
    if (state === "accepted") {
      return (
        <p>You are member of this Group</p>
      );
    }
    return (
        <div>
            <form onSubmit={handleRequestSubmit}>
              <input
                type="text"
                name="group_id"
                value={group.id}
                hidden
                readOnly
              />
              <button type="submit" name="request">{ state === "pending" ? "Accept Invite" : "Request Invite" }</button>
            </form>
        </div>
    );
};

export default RequestInvite;