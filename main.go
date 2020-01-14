package main

import (
  "net/http"
  "os"
  "strings"
)

var prefix string = ""

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = prefix + " " + message + "\n"

  w.Write([]byte(message))
}

func main() {
  var ok bool
  prefix, ok = os.LookupEnv("PREFIX")
  if !ok {
    prefix = "Hello"
  }
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
