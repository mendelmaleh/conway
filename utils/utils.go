package utils

import "time"

const ISO8601 = "2006-01-02"

// Mod that always returns positive ints
func Mod(a, b int) int {
	return (a%b + b) % b
}

// Date returns a time.Time struct with zero time and UTC location
func Date(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// Leap checks if the year is leap
func Leap(year int) bool {
	return year%4 == 0 && year%400 != 0
}
