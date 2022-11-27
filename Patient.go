package main

import "github.com/google/uuid"

type PatientResp struct {
	Status int         `json:"status"`
	Data   PatientData `json:"data"`
	Ticket AuthTicket  `json:"ticket"`
}

type PatientData struct {
	Patient                 Patient `json:"patient"`
	UnreadMessages          int     `json:"unreadMessages"`
	UnresolvedNotifications int     `json:"unresolvedNotifications"`
}

type Patient struct {
	ID                   uuid.UUID            `json:"id"`
	FirstName            string               `json:"firstName"`
	LastName             string               `json:"lastName"`
	Email                string               `json:"email"`
	LVAccount            bool                 `json:"lvAccount"`
	DateOfBirth          int                  `json:"dob"`
	LastReportViewTime   int                  `json:"lastReportViewTime"`
	Practices            Practices            `json:"practices"`
	TimeOffsetMinutes    int                  `json:"timeOffsetMinutes"`
	Details              interface{}          `json:"details"`
	Created              int64                `json:"created"`
	Messages             Messages             `json:"messages"`
	Notifications        PatientNotifications `json:"notifications"`
	Programs             interface{}          `json:"programs"`
	HighGlucoseThreshold int                  `json:"highGlucoseThreshold"`
	LowGlucoseThreshold  int                  `json:"lowGlucoseThreshold"`
	TargetRangeHigh      int                  `json:"targetRangeHigh"`
	TargetRangeLow       int                  `json:"targetRangeLow"`
}

type PatientNotifications struct {
	NotificationsOn bool        `json:"notifications_on"`
	EmailPreference string      `json:"emailPreference"`
	HighThreshold   int         `json:"highThreshold"`
	LowThreshold    int         `json:"lowThreshold"`
	TargetRange     TargetRange `json:"targetRange"`
}

type TargetRange struct {
	Max int `json:"max"`
	Min int `json:"min"`
}
