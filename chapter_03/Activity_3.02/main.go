package main

import (
	"fmt"
//	"errors"
)


func main() {
  credit_score := 500
  income := 1000
  loan_amount := 1000
  rate := 15
  loan_term := 24
  approved := true
  var monthly_payment float64

  fmt.Println("Applicant 1")
  fmt.Println("-----------")
  fmt.Println("Credit Score    :", credit_score)
  fmt.Println("Income          :", income)
  fmt.Println("Loan Amount     :", loan_amount)

  if credit_score < 450 {
    rate = 20
  }

  total_cost := loan_amount * rate / 100
  monthly_payment = (float64(loan_amount) + float64(total_cost))/float64(loan_term)

  if rate == 20 && monthly_payment > 0.1*float64(income) {
    approved = false
  } else if rate == 15 && monthly_payment > 0.2*float64(income) {
    approved = false
  }


  if loan_term%12 == 0 {
    fmt.Println("Loan Term       :", loan_term)
  } else if credit_score < 0 || monthly_payment < 0 || loan_amount < 0 || loan_term < 0 {
  //  return errors.New("credit_score or monthly_payment or loan_amount loan_term is less than 0")
    approved = false
  } else {
  //  return errors.New("Loan term must be divided by 12")
    approved = false
  }

  fmt.Println("Monthly Payment :", monthly_payment)
  fmt.Println("Rate            :", rate)
  fmt.Println("Total Cost      :", total_cost)
  fmt.Println("Approved        :", approved)

}
