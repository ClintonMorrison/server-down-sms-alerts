package main

import "flag"

func main() {
	var configFilePath string
	flag.StringVar(
		&configFilePath,
		"svar",
		"./config.json",
		"path to a config file (see config.sample.json for an example)")
	flag.Parse()

	config := ParseConfig(configFilePath)
	for _, site := range config.Sites {
		go MonitorSite(config.Twilio, site)
	}

	select {}
}
