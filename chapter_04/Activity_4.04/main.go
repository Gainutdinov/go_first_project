package main
import (
  "fmt"
)

func main() {
  week_days := []string{"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}
  week_days = append(week_days[6:], week_days[:6]...)
  //week_days = append(week_days,week_days[:len(week_days)-1]...)
  //week_days = week_days[len(week_days)-7:]
  fmt.Println(week_days)
}
