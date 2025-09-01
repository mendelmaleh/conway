package hebcal_test

import (
	"testing"
	"time"

	"github.com/mendelmaleh/conway/hebcal"
	"github.com/mendelmaleh/conway/utils"
)

func TestHebrewMonth_Partner(t *testing.T) {
	tests := []struct {
		month hebcal.HebrewMonth
		want  time.Month
	}{
		{hebcal.Tishrei, time.August},
		{hebcal.Cheshvan, time.September},
		{hebcal.Kislev, time.October},
		{hebcal.Tevet, time.November},
		{hebcal.Shevat, time.December},
		{hebcal.AdarI, time.January},
		{hebcal.AdarII, time.February},
		{hebcal.Nissan, time.March},
		{hebcal.Iyar, time.April},
		{hebcal.Sivan, time.May},
		{hebcal.Tamuz, time.June},
		{hebcal.Av, time.July},
		{hebcal.Elul, time.August},
	}
	for _, tt := range tests {
		t.Run(tt.month.String(), func(t *testing.T) {
			if got := tt.month.Partner(); got != tt.want {
				t.Errorf("Partner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHebrewDate_Height(t *testing.T) {
	tests := []struct {
		date hebcal.HebrewDate
		want int
		name string // description of this test case
	}{
		{hebcal.HebrewDate{5760, hebcal.AdarII, 7}, 57, "moshe"},
		{hebcal.HebrewDate{5779, hebcal.Iyar, 7}, 46, "paper example"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Height(); got != tt.want {
				t.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHebrewDate_Roman(t *testing.T) {
	tests := []struct {
		date hebcal.HebrewDate
		want time.Time
		name string // description of this test case
	}{
		{
			hebcal.HebrewDate{5760, hebcal.Tishrei, 1},
			utils.Date(1999, time.September, 11), "rosh hashana",
		},
		{
			hebcal.HebrewDate{5760, hebcal.AdarII, 7},
			utils.Date(2000, time.March, 14),
			"moshe",
		},
		{
			hebcal.HebrewDate{5779, hebcal.Iyar, 7},
			utils.Date(2019, time.May, 12),
			"paper example",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Roman(); got != tt.want {
				t.Errorf("Roman() = %v, want %v", got, tt.want)
			}
		})
	}
}
