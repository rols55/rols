package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"01.kood.tech/git/rols55/social-network/pkg/logger"

	"github.com/gofrs/uuid"
)

type User struct {
	Id        int64  `json:"id,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Gender    string `json:"sex,omitempty"`
	Age       string `json:"birthday,omitempty"`
	Email     string `json:"email,omitempty"`
	Public    bool   `json:"public"`
	Nickname  string `json:"nickname,omitempty"`
	AboutMe   string `json:"aboutme,omitempty"`
	Image     string `json:"image,omitempty"`
	Password  string `json:"-"`
}

// Create new user
func (u *User) Create(db *sql.DB) error {
	result, err := db.Exec(`INSERT INTO users(uuid,
		username, 
		firstname, 
		lastname, 
		sex, 
		birthday, 
		email,
		nickname, 
		aboutme, 
		image,
		public,
		password) values(?,?,?,?,?,?,?,?,?,?,?,?)`,
		u.UUID,
		u.Username,
		u.Firstname,
		u.Lastname,
		u.Gender,
		u.Age,
		u.Email,
		u.Nickname,
		u.AboutMe,
		u.Image,
		u.Public,
		u.Password)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	if u.Id, err = result.LastInsertId(); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

// Deletes the user
func (u *User) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE uuid = ?", u.UUID)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	u = nil // is this okay?
	// maybe return affected rows
	return nil
}

func (u *User) TogglePublic(db *sql.DB) error {
	_, err := db.Exec(`UPDATE users SET public = NOT public WHERE uuid = ?`, u.UUID)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

// Updates the user
func (u *User) UpdateUser(db *sql.DB) error {
	if _, err := db.Exec(`UPDATE users SET
		firstname = ?,
		lastname = ?,
		sex = ?,
		birthday = ?,
		email = ?,
		nickname = ?,
		aboutme = ?,
		image = ?,
		public = ?
		WHERE uuid = ?`,
		u.Firstname, u.Lastname, u.Gender, u.Age, u.Email, u.Nickname, u.AboutMe, u.Image, u.Public, u.UUID); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

// Returns the all posts by the user
func (u *User) GetPosts(db *sql.DB) ([]*Post, error) {
	return GetPostsByUserUuid(db, u.UUID)
}

// Adds a like to a post by post type (posts and comments) - Is this func needed?
func (u *User) AddLike(Post interface{}) error {
	return errors.New(logger.GetCurrentFuncName() + " not implemented")
}

// Adds a like by given post id (posts and comments) - Is this func needed?
func (u *User) AddLikeById(PostId int) error {
	return errors.New(logger.GetCurrentFuncName() + " not implemented")
}

// Queries last message partners
func (u *User) GetMessagePartners(db *sql.DB) ([]*MessagePartner, error) {
	partners, err := GetMessagePartnersByUserId(db, u.Id)
	if err != nil {
		return []*MessagePartner{}, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	for _, partner := range partners {
		err := db.QueryRow(`SELECT username FROM users WHERE id = ?`, partner.Id).Scan(&partner.Name)
		if err != nil {
			fmt.Println(err)
		}
	}
	return partners, nil
}

// Returns all the users
func GetUsers(db *sql.DB, exclude ...int64) ([]*User, error) {
	query := "SELECT * FROM users"
	if len(exclude) > 0 {
		ids := fmt.Sprintf("%v", exclude)
		ids = strings.Replace(ids, "[", "(", -1)
		ids = strings.Replace(ids, "]", ")", -1)
		query += " WHERE id NOT IN " + ids
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.UUID, &user.Username, &user.Firstname, &user.Lastname, &user.Gender, &user.Age, &user.Email, &user.Public, &user.Nickname, &user.AboutMe, &user.Image, &user.Password)
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

func GetUsersSorted(db *sql.DB, exclude *User) ([]*User, error) {
	query := `
    SELECT u.*
	FROM users u
	LEFT JOIN (
		SELECT reciver_id AS user_uuid, MAX(timestamp) AS message_date
		FROM messages
        WHERE sender_id = ?
		GROUP BY reciver_id

		UNION

		SELECT sender_id AS user_uuid, MAX(timestamp) AS message_date
		FROM messages
        WHERE reciver_id = ?
		GROUP BY sender_id

	) m ON u.uuid = m.user_uuid
	WHERE u.id <> ?
	GROUP BY u.username
	ORDER BY COALESCE(MAX(m.message_date), '1900-01-01') DESC, u.username COLLATE NOCASE;
	`

	rows, err := db.Query(query, exclude.UUID, exclude.UUID, exclude.Id)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.UUID, &user.Username, &user.Firstname, &user.Lastname, &user.Gender, &user.Age, &user.Email, &user.Public, &user.Nickname, &user.AboutMe, &user.Image, &user.Password)
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

// Returns a user by given id
func GetUserById(db *sql.DB, Id int64) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE id = ? LIMIT 1", Id)
	user := &User{}
	if err := row.Scan(&user.Id, &user.UUID, &user.Username, &user.Firstname, &user.Lastname, &user.Gender, &user.Age, &user.Email, &user.Public, &user.Nickname, &user.AboutMe, &user.Image, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return user, nil
}

// Returns user by its registrated email
// this method should not be used, github users with no public email will not have email set
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE email = ? LIMIT 1", email)
	user := &User{}
	if err := row.Scan(&user.Id, &user.UUID, &user.Username, &user.Firstname, &user.Lastname, &user.Gender, &user.Age, &user.Email, &user.Public, &user.Nickname, &user.AboutMe, &user.Image, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return user, nil
}

// Returns user by its registrated username
func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE username = ? OR email = ? LIMIT 1", username, username)
	user := &User{}
	if err := row.Scan(&user.Id, &user.UUID, &user.Username, &user.Firstname, &user.Lastname, &user.Gender, &user.Age, &user.Email, &user.Public, &user.Nickname, &user.AboutMe, &user.Image, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return user, nil
}

// Returns user by its UUID
func GetUserByUUID(db *sql.DB, uuid string) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE uuid = ? LIMIT 1", uuid)
	user := &User{}
	if err := row.Scan(&user.Id, &user.UUID, &user.Username, &user.Firstname, &user.Lastname, &user.Gender, &user.Age, &user.Email, &user.Public, &user.Nickname, &user.AboutMe, &user.Image, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return user, nil
}

// Deletes a user by given id
func DeleteUserById(db *sql.DB, Id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", Id)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	// maybe return affected rows
	return nil
}

// Deletes a user by given id
func DeleteUserByUUID(db *sql.DB, UUID int) error {
	_, err := db.Exec("DELETE FROM users WHERE uuid = ?", UUID)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	// maybe return affected rows
	return nil
}

// Deletes a user by given id
func DeleteUserByEmail(db *sql.DB, Email string) error {
	_, err := db.Exec("DELETE FROM users WHERE email = ?", Email)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	// maybe return affected rows
	return nil
}

// Creates a new user
func CreateUser(db *sql.DB, username string, firstname string, lastname string, gender string, age string, email string, nickname string, aboutme string, image string, public bool, password string) (*User, error) {
	user := &User{
		Username:  username,
		Firstname: firstname,
		Lastname:  lastname,
		Gender:    gender,
		Age:       age,
		Email:     email,
		Nickname:  nickname,
		AboutMe:   aboutme,
		Image:     image,
		Public:    public,
		Password:  password,
		UUID:      uuid.Must(uuid.NewV4()).String(),
	}
	if err := user.Create(db); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetGroupInvites(db *sql.DB) ([]*Group, error) {
	query := `
		SELECT group_id  
		FROM group_invites
		WHERE user_id = ? AND state = 'pending'
		`

	rows, err := db.Query(query, u.Id)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	groups := make([]*Group, 0)
	for rows.Next() {
		var group_id int64
		err = rows.Scan(&group_id)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		if group, err := GetGroupById(db, group_id); err == nil {

			groups = append(groups, group)
		} else {
			logger.Error("Could not find group: %v", err.Error())
		}
	}
	if len(groups) == 0 {
		return nil, ErrNotFound
	}

	return groups, nil
}

func (u *User) AcceptGroupInvite(db *sql.DB, groupId int64) error {
	group, err := GetGroupById(db, groupId)
	if err != nil {
		return err
	}
	return group.AcceptInvite(db, u.Id)
}

func (u *User) RejectGroupInvite(db *sql.DB, groupId int64) error {
	group, err := GetGroupById(db, groupId)
	if err != nil {
		return err
	}
	return group.RejectInvite(db, u.Id)
}

func (u *User) RequestGroupInvite(db *sql.DB, groupId int64) error {
	group, err := GetGroupById(db, groupId)
	if err != nil {
		return err
	}
	return group.RequestInvite(db, u.Id)
}

func (u *User) ApproveRequest(db *sql.DB, groupId int64, userId int64) error {
	group, err := GetGroupById(db, groupId)
	if err != nil {
		return err
	}
	return group.ApproveRequest(db, u.Id, userId)
}

func (u *User) RejectRequest(db *sql.DB, groupId int64, userId int64) error {
	group, err := GetGroupById(db, groupId)
	if err != nil {
		return err
	}
	return group.RejectRequest(db, u.Id, userId)
}

func (u *User) GetGroups(db *sql.DB) ([]*Group, error) {
	return GetGroupsByUser(db, u.Id)
}
