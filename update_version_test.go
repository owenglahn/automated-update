package automatedupdate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

var profile Profile
var httpServer *httptest.Server

func setup() {

	Configure()
}

func TestUpdateFromCSV(t *testing.T) {
	file, _ := ioutil.ReadFile("update.json")
	var receivedMacAddresses []string
	httpServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedMacAddress := strings.Split(r.URL.Path, "clientId:")[1]
		receivedMacAddresses = append(receivedMacAddresses, receivedMacAddress)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write(file)
	}))
	client := Client{httpServer.URL}
	client.UpdateFromCSV("test.csv", profile)

	var expectedMacAddresses []string = GetMacFromCSV("test.csv")
	if reflect.DeepEqual(expectedMacAddresses, receivedMacAddresses) {
		t.Logf("UpdateFromCSV PASSED. Mac Addressses sent correctly. Expected: \n%v\n And Got:\n%v",
			expectedMacAddresses, receivedMacAddresses)
	} else {
		t.Errorf("UpdateFromCSV FAILED. Mac Addresses not correct\nExpected:\n %v \n But Got: %v",
			expectedMacAddresses, receivedMacAddresses)
	}
}

func TestUpdateVersion(t *testing.T) {
	file, _ := ioutil.ReadFile("update.json")
	var requestBody Profile
	var url string
	var auth string
	httpServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url = r.URL.Path
		auth = r.Header.Get("Authorization")
		json.NewDecoder(r.Body).Decode(&requestBody)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write(file)
	}))
	var inputBody Profile
	client := Client{httpServer.URL}
	json.Unmarshal(file, &inputBody)
	client.UpdateVersion("testAddress", inputBody)

	if url == "/clientId:testAddress" {
		t.Log("1 / 3: URLs match")
	} else {
		t.Error("UpdateVersion FAILED. URLs do not match")
	}

	if auth == "Basic "+config.API_TOKEN {
		t.Log(" 2 / 3: Auth matches")
	} else {
		t.Errorf("UpdateVersion FAILED, invalid authentication.")
	}

	if reflect.DeepEqual(inputBody, requestBody) {
		t.Log("UpdateVersion PASSED. Input request and request received on the server are the same")
	} else {
		t.Error("UpdateVersion FAILED. Requests do not match")
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	httpServer.Close()
	os.Exit(code)
}
