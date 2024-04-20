//import React from "react";
import Link from 'next/link'
import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';

const DropDownMenu = ({ onCloseMenu, checkNotificationsStatus }) => {
    const router = useRouter();
      const [notifications, setNotifications] = useState([]);
      async function markNotificationAsRead(notificationID){
        router.push('/profile?referrer=followers');
          try {
              const response = await fetch(`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/notification/${notificationID}`);
              if (response.ok) {
                onCloseMenu(); 
                checkNotificationsStatus();       
              } else {
                  console.error('Failed to mark notification as read');
              }
          } catch (error) {
              console.error('Error marking notification as read:', error);
          }
      };
      useEffect(() => {
          // Fetch notifications data
          const fetchNotifications = async () => {
              try {
                  const cookies = document.cookie;
                  if (cookies) {
                      const response = await fetch(`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/notifications/`, { credentials: 'include' });
                      if (response.ok) {
                          const data = await response.json();
                          setNotifications(data.notifications);
                      } else {
                          console.error('Failed to fetch notifications');
                      }
                  } else {
                      console.error('No cookies found');
                  }
              } catch (error) {
                  console.error('Error fetching notifications:', error);
              }
          };
          fetchNotifications();
      }, []);

      const handleNotificationType = (notification) => {
        let basePath = "http://localhost:3000/"; // Base path of your application
        let link = "#";
        switch(notification.Type) {
            case "follow":
                link = "/profile?referrer=followers";
                break;
            case "invite":
                link = "/profile?referrer=groups";
                break;
            case "request":
                link = `${basePath}group/${notification.Group}?referrer=requests`;
                break;
            case "event":
                link = `${basePath}group/${notification.Group}?referrer=events`;
                break;
        }                
        return (
            <Link href={link} onClick={() => markNotificationAsRead(notification.ID)}>
                {notification.Status} {notification.Text}
            </Link>
        )
      }
      return (
        <div className='flex flex-col dropDownMenu'>
            <ul className="flex flex-col gap-4">
                <div>
                    <h1>Notifications</h1>
    
                    {notifications && notifications.length > 0 ? (
                        notifications.map(notification => (
                            <div className={notification.IsRead === false ? 'unread' : 'read'} key={notification.ID}>
                                { handleNotificationType(notification) }
                            </div>
                        ))
                    ) : (
                        <li>No notifications</li>
                    )}
                </div>
            </ul>
        </div>
    );
  };
  export default DropDownMenu;
  