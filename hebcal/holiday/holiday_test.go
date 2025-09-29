package holiday

import (
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/mendelmaleh/conway/hebcal"
	"github.com/mendelmaleh/conway/utils"
)

func fmtdate(t time.Time) string {
	return t.Format(utils.ISO8601)
}

func TestHebcal(t *testing.T) {
	for year := 5600; year < 9600; year++ {
		for _, e := range All {
			for _, d := range e.Fill(year, true) {
				got := d.Date.Roman()
				norm := hebcal.FromRoman(got)
				want := hdate.New(year, hdate.HMonth(norm.Month), norm.Day).Gregorian()
				if got := d.Date.Roman(); got != want {
					t.Errorf("%v.Roman() = %v, want %v (%v)", d.Date, fmtdate(got), fmtdate(want), norm)
				}

			}
		}
	}
}
