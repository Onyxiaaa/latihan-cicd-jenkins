package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    port := "8080"
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "HALO DARI GUI JENKINS!\n")
    fmt.Fprintf(w, "Sukses Deploy GIT baru 11:14\n")
}
