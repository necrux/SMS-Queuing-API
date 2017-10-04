package main

import (
  //"encoding/json"
  "fmt"
  "net/http"
)

const (
  port = ":5555"
)

func main() {
  mux := http.NewServeMux()
  fmt.Printf("Started server at http://localhost%v.\n", port)

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    fmt.Printf("NUMBER => %s\n", r.FormValue("number"))
    fmt.Printf("CARRIER => %s\n", r.FormValue("carrier"))
    fmt.Printf("HOUR => %s\n", r.FormValue("hour"))
    fmt.Printf("MINUTE => %s\n", r.FormValue("minute"))
    fmt.Printf("MESSAGE => %s\n", r.FormValue("message"))

    fmt.Fprintln(w, "Message queued.")
  })

  http.ListenAndServe(port, mux)
}
