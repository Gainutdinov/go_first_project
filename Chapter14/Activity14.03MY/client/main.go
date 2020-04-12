package main
import (
		"bytes"
		"fmt"
		"io"
		"io/ioutil"
		"log"
		"mime/multipart"
		"net/http"
		"os"
		"encoding/json"
)
type Names struct {
		Names []string `json:"names"`
}

func postFileAndReturnResponse(filename string) string {
    fileDataBuffer := bytes.Buffer{}
        multipartWritter := multipart.NewWriter(&fileDataBuffer)
        file, err := os.Open(filename)
        if err != nil {
                log.Fatal(err)
        }
        formFile, err := multipartWritter.CreateFormFile("myFile",file.Name())
        if err != nil {
                log.Fatal(err)
        }
        _, err = io.Copy(formFile, file)
        if err != nil {
                log.Fatal(err)
        }
        multipartWritter.Close()
        req, err := http.NewRequest("POST", "http://localhost:8080", &fileDataBuffer)
        if err != nil {
                log.Fatal(err)
        }
        req.Header.Set("Content-Type", multipartWritter.FormDataContentType())
        response, err := http.DefaultClient.Do(req)
        if err != nil {
                log.Fatal(err)
        }
        defer response.Body.Close()
        data, err := ioutil.ReadAll(response.Body)
        if err != nil {
                log.Fatal(err)
        }
        return string(data)
}

func getDataAndParseResponse() (int, int) {

        // send the GET request
        r, err := http.Get("http://localhost:8080")
        if err != nil {
                log.Fatal(err)
        }

        // get data from the response body
        defer r.Body.Close()
        data, err := ioutil.ReadAll(r.Body)
        if err != nil {
                log.Fatal(err)
        }

        names := Names{}
        err = json.Unmarshal(data, &names)
        if err != nil {
                log.Fatal(err)
        }

        electricCount := 0
        boogalooCount := 0
        for _, name := range names.Names {
                if name == "Electric" {
                        electricCount++
                } else if name == "Boogaloo" {
                        boogalooCount++
                }
        }

        // return the count data
        return electricCount, boogalooCount
}


func main() {
        electricCount, boogalooCount := getDataAndParseResponse()
        fmt.Println("Electric Count: ", electricCount)
        fmt.Println("Boogaloo Count: ", boogalooCount)
		data := postFileAndReturnResponse("./test.txt")
		fmt.Println(data)
}

