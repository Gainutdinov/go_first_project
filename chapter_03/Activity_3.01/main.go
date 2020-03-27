package main

import (
	"fmt"
)

func main() {
  goods := [3]string{"Cake", "Milk", "Butter"}

  cost := map[string]float32{
    "Cake": 0.99,
    "Milk": 2.75,
    "Butter": 0.87,
  }

  taxes := map[string]float32{
    "Cake": 7.5,
    "Milk": 1.5,
    "Butter": 2,
  }

  var tax_total float32

  for good := range goods{
    tax_total = tax_total + cost[goods[good]] * taxes[goods[good]]/100
  }
  fmt.Println("Sales Tax Total:", tax_total)
}
