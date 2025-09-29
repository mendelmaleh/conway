package hebcal

import (
	"time"

	"github.com/mendelmaleh/conway/utils"
)

type HebrewMonth int

const (
	Nissan HebrewMonth = 1 + iota // TODO: +2?
	Iyar
	Sivan
	Tamuz
	Av
	Elul
	Tishrei
	Cheshvan
	Kislev
	Tevet
	Shevat
	AdarI
	AdarII
	// TODO: regular Adar?
)

func (month HebrewMonth) String() string {
	names := [...]string{
		"Nissan",
		"Iyar",
		"Sivan",
		"Tammuz",
		"Av",
		"Elul",
		"Tishrei",
		"Cheshvan",
		"Kislev",
		"Tevet",
		"Shevat",
		"Adar I",
		"Adar II",
	}

	return names[int(month)-1]
}

// Number of roman partner month for hebrew month
func (month HebrewMonth) Number() int {
	switch m := int(month); {
	// nissan to elul is march to august
	case month <= Elul:
		return m + 2
	// tishrei to adar ii is august to february
	default:
		return m + 1
	}
}

// Partner roman month for hebrew month
func (month HebrewMonth) Partner() time.Month {
	return time.Month((month.Number()-1)%12 + 1)
}

type HebrewDate struct {
	Year  int
	Month HebrewMonth
	Day   int
}

// Height of hebrew month
func (month HebrewMonth) Height(he, she, it int) int {
	switch month {
	case Tishrei, Cheshvan:
		return he
	case Kislev:
		return max(he, she)
	case Tevet, Shevat, AdarI, AdarII:
		return she
	default:
		return it
	}
}

// Height of hebrew date
func (date *HebrewDate) Height() int {
	return date.Month.Height(HebrewHeSheIt(date.Year)) + date.Day
}

// Roman conversion from hebrew date
func (date HebrewDate) Roman() time.Time {
	// fix adar ii in non-leap year
	if date.Month == AdarII && Finder(date.Year-3760) != LeapPrev {
		date.Month = AdarI
	}

	height := date.Height()
	number := date.Month.Number()
	partner := date.Month.Partner()

	romanyear := date.Year - 3761
	// source: github.com/bendory/conway-hebrew-calendar
	if date.Month <= Elul || date.Month > Shevat || partner == time.January {
		romanyear++
	}

	return utils.Date(romanyear, partner, height-number)
}
