package cmd

import "time"


func GetTimeZone(zone string, flag string) (string, error) {
	loc, err := time.LoadLocation(zone)
    if err != nil {
		return "", err
	}

	timeNow := time.Now().In(loc)
	if flag != "" {
		return timeNow.Format(flag), nil
	} else {
		return timeNow.Format(time.RFC1123), nil
	}
}