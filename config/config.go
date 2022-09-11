package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	Token  string
	Prefix string
	config *configStruct
)

type configStruct struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func ReadConfig() error {
	file, err := os.ReadFile("./config.json") //we read the config file
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = json.Unmarshal(file, &config) // Unmarshal is the JSON.stringify of JavaScript
	if err != nil {
		log.Fatal(err)
		return err
	}

	Token = config.Token
	Prefix = config.Prefix

	return nil
}