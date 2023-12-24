package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Weekday())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	date := time.Date(2023, time.October, 0, 24, 0, 0, 0, time.UTC)
	fmt.Println(date)

	dateValue := "2023-10-19"
	parseDate, _ := time.Parse("2006-01-02", dateValue)
	parseDate_rfc822, err := time.Parse(time.RFC822, "19 Oct 23 20:19 UTC")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n%v\n%v\n", parseDate, parseDate_rfc822)
}
