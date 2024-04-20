"use client";
import styles from "./page.module.css";
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation'
import apiFetch from "../apiFetch";

const Events = ( {groupId}) => {
  const [events, setEvents] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const router = useRouter();

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        apiFetch(`/events?group=${groupId}`, (resp) => setEvents(resp.events));
      } catch (error) {
        setError(error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchEvents();
  }, []);

  const handleOnClick = (event) => {
    router.push('/event/'+event.id)
  }

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }
  return (
    <>
      <div className={styles.eventList}>
        {events ? events.map( (event, idx) => (
          <div key={idx} id={event.id} className={styles.eventListItem} onClick={() => handleOnClick(event)}>
              {event.title}
          </div>
        )) : null}
      </div>
    </>
  );
}

export default Events;