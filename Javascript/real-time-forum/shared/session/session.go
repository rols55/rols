package session

import (
	"errors"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

var (
	ErrNotFound          = errors.New("result not found")
	ErrExpired           = errors.New("session expired")
	CookieExpirationTime = 7200 * time.Second // 10 minutes
	secure               = false              //TODO: HTTPS
)

// Map of all the sessions, maped by the UUID of the session
var Sessions = map[int64]Session{}

// Basic Session type: email of the user this session belongs to and expire time when the session expires
type Session struct {
	UserId  int64
	Token   string
	Expires time.Time
}

// Check if the session has expired, returns true if the session has expired
func (s Session) IsExpired() bool {
	return s.Expires.Before(time.Now())
}

// Updates the session with new token
func (s *Session) Update(w *http.ResponseWriter) {

	//generate new session token
	sessionToken := uuid.Must(uuid.NewV4())

	//and update the session itself
	s.Token = sessionToken.String()
	s.Expires = time.Now().Add(CookieExpirationTime)

	//save the cookie with the token
	http.SetCookie(*w, &http.Cookie{
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		Name:     "session_token",
		Value:    s.Token,
		Expires:  s.Expires,
	})
}

// Creates new session for given user id and writes it to the ResponseWriter
func New(w *http.ResponseWriter, id int64) {

	// new values for the session
	sessionToken := uuid.Must(uuid.NewV4())

	// add the new session to the map
	Sessions[id] = Session{
		UserId:  id,
		Token:   sessionToken.String(),
		Expires: time.Now().Add(CookieExpirationTime),
	}

	//save the cookie with the token
	http.SetCookie(*w, &http.Cookie{
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		Name:     "session_token",
		Value:    sessionToken.String(),
		Expires:  time.Now().Add(CookieExpirationTime),
	})
}

// Returns session for the given request, errors ErrNotFound if server/client don't have it or if it's expired
func Get(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie("session_token")
	if err == nil { // if cookie exists
		sessionToken := cookie.Value
		for id, s := range Sessions {
			if s.Token == sessionToken {
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
		Name:     "session_token",
		Value:    "",
		MaxAge:   -1,
	}

	http.SetCookie(*w, cookie)
	return nil
}
