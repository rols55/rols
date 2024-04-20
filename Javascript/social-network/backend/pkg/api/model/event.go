package model

import (
	"database/sql"
	"errors"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type Event struct {
	Id          int64     `json:"id,omitempty"`
	GroupId     int64     `json:"group_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (e *Event) Create(db *sql.DB) error {
	result, err := db.Exec(`INSERT INTO events (
		group_id,
		title,
		description,
		date,
		timestamp
	)
	values(?,?,?,?,CURRENT_TIMESTAMP)`, e.GroupId, e.Title, e.Description, e.Date)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	if e.Id, err = result.LastInsertId(); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

func GetEvents(db *sql.DB) ([]*Event, error) {
	rows, err := db.Query(`SELECT
		id,
		group_id,
		title,
		description,
		date,
		timestamp
		FROM events ORDER BY date DESC`)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	events := make([]*Event, 0)
	for rows.Next() {
		var event Event
		err = rows.Scan(
			&event.Id,
			&event.GroupId,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Timestamp,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		events = append(events, &event)
	}
	if len(events) == 0 {
		return nil, ErrNotFound
	}

	return events, nil
}

func GetEventById(db *sql.DB, Id int64) (*Event, error) {
	row := db.QueryRow("SELECT * FROM events WHERE id =?", Id)
	event := &Event{}
	if err := row.Scan(
		&event.Id,
		&event.GroupId,
		&event.Title,
		&event.Description,
		&event.Date,
		&event.Timestamp,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return event, nil
}

func AttendEvent(db *sql.DB, userId int64, eventId int64, going bool) error {
	var state bool
	row := db.QueryRow("SELECT is_going FROM event_attendance WHERE event_id = ? AND user_id = ?", eventId, userId)
	if err := row.Scan(&state); err != nil {
		if err == sql.ErrNoRows {
			_, err := db.Exec(`INSERT INTO event_attendance (
				event_id,
				user_id,
				is_going
			)
			values(?,?,?)`, eventId, userId, going)
			if err != nil {
				return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
			}
			return nil
		}
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	if state != going {
		if _, err := db.Exec(`UPDATE event_attendance SET
		is_going = ?
		WHERE event_id = ? AND user_id = ?`,
			going, eventId, userId); err != nil {
			return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
	}

	return nil
}

func (e *Event) AttendEvent(db *sql.DB, userId int64, going bool) error {
	return AttendEvent(db, userId, e.Id, going)
}

func (e *Event) GetAttendees(db *sql.DB) ([]*User, error) {

	users := make([]*User, 0)

	query := `
	SELECT u.* 
	FROM users u
	JOIN event_attendance ea ON u.id = ea.user_id
	WHERE ea.is_going = TRUE AND ea.event_id = ?
	`

	rows, err := db.Query(query, e.Id)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.Id,
			&user.UUID,
			&user.Username,
			&user.Firstname,
			&user.Lastname,
			&user.Gender,
			&user.Age,
			&user.Email,
			&user.Public,
			&user.Nickname,
			&user.AboutMe,
			&user.Image,
			&user.Password,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		users = append(users, &user)
	}
	if len(users) == 0 {
		return nil, ErrNotFound
	}

	return users, nil
}

func (e *Event) GetNotAttending(db *sql.DB) ([]*User, error) {

	users := make([]*User, 0)

	query := `
	SELECT u.* 
	FROM users u
	JOIN event_attendance ea ON u.id = ea.user_id
	WHERE ea.is_going = FALSE AND ea.event_id = ?
	`

	rows, err := db.Query(query, e.Id)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.Id,
			&user.UUID,
			&user.Username,
			&user.Firstname,
			&user.Lastname,
			&user.Gender,
			&user.Age,
			&user.Email,
			&user.Public,
			&user.Nickname,
			&user.AboutMe,
			&user.Image,
			&user.Password,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		users = append(users, &user)
	}
	if len(users) == 0 {
		return nil, ErrNotFound
	}

	return users, nil
}

func GetEventsByGroupId(db *sql.DB, groupId int64) ([]*Event, error) {
	rows, err := db.Query(`SELECT
		id,
		group_id,
		title,
		description,
		date,
		timestamp
		FROM events WHERE group_id = ? ORDER BY date DESC`, groupId)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	events := make([]*Event, 0)
	for rows.Next() {
		var event Event
		err = rows.Scan(
			&event.Id,
			&event.GroupId,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Timestamp,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		events = append(events, &event)
	}
	if len(events) == 0 {
		return nil, ErrNotFound
	}

	return events, nil
}
