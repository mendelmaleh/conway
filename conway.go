package conway

import (
	"time"
)

// mod that always returns positive ints
func mod(a, b int) int {
	return (a%b + b) % b
}

func date(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func leap(year int) bool {
	return year%4 == 0 && year%400 != 0
}

// Doomsday calculates the weekday for a given date
func Doomsday(date time.Time) time.Weekday {
	// m := date.Month()

	// this month's doomsday occurance
	doomsday := DoomsdayMonth(date)

	// find distance from doomsday
	days := int(date.Sub(doomsday).Hours()) / 24 // TODO: can this break?

	// year doomsday
	weekday := DoomsdayYear(date.Year())

	return time.Weekday((int(weekday) + days) % 7)
}

// DoomsdayMonth returns the date of the month's doomsday occurance
func DoomsdayMonth(t time.Time) time.Time {
	switch m := t.Month(); {
	// the last of feb or jan will do (except in leap years it's jan 32)
	case m == time.January && leap(t.Year()):
		return date(t.Year(), t.Month()+1, 1)
	case m == time.January || m == time.February:
		return date(t.Year(), t.Month()+1, 0)
	}

	switch m := int(t.Month()); {
	// then for even months take the months's own day,
	case m%2 == 0:
		return date(t.Year(), t.Month(), m)
	// and for odd months, add four or take it away.
	case m <= 7:
		return date(t.Year(), t.Month(), m+4)
	// according to lenth - or simply remember
	// you only subtract for september or november.
	default:
		return date(t.Year(), t.Month(), m-4)
	}
	// americans: 9 to 5 job in a 7/11 store
}

// DoomsdayYear return the weekday of the years doomsday
func DoomsdayYear(year int) time.Weekday {
	// TODO: in julian times, lack-a-day, lack-a-day.
	// zero was sunday, centuries fell back a day.

	// but gregorian four-hundredths are always tuesday
	day := 2
	year %= 400

	// and now centuries take us back twos.
	day -= year / 100 * 2
	year %= 100

	// now to work out your doomsday the orthodox way,
	// three things you should add to the century day:
	// dozens, remainder, and four in the latter
	// (if you alter by sevens, of course it won't matter).
	day += year / 12
	rem := year % 12
	day += rem + rem/4

	return time.Weekday(mod(day, 7))
}
