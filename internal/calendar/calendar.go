// Package calendar is an internal packe to parse dates
// to be used for gocalendar
package calendar

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Layout is just a wrapper to time.RFC822
// used to parse "02 Jan 06 15:04 MST"
const Layout = time.RFC822

// Months variable contains the months shortname
var Months = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

// MonthLayout function used to format a layout used for time.Parse
// used for time.RFC822
func MonthLayout(m, y string) time.Time {
	s := "01 " + m + " " + y + " 00:00 BRT"
	t, err := time.Parse(Layout, s)

	if err != nil {
		log.Fatalln(err)
	}

	return t
}

func lastDay(t time.Time) int {

	d := strings.Split(t.AddDate(0, 1, 0).Sub(t).String(), "h")
	day, err := strconv.Atoi(d[0])

	if err != nil {
		log.Fatalln(err)
	}

	return day / 24
}

func parseWeek(t time.Time) [6][7]int {
	var wArray [6][7]int
	var wCount int

	ld := lastDay(t)

	for i := 0; i < ld; i++ {
		wday := int(t.Weekday())
		wArray[wCount][wday] = t.Day()
		t = t.AddDate(0, 0, 1)

		if wday == 6 {
			wCount++
		}
	}

	return wArray
}

func printCalendar(wArray [6][7]int) {

	for _, v := range wArray {
		for _, d := range v {
			if d > 0 {
				fmt.Printf("%5d", d)
			} else {
				fmt.Printf("%5v", "")
			}

		}
		fmt.Println()
	}

}

// Calendar print the calendar
func Calendar(m, y string) {

	t := MonthLayout(m, y)
	fmt.Println(m, t.Year())
	fmt.Printf("%5s%5s%5s%5s%5s%5s%5s", "Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat")
	fmt.Println()
	wArray := parseWeek(t)
	printCalendar(wArray)

}
