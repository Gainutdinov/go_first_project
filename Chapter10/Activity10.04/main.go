package main

import "fmt"
import "time"
//import "strconv"

func main() {
	now := time.Now()
	fmt.Println("The current time is (ANSIC): ",now.Format(time.ANSIC))
	after := now.Add(time.Second*6 + time.Minute*6 + time.Hour*6)
	fmt.Println("6 hours, 6 minutes and 6 seconds from now the time will be:",after.Format(time.ANSIC))
	//after = .Add(time.Minute * 6)
	//after = now.Add(time.Minute * 6)
}

