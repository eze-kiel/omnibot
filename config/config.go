package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// Token variable
	Token string

	config *configStruct
)

type configStruct struct {
	Token string `json:"Token"`
}

// ReadConfig reads the config.json file
func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token

	return nil
}
