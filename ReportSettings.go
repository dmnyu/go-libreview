package main

type ReportSettingsResponse struct {
	Status int                `json:"status"`
	Data   ReportSettingsData `json:"data"`
	Ticket AuthTicket         `json:"ticket"`
}

type ReportSettingsData struct {
	EndDate     int64       `json:"endDate"`
	Settings    Settings    `json:"settings"`
	DataSources DataSources `json:"dataSources"`
}

type Settings struct {
	LowThreshold  int         `json:"low_threshold"`
	HighThreshold int         `json:"high_threshold"`
	TargetRange   TargetRange `json:"targetRange"`
}
