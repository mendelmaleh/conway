package hebcal_test

import (
	"testing"
	"time"

	"github.com/mendelmaleh/conway/hebcal"
	"github.com/mendelmaleh/conway/utils"
)

func TestHeSheIt(t *testing.T) {
	tests := []struct {
		year int
		he   int
		she  int
		it   int
		name string // description of this test case
	}{
		{1999, 49, 60, 20, "1999"},
		{2000, 68, 50, 39, "2000"},
		{2016, 71, 53, 42, "paper example 2016"},
		{2017, 59, 70, 30, "paper example 2017"},
		{2018, 48, 59, 19, "paper example 2018"},
		{2019, 68, 49, 39, "paper example 2019"},
		{2025, 61, 72, 32, "2025"},
		{2026, 50, 61, 21, "2026"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			he, she, it := hebcal.HeSheIt(tc.year)
			if he != tc.he {
				t.Errorf("HeSheIt() = %v, want %v", he, tc.he)
			}
			if she != tc.she {
				t.Errorf("HeSheIt() = %v, want %v", she, tc.she)
			}
			if it != tc.it {
				t.Errorf("HeSheIt() = %v, want %v", it, tc.it)
			}
		})
	}
}

func TestHebrewHeSheIt(t *testing.T) {
	tests := []struct {
		year int
		he   int
		she  int
		it   int
		name string // description of this test case
	}{
		{5776, 52, 53, 42, "paper example 5776"},
		{5777, 71, 70, 30, "paper example 5777"},
		{5778, 59, 59, 19, "paper example 5778"},
		{5779, 48, 49, 39, "paper example 5779"},
		{5785, 71, 72, 32, "5785"},
		{5786, 61, 61, 21, "5786"},
		{5787, 50, 51, 41, "5787"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			he, she, it := hebcal.HebrewHeSheIt(tt.year)
			if he != tt.he {
				t.Errorf("HebrewHeSheIt() = %v, want %v", he, tt.he)
			}
			if she != tt.she {
				t.Errorf("HebrewHeSheIt() = %v, want %v", she, tt.she)
			}
			if it != tt.it {
				t.Errorf("HebrewHeSheIt() = %v, want %v", it, tt.it)
			}
		})
	}
}

func TestRomanHeight(t *testing.T) {
	tests := []struct {
		date time.Time
		want int
		name string // description of this test case
	}{
		{utils.Date(2019, time.April, 7), 11, "paper example"},
		{utils.Date(2020, time.January, 1), 14, "13th month"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hebcal.RomanHeight(tt.date.Month(), tt.date.Day()); got != tt.want {
				t.Errorf("RomanHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
