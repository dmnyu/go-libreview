package main

import (
	"github.com/google/uuid"
)

type UserResp struct {
	Status int  `json:"status"`
	Data   Data `json:"data"`
}

type Data struct {
	User          User              `json:"user"`
	Practices     Practices         `json:"practices"`
	Messages      Messages          `json:"messages"`
	Notifications UserNotifications `json:"notifications"`
	AuthTicket    AuthTicket        `json:"authTicket"`
	Invitations   *string           `json:"invitations"`
}

type User struct {
	ID                      uuid.UUID   `json:"id"`
	FirstName               string      `json:"firstName"`
	LastName                string      `json:"lastName"`
	DateOfBirth             int         `json:"dateOfBirth"`
	Email                   string      `json:"email"`
	Country                 string      `json:"country"`
	UILanguage              string      `json:"uiLanguage"`
	CommunicationLanguage   string      `json:"communicationLanguage"`
	AccountType             string      `json:"accountType"`
	UOM                     string      `json:"uom"`
	DateFormat              string      `json:"dateFormat"`
	TimeFormat              string      `json:"timeFormat"`
	EmailDay                []int       `json:"emailDay"`
	System                  System      `json:"system"`
	Details                 interface{} `json:"details"`
	TwoFactor               TwoFactor   `json:"twoFactor"`
	Created                 int64       `json:"created"`
	PrivatePracticeName     string      `json:"privatePracticeName"`
	PrivatePracticeAddress1 string      `json:"privatePracticeAddress1"`
	PrivatePracticeCity     string      `json:"privatePracticeCity"`
	PrivatePracticeState    string      `json:"privatePracticeState"`
	PrivatePracticeZip      string      `json:"privatePracticeZip"`
	PrivatePracticePhone    string      `json:"privatePracticePhone"`
}

type System struct {
	Messages SystemMessages `json:"messages"`
}

type SystemMessages struct {
	FirstUsePhoenix                  int64  `json:"firstUsePhoenix"`
	FirstUsePhoenixDashboard         int64  `json:"firstUsePhoenix-dashboard"`
	FirstUsePhoenixPatientProfile    int64  `json:"firstUsePhoenix-patientProfile"`
	FirstUsePhoenixReportsDataMerged int64  `json:"firstUsePhoenixReportsDataMerged"`
	LVWebPostRelease                 string `json:"lvWebPostRelease"`
	WebRegister                      int64  `json:"webRegister"`
	WelcomeScreen                    int64  `json:"welcomeScreen"`
}

type TwoFactor struct {
	Code            string `json:"code"`
	Fingerprint     string `json:"fingerprint"`
	PhoneNumber     string `json:"phoneNumber"`
	IsPrimaryMethod bool   `json:"isPrimaryMethod"`
	PrimaryMethod   string `json:"primaryMethod"`
	PrimaryValue    string `json:"primaryValue"`
	SecondaryMethod string `json:"secondaryMethod"`
	SecondaryValue  string `json:"secondaryValue"`
}

type UserNotifications struct {
	Unresolved int `json:"unresolved"`
}
