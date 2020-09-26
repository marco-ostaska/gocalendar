package main

import (
	"flag"

	"github.com/marco-ostaska/gocalendar/internal/calendar"
)

func main() {
	m := flag.String("month", "", "Month to display calendar using shortName")
	y := flag.String("year", "", "Year using RFC822.")

	flag.Parse()

	if len(*m) == 0 || len(*y) == 0 {
		flag.Usage()
		return
	}

	calendar.Calendar(*m, *y)

}
