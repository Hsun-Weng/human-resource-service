package util

import "time"

const dateLayout = "2006-01-02"
const dateTimeLayout = "2006-01-02T15:04:05Z"

func ParseDate(value string) (time.Time, error) {
	return time.Parse(dateLayout, value)
}

func FormatDate(date time.Time) string {
	return date.Format(dateLayout)
}

func FormatDateTime(time time.Time) string {
	return time.Format(dateTimeLayout)
}
