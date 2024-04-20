package session

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	gUUID "github.com/gofrs/uuid"
)

var (
	ErrNotFound          = errors.New("session not found")
	ErrExpired           = errors.New("session expired")
	CookieExpirationTime = 7200 * time.Second // 10 minutes
	secure               = false              //TODO: HTTPS
	SameSiteMode         = http.SameSiteNoneMode
	CookieName           = "session_token"
)

// Map of all the sessions, maped by the UUID of the session
var Sessions = map[int64]Session{}

// Basic Session type: email of the user this session belongs to and expire time when the session expires
type Session struct {
	UserId   int64
	UserUUID string
	Token    string
	Expires  time.Time
}

// Check if the session has expired, returns true if the session has expired
func (s Session) IsExpired() bool {
	return s.Expires.Before(time.Now())
}

// Updates the session with new token
func (s *Session) Update(w *http.ResponseWriter) {

	//generate new session token
	sessionToken := fmt.Sprintf("%v.%v", gUUID.Must(gUUID.NewV4()).String(), s.UserUUID)

	//and update the session itself
	s.Token = sessionToken
	s.Expires = time.Now().Add(CookieExpirationTime)

	//save the cookie with the token
	http.SetCookie(*w, &http.Cookie{
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		Name:     CookieName,
		Value:    s.Token,
		Expires:  s.Expires,
		SameSite: SameSiteMode,
	})
}

// Creates new session for given user id and writes it to the ResponseWriter
func New(w *http.ResponseWriter, id int64, uuid string) Session {

	// new values for the session
	sessionToken := fmt.Sprintf("%v.%v", gUUID.Must(gUUID.NewV4()).String(), uuid)

	// add the new session to the map
	Sessions[id] = Session{
		UserId:   id,
		UserUUID: uuid,
		Token:    sessionToken,
		Expires:  time.Now().Add(CookieExpirationTime),
	}

	//save the cookie with the token
	http.SetCookie(*w, &http.Cookie{
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		Name:     CookieName,
		Value:    sessionToken,
		Expires:  Sessions[id].Expires,
		SameSite: SameSiteMode,
	})

	return Sessions[id]
}

func GetUserIdByToken(token string) (int64, error) {
	// Iterate over sessions to find the session with the matching token
	for _, session := range Sessions {
		if session.Token == token {
			// Return the UserId if the token matches
			return session.UserId, nil
		}
	}
	// If no session with the matching token is found, return an error
	return 0, errors.New("session not found for token")
}

// Returns session for the given request, errors ErrNotFound if server/client don't have it or if it's expired
func Get(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie(CookieName)
	if err == nil { // if cookie exists
		for id, s := range Sessions {
			if s.Token == cookie.Value {
				if !s.IsExpired() {
					return &s, nil
				}
				delete(Sessions, id)
				return nil, ErrExpired
			}
		}
	} else if err != http.ErrNoCookie {
		return nil, err
	}
	return nil, ErrNotFound
}

func GetSession(sessionToken string) (*Session, error) {
	for id, s := range Sessions {
		if s.Token == sessionToken {
			if !s.IsExpired() {
				return &s, nil
			}
			delete(Sessions, id)
			return nil, ErrExpired
		}
	}
	return nil, ErrNotFound
}

// Delete the session for the given request, errors ErrNotFound if server/client don't have it (error can be ignored)
func Delete(w *http.ResponseWriter, r *http.Request) error {
	sess, err := Get(r)
	if err != nil { // if session does not exist
		return err
	}

	delete(Sessions, sess.UserId)

	cookie := &http.Cookie{
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		Name:     CookieName,
		Value:    "",
		MaxAge:   -1,
		SameSite: SameSiteMode,
	}

	http.SetCookie(*w, cookie)
	return nil
}
