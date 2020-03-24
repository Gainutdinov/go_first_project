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
  the_most_popular_word := ""
  count := 0
	for key, value := range words {
		if count < value {
			the_most_popular_word = key
			count = value
		}
	}
	fmt.Println("Most popular word:",  the_most_popular_word)
	fmt.Println("With a count of:",  count)
}
