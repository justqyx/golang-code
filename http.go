package main

import (
    "fmt"
    "log"
    "net/http"
    "regexp"
    "runtime/debug"
)

var visitors int

func handleHi(w http.ResponseWriter, r *http.Request) {
    if match, _ := regexp.MatchString(`^\w*$`, r.FormValue("color")); !match {
        http.Error(w, "Optional color is invalid", http.StatusBadRequest)
        return
    }

    visitors++

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write([]byte("Welcome! Your are visitor number: " + fmt.Sprint(visitors) + "!"))
}

func main() {
    log.Printf("Starting on port 4040")
    http.HandleFunc("/hi", handleHi)
    log.Fatal(http.ListenAndServe("127.0.0.1:4040", nil))
}
