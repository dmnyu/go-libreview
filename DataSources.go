package main

import "github.com/google/uuid"

type DataSources map[uuid.UUID]DataSource

type DataSource struct {
	ID                uuid.UUID `json:"id"`
	Nickname          string    `json:"nickname"`
	SN                string    `json:"sn"`
	Type              int       `json:"type"`
	UploadDate        int64     `json:"uploadDate"`
	DaysData          []int     `json:"daysData"`
	TimeOffsetMinutes int       `json:"timeOffsetMinutes"`
	FirmwareVersion   string    `json:"firmwareVersion"`
}

type PrimaryDevice struct {
	ID              uuid.UUID `json:"id"`
	Type            int       `json:"type"`
	FirmwareVersion string    `json:"firmwareVersion"`
}
