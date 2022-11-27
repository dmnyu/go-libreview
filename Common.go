package main

type Messages struct {
	Unread          int    `json:"unread,omitempty"`
	EmailPreference string `json:"emailPreference,omitempty"`
}

type AuthTicket struct {
	Token    string `json:"token"`
	Expires  int64  `json:"expires"`
	Duration int64  `json:"duration"`
}
