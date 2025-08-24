package hebcal_test

import (
	"testing"

	"github.com/mendelmaleh/conway/hebcal"
)

func TestHeSheIt(t *testing.T) {
	tests := []struct {
		year int
		he   int
		she  int
		it   int
		name string // description of this test case
	}{
		{2016, 71, 53, 42, "paper example 2016"},
		{2017, 59, 70, 30, "paper example 2017"},
		{2018, 48, 59, 19, "paper example 2018"},
		{2019, 68, 49, 39, "paper example 2019"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			he, she, it := hebcal.HeSheIt(tc.year)
			if it != tc.it {
				t.Errorf("HeSheIt() = %v, want %v", it, tc.it)
			}
			if he != tc.he {
				t.Errorf("HeSheIt() = %v, want %v", he, tc.he)
			}
			if she != tc.she {
				t.Errorf("HeSheIt() = %v, want %v", she, tc.she)
			}
		})
	}
}
