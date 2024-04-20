"use client";
import styles from "./page.module.css";
import apiFetch from "../apiFetch";
import { useState, useEffect } from 'react';

const ApproveRequest = ({groupId}) => {
  const [data, setData] = useState(null);
  const [requests, setRequests] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchRequests = async () => {
      try {
        apiFetch(`/groups/requests${groupId ? "?group="+groupId : ""}`, (resp) => {setData(resp);setRequests(resp.group_requests)});
      } catch (error) {
        setError(error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchRequests();
  }, []);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  let path = '';
  let index = -1;

  const handleClick = (e, idx) => {
    path = e.target.value;
    index = idx;
  };
  const handleApproveRequest = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.target);
    try {
      apiFetch(path, null, {
        method: 'POST',
        body: formData
      }).then(() => setRequests(requests.filter((req, i) => i != index )));
    } catch (error) {
      console.error(error);
    }
  };
  if (requests) {
    return (
      <div className={styles.groupList}>
        <div>
          {requests.map(req => (
            <div key={req.group_id}>
              <ul>
                {req.user_uuids.map((uuid, idx) => 
                <li key={uuid}>
                  <form onSubmit={handleApproveRequest}>
                    <p>{req.usernames[idx]}
                      <input
                          type="text"
                          name="group_id"
                          value={req.group_id}
                          hidden
                          readOnly
                        />
                        <input
                          type="text"
                          name="requester_uuid"
                          value={uuid}
                          hidden
                          readOnly
                        />
                        <button type="submit" onClick={(e) => handleClick(e, idx)} name="accept" value="/groups/request/approve">Accept</button>
                        <button type="submit" onClick={(e) => handleClick(e, idx)} name="reject" value="/groups/request/reject">Reject</button>
                    </p>
                  </form>
                </li>)}
              </ul>
            </div>
          ))}
        </div>
      </div>
    );
  }
  return (
    <div>
    </div>
  );
};

export default ApproveRequest;