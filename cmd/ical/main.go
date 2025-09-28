package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/emersion/go-ical"
	"github.com/google/uuid"
	"github.com/mendelmaleh/conway/hebcal"
	"github.com/mendelmaleh/conway/hebcal/holiday"
	"github.com/mendelmaleh/conway/utils"
)

func main() {
	year := flag.Int("year", hebcal.FromRoman(time.Now()).Year, "hebrew year")
	diaspora := flag.Bool("diaspora", false, "use diaspora holidays")
	test := flag.Bool("test", false, "testing mode")
	flag.Parse()

	cal := ical.NewCalendar()
	cal.Props.SetText(ical.PropProductID, "-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN")
	cal.Props.SetText(ical.PropVersion, "2.0")
	cal.Props.SetText("X-WR-CALNAME", fmt.Sprintf("%d Hebrew calendar", year))
	cal.Props.SetText("X-APPLE-CALENDAR-COLOR", "#FFCC00") // yellow

	for _, v := range holiday.All {
		for _, d := range v.Fill(*year, *diaspora) {
			if *test {
				fmt.Println(
					d.Date.Roman().Format(utils.ISO8601+" Mon:"),
					d.Name,
					d.Type.Is(holiday.Major),
					d.Start,
				)
				continue
			}

			e := ical.NewEvent()
			id, err := uuid.NewRandom()
			if err != nil {
				log.Fatal(err)
			}
			e.Props.SetText(ical.PropUID, id.String())
			e.Props.SetDateTime(ical.PropDateTimeStamp, time.Now())

			e.Props.SetDate(ical.PropDateTimeStart, d.Date.Roman())
			e.Props.SetText(ical.PropSummary, d.Name)
			e.Props.SetText(ical.PropDescription, fmt.Sprintf("%s, starts at %s", d.Type, d.Start))
			cal.Children = append(cal.Children, e.Component)
		}
	}

	enc := ical.NewEncoder(os.Stdout)
	if err := enc.Encode(cal); err != nil {
		log.Fatal(err)
	}
}
