package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
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
	log.Println(r.RemoteAddr + " requesting /env")
	for _, e := range os.Environ() {
		w.Write([]byte(e + "\n"))
	}
}

func connectPg(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr + " requesting /pg")
	connStr, ok := os.LookupEnv("connStr")
	if !ok {
		w.Write([]byte("'connStr' environment variable missing"))
		return
	}
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	} else {
		log.Println("Damn, no flush")
	}
	db, err := sql.Open("postgres", connStr)
    log.Println("Connected to db")
	defer db.Close()
	if err != nil {
		log.Fatal("error connecting to db")
		log.Fatal(err)
	}
	w.Write([]byte("success"))
	var result string
	err = db.QueryRow("SELECT 1;").Scan(&result)
    if err != nil {
        log.Fatal(err)
    }
	if result != "1" {
		log.Println("select (failed? )result is not 1???")
	}
	w.Write([]byte(result))
}

func returnEnv(w http.ResponseWriter, r *http.Request) {
	variable := strings.Trim(r.URL.Path, "/")
	log.Println(r.RemoteAddr + " requesting " + variable)
	result, ok := os.LookupEnv(variable)
	if ok {
		w.Write([]byte(result))
	} else {
		w.Write([]byte(notFound))
	}
}

func main() {
	log.Println("env tester starting")
	http.HandleFunc("/", returnEnv)
	http.HandleFunc("/ip", getIp)
	http.HandleFunc("/env", getEnv)
	http.HandleFunc("/pg", connectPg)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
