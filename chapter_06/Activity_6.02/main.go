package main

import (
	"fmt"
	"errors"
)

var (
  ErrInvalidLastName  = errors.New("invalid last name")
  ErrInvalidRoutingNum = errors.New("invalid routing number")
)

type directDeposit struct {
  firstName string
  lastName string
  bankName string
  routingNumber int
  accountNumber int
}

func (d *directDeposit) validateRoutingNumber() error {
  if d.routingNumber < 100 {
    return ErrInvalidRoutingNum
  }
  return nil
}

func (d *directDeposit) validateLastName() error {
  if d.lastName == "" {
    return ErrInvalidLastName
  }
  return nil
}

func (d *directDeposit) report() {
  fmt.Println("***********************")
  fmt.Println("Last Name:", d.lastName)
  fmt.Println("First Name:", d.firstName)
  fmt.Println("Bank Name:", d.bankName)
  fmt.Println("Routing Number:", d.routingNumber)
  fmt.Println("Account Number:", d.accountNumber)
}


func main() {
  dpst := directDeposit{"Abe","","XYZ Inc", 17, 1809}

  err := dpst.validateRoutingNumber()
  if err != nil {
    fmt.Println(err)
  }
  err = dpst.validateLastName()
  if err != nil {
    fmt.Println(err)
  }
  defer dpst.report()
}
