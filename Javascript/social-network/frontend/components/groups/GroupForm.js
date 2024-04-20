"use client";
import apiFetch from "../apiFetch";
import UserInviteInput from "./UserInviteInput";
import styles from "./page.module.css";
import { useRouter } from 'next/navigation'

const GroupForm = () => {
    const router = useRouter();
    const handleSubmit = async (event) => {
      event.preventDefault();
      const formData = new FormData(event.target);
      try {
        apiFetch("/group", (resp) => { console.log(resp); router.push('/group/'+resp.group.id)}, { method: 'POST', body: formData });
      } catch (error) {
        console.error('Error creating group:', error);
      }
    };
  
    return (
      <form className={styles.groupForm} onSubmit={handleSubmit}>
        <h2>Create New Group</h2>
        <label htmlFor="formGroupTitle">Group Title:</label>
        <input
          type="text"
          name="title"
          id="formGroupTitle"
          required
        />
        <label htmlFor="formGroupDescription">Group Description:</label>
        <textarea
          type="text"
          name="description"
          id="formGroupDescription"
          required
        ></textarea>
        <UserInviteInput />
        <button className="button-27" type="submit">Create Group</button>
      </form>
    );
  };
  
  export default GroupForm;