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

func nonLoggedHours() func(int) int {
	total := 0
	return func(nonlgdhrs int) int {
		total += nonlgdhrs
		return nonlgdhrs
	}
}

func (d *Developer) PayDay() (int, bool) {
	if d.HoursWorked() > 40 {
		hoursOver := d.HoursWorked() - 40
		overTime := hoursOver * 2
		regularPay := d.HoursWorked() * d.HourlyRate
		return regularPay + overTime, true
	}
	return d.HoursWorked() * d.HourlyRate, false
}

func (d *Developer) PayDetails() {
	for i, v := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ",v)
		case 1:
			fmt.Println("Monday hours: ",v)
		case 2:
			fmt.Println("Tuesday hours: ",v)
		case 3:
			fmt.Println("Wednesday hours: ",v)
		case 4:
			fmt.Println("Thursday hours: ",v)
		case 5:
			fmt.Println("Friday hours: ",v)
		case 6:
			fmt.Println("Saturday hours: ",v)
		}
	}
}


func main() {
	dev := Developer{
		Individual:Employee{Id: 1234, FirstName: "Marat", LastName: "Gainutdinov"},
		HourlyRate: 10,
	}
	fmt.Println("Hours worked on Monday: " ,dev.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday: " ,dev.WorkWeek[Tuesday])
	fmt.Printf("Hours worked this week: %d",dev.HoursWorked())
	dev.LogHours(Monday, 8)
	dev.LogHours(Tuesday, 10)
	dev.LogHours(Wednesday, 10)
	dev.LogHours(Thursday, 10)
	dev.LogHours(Friday, 6)
	dev.LogHours(Saturday, 8)
	dev.PayDetails()
}

