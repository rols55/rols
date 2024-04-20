"use client"
import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import DropDownMenu from './dropDown';
import { useWebSocket } from '@/app/WebSocketProvider';

export default function Notifications() {
  const [openMenu, setOpenMenu] = useState(false);
  const [hasReadNotifications, setHasReadNotifications] = useState(false);  
  const socket = useWebSocket();

  useEffect(() => {
    const saved = localStorage.getItem('hasReadNotifications');
    if (saved) {
      setHasReadNotifications(JSON.parse(saved));
    }
  }, [])

  useEffect(() => {
    if (!socket) return;
    socket.handleOnNotification = (data) => {
      handleNotificationReceived(data);
    }
  }, [socket]);

  useEffect(() => {
    if (typeof window !== 'undefined') {
      localStorage.setItem('hasReadNotifications', JSON.stringify(hasReadNotifications));
    }
  }, [hasReadNotifications]);

  const toggleMenu = () => {
    setOpenMenu(prev => !prev);
  };

  const handleNotificationReceived = (data) => {
    setHasReadNotifications(true);
  };

  const checkNotificationsStatus = async () => {
    try {
      const cookies = document.cookie;
      if (cookies) {
        const response = await fetch(`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/notifications/`, { credentials: 'include' });
        if (response.ok) {
          const data = await response.json();
          if (data.notifications && data.notifications.length > 0) {
            const hasReadNotifications = data.notifications.some(notification => !notification.IsRead);
            setHasReadNotifications(hasReadNotifications); 
          } else {
            console.log("There are no notifications.");
          }
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
 

  return (
    <div>
      <Link href="#" className="icon-button" onClick={toggleMenu}>
        <span className="material-symbols-outlined">notifications</span>
        <div className={hasReadNotifications === false ? 'icon-button__badge' : 'icon-button__badge visible'}></div>
        
      </Link>
      {openMenu && <DropDownMenu onCloseMenu={toggleMenu} checkNotificationsStatus={checkNotificationsStatus} />}
    </div>
  );
};
