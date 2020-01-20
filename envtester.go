package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	fmt.Print(r.RemoteAddr + " requesting " + variable + "\n")
	for _, e := range os.Environ() {
		w.Write([]byte(e + "\n"))
	}
}

func connectPg(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.RemoteAddr + " requesting " + variable + "\n")
    connStr, _ := os.LookupEnv("connStr")
    db, err := sql.Open("postgres", connStr)
    defer db.Close()
    if err != nil {
         w.Write([]byte(err.Error()))
    }
    w.Write([]byte("success"))
    row := db.QueryRow("SELECT 1")
    var result string
    row.Scan(&result)
    w.Write([]byte(result))
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
	fmt.Print("env tester starting\n")
	http.HandleFunc("/", returnEnv)
	http.HandleFunc("/ip", getIp)
	http.HandleFunc("/env", getEnv)
	http.HandleFunc("/pg", connectPg)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
