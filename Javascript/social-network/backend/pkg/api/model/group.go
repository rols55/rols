package model

import (
	"database/sql"
	"errors"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type Group struct {
	Id          int64     `json:"id,omitempty"`
	UserId      int64     `json:"user_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

// Create new user
func (g *Group) Create(db *sql.DB) error {
	result, err := db.Exec(`INSERT INTO groups (
		user_id,
		title,
		description,
		timestamp
	)
	values(?,?,?,CURRENT_TIMESTAMP)`, g.UserId, g.Title, g.Description)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	if g.Id, err = result.LastInsertId(); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

func GetGroups(db *sql.DB) ([]*Group, error) {
	rows, err := db.Query(`SELECT
		id,
		user_id,
		title,
		description,
		timestamp
		FROM groups ORDER BY timestamp DESC`)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	groups := make([]*Group, 0)
	for rows.Next() {
		var group Group
		err = rows.Scan(
			&group.Id,
			&group.UserId,
			&group.Title,
			&group.Description,
			&group.Timestamp,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		groups = append(groups, &group)
	}
	if len(groups) == 0 {
		return nil, ErrNotFound
	}

	return groups, nil
}

func GetGroupById(db *sql.DB, Id int64) (*Group, error) {
	row := db.QueryRow("SELECT * FROM groups WHERE id =?", Id)
	group := &Group{}
	if err := row.Scan(
		&group.Id,
		&group.UserId,
		&group.Title,
		&group.Description,
		&group.Timestamp,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return group, nil
}

func (g *Group) InviteUsers(db *sql.DB, inviterId int64, userUUIDs []string) error {

	if canIvite, err := g.CanInviteUsers(db, inviterId); !canIvite {
		if err != nil {
			return err
		}
		return errors.New(logger.GetCurrentFuncName() + " User cannot invite to the group")
	}

	if userUUIDs == nil {
		return errors.New(logger.GetCurrentFuncName() + " No user UUIDs provided for group invites")
	}

	for _, uuid := range userUUIDs {
		if user, err := GetUserByUUID(db, uuid); err == nil {
			if user.Id == g.UserId {
				continue
			}
			state, err := g.GetUserInviteState(db, user.Id)
			if err == ErrNotFound {
				if err := g._createUserInvite(db, user.Id, false); err != nil {
					logger.Error(logger.GetCurrentFuncName() + " No user with UUID = " + uuid)
					continue
				}
			}

			if state == "request" {
				if err := g._updateUserInvite(db, user.Id, "accepted"); err != nil {
					logger.Error(logger.GetCurrentFuncName() + " No user with UUID = " + uuid)
					continue
				}
			}
		} else {
			logger.Error(logger.GetCurrentFuncName() + " " + err.Error())
		}
	}

	return nil
}

func (g *Group) _createUserInvite(db *sql.DB, userId int64, request bool) error {

	state := "pending"
	if request {
		state = "requested"
	}

	_, err := db.Exec(`INSERT INTO group_invites (
		user_id,
		group_id,
		state
	)
	values(?,?,?)`, userId, g.Id, state)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return nil
}

func (g *Group) _updateUserInvite(db *sql.DB, userId int64, state string) error {
	if _, err := db.Exec(`UPDATE group_invites SET
		state = ?
		WHERE user_id = ? AND group_id = ?`,
		state, userId, g.Id); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return nil
}

func (g *Group) GetUserInviteState(db *sql.DB, userId int64) (string, error) {
	row := db.QueryRow("SELECT state FROM group_invites WHERE user_id = ? AND group_id = ?", userId, g.Id)
	var state string
	if err := row.Scan(&state); err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return state, nil
}

func (g *Group) CanInviteUsers(db *sql.DB, inviterId int64) (bool, error) {
	if g.UserId == inviterId {
		return true, nil
	}

	state, err := g.GetUserInviteState(db, inviterId)
	if err == nil && state == "accepted" {
		return true, nil
	} else if err == nil {
		return false, nil
	}

	return false, err
}

func (g *Group) AcceptInvite(db *sql.DB, userId int64) error {
	state, err := g.GetUserInviteState(db, userId)
	if err != nil {
		return err
	}
	if state == "pending" {
		return g._updateUserInvite(db, userId, "accepted")
	}
	return errors.New(logger.GetCurrentFuncName() + " User has no pending invite")
}

func (g *Group) RejectInvite(db *sql.DB, userId int64) error {
	state, err := g.GetUserInviteState(db, userId)
	if err != nil {
		return err
	}
	if state == "pending" {
		return g._updateUserInvite(db, userId, "rejected")
	}
	return errors.New(logger.GetCurrentFuncName() + " User has no pending invite")
}

func (g *Group) RequestInvite(db *sql.DB, userId int64) error {
	state, err := g.GetUserInviteState(db, userId)
	if err == ErrNotFound {
		return g._createUserInvite(db, userId, true)
	}
	if state == "pending" {
		return g._updateUserInvite(db, userId, "accepted")
	}
	return errors.New(logger.GetCurrentFuncName() + " User cannot request invite to the group")
}

func (g *Group) ApproveRequest(db *sql.DB, approverId int64, userId int64) error {
	if g.UserId != approverId {
		return errors.New(logger.GetCurrentFuncName() + " User cannot approve group invites that he is not the author of")
	}
	state, err := g.GetUserInviteState(db, userId)
	if err != nil {
		return err
	}
	if state != "requested" {
		return errors.New(logger.GetCurrentFuncName() + " User cannot approve none existing invites")
	}
	g._updateUserInvite(db, userId, "accepted")
	return nil
}

func (g *Group) RejectRequest(db *sql.DB, approverId int64, userId int64) error {
	if g.UserId != approverId {
		return errors.New(logger.GetCurrentFuncName() + " User cannot approve group invites that he is not the author of")
	}
	state, err := g.GetUserInviteState(db, userId)
	if err != nil {
		return err
	}
	if state != "requested" {
		return errors.New(logger.GetCurrentFuncName() + " User cannot approve none existing invites")
	}
	g._updateUserInvite(db, userId, "rejected")
	return nil
}

func (g *Group) GetRequests(db *sql.DB) ([]*User, error) {

	users := make([]*User, 0)

	query := `
	SELECT u.* 
	FROM users u
	JOIN group_invites gi ON u.id = gi.user_id
	WHERE gi.state = 'requested' AND gi.group_id = ?
	`

	rows, err := db.Query(query, g.Id)
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

func (g *Group) GetMemebers(db *sql.DB) ([]*User, error) {

	author, err := GetUserById(db, g.UserId)
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0)
	users = append(users, author)

	query := `
	SELECT u.* 
	FROM users u
	JOIN group_invites gi ON u.id = gi.user_id
	WHERE gi.state = 'accepted' AND gi.group_id = ?
	`

	rows, err := db.Query(query, g.Id)
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

func (g *Group) GetPosts(db *sql.DB) ([]*Post, error) {
	rows, err := db.Query(`SELECT * FROM posts WHERE group_id = ? ORDER BY creation_date DESC`, g.Id)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()
	posts := make([]*Post, 0)
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Author,
			&post.Title,
			&post.Text,
			&post.ImagePost,
			&post.Privacy,
			&post.Followers,
			&post.CreationDate,
			&post.GroupId,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		posts = append(posts, &post)
	}
	if len(posts) == 0 {
		return nil, ErrNotFound
	}

	return posts, nil
}

func GetGroupsByUser(db *sql.DB, userId int64) ([]*Group, error) {
	rows, err := db.Query(`SELECT * FROM groups WHERE id IN (
		SELECT id FROM groups WHERE user_id = ?
		UNION
		SELECT group_id FROM group_invites WHERE user_id = ? AND state = 'accepted'
		) ORDER BY timestamp DESC`, userId, userId)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	groups := make([]*Group, 0)
	for rows.Next() {
		var group Group
		err = rows.Scan(
			&group.Id,
			&group.UserId,
			&group.Title,
			&group.Description,
			&group.Timestamp,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		groups = append(groups, &group)
	}
	if len(groups) == 0 {
		return nil, ErrNotFound
	}

	return groups, nil

}

func (g *Group) IsMember(db *sql.DB, userId int64) bool {
	if g.UserId == userId {
		return true
	}
	query := `
	SELECT COUNT(*) FROM group_invites
	WHERE user_id = ? AND group_id = ? AND state = 'accepted'
	`
	var count int64
	err := db.QueryRow(query, userId, g.Id).Scan(&count)
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}
