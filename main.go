package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Location : ", t.Location(), " Time : ", t) //local time

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Location : ", location, "Time : ", t.In(location)) // America/New_York

	loc, _ := time.LoadLocation("Europe/London")
	now := time.Now().In(loc)
	fmt.Println("Location : ", loc, " Time : ", now) // Europe/London

	tomorrow := time.Date(2022, 2, 18, 0, 0, 0, 0, time.UTC)

	fmt.Println(tomorrow)

}
