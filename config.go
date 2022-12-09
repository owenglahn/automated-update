package automatedupdate

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	API_BASE_URL string `yaml:"API_BASE_URL"`
	API_TOKEN    string `yaml:"API_TOKEN"`
}

var config Config

func Configure() {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	yamlDecoder := yaml.NewDecoder(f)
	err = yamlDecoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
}
