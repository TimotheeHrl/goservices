package userbuilder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// hello copilot , are you here ?

func GetUser() User {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", "https://randomuser.me/api", nil)
	response, err := client.Do(req)
	respContent := response.Body
	body, err := ioutil.ReadAll(respContent)
	stringBody := string(body)
	var resCompain User
	json.Unmarshal([]byte(stringBody), &resCompain)

	fmt.Println(resCompain)
	fmt.Println(stringBody)
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	return resCompain
}
