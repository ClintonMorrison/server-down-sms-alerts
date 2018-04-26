package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type TwilioConfig struct {
	PhoneNumber string `json:"phoneNumber"`
	Sid         string `json:"sid"`
	Token       string `json:"token"`
}

type SiteConfig struct {
	Url                  string `json:"Url"`
	CheckInterval        string `json:"checkInterval"`
	RetryAttempts        int    `json:"retryAttempts"`
	DownTimeBeforeAlert  string `json:"downTimeBeforeAlert"`
	MaximumAlertInterval string `json:"maximumAlertInterval"`
	AlertNumber          string `json:"alertNumber"`
}

type Config struct {
	V      string       `json:"v"`
	Twilio TwilioConfig `json:"twilio"`
	Sites  []SiteConfig `json:"sites"`
}

func ParseConfig(path string) *Config {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Could not read config file: '%s'", path))
	}

	config := new(Config)
	err = json.Unmarshal(contents, config)

	fmt.Println(config.V)

	if err != nil {
		panic(fmt.Sprintf("Error parsing config file: '%s'", path))
	}

	return config
}

func unitToDuration(unit string) time.Duration {
	switch unit {
	case "second", "seconds":
		return time.Second
	case "minute", "minutes":
		return time.Minute
	case "hour", "hours":
		return time.Hour
	case "day", "days":
		return time.Duration(time.Hour * 24)
	}

	panic("invalid date unit: " + unit)
}

func ParseTime(s string) time.Duration {
	fields := strings.Fields(s)

	if len(fields) != 2 {
		panic("invalid time string: " + s)
	}

	quantity, err := strconv.Atoi(fields[0])
	if err != nil {
		panic("invalid time string: " + s)
	}

	unit := unitToDuration(fields[1])

	return time.Duration(quantity) * unit
}
