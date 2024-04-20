"use client";
import styles from "./page.module.css";
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation'
import apiFetch from "../apiFetch";

const Groups = ({uuid}) => {
  const [data, setData] = useState({groups: null});
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const router = useRouter();


  useEffect(() => {
    const fetchGroups = async () => {
      try {
        apiFetch(`/groups${uuid ? "?uuid="+uuid : ""}`, setData);
      } catch (error) {
        setError(error);
      } finally {
        setIsLoading(false);
      }
    };
    fetchGroups();
  }, []);

  const handleOnClick = (group) => {
    router.push('/group/'+group.id)
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <>
      <div className={styles.groupList}>
        {data.groups ? data.groups.map( (group) => (
          <div key={group.id} className={styles.groupListItem} onClick={() => handleOnClick(group)}>
              {group.title}
          </div>
        )) : null}
      </div>
    </>
  );
}

export default Groups;