# server-down-sms-alerts
This is a very basic golang application which can monitor one or more servers. It periodically requests URLs and logs when requests fail.

If multiple requests fail it sends an SMS message with Twilio.

Configuration
--
The Twilio credentials and information about the sites 
you want to monitor is provided with a JSON config file.

Below is an example:
```
{
  "twilio": {
    "phoneNumber": "+11234567890",
    "sid": "...",
    "token": "..."
  },
  "sites": [
    {
      "url": "https://mysite.com",
      "checkInterval": "10 seconds",
      "retryAttempts": 2,
      "downTimeBeforeAlert": "5 minutes",
      "maximumAlertInterval": "12 hours",
      "alertNumber": "+12223334444"
    }
  ]
}
```

The Twilio fields are:
- `phoneNumber`: the phone number to send texts from (in your Twilio account)
- `sid`: your Twilio account sid
- `token`: your Twilio account secret token

The site fields are:
- `url`: the url to request for the health check (only a 200 response is considered success)
- `checkInterval`: how often to try requesting the URL
- `retryAttempts`: how many times to retry the request, if the first attempt fails
- `downTimeBeforeAlert`: how long requests can fail before an alert is sent
- `maximumAlertInterval`: the minimum amount of time to wait before sending another alert
- `alertNumber`: the number to send SMS alerts to

For the `checkInternval`, `retryAttempts`, and `maximumAlertInterval` settings, these are examples of acceptable time formats:
```
1 second, 10 seconds, 1 minute, 7 minutes, 12 hours, 2 days
```


Running
---
You can build the application by running `go build` in the `server-down-sms-alerts` directory.


You can start the application with:
```
./server-down-sms-alerts --config="./config.json"
```
