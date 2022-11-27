package main

import "github.com/google/uuid"

type Practices map[uuid.UUID]Practice

type Practice struct {
	ID          uuid.UUID     `json:"id"`
	PracticeId  string        `json:"practiceId"`
	Name        string        `json:"name,omitempty"`
	Address1    string        `json:"address1,omitempty"`
	City        string        `json:"city,omitempty"`
	State       string        `json:"state,omitempty"`
	Zip         string        `json:"zip,omitempty"`
	PhoneNumber string        `json:"phoneNumber,omitempty"`
	Admin       bool          `json:"admin,omitempty"`
	Private     bool          `json:"private,omitempty"`
	CareTeam    []interface{} `json:"careTeam,omitempty"`
	Notes       string        `json:"notes,omitempty"`
	Records     []interface{} `json:"records"`
	Programs    *string       `json:"programs,omitempty"`
	InvitedBy   InvitedBy     `json:"invitedBy, omitempty"`
}

type InvitedBy struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}
