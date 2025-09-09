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
		utils.Date(2016, time.October, 3),
		utils.Date(2017, time.September, 21),
		utils.Date(2018, time.September, 10),
		utils.Date(2019, time.September, 30), // paper example
	}
	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.Year()), func(t *testing.T) {
			if got := hebcal.NewYear(tc.Year()); got != tc {
				t.Errorf("NewYear() = %v, want %v", fmtdate(got), fmtdate(tc))
			}
		})
	}
}
