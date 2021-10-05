package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	// Camera admin login
	username = "admin"
	password = "instar"
	// API endpoint to activate the PIR sensor
	cameraAPI = "http://192.168.2.77:8090/param.cgi?cmd=setpirattr&-pir_enable=1"
)

func main() {
	postCommand(cameraAPI, "POST")
}

func postCommand(url, method string) error {
	// Send request and catch error
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	// Authenticate and catch error
	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	return nil
}