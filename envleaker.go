package main

import (
  "net/http"
  "strings"
  "io/ioutil"
  "fmt"
  "os"
)

var notFound string = "Not found"

func getIp(w http.ResponseWriter, r *http.Request) {
  resp, err := http.Get("http://api.ipify.org/")
  if err != nil {
    w.Write([]byte(err.Error()))
  } else {
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      w.Write([]byte(err.Error()))
    }
    w.Write([]byte(body))
  }
}

func getEnv(w http.ResponseWriter, r *http.Request) {
  for _, e := range os.Environ() {
    w.Write([]byte(e + "\n"))
  }
}

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
  http.HandleFunc("/ip", getIp)
  http.HandleFunc("/env", getEnv)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
