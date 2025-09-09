package doomsday_test

import (
	"testing"
	"time"

	"github.com/mendelmaleh/conway/doomsday"
	"github.com/mendelmaleh/conway/utils"
)

var date = utils.Date

func fmtdate(t time.Time) string {
	return t.Format(utils.ISO8601)
}

func TestYear(t *testing.T) {
	tests := []struct {
		year int
		want time.Weekday
		name string // description of this test case
	}{
		{2000, time.Tuesday, "gregorian four-hundredths"},
		{2100, time.Sunday, "extra centuries"},
		{2200, time.Friday, "extra centuries rollover"},
		{2016, time.Monday, "paper example (text)"},
		{2019, time.Thursday, "paper example (num)"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if res := doomsday.Year(tc.year); res != tc.want {
				t.Errorf("DoomsdayYear() = %v, want %v", res, tc.want)
			}
		})
	}
}

func TestMonth(t *testing.T) {
	tests := []struct {
		t    time.Time
		want time.Time
		name string // description of this test case
	}{
		{
			date(2001, time.January, 22),
			date(2001, time.January, 31),
			"January non leap",
		},
		{
			date(2004, time.January, 22),
			date(2004, time.February, 1),
			"January leap",
		},
		{
			date(2007, time.February, 22),
			date(2007, time.February, 28),
			"February non leap",
		},
		{
			date(2008, time.February, 22),
			date(2008, time.February, 29),
			"February leap",
		},
		{
			date(2019, time.September, 19),
			date(2019, time.September, 5),
			time.September.String(),
		},
		{
			date(2016, time.October, 2),
			date(2016, time.October, 10),
			time.October.String(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if res := doomsday.Month(tc.t); res != tc.want {
				t.Errorf("DoomsdayMonth() = %v, want %v", fmtdate(res), fmtdate(tc.want))
			}
		})
	}
}

func TestDoomsday(t *testing.T) {
	tests := []struct {
		date time.Time
		want time.Weekday
		name string // description of this test case
	}{
		{
			date(2025, 8, 22),
			time.Friday,
			"today",
		},
		{
			date(2019, 9, 29),
			time.Sunday,
			"paper example (text)",
		},
		{
			date(2027, 10, 1),
			time.Friday,
			"failure from 2027 newyear test (-2)",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if res := doomsday.Doomsday(tc.date); res != tc.want {
				t.Errorf("Doomsday() = %v, want %v", res, tc.want)
			}
		})
	}
}
