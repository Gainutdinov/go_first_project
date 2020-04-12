package main
import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "math/rand"
        "encoding/json"
)
type server struct{}
type Names struct {
        Names []string `json:"names"`
}


func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   uploadedFile, uploadedFileHeader, err := r.FormFile("myFile")
   fmt.Println(r.Method)
   if r.Method == "POST" {
		if err != nil {
		   log.Fatal(err)
		}
		defer uploadedFile.Close()
		fileContent, err := ioutil.ReadAll(uploadedFile)
		if err != nil {
		   log.Fatal(err)
		}
		err = ioutil.WriteFile(fmt.Sprintf("./%s", uploadedFileHeader.Filename), fileContent, 0666)
		if err != nil {
		   log.Fatal(err)
		}
		w.Write([]byte(fmt.Sprintf("%s Uploaded!", uploadedFileHeader.Filename)))
   } else { // GET
	    names := Names{}
        // Generate random number of 'Electric' names
        for i := 0; i < rand.Intn(5)+1; i++ {
                names.Names = append(names.Names, "Electric")
        }
        // Generate random number of 'Boogaloo' names
        for i := 0; i < rand.Intn(5)+1; i++ {
                names.Names = append(names.Names, "Boogaloo")
        }
        // convert struct to bytes
        jsonBytes, _ := json.Marshal(names)
        log.Println(string(jsonBytes))
        w.Write(jsonBytes)
   }
}
func main() {
   log.Fatal(http.ListenAndServe(":8080", server{}))
}

