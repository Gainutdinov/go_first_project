package main
import (
  "os"
)
func main() {
  f, err := os.Create("test.txt")
  if err != nil {
  panic(err)
  }
  defer f.Close()
}
