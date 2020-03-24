package main
import (
  "fmt"
)

func main() {
  words := map[string]int{
    "Gonna": 3,
    "You": 3,
    "Give": 2,
    "Never": 1,
    "Up":  4,
  }
	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}
