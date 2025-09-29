//go:generate stringer -type=Leap

package hebcal

import (
	"time"

	"github.com/mendelmaleh/conway/doomsday"
	"github.com/mendelmaleh/conway/utils"
)

// adjustment for calendar length difference (hebrew 6939.69 vs gregorian 6939.60 days)
func bc(year int) int {
	// bc == 6 for 1900-2099
	by := year/100 - 11
	return by/4*3 + max(by%4-1, 0)
}

// NewYear calculates the roman date of rosh hashana
func NewYear(year int) time.Time {
	// find the roman year's place in the 19-year golden cycle
	golden := (year % 19) + 1

	// this tells us if the surrounding years are leap
	finder := float64(12 * golden % 19)

	y := float64(year - 1900)

	// newyear = sep A+B: c-d-E
	var a, b, c, d, e float64
	a = 1.5 * finder                          // acrobatic
	b = float64(bc(year)) + float64(year%4)/4 // bissextile
	c = finder + 1
	d = (2*y - 1) / 35
	e = (finder + 1) / 760 // can be ignored for 1762-2168

	day := a + b + (c-d)/18 - e
	// fmt.Println(golden, finder, y)
	// fmt.Println(a, b, c, d, e)
	// fmt.Println(day)

	// truncate, don't round! per david.slusky@ku.edu via email -- github.com/bendory/conway-hebrew-calendar
	date := utils.Date(year, time.September, int(day))

	// first postponement doesn't apply to our system
	weekday := doomsday.Doomsday(date)
	switch weekday {
	// second postponement: wed and fri for yom kippur, sun for hoshana rabba
	case time.Sunday, time.Wednesday, time.Friday:
		return date.AddDate(0, 0, 1)
	}

	leap := Finder(year)
	frac := day - float64(int(day))
	switch {
	// third postponement: avoid 356 day year by postponing start
	case weekday == time.Tuesday && frac > 0.633 && leap != LeapNext:
		return date.AddDate(0, 0, 2)
	// fourth postponement: avoid 382 day year by postponing start
	case weekday == time.Monday && frac > 0.898 && leap == LeapPrev:
		return date.AddDate(0, 0, 1)
	}

	return date
}

type Leap int

const (
	LeapNext Leap = iota
	LeapNone
	LeapPrev
)

func Finder(year int) Leap {
	// TODO: dedup
	// find the roman year's place in the 19-year golden cycle
	golden := (year % 19) + 1

	// this tells us if the surrounding years are leap
	finder := float64(12 * golden % 19)

	switch {
	case finder < 7:
		return LeapNext
	case finder < 12:
		return LeapNone
	default:
		return LeapPrev
	}
}
