package main
import "time"
import "fmt"
import "strconv"

func main() {
	now := time.Now()
	Hours := strconv.Itoa(int(now.Hour()))
	Minutes := strconv.Itoa(int(now.Minute()))
	Seconds := strconv.Itoa(int(now.Second()))
	Day := strconv.Itoa(int(now.Day()))
	Month := strconv.Itoa(int(now.Month()))
	Year := strconv.Itoa(int(now.Year()))
	fmt.Println(Hours+":"+Minutes+":"+Seconds+" "+Year+"/"+Month+"/"+Day)
}

