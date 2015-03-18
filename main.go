package main

import (
    "net/http"
    "flag"
    "fmt"
    "os/exec"
    "runtime"
)

func main() {
    var port = flag.Int("p", 8080, "Port to listen on")
    var err error

    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)

    fmt.Println(fmt.Sprintf("Listening on port: %d...\n", *port))
    http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)

    fmt.Sprintf("http://localhost:%d/", *port)

    fmt.Println(runtime.GOOS)

    switch runtime.GOOS {
        case "linux":
            err = exec.Command("xdg-open", fmt.Sprintf("http://localhost:%d/", *port)).Start()
        case "windows", "darwin":
            err = exec.Command("open", fmt.Sprintf("http://localhost:%d/", *port)).Start()
        default:
            err = fmt.Errorf("unsupported platform")
    }

    if err != nil {
        fmt.Print(err)
    }
}
