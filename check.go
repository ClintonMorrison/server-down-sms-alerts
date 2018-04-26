package main

import (
	"fmt"
	"net/http"
	"time"
)

func requestAndGetStatus(url string) bool {
	resp, err := http.Head(url)

	if err != nil {
		return false
	}

	resp.Body.Close()

	return resp.StatusCode == 200
}

func statusToText(isUp bool) string {
	if isUp {
		return "okay"
	}

	return "fail"
}

func formatTime(t time.Time) string {
	return t.Format("3:04:05 PM")
}

func logStatus(site SiteConfig, attempts int, isUp bool) {
	fmt.Printf(
		"[%s] %s (%d of %d): %s\n",
		formatTime(time.Now()),
		site.Url,
		attempts,
		site.RetryAttempts+1,
		statusToText(isUp))
}

func checkIfUp(site SiteConfig) bool {
	attempts := 0
	isUp := false

	for !isUp && attempts <= site.RetryAttempts {
		attempts++
		isUp = requestAndGetStatus(site.Url)
		logStatus(site, attempts, isUp)

		if !isUp && attempts <= site.RetryAttempts {
			time.Sleep(time.Duration(500 * time.Millisecond))
		}
	}

	return isUp
}

func downLongEnoughForAlert(downSince time.Time, site SiteConfig) bool {
	downTimeBeforeAlert := ParseTime(site.DownTimeBeforeAlert)
	return downSince.Add(downTimeBeforeAlert).Before(time.Now())
}

func canSendAlert(lastAlert time.Time, site SiteConfig) bool {
	maximumAlertInterval := ParseTime(site.MaximumAlertInterval)
	return lastAlert.Add(maximumAlertInterval).Before(time.Now())
}

func MonitorSite(twilio TwilioConfig, site SiteConfig) {
	waitTime := ParseTime(site.CheckInterval)

	downSinceTime := time.Time{}
	lastAlertTime := time.Time{}

	for {
		isUp := checkIfUp(site)

		if isUp {
			downSinceTime = time.Time{}
		} else {
			if downSinceTime.IsZero() {
				downSinceTime = time.Now()
			}

			if downLongEnoughForAlert(downSinceTime, site) {
				if canSendAlert(lastAlertTime, site) {
					sentSuccessfully := SendAlert(twilio, site, time.Now())
					if sentSuccessfully {
						lastAlertTime = time.Now()
					}
				}
			}
		}

		time.Sleep(waitTime)
	}
}
