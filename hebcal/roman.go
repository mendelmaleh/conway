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

	month := MonthPartner(m, before)
	hesheit := month.Height(HeSheIt(y))
	height := RomanHeight(m, d)

	// stretch height out if smaller than hesheit
	for hesheit >= height {
		m -= 1
		d += MonthDays(y, m)

		// new height, partner, hesheit
		month = MonthPartner(m, before)
		hesheit = month.Height(HeSheIt(y))
		height = RomanHeight(m, d)
	}

	day := height - hesheit
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
