package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Auth struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Fingerprint string `json:"fingerprint"`
}

var (
	agent  = "go-libreview/0.1.0a"
	client = &http.Client{}
	token  string

	auth Auth = Auth{
		Email:       "",
		Password:    "",
		Fingerprint: "",
	}

	root = "https://api-us.libreview.io"

	patientId = ""

	primaryDevice PrimaryDevice
)

func main() {
	log.Println("[INFO] Logging in to LibreView")
	token = getToken()
	primaryDevice = getPrimaryDevice()
}

func getPrimaryDevice() PrimaryDevice {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/patients/%s/reportSettings", root, patientId), nil)
	bearer := "Bearer " + token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("User-Agent", agent)
	req.Header.Add("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	b, _ := io.ReadAll(response.Body)
	reportSettingsResponse := ReportSettingsResponse{}
	json.Unmarshal(b, &reportSettingsResponse)
	if err != nil {
		panic(err)
	}

	primaryDevice := PrimaryDevice{}

	for _, v := range reportSettingsResponse.Data.DataSources {
		primaryDevice.ID = v.ID
		primaryDevice.Type = v.Type
		primaryDevice.FirmwareVersion = v.FirmwareVersion
	}

	log.Printf("[INFO] found primary device: %v", primaryDevice)
	return primaryDevice
}

func getPatient() {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/patients/%s", root, patientId), nil)
	bearer := "Bearer " + token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("User-Agent", agent)
	req.Header.Add("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	b, _ := io.ReadAll(response.Body)

	patientResponse := PatientResp{}
	err = json.Unmarshal(b, &patientResponse)
	if err != nil {
		panic(err)
	}

	patient := patientResponse.Data.Patient
	log.Printf("[INFO] found patient: %s %s [%v]", patient.FirstName, patient.LastName, patient.ID)

}

func getUser() {

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/user", root), nil)
	bearer := "Bearer " + token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("User-Agent", agent)
	req.Header.Add("Accept", "application/json")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	b, _ := io.ReadAll(response.Body)
	userResp := UserResp{}
	err = json.Unmarshal(b, &userResp)
	if err != nil {
		panic(err)
	}
	user := userResp.Data.User
	log.Printf("[INFO] found user: %s %s [%v]", user.FirstName, user.LastName, user.ID)

}

func getToken() string {
	authJ, err := json.Marshal(auth)
	if err != nil {
		panic(err)
	}

	authJBytes := bytes.NewBuffer(authJ)
	authURL := fmt.Sprintf("%s/auth/login", root)

	response, err := http.Post(authURL, "application/json", authJBytes)
	if err != nil {
		panic(err)
	}
	respBody, _ := io.ReadAll(response.Body)

	//TO DO map out login response
	body := make(map[string]map[string]map[string]string)
	json.Unmarshal(respBody, &body)
	token := body["data"]["authTicket"]["token"]
	log.Printf("[INFO] recieved token: %s...", token[:10])
	return token

}
