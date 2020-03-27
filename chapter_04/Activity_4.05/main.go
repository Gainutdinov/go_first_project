package main
import (
  "fmt"
)

func main() {
  slice := []string{"Good","Good","Bad","Good","Good"}
  slice = append(slice[:2], slice[3:]...)
  //week_days = append(week_days,week_days[:len(week_days)-1]...)
  //week_days = week_days[len(week_days)-7:]
  fmt.Println(slice)
}
