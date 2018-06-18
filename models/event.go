package msaevents

import (
	"encoding/json"
	"fmt"
)

type EventType string

const (
	EventTypeCreatedUser             EventType = "CREATED_USER"
	EventTypeUpdatedUser             EventType = "UPDATED_USER"
	EventTypeCreatedPasswordLost     EventType = "CREATED_PASSWORD_LOST"
	EventTypeCreatedWall             EventType = "CREATED_WALL"
	EventTypeCreatedPrivateMessage   EventType = "CREATED_PRIVATE_MESSAGE"
	EventTypeCreatedComment          EventType = "CREATED_COMMENT"
	EventTypeCreatedLike             EventType = "CREATED_LIKE"
	EventTypeCreatedFriendRequest    EventType = "CREATED_FRIEND_REQUEST"
	EventTypeCreatedFriend           EventType = "CREATED_FRIEND"
	EventTypeCreatedPhoto            EventType = "CREATED_PHOTO"
	EventTypeContentAbuseReport      EventType = "CONTENT_ABUSE_REPORT"
	EventTypeUserRegistrationGranted EventType = "USER_REGISTRATION_GRANTED"
	EventTypeUserAccountEnabled      EventType = "USER_ACCOUNT_ENABLED"
	EventTypeserAccountDisabled      EventType = "USER_ACCOUNT_DISABLED"
	EventTypeTrackingLinkClick       EventType = "TRACKING_LINK_CLICK"
)

var (
	Events = map[EventType]interface{}{
		EventTypeCreatedUser: EventCreatedUser{},
	}
)

type Event struct {
	ConfigId  string    `json:"config_id"`
	EventType EventType `json:"event_type"`
}

func (e *Event) Unmarshal(data []byte) (interface{}, error) {
	if t, ok := Events[e.EventType]; ok {
		err := json.Unmarshal(data, &t)
		return t, err
	}
	return nil, fmt.Errorf("fail to found interface for type '%s'", e.EventType)
}
