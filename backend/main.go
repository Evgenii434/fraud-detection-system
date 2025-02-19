package main

import (
    "log"
    "net/http"
    "github.com/yourusername/fraud-detection-system/api"
)

func main() {
    router := api.NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}