package hebcal

import (
	"time"

	"github.com/mendelmaleh/conway/utils"
)

func FromRoman(date time.Time) HebrewDate {
	y, m, d := date.Date()
	year := y + 3760

	before := date.Before(NewYear(y))
	if !before {
		year++
	}

	f := func(y int, m time.Month, d int, before bool) (HebrewMonth, int) {
		month := MonthPartner(m, before)
		day := RomanHeight(m, d) - month.Height(HeSheIt(y))
		return month, day
	}

	month, day := f(y, m, d, before)

	// stretch date to previous month if height is smaller than hesheit
	for day < 1 {
		m -= 1
		d += MonthDays(y, m)

		// new month and day
		month, day = f(y, m, d, before)
	}

	// TODO: extend into next month if day > month length

	return HebrewDate{year, month, day}

}

func MonthDays(year int, month time.Month) int {
	return utils.Date(year, month+1, 0).Day()
}

func MonthPartner(month time.Month, beforeNewYear bool) HebrewMonth {
	switch {
	case month <= time.February:
		return HebrewMonth(month - 2 + 13) // handle rollover
	case month < time.August:
		return HebrewMonth(month - 2)
	case month > time.August:
		return HebrewMonth(month - 1)
	case month == time.August && beforeNewYear:
		return Elul
	default:
		return Tishrei

	}
}
