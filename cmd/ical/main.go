package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/emersion/go-ical"
	"github.com/google/uuid"
	"github.com/mendelmaleh/conway/hebcal/holiday"
)

func main() {
	year := 5786
	diaspora := true

	cal := ical.NewCalendar()
	cal.Props.SetText(ical.PropProductID, "-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN")
	cal.Props.SetText(ical.PropVersion, "2.0")
	cal.Props.SetText("X-WR-CALNAME", fmt.Sprintf("%d Hebrew calendar", year))
	cal.Props.SetText("X-APPLE-CALENDAR-COLOR", "#FFCC00") // yellow

	for _, v := range holiday.All {
		for _, d := range v.Fill(year, diaspora) {
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
