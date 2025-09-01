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

func (month HebrewMonth) Partner() time.Month {
	return time.Month((month.Number()-1)%12 + 1)
}

type HebrewDate struct {
	Year  int
	Month HebrewMonth
	Day   int
}

func (date *HebrewDate) Height() int {
	romanyear := date.Year - 3760 // TODO: 3761?
	he, _, _ := HeSheIt(romanyear - 1)
	_, she, it := HeSheIt(romanyear)

	switch date.Month {
	case Tishrei, Cheshvan:
		return date.Day + he
	case Kislev:
		return date.Day + max(he, she)
	case Tevet, Shevat, AdarI, AdarII:
		return date.Day + she
	default:
		return date.Day + it
	}
}

func (date *HebrewDate) Roman() time.Time {
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
