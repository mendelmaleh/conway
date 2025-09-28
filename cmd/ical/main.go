package main

import (
	"fmt"

	"github.com/mendelmaleh/conway/hebcal/holiday"
	"github.com/mendelmaleh/conway/utils"
)

func main() {
	year := 5789
	diaspora := true

	for _, v := range holiday.All {
		for _, e := range v.Fill(year, diaspora) {
			fmt.Println(
				e.Date.Roman().Format(utils.ISO8601+" Mon:"),
				e.Name,
				e.Type.Is(holiday.Major),
				e.Start,
			)
		}
	}
}
