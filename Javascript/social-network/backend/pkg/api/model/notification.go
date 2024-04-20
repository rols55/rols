package model

import (
	"database/sql"
	"fmt"
	"time"
	//"errors"
	//"time"
	//"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type Notification struct {
	ID           int
	SenderUUID   string
	ReceiverUUID string
	Text         string
	Type         string
	Group        int
	Timestamp    time.Time
	IsRead       bool
}

func (p *Notification) InsertNotification(db *sql.DB, senderUUID string, receiverUUID string, text string, nType string, group int, timestamp time.Time, isRead bool) error {
	_, err := db.Exec("INSERT INTO notifications (sender_uuid, reciver_uuid, notification_text, notification_type, group_id, timestamp, is_read) VALUES (?, ?, ?, ?, ?, ?, ?)", senderUUID, receiverUUID, text, nType, group, timestamp, isRead)
	return err
}

// Retrieve notifications for a specific user
func (p *Notification) GetNotificationsForUser(db *sql.DB, UserUuid string) ([]*Notification, error) {
	rows, err := db.Query("SELECT id, sender_uuid, reciver_uuid, notification_text, notification_type, group_id, timestamp, is_read FROM notifications WHERE reciver_uuid = ?", UserUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*Notification
	for rows.Next() {
		var notification Notification
		err := rows.Scan(&notification.ID, &notification.SenderUUID, &notification.ReceiverUUID, &notification.Text, &notification.Type, &notification.Group, &notification.Timestamp, &notification.IsRead)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}
	return notifications, nil
}

// Retrieve notification by its ID
func (p *Notification) GetNotificationByID(db *sql.DB, notificationID int) (*Notification, error) {
	// Query to retrieve a notification by its ID
	row := db.QueryRow("SELECT id, sender_uuid, reciver_uuid, notification_text, notification_type, group_id, timestamp, is_read FROM notifications WHERE id = ?", notificationID)

	// Initialize a Notification struct to store the retrieved notification
	var notification Notification

	// Scan the query result into the Notification struct
	err := row.Scan(&notification.ID, &notification.SenderUUID, &notification.ReceiverUUID, &notification.Text, &notification.Type, &notification.Group, &notification.Timestamp, &notification.IsRead)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("notification with ID %d not found", notificationID)
		}
		return nil, err
	}

	// Return the retrieved notification
	return &notification, nil
}

// Update the IsRead field of a notification
func (p *Notification) UpdateNotificationIsRead(db *sql.DB, notificationID int, isRead bool) error {
	// Execute an UPDATE SQL statement to update the IsRead field of the notification
	_, err := db.Exec("UPDATE notifications SET is_read = ? WHERE id = ?", isRead, notificationID)
	return err
}
