package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func sendText(config TwilioConfig, to string, msg string) bool {
	base := "https://api.twilio.com/2010-04-01/Accounts"
	apiUrl := fmt.Sprintf("%s/%s/Messages", base, config.Sid)

	form := url.Values{}
	form.Set("To", to)
	form.Set("From", config.PhoneNumber)
	form.Set("Body", msg)

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Could not create alert request")
		return false
	}

	req.SetBasicAuth(config.Sid, config.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Twilio request failed")
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read Twilio request body")
		return false
	}

	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)

	return resp.StatusCode == 201
}

func SendAlert(twilio TwilioConfig, site SiteConfig, time time.Time) bool {
	msg := fmt.Sprintf("\n\nSite is down: %s\n\nCurrent time: %s", site.Url, time.Format("3:04:05 PM"))
	fmt.Println(">> Sending text to", site.AlertNumber)
	return sendText(twilio, site.AlertNumber, msg)
}
