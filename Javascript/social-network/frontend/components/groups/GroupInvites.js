"use client";
import apiFetch from "../apiFetch";
import styles from "./page.module.css";
import { useState, useEffect } from 'react';

const GroupInvites = () => {
  const [invites, setInvites] = useState([]);

  useEffect(() => {
    const fetchInvites = async () => {
      try {
        apiFetch('/groups/invites', (resp) => {
          if (resp.groups) {
            setInvites(resp.groups)
          }
        });
      } catch (error) {
        console.error('Error:', error);
      }
    };
    fetchInvites();
  }, []);

  const handleAccept = async (e, id) => {
    e.preventDefault();
    try {
      var formData = new FormData();
      formData.append('group_id', e.target.value);
      apiFetch('/groups/accept', (resp) => { console.log(resp)}, {
        method: 'POST',
        body: formData
      });
    } catch (error) {
      console.error('Error:', error);
    }
    const updatedItems = invites.filter((group) => group.id !== id);
    setInvites(updatedItems)
  };
  
  const handleReject = async (e, id) => {
    e.preventDefault();
    try {
      var formData = new FormData();
      formData.append('group_id', e.target.value);
      apiFetch('/groups/reject', (resp) => { console.log(resp)}, {
        method: 'POST',
        body: formData
      });
    } catch (error) {
      console.error('Error:', error);
    }
    const updatedItems = invites.filter((group) => group.id !== id);
    setInvites(updatedItems)
  };
  
  
  if (invites.length > 0) {
    return (
      <div className={styles.groupInvites}>
        <h3>My Group Invites</h3>
        {invites.map((group) => (
              <div key={group.id}>
                  <span>{group.title}</span>
                  <button value={group.id} onClick={(e) => handleAccept(e, group.id)}>Accept</button>
                  <button value={group.id} onClick={(e) => handleReject(e, group.id)}>Reject</button>
              </div>
        ))}
      </div>
    );
  }
  return <></>;
};

export default GroupInvites;