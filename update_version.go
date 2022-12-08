package automatedupdate

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

var httpClient = http.DefaultClient

type Application struct {
	ApplicationId string `json:"applicationId"`
	Version       string `json:"version"`
}

type Profile struct {
	Applications []Application `json:"applications"`
}

func UpdateFromCSV(pathToCSV string, profile Profile) {
	macAddresses := GetMacFromCSV(pathToCSV)
	for _, macAddress := range macAddresses {
		UpdateVersion(macAddress, profile)
	}
}

func GetMacFromCSV(pathToCSV string) []string {
	f, err := os.Open(pathToCSV)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	fileContent, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var macAddresses []string
	for i, line := range fileContent {
		// skip header
		if i != 0 {
			macAddresses = append(macAddresses, line[0])
		}
	}
	return macAddresses
}

func UpdateVersion(macAddress string, profile Profile) {
	jsonProfile, _ := json.Marshal(profile)
	req, err := http.NewRequest("PUT", config.API_BASE_URL+"/"+macAddress, bytes.NewBuffer(jsonProfile))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+config.API_TOKEN)
	res, err := httpClient.Do(req)
	if err != nil {
		print(err)
	} else if res.StatusCode != 204 {
		print("Error" + strconv.Itoa(res.StatusCode))
	}
	res.Body.Close()
}

func SetHttpClient(client *http.Client) {
	httpClient = client
}
