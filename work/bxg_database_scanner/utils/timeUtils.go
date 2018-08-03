package utils

import "time"

const (
	DateTimePattern = "2006-01-02 15:04:05"
)
func FormatDateTime(t *time.Time) string {
	return t.Format(DateTimePattern)
}
