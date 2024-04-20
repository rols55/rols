"use client";
import { useState, useEffect } from "react"
import styles from "./page.module.css";
import AttendEvent from "./AttendEvent";
import apiFetch from "../apiFetch";

const Event = ({ path }) => {
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);
    const [user, setUser] = useState(null);
    const [event, setEvent] = useState(null);
    const [userGoing, setUserGoing] = useState("pending");
    const [going, setGoing] = useState([]);
    const [notGoing, setNotGoing] = useState([]);

    useEffect(() => {
        if (!event) return
        const fetchAttendance = async () => {
            try {
              apiFetch("/events/attendance?event="+event.id, (resp) => {
                if (resp.going) {
                    setGoing(resp.going);
                }
                if (resp.not_going) {
                    setNotGoing(resp.not_going);
                }
                if (resp.is_going) {
                    setUserGoing(resp.is_going);
                }
              });
            } catch (error) {
              setError(error);
            }
        };
        fetchAttendance();
    }, [event]);

    useEffect(() => {
        const fetchEvent = async () => {
          try {
            apiFetch(path).then((resp) => {
                setEvent(resp.event);
                setUser(resp.user);
            });
          } catch (error) {
            setError(error);
          } finally {
            setIsLoading(false);
          }
        };
        fetchEvent();
        setIsLoading(false);
      }, [path]);
    
    const handleOnClick = (state) => {
        setUserGoing(state)
        if (state == "true") {
            setGoing([...going, user]);
        } else if (state === "false") {
            setNotGoing([...notGoing, user])
        }
    };

    if (isLoading) {
        return <div>Loading...</div>;
    }
    
    if (error) {
        return <div>Error: {error.message}</div>;
    }
  
    if (!event) {
        return <div>This Event Does Not Exist</div>;
    }
    return (
        <>
            {event && (
                <div className={styles.event}>
                    <div className={styles.eventTitleContainer}>
                        <div className={styles.eventTitle}>
                            <h2>{event.title}</h2>
                        </div>
                        <div>
                            <AttendEvent event={event} state={userGoing} onClick={(state) => handleOnClick(state)}/>
                        </div>
                    </div>
                    <p>{userGoing}</p>
                    <p>Description: {event.description}</p>
                    <p>Date: {new Date(event.date).toLocaleString()}</p>
                    <h3>Attendees</h3>
                    <p>Going:</p>
                    {going.map((m) => (
                        <div key={m.uuid}>
                            {m.username}
                        </div>
                    ))}
                    <p>Not Going:</p>
                    {notGoing.map((m) => (
                        <div key={m.uuid}>
                            {m.username}
                        </div>
                    ))}
                </div>
            )}
        </>
    );
  };
  
  export default Event;