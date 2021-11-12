package date_utils

import "time"

const (
	isoDateFormat = "2006-01-02T15:04:05-0700"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString(format string) string {
	return GetNow().Format(format);
}

func GetIsoString() string {
	return GetNowString(isoDateFormat)
}

