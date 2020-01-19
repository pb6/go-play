package main

import (
  "net/http"
  "strings"
  "fmt"
  "os"
)

var notFound string = "Not found"

func returnEnv(w http.ResponseWriter, r *http.Request) {
  variable := strings.Trim(r.URL.Path, "/")
  fmt.Print(r.RemoteAddr + " requesting " + variable + "\n")
  result, ok := os.LookupEnv(variable)
  if ok {
    w.Write([]byte(result))
  } else {
    w.Write([]byte(notFound))
  }
}

func main() {
  fmt.Print("env leaker starting\n")
  http.HandleFunc("/", returnEnv)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
