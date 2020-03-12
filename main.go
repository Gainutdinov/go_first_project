package main

import (
  "net/http"
  "first_project/controllers"
)

func main() {
    controllers.RegisterControllers()
    http.ListenAndServe(":3000", nil)
}
