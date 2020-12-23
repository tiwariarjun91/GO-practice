// different time formats
// Problem 6

package main 

import(
	"fmt"
	"time"
)

func main(){
	const format = "2020-23-12"
	timeset, _ := time.Parse(format, "1997-24-12")

	fmt.Println("unix format ;- ", timeset.Format(time.UnixDate))
}