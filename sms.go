package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Request struct {
  Number  string `json:"number"`
  Carrier string `json:"carrier"`
  Hour    string `json:"hour"`
  Minute  string `json:"minute"`
  Message string `json:"message"`
}

const (
  socket = "127.0.0.1:5555"
)

func parseBody(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()

  fmt.Printf("NUMBER => %s\n", r.FormValue("number"))
  fmt.Printf("CARRIER => %s\n", r.FormValue("carrier"))
  fmt.Printf("HOUR => %s\n", r.FormValue("hour"))
  fmt.Printf("MINUTE => %s\n", r.FormValue("minute"))
  fmt.Printf("MESSAGE => %s\n", r.FormValue("message"))

  fmt.Fprintln(w, "Message queued.")
  }

func parseJSONBody(w http.ResponseWriter, r *http.Request) {
  var request Request
  if r.Body == nil {
    http.Error(w, "Please send a request body", 400)
    return
  }

  err := json.NewDecoder(r.Body).Decode(&request)

  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  fmt.Fprintf(w, request.Number)
  fmt.Fprintln(w, "Message queued.")
  }

func main() {
  mux := http.NewServeMux()
  fmt.Printf("Started server at %v.\n", socket)

  mux.HandleFunc("/send", parseBody)

  http.ListenAndServe(socket, mux)
  mux.HandleFunc("/json", parseJSONBody)
}
