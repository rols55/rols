package controllers

import (
	"net/http"
	"path"
	"strconv"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/acl"
	"01.kood.tech/git/rols55/social-network/pkg/api/ws"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type EventsResponse struct {
	Status   string         `json:"status"`
	User     *model.User    `json:"user,omitempty"`
	Event    *model.Event   `json:"event,omitempty"`
	Events   []*model.Event `json:"events,omitempty"`
	IsGoing  string         `json:"is_going,omitempty"`
	Going    []*model.User  `json:"going,omitempty"`
	NotGoing []*model.User  `json:"not_going,omitempty"`
}

type EventRequests struct {
	EventId    int64    `json:"event_id,omitempty"`
	EventTitle string   `json:"event_title,omitempty"`
	Usernames  []string `json:"usernames,omitempty"`
	UserUUIDs  []string `json:"user_uuids,omitempty"`
}

func (h *BaseController) EventsGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var events []*model.Event

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	if events, err = model.GetEvents(h.db); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	res := EventsResponse{
		Status: "OK",
		User:   user,
		Events: events,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) GroupEventsGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var events []*model.Event

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	groupIdStr := r.URL.Query().Get("group")
	if groupIdStr == "" {
		logger.Error("Group ID missing")
		h.statusJSON(w, InvalidInput("Group ID missing"))
		return
	}

	groupId, err := strconv.ParseInt(groupIdStr, 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InvalidInput("Invalid Group ID"))
		return
	}

	group, err := model.GetGroupById(h.db, groupId)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	if !group.IsMember(h.db, user.Id) {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if events, err = model.GetEventsByGroupId(h.db, groupId); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	res := EventsResponse{
		Status: "OK",
		User:   user,
		Events: events,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) EventPOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("group_id") == "" {
		logger.Info("Failed to create event: No group provided")
		errMsg := "Group missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	groupId, err := strconv.ParseInt(r.FormValue("group_id"), 10, 64)
	if err != nil {
		logger.Info("Failed to create event: Invalid group")
		errMsg := "Invalid group"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	group, err := model.GetGroupById(h.db, groupId)
	if err != nil || group == nil {
		logger.Info("Failed to create event: Invalid group")
		errMsg := "Invalid group"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if !group.IsMember(h.db, user.Id) {
		logger.Info("Failed to create event: User is not a member of the group")
		errMsg := "Not a member of the group"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if r.FormValue("title") == "" {
		logger.Info("Failed to create event: No title provided")
		errMsg := "Title missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if len(r.FormValue("title")) > 60 {
		logger.Info("Failed to create event: Title too long")
		errMsg := "Title too long"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if r.FormValue("description") == "" {
		logger.Info("Failed to create event: No description provided")
		errMsg := "Description missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if r.FormValue("date") == "" {
		logger.Info("Failed to create event: No date provided")
		errMsg := "Date missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if r.FormValue("going") == "" {
		logger.Info("Failed to attend event: No attend state provided")
		errMsg := "Going missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	date, err := time.Parse("2006-01-02T15:04", r.FormValue("date"))
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	if since := time.Since(date); since >= 0 {
		logger.Info("Failed to create event: Cannot create past events (until: %v)", since)
		errMsg := "Date is in the past"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	event := &model.Event{
		GroupId:     group.Id,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Date:        date,
	}

	event.Create(h.db)
	logger.Info("Event %v created, id = %v", event.Title, event.Id)

	if err = event.AttendEvent(h.db, user.Id, r.FormValue("going") == "true"); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	members, err := group.GetMemebers(h.db)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	for _, member := range members {
		if member.Id == user.Id {
			continue
		}
		notification := model.Notification{}
		currentTime := time.Now()

		err = notification.InsertNotification(h.db, user.UUID, member.UUID, user.Username+" created new event!", "event", int(event.GroupId), currentTime, false)
		if err != nil {
			panic(err)
		}

		notificationData := ws.DataJSON{
			Type:         "notification",
			ReciverUUID:  member.UUID,
			Notification: user.Username + " created new event!",
			Timestamp:    currentTime,
		}

		h.cp.Notify(notificationData)
	}

	res := EventsResponse{
		Status: "OK",
		User:   user,
		Event:  event,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) EventGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var eventId int64
	var event *model.Event

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	eventId, err = strconv.ParseInt(path.Base(r.URL.Path), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	event, err = model.GetEventById(h.db, eventId)
	if err != nil || event == nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	group, err := model.GetGroupById(h.db, event.GroupId)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	if !group.IsMember(h.db, user.Id) {
		logger.Info("User is not a member of the group")
		errMsg := "Not a member of the group"
		h.statusJSON(w, InvalidInput(errMsg))
	}

	res := EventsResponse{
		Status: "OK",
		User:   user,
		Event:  event,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) EventAttendPOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("event_id") == "" {
		logger.Info("Failed to attend event: No event provided")
		errMsg := "Event missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	eventId, err := strconv.ParseInt(r.FormValue("event_id"), 10, 64)
	if err != nil {
		logger.Info("Failed to attend event: Invalid event")
		errMsg := "Invalid event"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	event, err := model.GetEventById(h.db, eventId)
	if err != nil || event == nil {
		logger.Info("Failed to attend event: Invalid event")
		errMsg := "Invalid event"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	group, err := model.GetGroupById(h.db, event.GroupId)
	if err != nil || group == nil {
		logger.Info("Failed to attend event: Invalid group")
		errMsg := "Invalid group"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if !group.IsMember(h.db, user.Id) {
		logger.Info("Failed to attend event: User is not a member of the group")
		errMsg := "Not a member of the group"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if r.FormValue("going") == "" {
		logger.Info("Failed to attend event: No attend state provided")
		errMsg := "Going missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if err = event.AttendEvent(h.db, user.Id, r.FormValue("going") == "true"); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	res := EventsResponse{
		Status: "OK",
		User:   user,
		Event:  event,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) EventAttendGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	eventIdStr := r.URL.Query().Get("event")
	if eventIdStr == "" {
		logger.Error("Event ID missing")
		h.statusJSON(w, InvalidInput("Event ID missing"))
		return
	}

	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InvalidInput("Invalid Event ID"))
		return
	}

	event, err := model.GetEventById(h.db, eventId)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	group, err := model.GetGroupById(h.db, event.GroupId)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	if !group.IsMember(h.db, user.Id) {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	going, err := event.GetAttendees(h.db)
	if err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	notGoing, err := event.GetNotAttending(h.db)
	if err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	isGoing := "pending"
	for _, u := range going {
		if u.Id == user.Id {
			isGoing = "true"
			break
		}
	}

	if isGoing != "true" && notGoing != nil {
		for _, u := range notGoing {
			if u.Id == user.Id {
				isGoing = "false"
				break
			}
		}
	}

	res := EventsResponse{
		Status:   "OK",
		IsGoing:  isGoing,
		Going:    going,
		NotGoing: notGoing,
	}

	h.writeJSON(w, res)
}
