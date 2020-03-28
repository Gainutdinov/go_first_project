package main

import (
	"fmt"
)

type Employee struct {
	Id int
	FirstName string
	LastName string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday

)

func (d *Developer) LogHours(wd Weekday, hours int) {
	d.WorkWeek[wd] = hours
}

func (d *Developer) HoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}


func main() {
	dev := Developer{
		Individual:Employee{Id: 1234, FirstName: "Marat", LastName: "Gainutdinov"},
		HourlyRate: 10,
	}
	dev.LogHours(Monday, 8)
	dev.LogHours(Tuesday, 10)
	fmt.Println("Hours worked on Monday: " ,dev.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday: " ,dev.WorkWeek[Tuesday])
	fmt.Printf("Hours worked this week: %d",dev.HoursWorked())
}

