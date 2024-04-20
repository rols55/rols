import React from 'react'
import styles from "./page.module.css";
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation'

export default function Members( { members, uuid }) {
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);
    const router = useRouter();
  
    useEffect(() => {
        setIsLoading(false);
    }, []);
  
    const handleOnClick = (member) => {
      if (uuid && uuid !== member.uuid) {
        router.push('/profile/'+member.uuid);
      } else {
        router.push('/profile');
      }
    }
  
    if (isLoading) {
      return <div>Loading...</div>;
    }
  
    if (error) {
      return <div>Error: {error.message}</div>;
    }

    return (
      <>
        <div className={styles.memberList}>
          {members ? members.map( (m, idx) => (
            <div key={m.uuid} className={styles.memberListItem} onClick={() => handleOnClick(m)}>
                <h3>{m.username}</h3>
            </div>
          )) : null}
        </div>
      </>
    );
}
