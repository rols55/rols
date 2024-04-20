"use client";
import apiFetch from "../apiFetch";
import styles from "./page.module.css";
import { useRouter } from 'next/navigation'

const EventForm = ({groupId}) => {
    const router = useRouter();
    const handleSubmit = async (event) => {
      event.preventDefault();
      const formData = new FormData(event.target);
      try {
        apiFetch("/event", (resp) => { router.push('/event/'+resp.event.id)}, { method: 'POST', body: formData });
      } catch (error) {
        console.error('Error creating event:', error);
      }
    };
  
    return (
      <form className={styles.eventForm} onSubmit={handleSubmit}>
        <h2>Create New Event</h2>
          <label htmlFor="formEventTitle">Event Title:</label>
          <input
            type="text"
            name="title"
            id="formEventTitle"
            required
          />
          <label htmlFor="formEventDescription">Event Description:</label>
          <textarea
            style={{ height: 100 }}
            type="text"
            name="description"
            id="formEventDescription"
            required
          ></textarea>
          <label htmlFor="formEventDate">Event date:</label>
          <input
            style={{ maxWidth: "200px" }}
            type="datetime-local"
            name="date"
            id="formEventDate"
            required
          />
          <select
          name="going"
          defaultValue="true"
          required
          >
            <option value="true" >Going</option>
            <option value="false">Not Going</option>
          </select>
          <input
            type="text"
            name="group_id"
            id="formEventGroupId"
            value={groupId}
            required
            hidden
            readOnly
          />
        <button className="button-27" type="submit">Create Event</button>
      </form>
    );
  };
  
  export default EventForm;