package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type chuckNorrisFact struct {
	Value string `json:"value"`
}

/*
	Fun commands
*/

// Hello returns hello with the mention of the sender
func Hello(senderUsername string) string {
	return fmt.Sprintf("Hello <@%s> :)", senderUsername)
}

// ChuckNorrisFact returns a random string from a chuck norris API
func ChuckNorrisFact() string {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		log.Printf("An error occured while GETting chuck fact : %v", err)
		return "An error occured, please ask the creator to investigate..."
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var fact chuckNorrisFact
	err = json.Unmarshal(body, &fact)

	if err != nil {
		log.Printf("An error occured while unmarshalling chuck fact : %v", err)
		return "An error occured, please ask the creator to investigate..."
	}

	return fact.Value
}
