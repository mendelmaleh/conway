package hebcal_test

import (
	"testing"
	"time"

	"github.com/mendelmaleh/conway/hebcal"
	"github.com/mendelmaleh/conway/utils"
)

func TestNewYear(t *testing.T) {
	tests := []struct {
		year int
		want time.Time
		name string // description of this test case
	}{
		{2016, utils.Date(2016, time.October, 3), "2016"},
		{2017, utils.Date(2017, time.September, 21), "2017"},
		{2018, utils.Date(2018, time.September, 10), "2018"},
		{2019, utils.Date(2019, time.September, 30), "paper example"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if res := hebcal.NewYear(tc.year); res != tc.want {
				t.Errorf(
					"NewYear() = %v, want %v",
					res.Format(utils.ISO8601),
					tc.want.Format(utils.ISO8601),
				)
			}
		})
	}
}
