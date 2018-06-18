package msaevents

import (
	"time"
)

const (
	fallbackLanguage = "EN"
)

type EventCreatedUser struct {
	Event
	Data CreatedUser `json:"data"`
}

type CreatedUser struct {
	AccountEnabled bool     `json:"account_enabled"`
	AccountExpired bool     `json:"account_expired"`
	Authorities    []string `json:"authorities"`
	CreatedDate    string   `json:"created_date"`
	CurrentStatus  struct {
		LastConnectionDate string `json:"last_connection_date"`
		State              string `json:"state"`
	} `json:"current_status"`
	DateOfBirth   string        `json:"date_of_birth"`
	DisplayedName string        `json:"displayed_name"`
	Email         string        `json:"email"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	FullName      string        `json:"full_name"`
	Gender        string        `json:"gender"`
	Id            int64         `json:"id"`
	IdStr         string        `json:"id_str"`
	PhoneNumber   string        `json:"phone_number"`
	Presentation  string        `json:"presentation"`
	CustomFields  []CustomField `json:"custom_fields"`
}

func (u *CreatedUser) GetReadableDateOfBirth() string {
	if len(u.DateOfBirth) < 10 {
		return ""
	}
	if t, err := time.Parse("2006-02-01", u.DateOfBirth[:10]); err == nil {
		return t.Format("01/02/2006")
	}
	return ""
}

func (u *CreatedUser) GetCustomFieldsValues(lang string) []CustomFieldMapLabelValue {
	if u.CustomFields == nil {
		return nil
	}
	m := make([]CustomFieldMapLabelValue, len(u.CustomFields))
	for i, field := range u.CustomFields {
		m[i] = CustomFieldMapLabelValue{Label: field.Field.Label(lang), Value: field.Data.StringValue()}
	}
	return m
}
