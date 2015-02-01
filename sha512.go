package main

import (
    "crypto/sha512"
    "fmt"
    "net/http"
    "io"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
    h512 := sha512.New()
    io.WriteString(h512, r.URL.Path[1:])
    fmt.Fprintf(w, "%x", h512.Sum(nil))
}

func main() {
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8080", nil)
}
