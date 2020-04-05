package main
import "time"
import "fmt"
import "strconv"

func timePassed(start time.Time, end time.Time) string {
	Elapsed := end.Sub(start)
	Seconds := strconv.Itoa(int(Elapsed.Seconds()))
	MicroSeconds := strconv.Itoa(int(Elapsed.Microseconds()))
	return "THe execution took exactly"+Seconds+"."+MicroSeconds+"seconds!"

}

func main() {
	Start := time.Now()
	time.Sleep(2 * time.Second)
	End := time.Now()
	fmt.Println(timePassed(Start, End))
}

