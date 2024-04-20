"use client";
import styles from "./page.module.css";
import apiFetch from "../apiFetch";
import { useState, useEffect, useContext } from 'react';
import getUserUUID from "@/app/util";

const AttendEvent = ({event, state, onClick}) => {
    const [going, setGoing] = useState(state);
    const [uuid, setUUID] = useState("");
    let formGoing = "pending";
    
    useEffect(() => {
      getUserUUID().then((res) => { setUUID(res); });
    }, []);

    useEffect(() => {
      onClick(going);
    }, [going]);

    const handleClick = (e) => {
      formGoing = e.target.value;
    };

    const handleAttendSubmit = async (e) => {
      e.preventDefault();
      if (formGoing !== "true" && formGoing !== "false") {
        return
      }
      const formData = new FormData(e.target);
      formData.append("going", formGoing);
      try {
        apiFetch("/events/attend", () => setGoing("true"), {
          method: 'POST',
          body: formData
        });
      } catch (error) {
        console.error(error);
        setGoing("pending");
      }
    };

    if (state === "true") {
      return (
        <div>
            <p>Going</p>
        </div>
      );
    } 
    
    if (state === "false") {
       return (
        <div>
            <p>Not Going</p>
        </div>
      );     
    }
    return (
        <div>
            <form onSubmit={handleAttendSubmit}>
                <input
                type="text"
                name="event_id"
                value={event.id}
                hidden
                readOnly
                />
                <button type="submit" onClick={handleClick} name="going" value="true">Going</button>
                <button type="submit" onClick={handleClick} name="going" value="false">Not going</button>
            </form>
        </div>
    );
};

export default AttendEvent;