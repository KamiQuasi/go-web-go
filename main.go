package main

import (
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)
    
    //log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}
