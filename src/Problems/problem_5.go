// adding a date to the current day

package main

import(
	"fmt"
	"time"
)

func main(){
	today := time.Date(2020, time.December, 23,0,0,0,0,time.UTC) // today
	// fotmat (int:year, time.month, int:day, int: hour, int:min, int: sec, int: nsec, * location,err)
	days_15 := today.AddDate(0,0,15) // fifteen days from today
	// format date.AddDate(+- year, +- month, +- day)
	months_15:= today.AddDate(0,15,0) // fifteen months from today
	years_15 := today.AddDate(15,0,0) // fifteen years from today

	fmt.Println("Today:= ", today)
	fmt.Println("15 days from today:= ", days_15)
	fmt.Println("15 months from today:= ", months_15)
	fmt.Println("15 years from today:= ", years_15)


}