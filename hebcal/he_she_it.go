package hebcal

import (
	"time"

	"github.com/mendelmaleh/conway/utils"
)

// HeSheIt from roman year
func HeSheIt(year int) (int, int, int) {
	// if rosh hashanah falls on the nth day of september, then IT = n + 9
	newyear := NewYear(year)
	start := utils.Date(newyear.Year(), time.September, 0)
	it := 9 + int(newyear.Sub(start).Hours())/24

	// to move from IT to HE, we ELevate IT by 29
	he := it + 29

	// and to get to SHE, EXtend IT by adding EX
	ex := 10
	if utils.Leap(year) {
		ex += 1
	}
	if Finder(year) != LeapPrev {
		ex += 30
	}

	she := it + ex

	return he, she, it
}

// RomanHeight returns the height of a roman date
func RomanHeight(date time.Time) int {
	switch date.Month() {
	case time.January, time.February:
		return int(date.Month()) + 12 + date.Day()
	default:
		return int(date.Month()) + date.Day()
	}
}
