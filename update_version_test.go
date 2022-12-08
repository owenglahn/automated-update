package automatedupdate

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var profile Profile

func setup() {
	file, _ := ioutil.ReadFile("update.json")
	json.Unmarshal([]byte(file), &profile)
}

func TestUpdateFromCSV(t *testing.T) {
	UpdateFromCSV("test.csv", profile)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
