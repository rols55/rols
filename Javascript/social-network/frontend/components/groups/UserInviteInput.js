"use client";
import apiFetch from "../apiFetch";
import styles from "./page.module.css";
import { useState, useEffect } from 'react';

const UserInviteInput = ({members}) => {
  const [allUsers, setAllUsers] = useState([]);
  const [availableUsers, setAvailableUsers] = useState([]);
  const [invitedUsers, setInvitedUsers] = useState([]);
  const [selectedUser, setSelectedUser] = useState('');

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        apiFetch("/users", (resp) => {
          if (members) {
            const filteredUsers = resp.users.filter(user => {
              return !members.some(mem => user.uuid === mem.uuid);
            });
            setAvailableUsers(filteredUsers);
            setAllUsers(filteredUsers);
          } else {
            setAvailableUsers(resp.users);
            setAllUsers(resp.users);
          }
        })
      } catch (error) {
        console.error('Error fetching users:', error);
      }
    };
    fetchUsers();
  }, []);

  const handleInvite = (e) => {
    e.preventDefault();
    if (selectedUser) {
      setInvitedUsers([...invitedUsers, selectedUser]);
      setSelectedUser('');
      const updatedItems = availableUsers.filter((user) => user.uuid !== selectedUser);
      setAvailableUsers(updatedItems)
    }
  };
  return (
    <div className={styles.inviteUsers}>
      <h3>Invite Users</h3>
      <div>
        <ul>
          {invitedUsers.map((uuid, index) => (
            <li key={index}>{allUsers.find(user => user.uuid === uuid).username}</li>
          ))}
        </ul>
        <input
          type="text"
          name="invite"
          id="formGroupInvite"
          value={invitedUsers}
          hidden
          readOnly
        />
      </div>
      <div>
        <label htmlFor="userInviteSelection">Select a user to invite:</label>
        <select id="userInviteSelection" value={selectedUser} onChange={e => setSelectedUser(e.target.value)}>
          <option value="">Select...</option>
          {availableUsers.map(user => (
            <option key={user.uuid} value={user.uuid}>{user.username}</option>
          ))}
        </select>
        <button onClick={e => handleInvite(e)}>Invite</button>
      </div>
    </div>
  );
};

export default UserInviteInput;