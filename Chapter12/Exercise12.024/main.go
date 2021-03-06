package main
import (
  "fmt"
  "os"
)
func main() {
  file, err := os.Stat("junk.txt")
  if err != nil {
  if os.IsNotExist((err)) {
    fmt.Println("junk.txt: File does not exist!")
    fmt.Println(file)
  }
  }
  fmt.Println()
  file, err = os.Stat("test.txt")
  if err != nil {
  if os.IsNotExist((err)) {
    fmt.Println("test.txt: File does not exist!")
  }
  }
  fmt.Printf("file name: %s\nIsDir: %t\nModTime: %v\nMode: %v\nSize: %d\n", file.Name(), file.IsDir(), file.ModTime(), file.Mode(), file.Size())
}
