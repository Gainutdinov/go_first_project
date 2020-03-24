package main
import (
  "fmt"
  "strconv"
)

func main() {
  //sum := 0
  for i := 1; i < 101; i++ {
    //value_str := strconv.Itoa(i)
    if  i%3 == 0 && i%5 == 0 {
      fmt.Println("FizzBuzz")
    } else if i%5 == 0 {
      fmt.Println("Buzz")
    } else if i%3 == 0 {
      fmt.Println("Fizz")
    } else {
      fmt.Println(strconv.Itoa(i))
    }
  }
}
