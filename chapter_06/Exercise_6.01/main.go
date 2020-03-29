package main

import (
	"fmt"
	"errors"
)

func main() {
	var InvalidLastName = errors.New("InvalidLastName text")
	var InvalidRoutingNumber = errors.New("InvalidRoutingNumber text")
	fmt.Println(InvalidLastName)
	fmt.Println(InvalidRoutingNumber)
}
