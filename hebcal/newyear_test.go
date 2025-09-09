package hebcal_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/mendelmaleh/conway/hebcal"
	"github.com/mendelmaleh/conway/utils"
)

func fmtdate(t time.Time) string {
	return t.Format(utils.ISO8601)
}

func TestNewYear(t *testing.T) {
	tests := []time.Time{
		utils.Date(1999, time.September, 11),

		utils.Date(2015, time.September, 14),
		utils.Date(2016, time.October, 3),
		utils.Date(2017, time.September, 21),

		// 19 year cycle -- credit github.com/bendory/conway-hebrew-calendar
		utils.Date(2018, time.September, 10),
		utils.Date(2019, time.September, 30), // paper example
		utils.Date(2020, time.September, 19),
		utils.Date(2021, time.September, 7),
		utils.Date(2022, time.September, 26),
		utils.Date(2023, time.September, 16),
		utils.Date(2024, time.October, 3),
		utils.Date(2025, time.September, 23),
		utils.Date(2026, time.September, 12),
		utils.Date(2027, time.October, 2),
		utils.Date(2028, time.September, 21),
		utils.Date(2029, time.September, 10),
		utils.Date(2030, time.September, 28),
		utils.Date(2031, time.September, 18),
		utils.Date(2032, time.September, 6),
		utils.Date(2033, time.September, 24),
		utils.Date(2034, time.September, 14),
		utils.Date(2035, time.October, 4),
		utils.Date(2036, time.September, 22),

		utils.Date(1899, time.September, 5),
	}

	for _, tc := range tests {
		// test new year calculation
		t.Run(strconv.Itoa(tc.Year()), func(t *testing.T) {
			if got := hebcal.NewYear(tc.Year()); got != tc {
				t.Errorf("NewYear() = %v, want %v", fmtdate(got), fmtdate(tc))
			}
		})

		hebrew := hebcal.HebrewDate{tc.Year() + 3761, hebcal.Tishrei, 1}
		// test hebrew to roman conversion
		if got := hebrew.Roman(); got != tc {
			t.Errorf("(HebrewDate).Roman() = %v, want %v", got, hebrew)
		}

		// test roman to hebrew conversion
		if got := hebcal.FromRoman(tc); got != hebrew {
			t.Errorf("FromRoman() = %v, want %v", got, hebrew)
		}
	}
}
