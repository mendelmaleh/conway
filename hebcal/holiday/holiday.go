//go:generate stringer -type=EventStart

package holiday

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/mendelmaleh/conway/hebcal"
	"github.com/mendelmaleh/conway/utils"
)

type Event struct {
	Name   string
	Date   hebcal.HebrewDate
	Length int
	Type   EventType
	Start  EventStart
}

type EventType uint

const (
	Holiday EventType = 1 << iota
	Fast
	Major
)

func (t EventType) String() string {
	var b []string

	if t.Is(Major) {
		b = append(b, "major")
	}

	if t.Is(Holiday) {
		b = append(b, "holiday")
	}

	if t.Is(Fast) {
		b = append(b, "fast")
	}

	if len(b) == 0 {
		return fmt.Sprintf("Unknown EventType(%d)", t)
	}

	r := []rune(strings.Join(b, " "))
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func (t EventType) Is(flag EventType) bool {
	return t&flag == flag
}

type EventStart uint

const (
	Sunset    EventStart = iota // default
	Nightfall                   // only for second days
	Dawn
)

func (e Event) Fill(year int, diaspora bool) []Event {
	day := e
	day.Date.Year = year
	if day.Length == 0 {
		day.Length = 1 // fix lazy default
	}

	// extend diaspora
	if diaspora && (e == SheminiAtzeret || e == Pesach || e == Shavuot) {
		day.Length += 1
	}

	if day.Length == 1 {
		return []Event{day} // early return
	}

	// rosh hashana is two major days everywhere
	// sukkot and pesach are multiple days and can be extended
	// chanuka is multiple days and isn't extended
	// shemini atzeret and shavuot can be extended
	//
	// sukkot isn't actually extended over shemini atzeret
	// pesach last days are major
	// yom kippur is major but and isn't extended

	var days []Event
	for i := 0; i < day.Length; i++ {
		day := day

		if e == Sukkot && (i > 0 && !diaspora) || (i > 1 && diaspora) {
			day.Type &= ^Major
		}

		if e == Pesach && i > (day.Length-e.Length) && i < 6 {
			day.Type &= ^Major
		}

		// delay start to nightfall for major holidays following another major or saturday
		if day.Type.Is(Major|Holiday) && (i > 0 && days[i-1].Type.Is(Major)) ||
			day.Date.Roman().Weekday() == time.Sunday {
			day.Start = Nightfall
		}

		day.Date.Day += i
		day.Name = fmt.Sprintf("%s %s", e.Name, utils.RomanNumeral(i+1))

		days = append(days, day)
	}

	return days
}
