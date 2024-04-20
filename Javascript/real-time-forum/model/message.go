package model

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/shared/logger"
	"time"
)

type Message struct {
	Id           int64     `json:"id,omitempty"`
	Sender_UUID  string    `json:"sender_uuid,omitempty"`
	Reciver_UUID string    `json:"reciver_uuid,omitempty"`
	Message_text string    `json:"message_text,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	Is_read      bool      `json:"is_read,omitempty"`
}

type MessagePartner struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (m *Message) Create(db *sql.DB) (*Message, error) {
	err := db.QueryRow(`INSERT INTO messages(sender_id, reciver_id, message_text, timestamp) 
		values( ?, ?, ?, CURRENT_TIMESTAMP) RETURNING id, timestamp`,
		m.Sender_UUID, m.Reciver_UUID, m.Message_text).Scan(&m.Id, &m.Timestamp)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	m.Is_read = false
	return m, nil
}

func (m *Message) Read(db *sql.DB) error {
	_, err := db.Exec("UPDATE messages SET is_read = TRUE WHERE message_id = ?", m.Id)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

// Returns messages between two users ordered by timestamp, pagination by 10
func GetMessages(db *sql.DB, user string, target string, page int) ([]*Message, error) {
	offset := page * 10
	rows, err := db.Query(`
	SELECT id, sender_id, reciver_id, message_text, timestamp, is_read
	FROM Messages
	WHERE (sender_id = ? AND reciver_id = ?) OR (sender_id = ? AND reciver_id = ?)
	ORDER BY timestamp DESC
	LIMIT 10 OFFSET ?
	`, user, target, target, user, offset)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()
	messages := make([]*Message, 0)
	for rows.Next() {
		var message Message
		err = rows.Scan(
			&message.Id,
			&message.Sender_UUID,
			&message.Reciver_UUID,
			&message.Message_text,
			&message.Timestamp,
			&message.Is_read)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		messages = append(messages, &message)
	}
	if len(messages) == 0 {
		return nil, ErrNotFound
	}
	return messages, nil
}

func GetMessagePartnersByUserId(db *sql.DB, userID int64) ([]*MessagePartner, error) {
	rows, err := db.Query(`
		SELECT MAX(timestamp) AS last_message_timestamp, 
			CASE WHEN sender_id = ? THEN recipient_id ELSE sender_id END AS chat_partner
		FROM Message
		WHERE sender_id = ? OR recipient_id = ?
		GROUP BY chat_partner
	`, userID, userID, userID)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	partners := []*MessagePartner{}
	for rows.Next() {
		var partner MessagePartner

		err := rows.Scan(&partner.Timestamp, &partner.Id)
		if err != nil {
			fmt.Println(err)
		}
		partners = append(partners, &partner)
	}
	return partners, nil
}
