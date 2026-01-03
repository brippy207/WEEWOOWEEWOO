package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    // API routes

    // Serve files from static folder
    http.Handle("/", http.FileServer(http.Dir("./static")))

    // Serve api /TonyRippyis
    http.HandleFunc("/TonyRippyis", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "old, weak, bald, and short")
    })

    port := ":5000"
    fmt.Println("Server is running on port" + port)

    // Start server on port specified above
    log.Fatal(http.ListenAndServe(port, nil))

}