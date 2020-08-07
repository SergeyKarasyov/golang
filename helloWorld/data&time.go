package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println(current.String())
	fmt.Println("MM-DD-YYYY :", current.Format("01-02-2006"))
	fmt.Println("YYYY-MM-DD hh:mm:ss", current.Format("2006-01-02 15:04:05"))

	septDate := current.AddDate(0, 1, 0)
	fmt.Println(current.String())
	fmt.Println(septDate.String())

	//subtracting from the date
	septDate.Sub(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println(septDate.Sub(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)).String())

	oneLessYears := septDate.AddDate(-1, 0, 0)
	fmt.Println(oneLessYears.String())

	tenMoreMinutes := septDate.Add(10 * time.Minute)
	fmt.Println(tenMoreMinutes)

	//diff between 2 dates
	first := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	second := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)

	difference := second.Sub(first)
	fmt.Printf("Difference %v", difference)

	// str := "2010-09-08T11:15:22.271Z"
	// layout := "2006-01-02T15:04:05.000Z"

	str := "2018-08-08T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.String())
}
